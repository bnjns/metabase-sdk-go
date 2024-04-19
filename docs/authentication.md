---
title: Authentication
nav_order: 2
---

# Authentication

Metabase supports 2 authentication methods:

- API keys (recommended)
- Session-based, using a username and password

{: .info }
Most administrative actions require that the API key or user are a "super user".

## Using an API key
{: .d-inline-block }

v0.49+
{: .label .label-purple }

Follow Metabase's documentation to [create an API key](https://www.metabase.com/docs/latest/people-and-groups/api-keys),
then use `NewApiKeyAuthenticator` to create the authenticator:

```go
authenticator, err := metabase.NewApiKeyAuthenticator("<your api key>")
client, err := metabase.NewClient("<host>", authenticator)
```

The API key will be included in the `X-Api-Key` header of every request.

## Using a username and password

If you are unable to use an API key to authenticate, you can still use session-based authentication by providing the
username and password of a user with `NewSessionAuthenticator`:

```go
authenticator, err := metabase.NewSessionAuthenticator("<username>", "<password>")
client, err := metabase.NewClient("<host>", authenticator)
```

This will automatically log into Metabase when the client is created, and include the session ID in
the `X-Metabase-Session` header for every request.
