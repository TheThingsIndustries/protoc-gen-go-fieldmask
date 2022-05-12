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

func (g *generator) fileHasFieldPathNormalizer() bool {
	for _, message := range appendSubMessages(nil, g.file.Messages...) {
		if g.messageHasFieldPathNormalizer(message) {
			return true
		}
	}
	return false
}

func (*generator) messageHasFieldPathNormalizer(message *protogen.Message) bool {
	// No code is generated for map entries, so we also don't need to generate field_path_normalizers.
	if message.Desc.IsMapEntry() {
		return false
	}

	var generateFieldPathNormalizer bool

	// If the file has the (thethings.fieldmask.file) option, and field_path_normalizer_all is set, we start with that.
	fileOpts := message.Desc.ParentFile().Options().(*descriptorpb.FileOptions)
	if proto.HasExtension(fileOpts, annotations.E_File) {
		if fileExt, ok := proto.GetExtension(fileOpts, annotations.E_File).(*annotations.FileOptions); ok {
			if fileExt.FieldPathNormalizerAll != nil {
				generateFieldPathNormalizer = *fileExt.FieldPathNormalizerAll
			}
		}
	}

	// If the message has the (thethings.fieldmask.message) option, the field_path_normalizer field can still override to true or false if explicitly set.
	messageOpts := message.Desc.Options().(*descriptorpb.MessageOptions)
	if proto.HasExtension(messageOpts, annotations.E_Message) {
		if messageExt, ok := proto.GetExtension(messageOpts, annotations.E_Message).(*annotations.MessageOptions); ok {
			if messageExt.FieldPathNormalizer != nil {
				generateFieldPathNormalizer = *messageExt.FieldPathNormalizer
			}
		}
	}

	return generateFieldPathNormalizer
}

func (g *generator) genFieldPathNormalizer(message *protogen.Message) { //nolint:gocyclo
	if !g.messageHasFieldPathNormalizer(message) {
		return
	}

	g.P("// NormalizeFieldPaths normalizes the field paths.")
	g.P("func (x *", message.GoIdent, ") NormalizeFieldPaths(paths ...string) ([]string, error) {")

	// This is the set of fields we'll use to skip duplicates.
	g.P("var (")
	g.P("normalizedPaths []string")
	g.P("fset = make(", pp.Ident("FieldSet"), ", ", len(message.Fields)+len(message.Oneofs), ")")
	g.P(")")

	// We first look at exact fields (without sub-fields).
	// That way, if paths contains both the field, and sub-fields, we can skip the logic for sub-paths.
	g.P("for _, field := range ", pp.Ident("TopLevelPaths"), "(paths) {")

	// Skip duplicate fields.
	g.P("if fset.Contains(field) {")
	g.P("continue")
	g.P("}")

	g.P("switch field {")

	// Error if we don't know the field.
	g.P("default:")
	g.P("return nil, ", pp.Ident("FieldErrorf"), `("`, message.GoIdent, `", field, "unknown field")`)

	for _, field := range message.Fields {
		// The "case" supports both the proto name and the JSON name, if those are different.
		if string(field.Desc.Name()) != field.Desc.JSONName() {
			g.P(`case "`, field.Desc.Name(), `", "`, field.Desc.JSONName(), `":`)
		} else {
			g.P(`case "`, field.Desc.Name(), `":`)
		}
		g.P(`normalizedPaths = append(normalizedPaths, "`, field.Desc.Name(), `")`)
	}

	for _, oneof := range message.Oneofs {
		// For each oneof, we can set the oneof directly, so we allow it here.
		g.P(`case "`, oneof.Desc.Name(), `":`) // NOTE: oneofs don't have a JSONName.
		g.P(`normalizedPaths = append(normalizedPaths, "`, oneof.Desc.Name(), `")`)
	}

	g.P("}") // end switch field
	g.P("fset.Add(field)")
	g.P("}") // end for _, field := range ...

	// After we've completed handling the "exact match" paths, we look at everything that can have nested fields.
	fieldsWithSubFields := filterFields(message.Fields, fieldIsMessage)
	if len(fieldsWithSubFields) > 0 || len(message.Oneofs) > 0 {
		g.P("for _, field := range ", pp.Ident("SubPaths"), "(paths) {")

		g.P("topLevelField := ", pp.Ident("TopLevelField"), "(field)")

		// Skip duplicate fields.
		g.P("if fset.Contains(topLevelField) {")
		g.P("continue")
		g.P("}")

		g.P("switch topLevelField {")

		// Error if we don't know the field.
		g.P("default:")
		g.P("return nil, ", pp.Ident("FieldErrorf"), `("`, message.GoIdent, `", field, "unknown field")`)

		for _, field := range fieldsWithSubFields {
			// The "case" supports both the proto name and the JSON name, if those are different.
			if string(field.Desc.Name()) != field.Desc.JSONName() {
				g.P(`case "`, field.Desc.Name(), `", "`, field.Desc.JSONName(), `":`)
			} else {
				g.P(`case "`, field.Desc.Name(), `":`)
			}

			g.P("normalizedSubFields, err := ((*", field.Message.GoIdent, ")(nil)).NormalizeFieldPaths(", pp.Ident("SubPathsOf"), "(paths, topLevelField)...)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("normalizedPaths = append(normalizedPaths, ", pp.Ident("PrefixPaths"), "(normalizedSubFields, topLevelField)...)")
		}

		for _, oneof := range message.Oneofs {
			// Even though the name of the oneof should not be part of a field mask sub-path, we allow it anyway.
			g.P(`case "`, oneof.Desc.Name(), `":`) // NOTE: oneofs don't have a JSONName.
			g.P("normalizedSubFields, err := x.NormalizeFieldPaths(", pp.Ident("SubPathsOf"), "(paths, topLevelField)...)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("normalizedPaths = append(normalizedPaths, normalizedSubFields...)")
		}

		g.P("}") // end switch topLevelField
		g.P("fset.Add(topLevelField)")
		g.P("}") // end for _, field := range ...
	}

	g.P("return normalizedPaths, nil")
	g.P("}") // end func NormalizeFieldPaths
}
