package pet

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/leshless/golibrary/auto_generator"
	_ "github.com/leshless/golibrary/constructor_generator"
	_ "github.com/leshless/golibrary/enum_generator"
	_ "github.com/leshless/golibrary/error_generator"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/oapi-codegen/runtime"
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
