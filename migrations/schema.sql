--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: events; Type: TABLE; Schema: public; Owner: genos; Tablespace: 
--

CREATE TABLE public.events (
    id uuid NOT NULL,
    name character varying(50) NOT NULL,
    description text,
    schema text NOT NULL,
    output character varying(255) NOT NULL,
    status integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.events OWNER TO genos;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: genos; Tablespace: 
--

CREATE TABLE public.schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO genos;

--
-- Name: services; Type: TABLE; Schema: public; Owner: genos; Tablespace: 
--

CREATE TABLE public.services (
    id uuid NOT NULL,
    name character varying(50) NOT NULL,
    description text,
    prefix character varying(100) NOT NULL,
    status integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.services OWNER TO genos;

--
-- Name: events_pkey; Type: CONSTRAINT; Schema: public; Owner: genos; Tablespace: 
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: services_pkey; Type: CONSTRAINT; Schema: public; Owner: genos; Tablespace: 
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: genos; Tablespace: 
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

