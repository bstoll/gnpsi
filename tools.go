//go:build tools

package gnpsi

// These imports only exist to keep go.mod entries for packages that are referenced in BUILD files,
// but not in Go code.

import (
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
)
