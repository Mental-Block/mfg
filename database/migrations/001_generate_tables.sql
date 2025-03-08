-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS public.user
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    username TEXT NOT NULL CHECK (LENGTH(username) <= 100),
    password TEXT NOT NULL DEFAULT 'password',
    email TEXT UNIQUE CHECK (LENGTH(email) <= 255),
    updated_by TEXT,
    updated_dt TIMESTAMPTZ,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

COMMENT ON TABLE public.user
    IS 'demoralizing user is added or updated on any table';

CREATE TABLE IF NOT EXISTS public.package
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMPTZ,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.order
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    board_id INT NOT NULL,
    name TEXT NOT NULL,
    due_date DATE,
    quantity INT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMPTZ,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.order_job
(
    order_id INT NOT NULL,
    job_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS public.machine
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    line_id INT,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMPTZ,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.line
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMPTZ, 
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.line_job
(
    job_id INT NOT NULL,
    line_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS public.job
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMPTZ, 
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.feeder
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    cart_id INT,
    package_id INT,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMPTZ,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.cart
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    machine_id INT,
    name TEXT NOT NULL,
    slots INT NOT NULL,
    direction BIT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMPTZ,
    created_by TEXT NOT NULL DEFAULT SESSION_USER, 
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.board
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT, 
    updated_dt TIMESTAMPTZ,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.cart
    ADD CONSTRAINT fk_cart_machine
    FOREIGN KEY (machine_id)
    REFERENCES public.machine (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.order
    ADD CONSTRAINT fk_order_board
    FOREIGN KEY (board_id)
    REFERENCES public.board (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.order_job
    ADD CONSTRAINT fk_order_job_order
    FOREIGN KEY (order_id)
    REFERENCES public.order (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.order_job
    ADD CONSTRAINT fk_order_job_job
    FOREIGN KEY (job_id)
    REFERENCES public.job (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.machine
    ADD CONSTRAINT fk_machine_line
    FOREIGN KEY (line_id)
    REFERENCES public.line (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.line_job
    ADD CONSTRAINT fk_line_job_job
    FOREIGN KEY (job_id)
    REFERENCES public.job (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.line_job
    ADD CONSTRAINT fk_line_job_line
    FOREIGN KEY (line_id)
    REFERENCES public.line (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.feeder
    ADD CONSTRAINT fk_feeder_cart
    FOREIGN KEY (cart_id)
    REFERENCES public.cart (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.feeder
    ADD CONSTRAINT fk_feeder_package 
    FOREIGN KEY (package_id)
    REFERENCES public.package (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

---- create above / drop below ----

DROP TABLE public.board;
DROP TABLE public.cart;
DROP TABLE public.feeder;
DROP TABLE public.job;
DROP TABLE public.line_job;
DROP TABLE public.line;
DROP TABLE public.machine;
DROP TABLE public.order_job;
DROP TABLE public.order;
DROP TABLE public.packge;


-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
