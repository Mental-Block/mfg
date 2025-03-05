CREATE TABLE IF NOT EXISTS public.package
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name text,
    updatedBy text,
    createdBy text NOT NULL,
    updatedDT timestamp with time zone,
    createdDT timestamp with time zone NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.package
    ADD FOREIGN KEY (id)
    REFERENCES public.feeder (package_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
