<div align="center">

<a href="[Speakeasy](https://speakeasyapi.dev/)">
  <img src="https://github.com/speakeasy-api/speakeasy/assets/68016351/e959f81a-b250-4003-8c5c-a45b9463fc95" alt="Speakeasy Logo" width="400">
<h2>Speakeasy gRPC Gateway OpenAPI Example</h2>
</a>

</div>

This example gRPC Gateway app demonstrates Speakeasy-recommended practices for generating clear OpenAPI specifications and SDKs.

## Installation

To install the application on your local machine:

Clone the repository:
```bash
git clone https://github.com/ritza-co/speakeasy-grpc-gateway-example.git
```

Navigate into the directory:
```bash
cd speakeasy-grpc-gateway-example
```

### Install Go

On macOS, install Go by running:

```bash
brew install go
```

Alternatively follow the [Go installation instructions](https://go.dev/doc/install) for your platform.

### Install Buf

On macOS, install Buf by running:

```bash
brew install bufbuild/buf/buf
```

Alternatively, follow the [Buf CLI installation instructions](https://buf.build/docs/installation) for your platform.

### Install Buf Modules

We'll use Buf modules to manage our dependencies.

```bash
cd proto
buf mod update
cd ..
```

### Install protoc-gen-go

Buf requires the `protoc-gen-go` plugin to generate Go code from protobuf files.

Install `protoc-gen-go` by running:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Make sure that the `protoc-gen-go` binary is in your `$PATH`.

### Install Go Requirements

Run the following in the terminal from the project's root:

```bash
go mod tidy
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### Install Speakeasy CLI

To create an SDK, you'll also need the Speakeasy CLI installed, or use the Speakeasy dashboard.

[Install Speakeasy CLI](https://github.com/speakeasy-api/speakeasy#installation):
```bash
brew install speakeasy-api/homebrew-tap/speakeasy
```

## Generate OpenAPI Document and Create SDK

```bash
buf generate && go run convert/convert.go && speakeasy generate sdk \
    --schema openapi/speakeasy/v1/speakeasy.openapi.json \
    --lang typescript \
    --out ./sdk
```

## License

This project is licensed under the terms of the Apache 2.0 license.