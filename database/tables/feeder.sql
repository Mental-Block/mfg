CREATE TABLE IF NOT EXISTS public.feeder
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    cart_id INT,
    package_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name text,
    type text,
    updatedDT timestamp with time zone,
    createdDT timestamp with time zone NOT NULL,
    createdBy text NOT NULL,
    updatedBy text,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.feeder
    ADD FOREIGN KEY (cart_id)
    REFERENCES public.cart (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;