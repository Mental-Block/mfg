ALTER TABLE IF EXISTS public.cart
    ADD CONSTRAINT fk_cart_machine
    FOREIGN KEY (machine_id)
    REFERENCES public.machine (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;