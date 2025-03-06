# Database

## TODO

- Move the creds to docker secrets. fine for now as its local database.
- Storing DB users password in plain text. use a hashing algorithms

## Getting started

### Prerequisite

- Docker version 28+
- Docker Compose version v2.33.1+

### Setup

- Generate docker container from image `sudo docker-compose up`
- Find ip adress of postgres container `sudo docker container ls` then inspect the container id. `sudo docker inspect {YOUR CONTAINER ID}`
- Go to `http://localhost:5050` or whatever PORT is specified
- Login to PGAdmin using the enviroment variables. PGADMIN_DEFAULT_EMAIL and PGADMIN_DEFAULT_PASSWORD
- Add hostname/address which is the ip adress from the postgres container in my case the ip was <strong>172.18.0.3</strong> and add other creds from enviroment variables such as POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB.

### Generate Tables

- `source ./tables/insert.sh`

### Add Mock data

- `source ./mockdata/insert.sh`

Theres also the option to truncate all data.

- `source ./mockdata/truncate.sh`

## Database ERD

![alt text](ERD.png)
