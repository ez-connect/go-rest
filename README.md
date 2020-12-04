# Golang RESTful

A collection of helpers for building a RESTful API server.

- Generate source code from configs
- Genrate Open API definitions

### Install

Get the package

```bash
go get github.com/ez-connect/go-rest
```

Install the binary

```bash
go install github.com/ez-connect/go-rest/cmd/go-server
```

### API Documentation

https://pkg.go.dev/github.com/ez-connect/go-rest

### Example

Create a new service naming  `hello`

```bash
go-rest-gen -new hello
```

Generate source code in `/generated`

```bash
go-rest-gen
```
