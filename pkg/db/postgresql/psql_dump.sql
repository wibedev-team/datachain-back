SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE public.about (
    title text DEFAULT ''::text NOT NULL,
    description text DEFAULT ''::text NOT NULL,
    img text DEFAULT ''::text NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);

ALTER TABLE public.about OWNER TO postgres;

CREATE TABLE public.features (
    title text,
    descr text NOT NULL
);

ALTER TABLE public.features OWNER TO postgres;

CREATE TABLE public.footer (
    email text NOT NULL,
    telephone text NOT NULL,
    address text NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);

ALTER TABLE public.footer OWNER TO postgres;

CREATE TABLE public.solutions (
    title text NOT NULL,
    link text NOT NULL,
    file text NOT NULL
);

ALTER TABLE public.solutions OWNER TO postgres;

CREATE TABLE public.stacks (
    img text NOT NULL
);

ALTER TABLE public.stacks OWNER TO postgres;

CREATE TABLE public.team (
    name text NOT NULL,
    "position" text NOT NULL,
    link text NOT NULL,
    img text NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);

ALTER TABLE public.team OWNER TO postgres;

CREATE TABLE public.users (
    login text,
    password text,
    role text DEFAULT 'ADMIN'::text
);

ALTER TABLE public.users OWNER TO postgres;

COPY public.users (login, password, role) FROM stdin;
admin	admin	ADMIN
test	test	ADMIN
\.

ALTER TABLE ONLY public.solutions
    ADD CONSTRAINT solutions_pkey PRIMARY KEY (title);

ALTER TABLE ONLY public.features
    ADD CONSTRAINT features_title_fkey FOREIGN KEY (title) REFERENCES public.solutions(title);
