ALTER TABLE IF EXISTS public.machine
    ADD CONSTRAINT fk_machine_line
    FOREIGN KEY (line_id)
    REFERENCES public.line (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;