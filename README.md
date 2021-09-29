# Echo template for [Create Go App CLI](https://github.com/create-go-app/cli)


[Echo](https://github.com/labstack/echo) is a high performance, minimalist Go web framework.

## ⚡️ Quick start

1. Create a new project with Echo:

```bash
cgapp create

# Choose a backend framework:
# > echo
#  fiber
#  net/http

```

## Development

Prerequisites:
- [Go](https://golang.org)
- [Docker](https://github.com/google/protobuf)

### Run tests
```
    make test
```

### Run static code analysis
```bash
    make lint
```

### Run code coverage
```bash
    make coverage
```

### Run the server
```bash
    ## with Go
    make local
```
```bash
    ## with Docker
    make docker.run
```


5. Go to API Docs page (Swagger): [127.0.0.1:5000/swagger/index.html](http://127.0.0.1:5000/swagger/index.html)

![Screenshot](https://user-images.githubusercontent.com/11155743/112715187-07dab100-8ef0-11eb-97ea-68d34f2178f6.png)


## 🗄 Template structure



### ./cmd

```
├── cmd
│     └── template
│             └── main.go
│...
```
**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/models` folder for describe business models and methods of your project
- `./app/queries` folder for describe queries for models of your project

### ./docs

**Folder with API Documentation**. This directory contains config files for auto-generated API Docs by Swagger.

### ./internal

```
├── internal
│   ├── api
│   │   ├── models
│   ├── cache
│   ├── datastore
│   ├── middleware
│   ├── migrations
│   ├── redis
│   └── token
```

[comment]: <> (**Folder with project-specific functionality**. This directory contains all the project-specific code tailored only for your business use case, like _configs_, _middleware_, _routes_ or _utils_.)

[comment]: <> (- `./pkg/configs` folder for configuration functions)

[comment]: <> (- `./pkg/middleware` folder for add middleware &#40;Fiber built-in and yours&#41;)

[comment]: <> (- `./pkg/repository` folder for describe `const` of your project)

[comment]: <> (- `./pkg/routes` folder for describe routes of your project)

[comment]: <> (- `./pkg/utils` folder with utility functions &#40;server starter, error checker, etc&#41;)

[comment]: <> (### ./platform)

[comment]: <> (**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project, like _setting up the database_ or _cache server instance_ and _storing migrations_.)

[comment]: <> (- `./platform/cache` folder with in-memory cache setup functions &#40;by default, Redis&#41;)

[comment]: <> (- `./platform/database` folder with database setup functions &#40;by default, PostgreSQL&#41;)

[comment]: <> (- `./platform/migrations` folder with migration files &#40;used with [golang-migrate/migrate]&#40;https://github.com/golang-migrate/migrate&#41; tool&#41;)

[comment]: <> (## ⚙️ Configuration)

[comment]: <> (```ini)

[comment]: <> (# .env)

[comment]: <> (# Stage status to start server:)

[comment]: <> (#   - "dev", for start server without graceful shutdown)

[comment]: <> (#   - "prod", for start server with graceful shutdown)

[comment]: <> (STAGE_STATUS="dev")

[comment]: <> (# Server settings:)

[comment]: <> (SERVER_HOST="0.0.0.0")

[comment]: <> (SERVER_PORT=5000)

[comment]: <> (SERVER_READ_TIMEOUT=60)

[comment]: <> (# JWT settings:)

[comment]: <> (JWT_SECRET_KEY="secret")

[comment]: <> (JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=15)

[comment]: <> (JWT_REFRESH_KEY="refresh")

[comment]: <> (JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT=720)

[comment]: <> (# Database settings:)

[comment]: <> (DB_HOST="cgapp-postgres")

[comment]: <> (DB_PORT=5432)

[comment]: <> (DB_USER="postgres")

[comment]: <> (DB_PASSWORD="password")

[comment]: <> (DB_NAME="postgres")

[comment]: <> (DB_SSL_MODE="disable")

[comment]: <> (DB_MAX_CONNECTIONS=100)

[comment]: <> (DB_MAX_IDLE_CONNECTIONS=10)

[comment]: <> (DB_MAX_LIFETIME_CONNECTIONS=2)

[comment]: <> (# Redis settings:)

[comment]: <> (REDIS_HOST="cgapp-redis")

[comment]: <> (REDIS_PORT=6379)

[comment]: <> (REDIS_PASSWORD="")

[comment]: <> (REDIS_DB_NUMBER=0)

[comment]: <> (```)

## ⚠️ License

Apache 2.0 &copy; [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).