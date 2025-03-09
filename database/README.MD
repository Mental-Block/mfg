# Database

## Requirements

Install the `tern` tool using [`install_tools`](https://github.com/jackc/tern)

## Local PostgreSQL

It is recommend to use `docker-compose.yml` file for the database. but if you just want to test migrations locally just spin up a new docker container.

```
docker run \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=USER \
  -e POSTGRES_PASSWORD=PASSWORD \
  -e POSTGRES_DB=DATABASE_NAME \
  -p 5432:5432 \
  postgres:17.4-bullseye
```

## Migrations

Run:

```
tern migrate \
    --destination 1 \
    --migrations "./migrations/" \
    --conn-string "postgres://USER:PASSWORD@localhost:5432/DATABASE_NAME?sslmode=disable"
```

Create:

```
tern new -m /migrations/ <migration name>
```
