# Maintenance Log Tracker

Vehicle maintencene log application written in Go with GraphQL and a PostgreSQL database.

Front end is written in Vue.js with the Nuxt.js framework.

# Running a Development Version

## Requirements

Go (version 1.15), Node (version 1.14) and Yarn (version 1.22) must be installed.

### Running development sever
> Note: a postgresql database running on port 5432 is required.

```
$ cd server/
$ go install
$ go run dev
```

### Running development web

```
$ cd web/
$ yarn install
$ yarn dev
```


## Building for production

Dockerfiles have been made and are available in the [server](server/Dockerfile) and [web](web/Dockerfile) folders.

### Building and creating Docker containers

```
$ docker-compose build
```
