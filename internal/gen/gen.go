// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package gen

import (
	"fmt"
	"strings"

	"github.com/TheThingsIndustries/protoc-gen-go-fieldmask/internal/gogoproto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Version is the version of the generator.
var Version = "0.0.0-dev"

// Params are the parameters for the generator.
var Params struct {
	Lang string
}

const (
	fmtPackage = protogen.GoImportPath("fmt")
	pp         = protogen.GoImportPath("github.com/TheThingsIndustries/protoc-gen-go-fieldmask/fieldmaskplugin")
)

type generator struct {
	gen  *protogen.Plugin
	file *protogen.File
	*protogen.GeneratedFile
}

// GenerateFile generates a file with FieldSetters.
func GenerateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	g := &generator{
		gen:  gen,
		file: file,
	}

	// If the file doesn't have FieldSetters, we can skip it.
	if !g.fileHasFieldSetter() && !g.fileHasFieldPaths() {
		return nil
	}

	// Generate a new file that ends with `_fieldmask.pb.go`.
	filename := file.GeneratedFilenamePrefix + "_fieldmask.pb.go"
	g.GeneratedFile = gen.NewGeneratedFile(filename, file.GoImportPath)

	// The standard header for generated files.
	g.P("// Code generated by protoc-gen-go-fieldmask. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-fieldmask v", Version)
	g.P("// - protoc             ", g.protocVersion())
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}

	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	g.generateFileContent()

	return g.GeneratedFile
}

func containsMessage(s []*protogen.Message, search *protogen.Message) bool {
	for _, el := range s {
		if el == search {
			return true
		}
	}
	return false
}

func appendSubMessages(dst []*protogen.Message, msgs ...*protogen.Message) []*protogen.Message {
	for _, msg := range msgs {
		if !containsMessage(dst, msg) {
			dst = append(dst, msg)
		}
		for _, subMsg := range msg.Messages {
			if !containsMessage(dst, subMsg) {
				dst = appendSubMessages(dst, subMsg)
			}
		}
	}
	return dst
}

func (g *generator) protocVersion() string {
	v := g.gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

func (g *generator) generateFileContent() {
	for _, message := range g.file.Messages {
		g.genMessage(message)
	}
}

func fieldGoName(field *protogen.Field) interface{} {
	var fieldGoName interface{} = field.GoName
	if Params.Lang == "gogo" {
		// If we use gogo, the GoName of a field can be overridden with the (gogoproto.customname) option.
		if proto.HasExtension(field.Desc.Options(), gogoproto.E_Customname) {
			fieldGoName = proto.GetExtension(field.Desc.Options(), gogoproto.E_Customname).(string)
		}
		// Fields with the (gogoproto.embed) option should use the name of the embedded message.
		if proto.HasExtension(field.Desc.Options(), gogoproto.E_Embed) {
			if proto.GetExtension(field.Desc.Options(), gogoproto.E_Embed).(bool) {
				fieldGoName = field.Message.GoIdent
			}
		}
	}
	return fieldGoName
}

func fieldIsNullable(field *protogen.Field) bool {
	// Typically, only message fields are nullable (use pointers).
	nullable := field.Desc.Kind() == protoreflect.MessageKind

	// If we use gogo, the nullability of fields (messages, but also bytes fields with custom types) is controlled with the (gogoproto.nullable) option.
	if Params.Lang == "gogo" {
		if proto.HasExtension(field.Desc.Options(), gogoproto.E_Customtype) {
			nullable = true
		}
		if proto.HasExtension(field.Desc.Options(), gogoproto.E_Nullable) {
			nullable = proto.GetExtension(field.Desc.Options(), gogoproto.E_Nullable).(bool)
		}
	}

	return nullable
}

func filterFields(fields []*protogen.Field, selector func(f *protogen.Field) bool) []*protogen.Field {
	out := make([]*protogen.Field, 0, len(fields))
	for _, path := range fields {
		if selector(path) {
			out = append(out, path)
		}
	}
	return out
}

func fieldIsMessage(f *protogen.Field) bool {
	if f.Message == nil || f.Desc.IsList() || f.Desc.IsMap() {
		return false
	}
	if strings.HasPrefix(string(f.Message.Desc.FullName()), "google.protobuf.") {
		return false
	}
	return true
}

func (g *generator) genMessage(message *protogen.Message) { //nolint:gocyclo
	// Generate for all sub-messages defined in the message.
	for _, message := range message.Messages {
		g.genMessage(message)
	}

	g.genFieldPaths(message)

	g.genFieldSetter(message)
}
