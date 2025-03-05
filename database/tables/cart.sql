CREATE TABLE IF NOT EXISTS public.cart
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    machine_id INT,
    name text,
    slots integer,
    direction bit,
    updatedDT timestamp with time zone,
    createdDT timestamp with time zone NOT NULL,
    createdBy text NOT NULL,
    updatedBy text,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.cart
    ADD FOREIGN KEY (machine_id)
    REFERENCES public.machine (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;