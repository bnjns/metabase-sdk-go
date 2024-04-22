---
title: Introduction
nav_order: 1
---

# Metabase SDK for Go

`metabase-sdk-go` is an SDK for the Go programming language for interacting with the Metabase REST API. Please note,
this is not an official SDK or endorsed by Metabase.

{: .warning }
This SDK is still a work-in-progress and does not yet cover the full API.

## API compatibility

<p style="display: flex; flex-direction: row">
   <strong>Supported Metabase Version:</strong>
   <span class="label" style="text-transform:none;">v0.49</span>
</p>

While Metabase is versioned, its API is not and is constantly evolving; as such the SDK may only work with specific
versions of Metabase.

Best efforts will be made to make the SDK generic enough that it can work across multiple versions of Metabase, however
each version of the SDK will be designed to target a specific Metabase version.

## Getting started

1. Add the SDK to your project using `go get`

   ```sh
   go get github.com:bnjns/metabase-sdk-go
   ```

2. Create [an authenticator](authentication.md), to tell the client how to authenticate with the API

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

## Customising the client

### Timeout

You can customise the global request timeout (default is 10 seconds):

```go
// ...

client, err := metabase.NewClient("<host>", authenticator, func (opt *metabase.Options) {
    opt.Timeout = 30 * time.Second
})
```

Or using the `metabase.WithTimeout` function:

```go
// ...

client, err := metabase.NewClient("<host>", authenticator, metabase.WithTimeout(30 * time.Second))
```

### Additional headers

You can configure the client to add custom headers to all requests made to Metabase:

```go
// ...

client, err := metabase.NewClient("<host>", authenticator, func (opt *metabase.Options) {
    opt.AdditionalHeaders = map[string]string{
        "Authorizer": "Bearer <id token>",
    }
})
```

Or using the `metabase.WithHeader` function (this will append to any existing headers):

```go
// ...

client, err := metabase.NewClient(
    "<host>",
    authenticator,
    metabase.WithHeader("Authorizer", "Bearer <id token>"),
    metabase.WithHeader("X-Example", "example value"),
)
```

If you want to configure multiple headers, you can also use `metabase.WithHeaders`:

```go
client, err := metabase.NewClient("<host>", authenticator, metabase.WithHeaders(map[string]string{
    "Authorizer": "Bearer <id token>",
    "X-Example": "example value",
}))
```

{: .warning }
> You cannot use this to set any of the following headers:
>
> - `Content-Type`
> - `X-Api-Key`
> - `X-Metabase-Session`
