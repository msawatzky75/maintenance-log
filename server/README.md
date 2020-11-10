# Maintenance Log Tracker Server

## Requirments

GoLang GraphQL server for the Vehicle Maintenance Log application.

GoLang is required to run and compile the application, and GQLGen is required to further development.

## Running Development

To run the server, configure the [.env](.env) to your environment.

```
PGADMIN_DEFAULT_EMAIL=email@example.com
PGADMIN_DEFAULT_PASSWORD=password

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_DB=maintenance_log
POSTGRES_PASSWORD=password

HOST_DOMAIN=log.yourdomain.com
JWT_SECRET=top-secret-here # generate a key with RSA or HS256
PASSWORD_HASH_COST=13 # idealy 250ms for bcrypt to hash

CORS_ORIGIN=http://localhost:3000
```

and run the [server.go](server.go) with go with

```sh
go run server.go
```

## Developing

This is a schema-first application, so any API changes must be made in the [schema](graph/schema) files first, then generated into resolvers with `GQLGen`.
New models will be placed into the [models](graph/models) directory, and new resolvers will be in the [resolvers](graph/resolvers) directory.
Anything else that is required will appear in the [generated](graph/generated/generated.go) file. If you need to modify something from here, move it to the [models](graph/models) and regenerate using `GQLGen`.
