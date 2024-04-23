---
title: User
---

# User API

[Metabase docs](https://www.metabase.com/docs/latest/api/user){: .md-button }

The request and response types can be imported from `github.com/bnjns/metabase-sdk-go/service/user`.

### Create a user

```go
ctx := context.Background()

email := "example@example.com"
firstName := "Example"
lastName := "User"
userId, err := client.User.Create(ctx, &user.CreateRequest{
    Email:     email,
    FirstName: &firstName,
    LastName:  &lastName,
    GroupMemberships: &[]user.GroupMembership{
        {Id: 3},
    },
})
```

!!! warning

    When specifying the group memberships, you should not specify the `All Users` or `Administrators` groups (ID 1 and
    2 respectively); all users are automatically added to `All Users`, and you can add a user to `Administrators` by
    updating the user with `IsSuperuser` set to `true`.

### Get a user

```go
ctx := context.Background()
currentUser, err := client.User.Get(ctx, 1)
```

See [`user.User`](https://pkg.go.dev/github.com/bnjns/metabase-sdk-go/service/user#User) for the fields returned by the SDK.

### Get the current user

```go
ctx := context.Background()
currentUser, err := client.User.GetCurrentUser(ctx)
```

See [`user.User`](https://pkg.go.dev/github.com/bnjns/metabase-sdk-go/service/user#User) for the fields returned by the SDK.

### Update a user

```go
ctx := context.Background()

email := "example@example.com"
firstName := "Example"
lastName := "User"
locale := "eu"
isSuperuser := true
err := client.User.Update(ctx, 1, &user.UpdateRequest{
    Email:     &email,
    FirstName: &firstName,
    LastName:  &lastName,
    GroupMemberships: &[]user.GroupMembership{
        {Id: 3},
    },
    Locale:      &locale,
    IsSuperuser: &isSuperuser,
})
```

!!! tip

    As this API uses a `PUT` operation, you may need to fetch the user before updating so that other properties are not
    lost:

    ```go
    isSuperuser := true
    
    usr, _ := client.User.Get(ctx, 1)
    err := client.User.Update(ctx, 1, &user.UpdateRequest{
        Email:            &usr.Email,
        FirstName:        usr.FirstName,
        LastName:         usr.LastName,
        Locale:           usr.Locale,
        IsSuperuser:      &isSuperuser,
        LoginAttributes:  usr.LoginAttributes,
        GroupMemberships: &usr.GroupMemberships,
    })
    ```

### Disable a user

You cannot delete users in Metabase; instead you can disable them so that they cannot be logged into:

```go
ctx := context.Background()
err := client.User.Disable(ctx, 1)
```

### Reactivate a user

```go
ctx := context.Background()
err := client.User.Reactivate(ctx, 1)
```

This will gracefully reactivate the specified user (ie, reactivating a user that's already enabled will not return an
error).
