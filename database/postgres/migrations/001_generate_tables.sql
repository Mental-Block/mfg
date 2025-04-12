-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS {{.schema}}.auth
(
    auth_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    password TEXT,
    email TEXT UNIQUE,
    verified BOOLEAN NOT NULL DEFAULT false,
    password_reset_token TEXT, -- delete in the future once redis is added
    email_verification_token TEXT, -- delete in the future once redis is added
    oauth BOOLEAN NOT NULL DEFAULT false, -- see if user is using oauth to login
    updated_by TEXT,
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (auth_id)
);

CREATE TABLE IF NOT EXISTS {{.schema}}.user
(
    user_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    auth_id INT NOT NULL,
    username TEXT NOT NULL CHECK (LENGTH(username) <= 30),
    updated_by TEXT,
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS {{.schema}}.package
(
    package_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (package_id)
);

CREATE TABLE IF NOT EXISTS {{.schema}}.order
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

CREATE TABLE IF NOT EXISTS {{.schema}}.order_job
(
    order_id INT NOT NULL,
    job_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS {{.schema}}.machine
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

CREATE TABLE IF NOT EXISTS {{.schema}}.line
(
    line_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMP, 
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (line_id)
);

CREATE TABLE IF NOT EXISTS {{.schema}}.line_job
(
    job_id INT NOT NULL,
    line_id INT NOT NULL
);

CREATE TABLE IF NOT EXISTS {{.schema}}.job
(
    job_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT,
    updated_dt TIMESTAMP, 
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (job_id)
);

CREATE TABLE IF NOT EXISTS {{.schema}}.feeder
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

CREATE TABLE IF NOT EXISTS {{.schema}}.cart
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

CREATE TABLE IF NOT EXISTS {{.schema}}.board
(
    board_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name TEXT NOT NULL,
    updated_by TEXT, 
    updated_dt TIMESTAMP,
    created_by TEXT NOT NULL DEFAULT SESSION_USER,
    created_dt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (board_id)
);

ALTER TABLE IF EXISTS {{.schema}}.cart
    ADD CONSTRAINT fk_cart_machine
    FOREIGN KEY (machine_id)
    REFERENCES {{.schema}}.machine (machine_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.order
    ADD CONSTRAINT fk_order_board
    FOREIGN KEY (board_id)
    REFERENCES {{.schema}}.board (board_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.order_job
    ADD CONSTRAINT fk_order_job_order
    FOREIGN KEY (order_id)
    REFERENCES {{.schema}}.order (order_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.order_job
    ADD CONSTRAINT fk_order_job_job
    FOREIGN KEY (job_id)
    REFERENCES {{.schema}}.job (job_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.machine
    ADD CONSTRAINT fk_machine_line
    FOREIGN KEY (line_id)
    REFERENCES {{.schema}}.line (line_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.line_job
    ADD CONSTRAINT fk_line_job_job
    FOREIGN KEY (job_id)
    REFERENCES {{.schema}}.job (job_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.line_job
    ADD CONSTRAINT fk_line_job_line
    FOREIGN KEY (line_id)
    REFERENCES {{.schema}}.line (line_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.feeder
    ADD CONSTRAINT fk_feeder_cart
    FOREIGN KEY (cart_id)
    REFERENCES {{.schema}}.cart (cart_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.feeder
    ADD CONSTRAINT fk_feeder_package 
    FOREIGN KEY (package_id)
    REFERENCES {{.schema}}.package (package_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS {{.schema}}.user
    ADD CONSTRAINT fk_user_auth
    FOREIGN KEY (auth_id)
    REFERENCES {{.schema}}.auth (auth_id) MATCH SIMPLE
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
DROP TABLE {{.schema}}.auth;

DROP TYPE IF EXISTS CartDirection;