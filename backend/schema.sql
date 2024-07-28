--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3 (Debian 16.3-1.pgdg120+1)
-- Dumped by pg_dump version 16.3

-- Started on 2024-07-28 22:28:00 IST

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

--
-- TOC entry 857 (class 1247 OID 16434)
-- Name: content_type; Type: TYPE; Schema: public; Owner: lms_admin
--

CREATE TYPE public.content_type AS ENUM (
    'ppt',
    'video'
);


ALTER TYPE public.content_type OWNER TO lms_admin;

--
-- TOC entry 848 (class 1247 OID 16396)
-- Name: emp_type; Type: TYPE; Schema: public; Owner: lms_admin
--

CREATE TYPE public.emp_type AS ENUM (
    'INTERN',
    'FULL',
    'CONTRACT'
);


ALTER TYPE public.emp_type OWNER TO lms_admin;

--
-- TOC entry 845 (class 1247 OID 16391)
-- Name: role_type; Type: TYPE; Schema: public; Owner: lms_admin
--

CREATE TYPE public.role_type AS ENUM (
    'ADMIN',
    'STUDENT'
);


ALTER TYPE public.role_type OWNER TO lms_admin;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 220 (class 1259 OID 16440)
-- Name: content; Type: TABLE; Schema: public; Owner: lms_admin
--

CREATE TABLE public.content (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    content_type public.content_type NOT NULL,
    content_url character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at timestamp without time zone
);


ALTER TABLE public.content OWNER TO lms_admin;

--
-- TOC entry 219 (class 1259 OID 16439)
-- Name: content_id_seq; Type: SEQUENCE; Schema: public; Owner: lms_admin
--

CREATE SEQUENCE public.content_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.content_id_seq OWNER TO lms_admin;

--
-- TOC entry 3384 (class 0 OID 0)
-- Dependencies: 219
-- Name: content_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: lms_admin
--

ALTER SEQUENCE public.content_id_seq OWNED BY public.content.id;


--
-- TOC entry 218 (class 1259 OID 16414)
-- Name: user_auth; Type: TABLE; Schema: public; Owner: lms_admin
--

CREATE TABLE public.user_auth (
    auth_id integer NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    role_type character varying(255) NOT NULL,
    registered_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.user_auth OWNER TO lms_admin;

--
-- TOC entry 217 (class 1259 OID 16413)
-- Name: user_auth_auth_id_seq; Type: SEQUENCE; Schema: public; Owner: lms_admin
--

CREATE SEQUENCE public.user_auth_auth_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_auth_auth_id_seq OWNER TO lms_admin;

--
-- TOC entry 3385 (class 0 OID 0)
-- Dependencies: 217
-- Name: user_auth_auth_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: lms_admin
--

ALTER SEQUENCE public.user_auth_auth_id_seq OWNED BY public.user_auth.auth_id;


--
-- TOC entry 216 (class 1259 OID 16404)
-- Name: users; Type: TABLE; Schema: public; Owner: lms_admin
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255) NOT NULL,
    role public.role_type NOT NULL,
    emp_type public.emp_type NOT NULL,
    team character varying(255),
    training_assigned character varying(255)[] NOT NULL,
    training_completed character varying(255)[] NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    email character varying(255) NOT NULL
);


ALTER TABLE public.users OWNER TO lms_admin;

--
-- TOC entry 215 (class 1259 OID 16403)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: lms_admin
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO lms_admin;

--
-- TOC entry 3386 (class 0 OID 0)
-- Dependencies: 215
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: lms_admin
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3226 (class 2604 OID 16443)
-- Name: content id; Type: DEFAULT; Schema: public; Owner: lms_admin
--

ALTER TABLE ONLY public.content ALTER COLUMN id SET DEFAULT nextval('public.content_id_seq'::regclass);


--
-- TOC entry 3224 (class 2604 OID 16417)
-- Name: user_auth auth_id; Type: DEFAULT; Schema: public; Owner: lms_admin
--

ALTER TABLE ONLY public.user_auth ALTER COLUMN auth_id SET DEFAULT nextval('public.user_auth_auth_id_seq'::regclass);


--
-- TOC entry 3222 (class 2604 OID 16407)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: lms_admin
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3235 (class 2606 OID 16448)
-- Name: content content_pkey; Type: CONSTRAINT; Schema: public; Owner: lms_admin
--

ALTER TABLE ONLY public.content
    ADD CONSTRAINT content_pkey PRIMARY KEY (id);


--
-- TOC entry 3229 (class 2606 OID 16432)
-- Name: users unique_email; Type: CONSTRAINT; Schema: public; Owner: lms_admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT unique_email UNIQUE (email);


--
-- TOC entry 3233 (class 2606 OID 16422)
-- Name: user_auth user_auth_pkey; Type: CONSTRAINT; Schema: public; Owner: lms_admin
--

ALTER TABLE ONLY public.user_auth
    ADD CONSTRAINT user_auth_pkey PRIMARY KEY (auth_id);


--
-- TOC entry 3231 (class 2606 OID 16412)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: lms_admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2024-07-28 22:28:00 IST

--
-- PostgreSQL database dump complete
--

