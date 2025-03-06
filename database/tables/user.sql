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