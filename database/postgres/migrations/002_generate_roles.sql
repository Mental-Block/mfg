-- Write your migrate up statements here

CREATE SCHEMA IF NOT EXISTS auth;

ALTER TABLE public.auth SET SCHEMA auth;

CREATE TABLE IF NOT EXISTS auth.resource
(
    resource_id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    name text NOT NULL,
    updated_by text,
    updated_dt timestamp without time zone,
    created_by text NOT NULL DEFAULT SESSION_USER,
    created_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (resource_id)
);

CREATE TABLE IF NOT EXISTS auth.permission
(
    permission_id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    name text NOT NULL,
    updated_by text,
    updated_dt timestamp without time zone,
    created_by text NOT NULL DEFAULT SESSION_USER,
    created_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (permission_id)
);

CREATE TABLE IF NOT EXISTS auth.role
(
    role_id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    updated_by text,
    updated_dt timestamp without time zone,
    created_by text NOT NULL DEFAULT SESSION_USER,
    created_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (role_id)
);

CREATE TABLE IF NOT EXISTS auth.role_resource_permission
(
    role_id integer NOT NULL,
    resource_permission_id integer NOT NULL,
    assigned_by text NOT NULL DEFAULT SESSION_USER,
    assigned_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS auth.auth_role
(
    auth_id integer NOT NULL,
    role_id integer NOT NULL,
    assigned_by text NOT NULL DEFAULT SESSION_USER,
    assigned_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS auth.resource_permission
(
    resource_permission_id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    resource_id integer NOT NULL,
    permission_id integer NOT NULL,
    updated_by text,
    updated_dt timestamp without time zone,
    created_by text NOT NULL DEFAULT SESSION_USER,
    created_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (resource_permission_id)
);

ALTER TABLE IF EXISTS auth.role_resource_permission
    ADD CONSTRAINT fk_role_resource_permission FOREIGN KEY (role_id)
    REFERENCES auth.role (role_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS auth.role_resource_permission
    ADD CONSTRAINT fk_resource_permission_role FOREIGN KEY (resource_permission_id)
    REFERENCES auth.resource_permission (resource_permission_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS auth.auth_role
    ADD CONSTRAINT fk_auth_role FOREIGN KEY (auth_id)
    REFERENCES auth.auth (auth_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS auth.auth_role
    ADD CONSTRAINT fk_role_auth FOREIGN KEY (role_id)
    REFERENCES auth.role (role_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS auth.resource_permission
    ADD CONSTRAINT fk_resource_resource_permission FOREIGN KEY (resource_id)
    REFERENCES auth.resource (resource_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS auth.resource_permission
    ADD CONSTRAINT fk_permission_resource_permission FOREIGN KEY (permission_id)
    REFERENCES auth.permission (permission_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

---- create above / drop below ----

ALTER TABLE auth.auth SET SCHEMA public;

DROP TABLE auth.resource;
DROP TABLE auth.permission;
DROP TABLE auth.role;
DROP TABLE auth.role_resource_permission;
DROP TABLE auth.auth_role;
DROP TABLE auth.resource_permission;

DROP SCHEMA auth;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
