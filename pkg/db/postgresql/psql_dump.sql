--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

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

--
-- Name: about; Type: TABLE; Schema: public; Owner: chechyotka
--

CREATE TABLE public.about (
    title text DEFAULT ''::text NOT NULL,
    description text DEFAULT ''::text NOT NULL,
    img text DEFAULT ''::text NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.about OWNER TO chechyotka;

--
-- Name: footer; Type: TABLE; Schema: public; Owner: chechyotka
--

CREATE TABLE public.footer (
    email text NOT NULL,
    telephone text NOT NULL,
    address text NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.footer OWNER TO chechyotka;

--
-- Name: stacks; Type: TABLE; Schema: public; Owner: chechyotka
--

CREATE TABLE public.stacks (
    img text NOT NULL
);


ALTER TABLE public.stacks OWNER TO chechyotka;

--
-- Name: team; Type: TABLE; Schema: public; Owner: chechyotka
--

CREATE TABLE public.team (
    name text NOT NULL,
    "position" text NOT NULL,
    link text NOT NULL,
    img text NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.team OWNER TO chechyotka;

--
-- Name: users; Type: TABLE; Schema: public; Owner: chechyotka
--

CREATE TABLE public.users (
    login text,
    password text,
    role text DEFAULT 'ADMIN'::text
);


ALTER TABLE public.users OWNER TO chechyotka;

--
-- Data for Name: about; Type: TABLE DATA; Schema: public; Owner: chechyotka
--

COPY public.about (title, description, img, created_at) FROM stdin;
test	test	9ffc220b-8435-441c-b04c-6df33212513e.png	2023-07-22 13:24:40.05576
		41822617-b844-4145-baa7-47c5dc36a3cc.png	2023-07-22 13:24:40.05576
testawdadad	testawdawdda	c1525a77-6c55-4b2f-9ba5-23a5eda3bc9c.png	2023-07-22 16:26:38.598898
testawd	testawdawddatestawdawddatestawdawddatestawdawddatestawdawddatestawdawddatestawdawdda	d30eb49d-f929-4890-8993-a388e4d6cd7a.png	2023-07-22 16:27:20.39954
testawd	sdsfsf	b97c74aa-6feb-4f21-a952-f8a78a80fe87.png	2023-07-22 16:28:32.648284
testawd	sdsfsf122222222222222222	5b859a83-0306-495d-9794-60376cc7555a.png	2023-07-22 16:29:03.854327
testawd	sdsfsf122222222222222222	916b1ae5-7a30-4d0b-a849-8b31eb3d62ca.png	2023-07-22 16:29:14.670907
testawd	<h1>Hello world</h1>	0d00b577-c46c-4c66-a555-47ccda028efa.png	2023-07-22 16:32:39.968832
testawd	<h3>Hello world</h3>	4335c1f1-abc4-42ec-9c7f-e5b66b2b754c.png	2023-07-22 16:40:29.600167
testawd	<h1>Hello world</h1>	47ddaed5-52fe-41a6-8cad-1ca2ce34d485.png	2023-07-22 16:40:44.05793
testawd	<h1>Hello world</h1>	4fd19205-3c3e-45f0-8e2c-d3319daa243f.png	2023-07-22 16:49:08.371748
finish	<div>finisi</div>	12a5ccee-2cde-4c28-baed-811b33563ef3.png	2023-07-22 16:49:33.444055
finishrth	<div>finisrthrhrhi</div>	6406cadb-9c1a-4751-bd9d-84a2551348d0.png	2023-07-22 21:24:02.210504
test	<h1>finisrthrhrhi</h1>	5bde1af6-3798-4bd5-9bf2-2d9f770f3698.png	2023-07-22 21:30:19.895081
\.


--
-- Data for Name: footer; Type: TABLE DATA; Schema: public; Owner: chechyotka
--

COPY public.footer (email, telephone, address, created_at) FROM stdin;
roman@test	+23131`	aweaeaedadad	2023-07-22 19:36:28.991249
roman@test	+23123tkjbnask	aweaeaedadad	2023-07-22 19:52:31.014591
roman@test	+2222222222222	aweaeaedadad	2023-07-22 19:52:49.839425
roman@test	+12313123213123	aweaeaedadad	2023-07-22 19:52:58.8938
rom	=	a	2023-07-22 20:30:09.685234
rom	=====	a	2023-07-22 20:31:16.093226
aw	aw	aw	2023-07-22 21:31:31.348139
\.


--
-- Data for Name: stacks; Type: TABLE DATA; Schema: public; Owner: chechyotka
--

COPY public.stacks (img) FROM stdin;
\.


--
-- Data for Name: team; Type: TABLE DATA; Schema: public; Owner: chechyotka
--

COPY public.team (name, "position", link, img, created_at) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: chechyotka
--

COPY public.users (login, password, role) FROM stdin;
admin	admin	ADMIN
test	test	ADMIN
\.


--
-- PostgreSQL database dump complete
--

