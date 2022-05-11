// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package gen

import (
	"github.com/TheThingsIndustries/protoc-gen-go-fieldmask/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (g *generator) fileHasFieldPaths() bool {
	for _, message := range appendSubMessages(nil, g.file.Messages...) {
		if g.messageHasFieldPaths(message) {
			return true
		}
	}
	return false
}

func (*generator) messageHasFieldPaths(message *protogen.Message) bool {
	// No code is generated for map entries, so we also don't need to generate field_path_normalizers.
	if message.Desc.IsMapEntry() {
		return false
	}

	var generateFieldPaths bool

	// If the file has the (thethings.fieldmask.file) option, and field_path_normalizer_all is set, we start with that.
	fileOpts := message.Desc.ParentFile().Options().(*descriptorpb.FileOptions)
	if proto.HasExtension(fileOpts, annotations.E_File) {
		if fileExt, ok := proto.GetExtension(fileOpts, annotations.E_File).(*annotations.FileOptions); ok {
			if fileExt.FieldPathsAll != nil {
				generateFieldPaths = *fileExt.FieldPathsAll
			}
		}
	}

	// If the message has the (thethings.fieldmask.message) option, the field_path_normalizer field can still override to true or false if explicitly set.
	messageOpts := message.Desc.Options().(*descriptorpb.MessageOptions)
	if proto.HasExtension(messageOpts, annotations.E_Message) {
		if messageExt, ok := proto.GetExtension(messageOpts, annotations.E_Message).(*annotations.MessageOptions); ok {
			if messageExt.FieldPaths != nil {
				generateFieldPaths = *messageExt.FieldPaths
			}
		}
	}

	return generateFieldPaths
}

func (g *generator) genFieldPaths(message *protogen.Message) { //nolint:gocyclo
	if !g.messageHasFieldPaths(message) {
		return
	}

	fieldsWithSubFields := filterFields(message.Fields, fieldIsMessage)

	g.P("// FieldPaths returns the field paths up to the given maximum depth.")
	g.P("func (x *", message.GoIdent, ") FieldPaths(maxDepth int) []string {")
	g.P("if maxDepth == 0 {")
	g.P("return nil")
	g.P("}")

	if len(fieldsWithSubFields) > 0 {
		g.P("if maxDepth > 1 {")
		g.P("return ", pp.Ident("Flatten"), "([][]string{")
		for _, field := range message.Fields {
			g.P(`{"`, field.Desc.Name(), `"},`)
			if fieldIsMessage(field) {
				g.P(pp.Ident("PrefixPaths"), "(((*", field.Message.GoIdent, ")(nil)).FieldPaths(maxDepth-1), ", `"`, field.Desc.Name(), `"),`)
			}
		}
		g.P("})")
		g.P("}")
	}

	g.P("return []string{")
	for _, field := range message.Fields {
		g.P(`"`, field.Desc.Name(), `",`)
	}
	g.P("}")

	g.P("}")
}
