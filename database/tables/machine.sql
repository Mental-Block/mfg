CREATE TABLE IF NOT EXISTS public.machine
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    line_id INT,
    name text,
    updatedDT timestamp with time zone,
    createdDT timestamp with time zone NOT NULL,
    createdBy text NOT NULL,
    updatedBy text,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.machine
    ADD FOREIGN KEY (line_id)
    REFERENCES public.line (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
