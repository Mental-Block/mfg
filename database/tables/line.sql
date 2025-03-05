CREATE TABLE IF NOT EXISTS public.line
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name text,
    updatedDT timestamp with time zone,
    createdDT timestamp with time zone NOT NULL,
    createdBy text NOT NULL,
    updatedBy text,
    PRIMARY KEY (id)
);