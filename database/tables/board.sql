CREATE TABLE IF NOT EXISTS public.board
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name text NOT NULL,
    createdDT timestamp without time zone NOT NULL,
    updatedDT timestamp without time zone,
    updatedBy text,
    createdBy text NOT NULL,
    PRIMARY KEY (id)
);