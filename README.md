# protoc-gen-go-fieldmask

> Protoc plugin for generating fieldmask-aware methods for proto messages in Go

## Background

The Things Stack works a lot with fieldmasks. This plugin generates methods that help us work with those field masks:

- A `FieldPaths` method that lets us get the full list of field paths supported by the message (and its sub-messages).  
    For a specific message: `(thethings.fieldmask.message) = { field_paths: true }`  
    For the entire file: `option (thethings.fieldmask.file) = { field_paths_all: true }`
- A `NormalizeFieldPaths` method that lets us normalize a list of field paths for the message.  
    For a specific message: `(thethings.fieldmask.message) = { field_path_normalizer: true }`  
    For the entire file: `option (thethings.fieldmask.file) = { field_path_normalizer_all: true }`
- A `SetFields` method that lets us set fields specified by such a field mask from a source struct into a destination struct.  
    For a specific message: `(thethings.fieldmask.message) = { field_setter: true }`  
    For the entire file: `option (thethings.fieldmask.file) = { field_setters_all: true }`

## Usage in Proto Code

```proto
syntax = "proto3";

import "github.com/TheThingsIndustries/protoc-gen-go-fieldmask/annotations.proto";

package thethings.fieldmask.example;

option go_package = "github.com/TheThingsIndustries/protoc-gen-go-fieldmask/example";

option (thethings.fieldmask.file) = {
  field_paths_all: true,             // Generate FieldPaths methods for everything in the file.
  field_path_normalizer_all: true,   // Generate field path normalizers for everything in the file.
  field_setters_all: true,           // Generate field setters for everything in the file.
};

message SomeMessage {
  string some_field = 1;
  string other_field = 2;
}

message MessageWithoutFieldPathsMethod {
  option (thethings.fieldmask.message) = { field_paths: false };

  string some_field = 1;
}

message MessageWithoutFieldSetter {
  option (thethings.fieldmask.message) = { field_setter: false };

  string some_field = 1;
}

message MessageWithoutFieldPathNormalizer {
  option (thethings.fieldmask.message) = { field_path_normalizer: false };

  string some_field = 1;
}
```

## Generating Go Code

```bash
$ protoc -I ./path/to -I . \
  --go_opt=paths=source_relative --go_out=./path/to \
  --go-fieldmask_opt=paths=source_relative --go-fieldmask_out=./path/to \
  ./path/to/*.proto
```

## Usage in Go Code

```go
dst := &SomeMessage{
  SomeField: "some value",
  OtherField: "other value",
}

src := &SomeMessage{
  SomeField: "new value",
  OtherField: "skipped value",
}

var src, dst *SomeMessage

normalizedPaths, err := dst.NormalizeFieldPaths("some_field")
if err != nil {
  return err
}

if err = dst.SetFields(src, normalizedPaths); err != nil {
  return err
}
```

## Contributing

We do not accept external issues with feature requests. This plugin only supports what we actually use ourselves at The Things Industries.

We do not accept external pull requests with new features, but everyone is free to fork it and add features in their own fork.

We do accept issues with bug reports and pull requests with bug fixes.
