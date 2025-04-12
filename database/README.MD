# Database

### Prerequisite

Install the `tern` tool using [`install_tools`](https://github.com/jackc/tern)

### Setup PGADMIN GUI 

- Find ip adress of postgres container `docker ps -a` then inspect the container id. `docker inspect <CONTAINER_ID>`
- Go to `http://localhost:5050` or whatever PGADMIN_PORT_EXTERNAL and PGADMIN_HOST is specified
- Login to PGAdmin using the enviroment variables. PGADMIN_EMAIL and PGADMIN_PWD in the .env file
- Add hostname/address which is the ip adress from the postgres container in my case the ip was <strong>172.18.0.3</strong> and add other creds from enviroment variables such as POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB.

## Local PostgreSQL 

It is recommend to use `docker-compose.yml` file for the database. but if you just want to test migrations locally just spin up a new docker container.

```
docker run \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=<USER> \
  -e POSTGRES_PASSWORD=<PASSWORD> \
  -e POSTGRES_DB=<DATABASE_NAME> \
  -p 5432:5432 \
  postgres:17.4-bullseye
```

## Migrations

Run:

```
tern migrate \
    --destination 1 \
    --migrations "./migrations/" \
    --conn-string "postgres://<USER>:<PASSWORD>@<HOST>:<EXTERNAL_PORT>/<DATABASE_NAME>?sslmode=disable"
```

Create:

```
tern new -m /migrations/ <MIGRATION NAME>
```

## Seed

Insert:

```
  cat ./postgres/seed/generate_mock_data.sql | docker exec -i <CONTAINER NAME> psql -U <USER> -d <DATABASE>
```

Remove:

```
  cat ./postgres/seed/remove_data.sql | docker exec -i <CONTAINER NAME> psql -U <USER> -d <DATABASE>
```
