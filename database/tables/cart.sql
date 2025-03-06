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