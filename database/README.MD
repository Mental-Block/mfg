# Database

## Requirements

Install the `tern` tool using [`install_tools`](https://github.com/jackc/tern)

## Local PostgreSQL

It is recommend to use `docker-compose.yml` file for the database. but if you just want to test migrations locally just spin up a new docker container.

```
docker run \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=root \
  -e POSTGRES_PASSWORD=root \
  -e POSTGRES_DB=mfg \
  -p 5432:5432 \
  postgres:17.4-bullseye
```

## Migrations

Run:

```
tern migrate \
    --destination 1 \
    --migrations "./migrations/" \
    --conn-string "postgres://root:root@localhost:5432/mfg?sslmode=disable"
```

Create:

```
tern new -m /migrations/ <migration name>
```

## TODO

- Move the creds to docker secrets. fine for now as its local database.
- Storing DB users password in plain text. use a hashing algorithms
