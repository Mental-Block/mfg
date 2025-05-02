-- Write your migrate up statements here

ALTER TABLE public.auth
ADD color VARCHAR(255); 

CREATE TABLE public.role
(
    role_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name  varchar(50)


)

CREATE TABLE public.group
(
    group_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name varchar(50)

)

CREATE TABLE public.permission
(
    permission_id INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name varchar(50)

)



---- create above / drop below ----

ALTER TABLE public.auth
DROP color;

DROP TABLE public.role
DROP TABLE public.permission
DROP TABLE public.group


-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
