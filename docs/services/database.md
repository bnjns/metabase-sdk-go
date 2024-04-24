---
title: Database
---

# Database API

[Metabase docs](https://www.metabase.com/docs/latest/api/database){ .md-button }

The request and response types can be imported from `github.com/bnjns/metabase-sdk-go/service/database`.

## Add a database

```go
ctx := context.Background()
databaseId, err := client.Database.Create(ctx, &database.CreateRequest{
    Engine: database.EnginePostgres,
    Name:   "Database name in Metabase",
    Details: database.Details{
        "host":     "<hostname>",
        "port":     5432,
        "user":     "<username>",
        "password": "<password>",
        "dbname":   "<database>",
    },
})
```

!!! warning

    The attributes needed in `Details` depends on the database engine. Unfortunately, the Metabase documentation does
    not provide any details; the best way to identify the fields needed is to add the database in your browser and
    inspect the request made to the API.

If your desired database engine isn't available within the SDK (`database.EngineXYZ`), you can simply cast it:

```go
databaseId, err := client.Database.Create(ctx, &database.CreateRequest{
    Engine: database.Engine("unsupported"),
    // ...
})
```

## Get a database

```go
ctx := context.Background()
db, err := client.Database.Get(ctx, 1)
```

See [`database.Database`](https://pkg.go.dev/github.com/bnjns/metabase-sdk-go/service/database#Database) for the fields
returned by the SDK.

## Update a database

```go
ctx := context.Background()

engine := database.EnginePostgres
updatedName := "Updated database name"
err := client.Database.Update(ctx, 1, &database.UpdateRequest{
    Engine: &engine,
    Name:   &updatedName,
    Details: &database.Details{
        "host":     "<hostname>",
        "port":     5432,
        "user":     "<username>",
        "password": "<password>",
        "dbname":   "<database>",
    },
})
```

!!! tip

    As this API uses a `PUT` operation, you may need to fetch the database before updating so that other properties are
    not lost:

    ```go
    updatedName := "Updated database name"
    
    db, _ := client.Database.Get(ctx, 1)
    err := client.Database.Update(ctx, 1, &database.UpdateRequest{
        Engine:           &db.Engine,
        Name:             &updatedName,
        Details:          db.Details,
        Refingerprint:    &db.Refingerprint,
        Schedules:        db.Schedules,
        Caveats:          db.Caveats,
        PointsOfInterest: db.PointsOfInterest,
        AutoRunQueries:   &db.AutoRunQueries,
        CacheTTL:         db.CacheTTL,
    })
    ```

## Remove a database

```go
ctx := context.Background()
err := client.Database.Delete(ctx, 1)
```
