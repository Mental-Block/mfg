-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS {{.schema}}.user
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

COMMENT ON TABLE {{.schema}}.user
    IS 'demoralizing user is added or updated on any table';

CREATE TABLE IF NOT EXISTS {{.schema}}.package
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMPTZ,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS {{.schema}}.order
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

CREATE TABLE IF NOT EXISTS {{.schema}}.order_job
(
    order_id INT NOT NULL,
    job_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS {{.schema}}.machine
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

CREATE TABLE IF NOT EXISTS {{.schema}}.line
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMPTZ, 
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS {{.schema}}.line_job
(
    job_id INT NOT NULL,
    line_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS {{.schema}}.job
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMPTZ, 
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS {{.schema}}.feeder
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

CREATE TABLE IF NOT EXISTS {{.schema}}.cart
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

CREATE TABLE IF NOT EXISTS {{.schema}}.board
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT, 
    updated_dt TIMESTAMPTZ,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS {{.schema}}.cart
    ADD CONSTRAINT fk_cart_machine
    FOREIGN KEY (machine_id)
    REFERENCES {{.schema}}.machine (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.order
    ADD CONSTRAINT fk_order_board
    FOREIGN KEY (board_id)
    REFERENCES {{.schema}}.board (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.order_job
    ADD CONSTRAINT fk_order_job_order
    FOREIGN KEY (order_id)
    REFERENCES {{.schema}}.order (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.order_job
    ADD CONSTRAINT fk_order_job_job
    FOREIGN KEY (job_id)
    REFERENCES {{.schema}}.job (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.machine
    ADD CONSTRAINT fk_machine_line
    FOREIGN KEY (line_id)
    REFERENCES {{.schema}}.line (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.line_job
    ADD CONSTRAINT fk_line_job_job
    FOREIGN KEY (job_id)
    REFERENCES {{.schema}}.job (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.line_job
    ADD CONSTRAINT fk_line_job_line
    FOREIGN KEY (line_id)
    REFERENCES {{.schema}}.line (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.feeder
    ADD CONSTRAINT fk_feeder_cart
    FOREIGN KEY (cart_id)
    REFERENCES {{.schema}}.cart (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.feeder
    ADD CONSTRAINT fk_feeder_package 
    FOREIGN KEY (package_id)
    REFERENCES {{.schema}}.package (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

---- create above / drop below ----

DROP TABLE {{.schema}}.line_job; 
DROP TABLE {{.schema}}.order_job;
DROP TABLE {{.schema}}.order;
DROP TABLE {{.schema}}.board;
DROP TABLE {{.schema}}.feeder;
DROP TABLE {{.schema}}.cart;
DROP TABLE {{.schema}}.machine;
DROP TABLE {{.schema}}.line;
DROP TABLE {{.schema}}.job;
DROP TABLE {{.schema}}.package;
DROP TABLE {{.schema}}.user;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
