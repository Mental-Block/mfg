ALTER TABLE IF EXISTS public.line_job
    ADD CONSTRAINT fk_line_job_job
    FOREIGN KEY (job_id)
    REFERENCES public.job (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

ALTER TABLE IF EXISTS public.line_job
    ADD CONSTRAINT fk_line_job_line
    FOREIGN KEY (line_id)
    REFERENCES public.line (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;
