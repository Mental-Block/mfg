ALTER TABLE IF EXISTS public.order
    ADD CONSTRAINT fk_order_board
    FOREIGN KEY (board_id)
    REFERENCES public.board (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;