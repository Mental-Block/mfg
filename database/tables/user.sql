CREATE TABLE IF NOT EXISTS public.user
(
    id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    username text NOT NULL,
    password text NOT NULL DEFAULT password,
    email text,
    createdDT timestamp without time zone NOT NULL,
    updatedDT timestamp without time zone,
    updatedBy text,
    createdBy text NOT NULL,
    PRIMARY KEY (id)
);

COMMENT ON TABLE public.user
    IS 'We are going to demoralize when added updated and created by user ';