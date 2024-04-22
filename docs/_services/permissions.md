---
title: Permissions
---

# Permissions API

[Metabase docs](https://www.metabase.com/docs/latest/api/permissions){: .btn .btn-purple }

The request and response types can be imported from `github.com/bnjns/metabase-sdk-go/service/permissions`.

## Permissions Group

### Create a permissions group

```go
ctx := context.Background()
groupId, err := client.Permissions.CreateGroup(ctx, &permissions.CreateGroupRequest{
    Name: "Permissions group name",
})
```

### Get a permissions group

```go
ctx := context.Background()
group, err := client.Permissions.GetGroup(ctx, 1)
```

See `permissions.Group` for the fields returned by the SDK.

### Update a permissions group

```go
ctx := context.Background()
groupId, err := client.Permissions.UpdateGroup(ctx, &permissions.UpdateGroupRequest{
    Name: "Updated group name",
})
```

### Delete a permissions group

```go
ctx := context.Background()
err := client.Permissions.DeleteGroup(ctx, 1)
```
