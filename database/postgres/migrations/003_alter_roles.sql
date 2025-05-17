ALTER TABLE public.role
ADD name TEXT NOT NULL;

ALTER TABLE public.user
ADD attribute JSONB DEFAULT NULL;

ALTER TABLE public.user
ADD active BOOLEAN DEFAULT true;

ALTER TABLE public.resource
ADD attribute JSONB DEFAULT NULL;

---- create above / drop below ----

ALTER TABLE public.role
DROP COLUMN name;

ALTER TABLE public.user
DROP COLUMN attribute;

ALTER TABLE public.user
DROP COLUMN active;

ALTER TABLE public.resource
DROP COLUMN attribute;