
CREATE TABLE IF NOT EXISTS public.job
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name text,
    createdDT timestamp with time zone NOT NULL,
    updatedDT timestamp with time zone NOT NULL,
    updatedBy text,
    createdBy text,
    PRIMARY KEY (id)
);
