
CREATE TABLE IF NOT EXISTS public.package
(
    package_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (package_id)
);

CREATE TABLE IF NOT EXISTS public.order
(
    order_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    board_id INT NOT NULL,
    name TEXT NOT NULL,
    due_date DATE,
    quantity INT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (order_id)
);

CREATE TABLE IF NOT EXISTS public.order_job
(
    order_id INT NOT NULL,
    job_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS public.machine
(
    machine_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    line_id INT,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (machine_id)
);

CREATE TABLE IF NOT EXISTS public.line
(
    line_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMP, 
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (line_id)
);

CREATE TABLE IF NOT EXISTS public.line_job
(
    job_id INT NOT NULL,
    line_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS public.job
(
    job_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMP, 
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (job_id)
);

CREATE TABLE IF NOT EXISTS public.feeder
(
    feeder_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    cart_id INT,
    package_id INT,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (feeder_id)
);

CREATE TYPE CartDirection AS ENUM ('left', 'right');

CREATE TABLE IF NOT EXISTS public.cart
(
    cart_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    machine_id INT,
    name TEXT NOT NULL,
    slots INT NOT NULL,
    direction CartDirection NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER, 
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (cart_id)
);

CREATE TABLE IF NOT EXISTS public.board
(
    board_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT, 
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (board_id)
);

ALTER TABLE IF EXISTS public.cart
    ADD CONSTRAINT fk_cart_machine
    FOREIGN KEY (machine_id)
    REFERENCES public.machine (machine_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.order
    ADD CONSTRAINT fk_order_board
    FOREIGN KEY (board_id)
    REFERENCES public.board (board_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.order_job
    ADD CONSTRAINT fk_order_job_order
    FOREIGN KEY (order_id)
    REFERENCES public.order (order_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.order_job
    ADD CONSTRAINT fk_order_job_job
    FOREIGN KEY (job_id)
    REFERENCES public.job (job_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.machine
    ADD CONSTRAINT fk_machine_line
    FOREIGN KEY (line_id)
    REFERENCES public.line (line_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.line_job
    ADD CONSTRAINT fk_line_job_job
    FOREIGN KEY (job_id)
    REFERENCES public.job (job_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.line_job
    ADD CONSTRAINT fk_line_job_line
    FOREIGN KEY (line_id)
    REFERENCES public.line (line_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.feeder
    ADD CONSTRAINT fk_feeder_cart
    FOREIGN KEY (cart_id)
    REFERENCES public.cart (cart_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.feeder
    ADD CONSTRAINT fk_feeder_package 
    FOREIGN KEY (package_id)
    REFERENCES public.package (package_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

DROP TABLE public.line_job; 
DROP TABLE public.order_job;
DROP TABLE public.order;
DROP TABLE public.board;
DROP TABLE public.feeder;
DROP TABLE public.cart;
DROP TABLE public.machine;
DROP TABLE public.line;
DROP TABLE public.job;
DROP TABLE public.package;