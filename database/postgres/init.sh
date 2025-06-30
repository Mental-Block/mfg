#!/usr/bin/env bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
	CREATE DATABASE "authorization";
	CREATE DATABASE "metrics";
	GRANT ALL PRIVILEGES ON DATABASE "authorization" TO "$POSTGRES_USER";
	GRANT ALL PRIVILEGES ON DATABASE "metrics" TO "$POSTGRES_USER";
EOSQL