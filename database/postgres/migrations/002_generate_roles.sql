
ALTER TABLE public.auth 
ADD version INT NOT NULL DEFAULT 1;

CREATE TABLE IF NOT EXISTS public.resource
(
    resource_id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    name text NOT NULL,
    updated_by text,
    updated_dt timestamp without time zone,
    created_by text NOT NULL DEFAULT SESSION_USER,
    created_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (resource_id)
);


CREATE TABLE IF NOT EXISTS public.permission
(
    permission_id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    name text NOT NULL,
    updated_by text,
    updated_dt timestamp without time zone,
    created_by text NOT NULL DEFAULT SESSION_USER,
    created_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (permission_id)
);

CREATE TABLE IF NOT EXISTS public.role
(
    role_id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    updated_by text,
    updated_dt timestamp without time zone,
    created_by text NOT NULL DEFAULT SESSION_USER,
    created_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (role_id)
);

CREATE TABLE IF NOT EXISTS public.role_resource_permission
(
    role_id integer NOT NULL,
    resource_permission_id integer NOT NULL,
    assigned_by text NOT NULL DEFAULT SESSION_USER,
    assigned_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS public.user_role
(
    user_id integer NOT NULL,
    role_id integer NOT NULL,
    assigned_by text NOT NULL DEFAULT SESSION_USER,
    assigned_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS public.resource_permission
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

CREATE TABLE IF NOT EXISTS public.user_resource_permission
(
    user_id integer NOT NULL,
    resource_permission_id integer NOT NULL,
    assigned_by text NOT NULL DEFAULT SESSION_USER,
    assigned_dt timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE IF EXISTS public.role_resource_permission
    ADD CONSTRAINT fk_role_resource_permission FOREIGN KEY (role_id)
    REFERENCES public.role (role_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.role_resource_permission
    ADD CONSTRAINT fk_resource_permission_role FOREIGN KEY (resource_permission_id)
    REFERENCES public.resource_permission (resource_permission_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.user_role
    ADD CONSTRAINT fk_user_role FOREIGN KEY (user_id)
    REFERENCES public.user (user_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.user_role
    ADD CONSTRAINT fk_role_user FOREIGN KEY (role_id)
    REFERENCES public.role (role_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.resource_permission
    ADD CONSTRAINT fk_resource_resource_permission FOREIGN KEY (resource_id)
    REFERENCES public.resource (resource_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.resource_permission
    ADD CONSTRAINT fk_permission_resource_permission FOREIGN KEY (permission_id)
    REFERENCES public.permission (permission_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.user_resource_permission
    ADD CONSTRAINT fk_user_user_resource_permission FOREIGN KEY (user_id)
    REFERENCES public.user (user_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.user_resource_permission
    ADD CONSTRAINT fk_resource_permission_user_resource_permission  FOREIGN KEY (resource_permission_id)
    REFERENCES public.resource_permission (resource_permission_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

---- create above / drop below ----

DROP TABLE public.user_role;
DROP TABLE public.role_resource_permission;
DROP TABLE public.user_resource_permission;
DROP TABLE public.role;
DROP TABLE public.resource_permission;
DROP TABLE public.permission;
DROP TABLE public.resource;

ALTER TABLE public.auth 
DROP COLUMN version;
