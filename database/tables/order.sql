CREATE TABLE IF NOT EXISTS public.order
(
    id uuid NOT NULL GENERATED ALWAYS AS IDENTITY,
    board_id uuid NOT NULL,
    name text NOT NULL,
    dueDate date,
    quantity integer NOT NULL,
    createdDT timestamp with time zone NOT NULL,
    updatedDT timestamp with time zone,
    createdBy text NOT NULL,
    updatedBy text,
    PRIMARY KEY (id)
);

