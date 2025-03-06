ALTER TABLE IF EXISTS public.order_job
    ADD CONSTRAINT fk_order_job_order
    FOREIGN KEY (order_id)
    REFERENCES public.order (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.order_job
    ADD CONSTRAINT fk_order_job_job
    FOREIGN KEY (job_id)
    REFERENCES public.job (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;