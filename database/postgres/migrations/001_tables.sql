-- Write your migrate up statements here

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.auth
( 
    auth_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    password TEXT,
    email TEXT UNIQUE,
    version INT NOT NULL DEFAULT 1,
    updated_by TEXT,
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (auth_id)
);

CREATE TABLE IF NOT EXISTS public.user
(
    user_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    username TEXT,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    title TEXT,
    avatar TEXT,
    metadata JSONB, 
    deleted_by TEXT,
    deleted_dt TIMESTAMP,
    updated_by TEXT,
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS public.user_auth
(
    auth_id INT,
    user_id INT
);





---- create above / drop below ----

DROP TABLE public.user;
DROP TABLE public.auth;
DROP TABLE public.user_auth;
