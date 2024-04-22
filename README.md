<div align="center">

### Metabase SDK Go

---

An SDK for interacting with the Metabase REST API for the Go programming language.

</div>

## 🧐 About

[Metabase](https://www.metabase.com/) is an analytics tool which allows anyone to easily learn and make decisions from
their company's data. It allows you to query data directly from your databases (called "questions"), which you can store
and share with others, as well as generate reports.

This module provides an SDK to interact with the Metabase API for the Go programming language.

> [!NOTE]
> This SDK is not officially maintained or endorsed by Metabase.

## 🎈 Usage

1. Add the SDK to your project using `go get`

   ```sh
   go get github.com:bnjns/metabase-sdk-go
   ```

2. Create an authenticator, to tell the client how to authenticate with the API

   ```go
    package main
    
    import "github.com/bnjns/metabase-sdk-go/metabase"
    
    func main()  { 
        // Both session-based and API key-based authentication is supported
        authenticator, err := metabase.NewApiKeyAuthenticator("<api key>")
        if err != nil {
            panic(err)
        }
    }
   ```

3. Create the client

   ```go
    func main() {
        // ...
        client, err := metabase.NewClient("<host>", authenticator)
        if err != nil {
            panic(err)
        }	
    }
   ```

See [the docs](#) for more details.

## 🔧 Contributing

### Prerequisites

- Go >= 1.22
- Docker and Docker Compose
- [Task](https://taskfile.dev/installation/)
- [golangci-lint](https://golangci-lint.run/welcome/install/)

If you wish to update the documentation, you will also need:

- Ruby 3.x
- [Jekyll](https://jekyllrb.com/docs/installation/)

### Install

To get started, simply clone the repository:

```sh
git clone git@github.com:bnjns/metabase-sdk-go.git
```

Then install the Go dependencies:

```sh
go mod download
```

### Running Metabase

Use Task to run and set up Metabase using Docker:

```sh
task run:metabase
task setup:metabase
```

### Running tests

```sh
go test -v ./...
```

or, using Task:

```sh
task check:test
```

### Linting

```sh
task check:lint
```

## 📝 Additional Documentation

- [Metabase API documentation](https://www.metabase.com/docs/latest/api-documentation)

## ✍️ Authors

- [@bnjns](https://github.com/bnjns)
