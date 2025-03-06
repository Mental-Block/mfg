ALTER TABLE IF EXISTS public.feeder
    ADD CONSTRAINT fk_feeder_cart
    FOREIGN KEY (cart_id)
    REFERENCES public.cart (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

 ALTER TABLE IF EXISTS public.feeder
    ADD CONSTRAINT fk_feeder_package 
    FOREIGN KEY (package_id)
    REFERENCES public.package (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;