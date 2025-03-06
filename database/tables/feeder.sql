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