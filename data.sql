-- Table: public.detailed

-- DROP TABLE public.detailed;

CREATE TABLE IF NOT EXISTS public.detailed
(
    id integer,
    season smallint NOT NULL,
    episodes integer NOT NULL,
    year smallint NOT NULL,
    CONSTRAINT detailed_id_fkey FOREIGN KEY (id)
        REFERENCES public.tvseries_info (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE public.detailed
    OWNER to postgres;

-- Table: public.tvseries_info

-- DROP TABLE public.tvseries_info;

CREATE TABLE IF NOT EXISTS public.tvseries_info
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 214564 CACHE 1 ),
    title character varying COLLATE pg_catalog."default" NOT NULL,
    producer character varying COLLATE pg_catalog."default",
    CONSTRAINT tvseries_info_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.tvseries_info
    OWNER to postgres;