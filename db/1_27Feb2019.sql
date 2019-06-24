--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.3
-- Dumped by pg_dump version 9.6.3

-- Started on 2019-02-27 20:57:55 EST

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 12655)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 3210 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- TOC entry 597 (class 1247 OID 16386)
-- Name: admin_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE admin_type AS ENUM (
    'SUPER_ADMIN',
    'ADMIN',
    'SUPPORT'
);


ALTER TYPE admin_type OWNER TO postgres;

--
-- TOC entry 600 (class 1247 OID 16400)
-- Name: buy_sell; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE buy_sell AS ENUM (
    'BUY',
    'SELL'
);


ALTER TYPE buy_sell OWNER TO postgres;

--
-- TOC entry 831 (class 1247 OID 17409)
-- Name: currency_pair_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE currency_pair_type AS ENUM (
    'FIAT_TO_FIAT',
    'FIAT_TO_CRYPTO',
    'CRYPTO_TO_FIAT',
    'CRYPTO_TO_CRYPTO'
);


ALTER TYPE currency_pair_type OWNER TO postgres;

--
-- TOC entry 813 (class 1247 OID 17334)
-- Name: currency_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE currency_type AS ENUM (
    'FIAT',
    'CRYPTO'
);


ALTER TYPE currency_type OWNER TO postgres;

--
-- TOC entry 727 (class 1247 OID 18234)
-- Name: data_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE data_type AS ENUM (
    'TEXT',
    'INTEGER',
    'FLOAT',
    'CATEGORY',
    'TIMESTAMP',
    'FILE'
);


ALTER TYPE data_type OWNER TO postgres;

--
-- TOC entry 878 (class 1247 OID 17692)
-- Name: field_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE field_type AS ENUM (
    'FILE',
    'BANK',
    'ADDRESS'
);


ALTER TYPE field_type OWNER TO postgres;

--
-- TOC entry 776 (class 1247 OID 17135)
-- Name: partial_full; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE partial_full AS ENUM (
    'PARTIAL',
    'FULL',
    'NOT_PROCESSED'
);


ALTER TYPE partial_full OWNER TO postgres;

--
-- TOC entry 789 (class 1247 OID 18293)
-- Name: transfer_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE transfer_type AS ENUM (
    'WALLET',
    'PAYEE',
    'BANK'
);


ALTER TYPE transfer_type OWNER TO postgres;

--
-- TOC entry 686 (class 1247 OID 16422)
-- Name: yes_no; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE yes_no AS ENUM (
    'YES',
    'NO'
);


ALTER TYPE yes_no OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 266 (class 1259 OID 18396)
-- Name: admin_email; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE admin_email (
    id bigint NOT NULL,
    admin_id bigint NOT NULL,
    email character varying(250) NOT NULL,
    "primary" boolean NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE admin_email OWNER TO postgres;

--
-- TOC entry 265 (class 1259 OID 18394)
-- Name: AdminEmail_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE "AdminEmail_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE "AdminEmail_id_seq" OWNER TO postgres;

--
-- TOC entry 3211 (class 0 OID 0)
-- Dependencies: 265
-- Name: AdminEmail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE "AdminEmail_id_seq" OWNED BY admin_email.id;


--
-- TOC entry 187 (class 1259 OID 16440)
-- Name: admin; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE admin (
    id bigint NOT NULL,
    name character varying(250) NOT NULL,
    adminname character varying(250) NOT NULL,
    admin_group_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    edit_name_times integer NOT NULL
);


ALTER TABLE admin OWNER TO postgres;

--
-- TOC entry 188 (class 1259 OID 16446)
-- Name: admin_access; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE admin_access (
    id bigint NOT NULL,
    admin_id bigint NOT NULL,
    access_key character varying(250) NOT NULL,
    allow yes_no NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE admin_access OWNER TO postgres;

--
-- TOC entry 189 (class 1259 OID 16449)
-- Name: admin_access_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE admin_access_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE admin_access_id_seq OWNER TO postgres;

--
-- TOC entry 3212 (class 0 OID 0)
-- Dependencies: 189
-- Name: admin_access_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE admin_access_id_seq OWNED BY admin_access.id;


--
-- TOC entry 270 (class 1259 OID 18424)
-- Name: admin_country; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE admin_country (
    id bigint NOT NULL,
    admin_id bigint NOT NULL,
    country_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE admin_country OWNER TO postgres;

--
-- TOC entry 269 (class 1259 OID 18422)
-- Name: admin_country_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE admin_country_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE admin_country_id_seq OWNER TO postgres;

--
-- TOC entry 3213 (class 0 OID 0)
-- Dependencies: 269
-- Name: admin_country_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE admin_country_id_seq OWNED BY admin_country.id;


--
-- TOC entry 268 (class 1259 OID 18406)
-- Name: admin_group; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE admin_group (
    id bigint NOT NULL,
    key character varying(250) NOT NULL,
    name character varying(250) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    description text
);


ALTER TABLE admin_group OWNER TO postgres;

--
-- TOC entry 267 (class 1259 OID 18404)
-- Name: admin_group_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE admin_group_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE admin_group_id_seq OWNER TO postgres;

--
-- TOC entry 3214 (class 0 OID 0)
-- Dependencies: 267
-- Name: admin_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE admin_group_id_seq OWNED BY admin_group.id;


--
-- TOC entry 190 (class 1259 OID 16451)
-- Name: admin_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE admin_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE admin_id_seq OWNER TO postgres;

--
-- TOC entry 3215 (class 0 OID 0)
-- Dependencies: 190
-- Name: admin_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE admin_id_seq OWNED BY admin.id;


--
-- TOC entry 262 (class 1259 OID 18369)
-- Name: admin_password; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE admin_password (
    id bigint NOT NULL,
    admin_id bigint NOT NULL,
    password text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE admin_password OWNER TO postgres;

--
-- TOC entry 261 (class 1259 OID 18367)
-- Name: admin_password_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE admin_password_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE admin_password_id_seq OWNER TO postgres;

--
-- TOC entry 3216 (class 0 OID 0)
-- Dependencies: 261
-- Name: admin_password_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE admin_password_id_seq OWNED BY admin_password.id;


--
-- TOC entry 264 (class 1259 OID 18380)
-- Name: admin_token; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE admin_token (
    id bigint NOT NULL,
    admin_id bigint NOT NULL,
    token text NOT NULL,
    expiration_time timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE admin_token OWNER TO postgres;

--
-- TOC entry 263 (class 1259 OID 18378)
-- Name: admin_token_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE admin_token_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE admin_token_id_seq OWNER TO postgres;

--
-- TOC entry 3217 (class 0 OID 0)
-- Dependencies: 263
-- Name: admin_token_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE admin_token_id_seq OWNED BY admin_token.id;


--
-- TOC entry 224 (class 1259 OID 17431)
-- Name: api_key; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE api_key (
    id bigint NOT NULL,
    app_name character varying NOT NULL,
    description text NOT NULL,
    token text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    domain_name character varying(250) NOT NULL,
    ip_address character varying(20) NOT NULL
);


ALTER TABLE api_key OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 17429)
-- Name: api_key_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE api_key_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE api_key_id_seq OWNER TO postgres;

--
-- TOC entry 3218 (class 0 OID 0)
-- Dependencies: 223
-- Name: api_key_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE api_key_id_seq OWNED BY api_key.id;


--
-- TOC entry 191 (class 1259 OID 16453)
-- Name: auth_token; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE auth_token (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    token text NOT NULL,
    expiration_time timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE auth_token OWNER TO postgres;

--
-- TOC entry 192 (class 1259 OID 16459)
-- Name: auth_token_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE auth_token_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE auth_token_id_seq OWNER TO postgres;

--
-- TOC entry 3219 (class 0 OID 0)
-- Dependencies: 192
-- Name: auth_token_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE auth_token_id_seq OWNED BY auth_token.id;


--
-- TOC entry 193 (class 1259 OID 16474)
-- Name: country; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE country (
    id bigint NOT NULL,
    name character varying(250) NOT NULL,
    iso_code character varying(20) NOT NULL,
    dial_code character varying(20) NOT NULL,
    code character varying(20) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE country OWNER TO postgres;

--
-- TOC entry 194 (class 1259 OID 16477)
-- Name: country_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE country_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE country_id_seq OWNER TO postgres;

--
-- TOC entry 3220 (class 0 OID 0)
-- Dependencies: 194
-- Name: country_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE country_id_seq OWNED BY country.id;


--
-- TOC entry 220 (class 1259 OID 17341)
-- Name: currency; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE currency (
    id bigint NOT NULL,
    code character varying(20) NOT NULL,
    name character varying(150) NOT NULL,
    description character varying(250) NOT NULL,
    type currency_type NOT NULL,
    country_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE currency OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 17339)
-- Name: currency_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE currency_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE currency_id_seq OWNER TO postgres;

--
-- TOC entry 3221 (class 0 OID 0)
-- Dependencies: 219
-- Name: currency_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE currency_id_seq OWNED BY currency.id;


--
-- TOC entry 222 (class 1259 OID 17389)
-- Name: currency_pair; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE currency_pair (
    id bigint NOT NULL,
    from_currency_id bigint NOT NULL,
    to_currency_id bigint NOT NULL,
    created_at character varying NOT NULL,
    updated_at character varying NOT NULL,
    deleted_at character varying,
    pair_type currency_pair_type NOT NULL,
    supported yes_no NOT NULL
);


ALTER TABLE currency_pair OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 17387)
-- Name: currency_pair_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE currency_pair_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE currency_pair_id_seq OWNER TO postgres;

--
-- TOC entry 3222 (class 0 OID 0)
-- Dependencies: 221
-- Name: currency_pair_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE currency_pair_id_seq OWNED BY currency_pair.id;


--
-- TOC entry 232 (class 1259 OID 17607)
-- Name: data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE data (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    country_id bigint NOT NULL,
    nickname character varying(250),
    "primary" boolean NOT NULL,
    active boolean NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    field_type field_type
);


ALTER TABLE data OWNER TO postgres;

--
-- TOC entry 238 (class 1259 OID 17664)
-- Name: data_category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE data_category (
    id bigint NOT NULL,
    data_id bigint NOT NULL,
    field_id bigint NOT NULL,
    field_category_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE data_category OWNER TO postgres;

--
-- TOC entry 237 (class 1259 OID 17662)
-- Name: data_category_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE data_category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE data_category_id_seq OWNER TO postgres;

--
-- TOC entry 3223 (class 0 OID 0)
-- Dependencies: 237
-- Name: data_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE data_category_id_seq OWNED BY data_category.id;


--
-- TOC entry 236 (class 1259 OID 17646)
-- Name: data_file; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE data_file (
    id bigint NOT NULL,
    data_id bigint NOT NULL,
    field_id bigint NOT NULL,
    name character varying(150) NOT NULL,
    extension character varying(20) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE data_file OWNER TO postgres;

--
-- TOC entry 235 (class 1259 OID 17644)
-- Name: data_file_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE data_file_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE data_file_id_seq OWNER TO postgres;

--
-- TOC entry 3224 (class 0 OID 0)
-- Dependencies: 235
-- Name: data_file_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE data_file_id_seq OWNED BY data_file.id;


--
-- TOC entry 231 (class 1259 OID 17605)
-- Name: data_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE data_id_seq OWNER TO postgres;

--
-- TOC entry 3225 (class 0 OID 0)
-- Dependencies: 231
-- Name: data_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE data_id_seq OWNED BY data.id;


--
-- TOC entry 234 (class 1259 OID 17615)
-- Name: data_text; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE data_text (
    id bigint NOT NULL,
    data_id bigint NOT NULL,
    field_id bigint NOT NULL,
    text text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE data_text OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 17613)
-- Name: data_text_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE data_text_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE data_text_id_seq OWNER TO postgres;

--
-- TOC entry 3226 (class 0 OID 0)
-- Dependencies: 233
-- Name: data_text_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE data_text_id_seq OWNED BY data_text.id;


--
-- TOC entry 209 (class 1259 OID 16939)
-- Name: email; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE email (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    email character varying(250) NOT NULL,
    "primary" yes_no NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE email OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 16950)
-- Name: email_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE email_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE email_id_seq OWNER TO postgres;

--
-- TOC entry 3227 (class 0 OID 0)
-- Dependencies: 210
-- Name: email_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE email_id_seq OWNED BY email.id;


--
-- TOC entry 197 (class 1259 OID 16487)
-- Name: field; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE field (
    id bigint NOT NULL,
    country_id bigint NOT NULL,
    name character varying(250) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    description text NOT NULL,
    "order" bigint NOT NULL,
    field_type field_type,
    key character varying(250),
    data_type data_type,
    has_category boolean,
    has_input_text boolean,
    has_file boolean
);


ALTER TABLE field OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 17590)
-- Name: field_category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE field_category (
    id bigint NOT NULL,
    field_id bigint NOT NULL,
    name character varying(250) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE field_category OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 17588)
-- Name: field_category_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE field_category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE field_category_id_seq OWNER TO postgres;

--
-- TOC entry 3228 (class 0 OID 0)
-- Dependencies: 229
-- Name: field_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE field_category_id_seq OWNED BY field_category.id;


--
-- TOC entry 198 (class 1259 OID 16490)
-- Name: field_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE field_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE field_id_seq OWNER TO postgres;

--
-- TOC entry 3229 (class 0 OID 0)
-- Dependencies: 198
-- Name: field_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE field_id_seq OWNED BY field.id;


--
-- TOC entry 258 (class 1259 OID 18249)
-- Name: field_subcategory; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE field_subcategory (
    id bigint NOT NULL,
    field_id bigint NOT NULL,
    sub_field_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE field_subcategory OWNER TO postgres;

--
-- TOC entry 257 (class 1259 OID 18247)
-- Name: field_subcategory_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE field_subcategory_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE field_subcategory_id_seq OWNER TO postgres;

--
-- TOC entry 3230 (class 0 OID 0)
-- Dependencies: 257
-- Name: field_subcategory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE field_subcategory_id_seq OWNED BY field_subcategory.id;


--
-- TOC entry 218 (class 1259 OID 17237)
-- Name: global_setting; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE global_setting (
    id bigint NOT NULL,
    country_id bigint NOT NULL,
    key character varying(250) NOT NULL,
    value text NOT NULL,
    admin_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE global_setting OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 17235)
-- Name: global_setting_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE global_setting_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE global_setting_id_seq OWNER TO postgres;

--
-- TOC entry 3231 (class 0 OID 0)
-- Dependencies: 217
-- Name: global_setting_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE global_setting_id_seq OWNED BY global_setting.id;


--
-- TOC entry 199 (class 1259 OID 16508)
-- Name: mobile; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE mobile (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    number character varying(20) NOT NULL,
    "primary" yes_no NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE mobile OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 16511)
-- Name: mobile_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE mobile_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE mobile_id_seq OWNER TO postgres;

--
-- TOC entry 3232 (class 0 OID 0)
-- Dependencies: 200
-- Name: mobile_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE mobile_id_seq OWNED BY mobile.id;


--
-- TOC entry 201 (class 1259 OID 16513)
-- Name: order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE "order" (
    id bigint NOT NULL,
    seller_user_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    amount double precision NOT NULL,
    rate double precision NOT NULL,
    total_amount double precision NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    buyer_user_id bigint NOT NULL,
    seller_currency_wallet_id bigint NOT NULL,
    seller_rate_currency_wallet_id bigint NOT NULL,
    buyer_currency_wallet_id bigint NOT NULL,
    buyer_rate_currency_wallet_id bigint NOT NULL
);


ALTER TABLE "order" OWNER TO postgres;

--
-- TOC entry 280 (class 1259 OID 18967)
-- Name: order_broker; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_broker (
    id bigint NOT NULL,
    order_id bigint NOT NULL,
    seller_broker_id bigint NOT NULL,
    buyer_broker_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE order_broker OWNER TO postgres;

--
-- TOC entry 279 (class 1259 OID 18965)
-- Name: order_broker_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_broker_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_broker_id_seq OWNER TO postgres;

--
-- TOC entry 3233 (class 0 OID 0)
-- Dependencies: 279
-- Name: order_broker_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_broker_id_seq OWNED BY order_broker.id;


--
-- TOC entry 272 (class 1259 OID 18804)
-- Name: order_buy; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_buy (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    amount double precision NOT NULL,
    rate double precision NOT NULL,
    total_amount double precision NOT NULL,
    created_at timestamp with time zone NOT NULL,
    currency_wallet_id bigint NOT NULL,
    rate_currency_wallet_id bigint NOT NULL,
    lock boolean NOT NULL
);


ALTER TABLE order_buy OWNER TO postgres;

--
-- TOC entry 282 (class 1259 OID 18991)
-- Name: order_buy_broker; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_buy_broker (
    id bigint NOT NULL,
    order_buy_id bigint NOT NULL,
    broker_id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE order_buy_broker OWNER TO postgres;

--
-- TOC entry 281 (class 1259 OID 18989)
-- Name: order_buy_broker_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_buy_broker_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_buy_broker_id_seq OWNER TO postgres;

--
-- TOC entry 3234 (class 0 OID 0)
-- Dependencies: 281
-- Name: order_buy_broker_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_buy_broker_id_seq OWNED BY order_buy_broker.id;


--
-- TOC entry 271 (class 1259 OID 18802)
-- Name: order_buy_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_buy_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_buy_id_seq OWNER TO postgres;

--
-- TOC entry 3235 (class 0 OID 0)
-- Dependencies: 271
-- Name: order_buy_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_buy_id_seq OWNED BY order_buy.id;


--
-- TOC entry 276 (class 1259 OID 18876)
-- Name: order_currency; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_currency (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE order_currency OWNER TO postgres;

--
-- TOC entry 275 (class 1259 OID 18874)
-- Name: order_currency_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_currency_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_currency_id_seq OWNER TO postgres;

--
-- TOC entry 3236 (class 0 OID 0)
-- Dependencies: 275
-- Name: order_currency_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_currency_id_seq OWNED BY order_currency.id;


--
-- TOC entry 300 (class 1259 OID 19339)
-- Name: order_graph_12h; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_graph_12h (
    id bigint NOT NULL,
    last_order_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    open double precision NOT NULL,
    high double precision NOT NULL,
    low double precision NOT NULL,
    close double precision NOT NULL,
    volume double precision NOT NULL,
    split double precision NOT NULL,
    dividend double precision NOT NULL,
    absolute_change double precision NOT NULL,
    percent_change double precision NOT NULL,
    date timestamp with time zone NOT NULL
);


ALTER TABLE order_graph_12h OWNER TO postgres;

--
-- TOC entry 299 (class 1259 OID 19337)
-- Name: order_graph_12h_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_graph_12h_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_graph_12h_id_seq OWNER TO postgres;

--
-- TOC entry 3237 (class 0 OID 0)
-- Dependencies: 299
-- Name: order_graph_12h_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_graph_12h_id_seq OWNED BY order_graph_12h.id;


--
-- TOC entry 292 (class 1259 OID 19247)
-- Name: order_graph_15m; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_graph_15m (
    id bigint NOT NULL,
    last_order_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    open double precision NOT NULL,
    high double precision NOT NULL,
    low double precision NOT NULL,
    close double precision NOT NULL,
    volume double precision NOT NULL,
    split double precision NOT NULL,
    dividend double precision NOT NULL,
    absolute_change double precision NOT NULL,
    percent_change double precision NOT NULL,
    date timestamp with time zone NOT NULL
);


ALTER TABLE order_graph_15m OWNER TO postgres;

--
-- TOC entry 291 (class 1259 OID 19245)
-- Name: order_graph_15m_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_graph_15m_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_graph_15m_id_seq OWNER TO postgres;

--
-- TOC entry 3238 (class 0 OID 0)
-- Dependencies: 291
-- Name: order_graph_15m_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_graph_15m_id_seq OWNED BY order_graph_15m.id;


--
-- TOC entry 302 (class 1259 OID 19362)
-- Name: order_graph_1d; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_graph_1d (
    id bigint NOT NULL,
    last_order_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    open double precision NOT NULL,
    high double precision NOT NULL,
    low double precision NOT NULL,
    close double precision NOT NULL,
    volume double precision NOT NULL,
    split double precision NOT NULL,
    dividend double precision NOT NULL,
    absolutue_change double precision NOT NULL,
    percent_change double precision NOT NULL,
    date timestamp with time zone NOT NULL
);


ALTER TABLE order_graph_1d OWNER TO postgres;

--
-- TOC entry 301 (class 1259 OID 19360)
-- Name: order_graph_1d_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_graph_1d_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_graph_1d_id_seq OWNER TO postgres;

--
-- TOC entry 3239 (class 0 OID 0)
-- Dependencies: 301
-- Name: order_graph_1d_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_graph_1d_id_seq OWNED BY order_graph_1d.id;


--
-- TOC entry 296 (class 1259 OID 19293)
-- Name: order_graph_1h; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_graph_1h (
    id bigint NOT NULL,
    last_order_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    open double precision NOT NULL,
    high double precision NOT NULL,
    low double precision NOT NULL,
    close double precision NOT NULL,
    volume double precision NOT NULL,
    split double precision NOT NULL,
    dividend double precision NOT NULL,
    absolute_change double precision NOT NULL,
    percent_change double precision NOT NULL,
    date timestamp with time zone NOT NULL
);


ALTER TABLE order_graph_1h OWNER TO postgres;

--
-- TOC entry 295 (class 1259 OID 19291)
-- Name: order_graph_1h_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_graph_1h_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_graph_1h_id_seq OWNER TO postgres;

--
-- TOC entry 3240 (class 0 OID 0)
-- Dependencies: 295
-- Name: order_graph_1h_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_graph_1h_id_seq OWNED BY order_graph_1h.id;


--
-- TOC entry 290 (class 1259 OID 19224)
-- Name: order_graph_1m; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_graph_1m (
    id bigint NOT NULL,
    last_order_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    open double precision NOT NULL,
    high double precision NOT NULL,
    low double precision NOT NULL,
    close double precision NOT NULL,
    volume double precision NOT NULL,
    split double precision NOT NULL,
    dividend double precision NOT NULL,
    absolute_change double precision NOT NULL,
    percent_change double precision NOT NULL,
    date timestamp with time zone NOT NULL
);


ALTER TABLE order_graph_1m OWNER TO postgres;

--
-- TOC entry 289 (class 1259 OID 19222)
-- Name: order_graph_1m_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_graph_1m_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_graph_1m_id_seq OWNER TO postgres;

--
-- TOC entry 3241 (class 0 OID 0)
-- Dependencies: 289
-- Name: order_graph_1m_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_graph_1m_id_seq OWNED BY order_graph_1m.id;


--
-- TOC entry 294 (class 1259 OID 19270)
-- Name: order_graph_30m; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_graph_30m (
    id bigint NOT NULL,
    last_order_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    open double precision NOT NULL,
    high double precision NOT NULL,
    low double precision NOT NULL,
    close double precision NOT NULL,
    volume double precision NOT NULL,
    split double precision NOT NULL,
    dividend double precision NOT NULL,
    absolutue_change double precision NOT NULL,
    percent_change double precision NOT NULL,
    date timestamp with time zone NOT NULL
);


ALTER TABLE order_graph_30m OWNER TO postgres;

--
-- TOC entry 293 (class 1259 OID 19268)
-- Name: order_graph_30m_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_graph_30m_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_graph_30m_id_seq OWNER TO postgres;

--
-- TOC entry 3242 (class 0 OID 0)
-- Dependencies: 293
-- Name: order_graph_30m_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_graph_30m_id_seq OWNED BY order_graph_30m.id;


--
-- TOC entry 286 (class 1259 OID 19041)
-- Name: order_graph_5m; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_graph_5m (
    id bigint NOT NULL,
    open double precision NOT NULL,
    high double precision NOT NULL,
    low double precision NOT NULL,
    close double precision NOT NULL,
    volume double precision NOT NULL,
    split character varying(250),
    dividend character varying(250),
    absolute_change double precision NOT NULL,
    percent_change double precision NOT NULL,
    date timestamp with time zone NOT NULL,
    last_order_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL
);


ALTER TABLE order_graph_5m OWNER TO postgres;

--
-- TOC entry 298 (class 1259 OID 19316)
-- Name: order_graph_6h; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_graph_6h (
    id bigint NOT NULL,
    last_order_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    open double precision NOT NULL,
    high double precision NOT NULL,
    low double precision NOT NULL,
    close double precision NOT NULL,
    volume double precision NOT NULL,
    split double precision NOT NULL,
    dividend double precision NOT NULL,
    absolute_change double precision NOT NULL,
    percent_change double precision NOT NULL,
    date timestamp with time zone NOT NULL
);


ALTER TABLE order_graph_6h OWNER TO postgres;

--
-- TOC entry 297 (class 1259 OID 19314)
-- Name: order_graph_6h_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_graph_6h_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_graph_6h_id_seq OWNER TO postgres;

--
-- TOC entry 3243 (class 0 OID 0)
-- Dependencies: 297
-- Name: order_graph_6h_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_graph_6h_id_seq OWNED BY order_graph_6h.id;


--
-- TOC entry 304 (class 1259 OID 19385)
-- Name: order_graph_7d; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_graph_7d (
    id bigint NOT NULL,
    last_order_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    open double precision NOT NULL,
    high double precision NOT NULL,
    low double precision NOT NULL,
    close double precision NOT NULL,
    volume double precision NOT NULL,
    split double precision NOT NULL,
    dividend double precision NOT NULL,
    absolute_change double precision NOT NULL,
    percent_change double precision NOT NULL,
    date timestamp with time zone NOT NULL
);


ALTER TABLE order_graph_7d OWNER TO postgres;

--
-- TOC entry 303 (class 1259 OID 19383)
-- Name: order_graph_7d_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_graph_7d_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_graph_7d_id_seq OWNER TO postgres;

--
-- TOC entry 3244 (class 0 OID 0)
-- Dependencies: 303
-- Name: order_graph_7d_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_graph_7d_id_seq OWNED BY order_graph_7d.id;


--
-- TOC entry 285 (class 1259 OID 19039)
-- Name: order_graph_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_graph_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_graph_id_seq OWNER TO postgres;

--
-- TOC entry 3245 (class 0 OID 0)
-- Dependencies: 285
-- Name: order_graph_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_graph_id_seq OWNED BY order_graph_5m.id;


--
-- TOC entry 202 (class 1259 OID 16516)
-- Name: order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_id_seq OWNER TO postgres;

--
-- TOC entry 3246 (class 0 OID 0)
-- Dependencies: 202
-- Name: order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_id_seq OWNED BY "order".id;


--
-- TOC entry 274 (class 1259 OID 18812)
-- Name: order_sell; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_sell (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    amount double precision NOT NULL,
    rate double precision NOT NULL,
    total_amount double precision NOT NULL,
    created_at timestamp with time zone NOT NULL,
    currency_wallet_id bigint NOT NULL,
    rate_currency_wallet_id bigint NOT NULL,
    lock boolean NOT NULL
);


ALTER TABLE order_sell OWNER TO postgres;

--
-- TOC entry 284 (class 1259 OID 19009)
-- Name: order_sell_broker; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_sell_broker (
    id bigint NOT NULL,
    order_sell_id bigint NOT NULL,
    broker_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone NOT NULL
);


ALTER TABLE order_sell_broker OWNER TO postgres;

--
-- TOC entry 283 (class 1259 OID 19007)
-- Name: order_sell_broker_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_sell_broker_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_sell_broker_id_seq OWNER TO postgres;

--
-- TOC entry 3247 (class 0 OID 0)
-- Dependencies: 283
-- Name: order_sell_broker_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_sell_broker_id_seq OWNED BY order_sell_broker.id;


--
-- TOC entry 273 (class 1259 OID 18810)
-- Name: order_sell_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_sell_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_sell_id_seq OWNER TO postgres;

--
-- TOC entry 3248 (class 0 OID 0)
-- Dependencies: 273
-- Name: order_sell_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_sell_id_seq OWNED BY order_sell.id;


--
-- TOC entry 288 (class 1259 OID 19064)
-- Name: order_wallet; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE order_wallet (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    rate_currency_id bigint NOT NULL,
    currency_wallet_id bigint NOT NULL,
    rate_currency_wallet_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE order_wallet OWNER TO postgres;

--
-- TOC entry 287 (class 1259 OID 19062)
-- Name: order_wallet_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE order_wallet_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE order_wallet_id_seq OWNER TO postgres;

--
-- TOC entry 3249 (class 0 OID 0)
-- Dependencies: 287
-- Name: order_wallet_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE order_wallet_id_seq OWNED BY order_wallet.id;


--
-- TOC entry 212 (class 1259 OID 17016)
-- Name: password; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE password (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    password text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE password OWNER TO postgres;

--
-- TOC entry 248 (class 1259 OID 18119)
-- Name: payee; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE payee (
    id bigint NOT NULL,
    payee_master_id bigint NOT NULL,
    name character varying(250) NOT NULL,
    nickname character varying(250) NOT NULL,
    email character varying(250) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE payee OWNER TO postgres;

--
-- TOC entry 250 (class 1259 OID 18132)
-- Name: payee_crypto; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE payee_crypto (
    id bigint NOT NULL,
    payee_id bigint NOT NULL,
    address text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE payee_crypto OWNER TO postgres;

--
-- TOC entry 249 (class 1259 OID 18130)
-- Name: payee_crypto_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE payee_crypto_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE payee_crypto_id_seq OWNER TO postgres;

--
-- TOC entry 3250 (class 0 OID 0)
-- Dependencies: 249
-- Name: payee_crypto_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE payee_crypto_id_seq OWNED BY payee_crypto.id;


--
-- TOC entry 256 (class 1259 OID 18217)
-- Name: payee_fiat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE payee_fiat (
    id bigint NOT NULL,
    payee_id bigint NOT NULL,
    data_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE payee_fiat OWNER TO postgres;

--
-- TOC entry 255 (class 1259 OID 18215)
-- Name: payee_fiat_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE payee_fiat_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE payee_fiat_id_seq OWNER TO postgres;

--
-- TOC entry 3251 (class 0 OID 0)
-- Dependencies: 255
-- Name: payee_fiat_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE payee_fiat_id_seq OWNED BY payee_fiat.id;


--
-- TOC entry 247 (class 1259 OID 18117)
-- Name: payee_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE payee_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE payee_id_seq OWNER TO postgres;

--
-- TOC entry 3252 (class 0 OID 0)
-- Dependencies: 247
-- Name: payee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE payee_id_seq OWNED BY payee.id;


--
-- TOC entry 246 (class 1259 OID 18092)
-- Name: payee_master; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE payee_master (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    currency_type currency_type NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE payee_master OWNER TO postgres;

--
-- TOC entry 245 (class 1259 OID 18090)
-- Name: payee_master_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE payee_master_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE payee_master_id_seq OWNER TO postgres;

--
-- TOC entry 3253 (class 0 OID 0)
-- Dependencies: 245
-- Name: payee_master_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE payee_master_id_seq OWNED BY payee_master.id;


--
-- TOC entry 278 (class 1259 OID 18903)
-- Name: site_bank; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE site_bank (
    id bigint NOT NULL,
    updated_admin_id bigint NOT NULL,
    country_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    beneficiary_name character varying(250) NOT NULL,
    beneficiary_address character varying(500) NOT NULL,
    beneficiary_city character varying(250) NOT NULL,
    beneficiary_province character varying(250) NOT NULL,
    beneficiary_postal_code character varying(20) NOT NULL,
    beneficiary_account_number character varying(100) NOT NULL,
    reference_text character varying(500) NOT NULL,
    bank_name character varying(250) NOT NULL,
    bank_address character varying(500) NOT NULL,
    bank_province character varying(250) NOT NULL,
    bank_postal_code character varying(20) NOT NULL,
    bank_phone_number character varying(20) NOT NULL,
    bank_institution_number character varying(20) NOT NULL,
    bank_transit_number character varying(20) NOT NULL,
    bank_swift_code character varying(20) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE site_bank OWNER TO postgres;

--
-- TOC entry 277 (class 1259 OID 18901)
-- Name: site_bank_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE site_bank_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE site_bank_id_seq OWNER TO postgres;

--
-- TOC entry 3254 (class 0 OID 0)
-- Dependencies: 277
-- Name: site_bank_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE site_bank_id_seq OWNED BY site_bank.id;


--
-- TOC entry 203 (class 1259 OID 16518)
-- Name: transaction; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE transaction (
    id bigint NOT NULL,
    order_id bigint NOT NULL,
    from_user bigint NOT NULL,
    to_user bigint NOT NULL,
    amount double precision NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    currency_id bigint
);


ALTER TABLE transaction OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 16521)
-- Name: transaction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE transaction_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transaction_id_seq OWNER TO postgres;

--
-- TOC entry 3255 (class 0 OID 0)
-- Dependencies: 204
-- Name: transaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE transaction_id_seq OWNED BY transaction.id;


--
-- TOC entry 216 (class 1259 OID 17224)
-- Name: transfer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE transfer (
    id bigint NOT NULL,
    transfer_master_id bigint NOT NULL,
    amount double precision NOT NULL,
    transfer_type transfer_type NOT NULL,
    from_user_id bigint NOT NULL,
    to_user_id bigint NOT NULL,
    processed_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE transfer OWNER TO postgres;

--
-- TOC entry 260 (class 1259 OID 18271)
-- Name: transfer_bank; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE transfer_bank (
    id bigint NOT NULL,
    transfer_id bigint NOT NULL,
    from_wallet_id bigint NOT NULL,
    to_data_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE transfer_bank OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 17222)
-- Name: transfer_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE transfer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transfer_id_seq OWNER TO postgres;

--
-- TOC entry 3256 (class 0 OID 0)
-- Dependencies: 215
-- Name: transfer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE transfer_id_seq OWNED BY transfer.id;


--
-- TOC entry 244 (class 1259 OID 18060)
-- Name: transfer_master; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE transfer_master (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    currency_type currency_type NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE transfer_master OWNER TO postgres;

--
-- TOC entry 243 (class 1259 OID 18058)
-- Name: transfer_master_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE transfer_master_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transfer_master_id_seq OWNER TO postgres;

--
-- TOC entry 3257 (class 0 OID 0)
-- Dependencies: 243
-- Name: transfer_master_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE transfer_master_id_seq OWNED BY transfer_master.id;


--
-- TOC entry 254 (class 1259 OID 18194)
-- Name: transfer_payee; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE transfer_payee (
    id bigint NOT NULL,
    transfer_id bigint NOT NULL,
    from_wallet_id bigint NOT NULL,
    to_payee_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE transfer_payee OWNER TO postgres;

--
-- TOC entry 253 (class 1259 OID 18192)
-- Name: transfer_payee_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE transfer_payee_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transfer_payee_id_seq OWNER TO postgres;

--
-- TOC entry 3258 (class 0 OID 0)
-- Dependencies: 253
-- Name: transfer_payee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE transfer_payee_id_seq OWNED BY transfer_payee.id;


--
-- TOC entry 252 (class 1259 OID 18171)
-- Name: transfer_wallet; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE transfer_wallet (
    id bigint NOT NULL,
    transfer_id bigint NOT NULL,
    from_wallet_id bigint NOT NULL,
    to_wallet_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE transfer_wallet OWNER TO postgres;

--
-- TOC entry 251 (class 1259 OID 18169)
-- Name: transfer_wallet_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE transfer_wallet_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transfer_wallet_id_seq OWNER TO postgres;

--
-- TOC entry 3259 (class 0 OID 0)
-- Dependencies: 251
-- Name: transfer_wallet_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE transfer_wallet_id_seq OWNED BY transfer_wallet.id;


--
-- TOC entry 259 (class 1259 OID 18269)
-- Name: transfer_withdraw_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE transfer_withdraw_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE transfer_withdraw_id_seq OWNER TO postgres;

--
-- TOC entry 3260 (class 0 OID 0)
-- Dependencies: 259
-- Name: transfer_withdraw_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE transfer_withdraw_id_seq OWNED BY transfer_bank.id;


--
-- TOC entry 205 (class 1259 OID 16523)
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE "user" (
    id bigint NOT NULL,
    name character varying(250) NOT NULL,
    username character varying(250) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    edit_name_times integer DEFAULT 0 NOT NULL
);


ALTER TABLE "user" OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 17564)
-- Name: user_country; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE user_country (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    country_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    eligible yes_no NOT NULL
);


ALTER TABLE user_country OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 17562)
-- Name: user_country_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE user_country_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_country_id_seq OWNER TO postgres;

--
-- TOC entry 3261 (class 0 OID 0)
-- Dependencies: 227
-- Name: user_country_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_country_id_seq OWNED BY user_country.id;


--
-- TOC entry 206 (class 1259 OID 16529)
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_id_seq OWNER TO postgres;

--
-- TOC entry 3262 (class 0 OID 0)
-- Dependencies: 206
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- TOC entry 211 (class 1259 OID 17014)
-- Name: user_password_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE user_password_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_password_id_seq OWNER TO postgres;

--
-- TOC entry 3263 (class 0 OID 0)
-- Dependencies: 211
-- Name: user_password_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_password_id_seq OWNED BY password.id;


--
-- TOC entry 214 (class 1259 OID 17027)
-- Name: user_setting; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE user_setting (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    key character varying(250) NOT NULL,
    value text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE user_setting OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 17025)
-- Name: user_setting_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE user_setting_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_setting_id_seq OWNER TO postgres;

--
-- TOC entry 3264 (class 0 OID 0)
-- Dependencies: 213
-- Name: user_setting_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_setting_id_seq OWNED BY user_setting.id;


--
-- TOC entry 207 (class 1259 OID 16531)
-- Name: wallet; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wallet (
    id bigint NOT NULL,
    wallet_master_id bigint NOT NULL,
    amount double precision NOT NULL,
    amount_locked double precision NOT NULL,
    nickname character varying(250) NOT NULL,
    "primary" boolean NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE wallet OWNER TO postgres;

--
-- TOC entry 195 (class 1259 OID 16479)
-- Name: wallet_crypto; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wallet_crypto (
    id bigint NOT NULL,
    wallet_id bigint NOT NULL,
    address text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE wallet_crypto OWNER TO postgres;

--
-- TOC entry 196 (class 1259 OID 16485)
-- Name: wallet_crypto_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE wallet_crypto_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE wallet_crypto_id_seq OWNER TO postgres;

--
-- TOC entry 3265 (class 0 OID 0)
-- Dependencies: 196
-- Name: wallet_crypto_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE wallet_crypto_id_seq OWNED BY wallet_crypto.id;


--
-- TOC entry 240 (class 1259 OID 17903)
-- Name: wallet_fiat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wallet_fiat (
    id bigint NOT NULL,
    wallet_id bigint NOT NULL,
    data_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE wallet_fiat OWNER TO postgres;

--
-- TOC entry 239 (class 1259 OID 17901)
-- Name: wallet_fiat_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE wallet_fiat_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE wallet_fiat_id_seq OWNER TO postgres;

--
-- TOC entry 3266 (class 0 OID 0)
-- Dependencies: 239
-- Name: wallet_fiat_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE wallet_fiat_id_seq OWNED BY wallet_fiat.id;


--
-- TOC entry 208 (class 1259 OID 16534)
-- Name: wallet_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE wallet_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE wallet_id_seq OWNER TO postgres;

--
-- TOC entry 3267 (class 0 OID 0)
-- Dependencies: 208
-- Name: wallet_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE wallet_id_seq OWNED BY wallet.id;


--
-- TOC entry 242 (class 1259 OID 17987)
-- Name: wallet_master; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wallet_master (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    currency_type currency_type NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE wallet_master OWNER TO postgres;

--
-- TOC entry 241 (class 1259 OID 17985)
-- Name: wallet_master_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE wallet_master_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE wallet_master_id_seq OWNER TO postgres;

--
-- TOC entry 3268 (class 0 OID 0)
-- Dependencies: 241
-- Name: wallet_master_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE wallet_master_id_seq OWNED BY wallet_master.id;


--
-- TOC entry 226 (class 1259 OID 17442)
-- Name: wallet_passphrase; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE wallet_passphrase (
    id bigint NOT NULL,
    wallet_id bigint NOT NULL,
    passphrase text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE wallet_passphrase OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 17440)
-- Name: wallet_passphrase_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE wallet_passphrase_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE wallet_passphrase_id_seq OWNER TO postgres;

--
-- TOC entry 3269 (class 0 OID 0)
-- Dependencies: 225
-- Name: wallet_passphrase_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE wallet_passphrase_id_seq OWNED BY wallet_passphrase.id;


--
-- TOC entry 2664 (class 2604 OID 16538)
-- Name: admin id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin ALTER COLUMN id SET DEFAULT nextval('admin_id_seq'::regclass);


--
-- TOC entry 2665 (class 2604 OID 16539)
-- Name: admin_access id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_access ALTER COLUMN id SET DEFAULT nextval('admin_access_id_seq'::regclass);


--
-- TOC entry 2706 (class 2604 OID 18427)
-- Name: admin_country id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_country ALTER COLUMN id SET DEFAULT nextval('admin_country_id_seq'::regclass);


--
-- TOC entry 2704 (class 2604 OID 18399)
-- Name: admin_email id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_email ALTER COLUMN id SET DEFAULT nextval('"AdminEmail_id_seq"'::regclass);


--
-- TOC entry 2705 (class 2604 OID 18409)
-- Name: admin_group id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_group ALTER COLUMN id SET DEFAULT nextval('admin_group_id_seq'::regclass);


--
-- TOC entry 2702 (class 2604 OID 18372)
-- Name: admin_password id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_password ALTER COLUMN id SET DEFAULT nextval('admin_password_id_seq'::regclass);


--
-- TOC entry 2703 (class 2604 OID 18383)
-- Name: admin_token id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_token ALTER COLUMN id SET DEFAULT nextval('admin_token_id_seq'::regclass);


--
-- TOC entry 2683 (class 2604 OID 17434)
-- Name: api_key id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY api_key ALTER COLUMN id SET DEFAULT nextval('api_key_id_seq'::regclass);


--
-- TOC entry 2666 (class 2604 OID 16540)
-- Name: auth_token id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY auth_token ALTER COLUMN id SET DEFAULT nextval('auth_token_id_seq'::regclass);


--
-- TOC entry 2667 (class 2604 OID 16543)
-- Name: country id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY country ALTER COLUMN id SET DEFAULT nextval('country_id_seq'::regclass);


--
-- TOC entry 2681 (class 2604 OID 17344)
-- Name: currency id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY currency ALTER COLUMN id SET DEFAULT nextval('currency_id_seq'::regclass);


--
-- TOC entry 2682 (class 2604 OID 17392)
-- Name: currency_pair id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY currency_pair ALTER COLUMN id SET DEFAULT nextval('currency_pair_id_seq'::regclass);


--
-- TOC entry 2687 (class 2604 OID 17610)
-- Name: data id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data ALTER COLUMN id SET DEFAULT nextval('data_id_seq'::regclass);


--
-- TOC entry 2690 (class 2604 OID 17667)
-- Name: data_category id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_category ALTER COLUMN id SET DEFAULT nextval('data_category_id_seq'::regclass);


--
-- TOC entry 2689 (class 2604 OID 17649)
-- Name: data_file id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_file ALTER COLUMN id SET DEFAULT nextval('data_file_id_seq'::regclass);


--
-- TOC entry 2688 (class 2604 OID 17618)
-- Name: data_text id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_text ALTER COLUMN id SET DEFAULT nextval('data_text_id_seq'::regclass);


--
-- TOC entry 2676 (class 2604 OID 16952)
-- Name: email id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY email ALTER COLUMN id SET DEFAULT nextval('email_id_seq'::regclass);


--
-- TOC entry 2669 (class 2604 OID 16545)
-- Name: field id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY field ALTER COLUMN id SET DEFAULT nextval('field_id_seq'::regclass);


--
-- TOC entry 2686 (class 2604 OID 17593)
-- Name: field_category id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY field_category ALTER COLUMN id SET DEFAULT nextval('field_category_id_seq'::regclass);


--
-- TOC entry 2700 (class 2604 OID 18252)
-- Name: field_subcategory id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY field_subcategory ALTER COLUMN id SET DEFAULT nextval('field_subcategory_id_seq'::regclass);


--
-- TOC entry 2680 (class 2604 OID 17240)
-- Name: global_setting id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY global_setting ALTER COLUMN id SET DEFAULT nextval('global_setting_id_seq'::regclass);


--
-- TOC entry 2670 (class 2604 OID 16548)
-- Name: mobile id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY mobile ALTER COLUMN id SET DEFAULT nextval('mobile_id_seq'::regclass);


--
-- TOC entry 2671 (class 2604 OID 16549)
-- Name: order id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "order" ALTER COLUMN id SET DEFAULT nextval('order_id_seq'::regclass);


--
-- TOC entry 2711 (class 2604 OID 18970)
-- Name: order_broker id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_broker ALTER COLUMN id SET DEFAULT nextval('order_broker_id_seq'::regclass);


--
-- TOC entry 2707 (class 2604 OID 18807)
-- Name: order_buy id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_buy ALTER COLUMN id SET DEFAULT nextval('order_buy_id_seq'::regclass);


--
-- TOC entry 2712 (class 2604 OID 18994)
-- Name: order_buy_broker id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_buy_broker ALTER COLUMN id SET DEFAULT nextval('order_buy_broker_id_seq'::regclass);


--
-- TOC entry 2709 (class 2604 OID 18879)
-- Name: order_currency id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_currency ALTER COLUMN id SET DEFAULT nextval('order_currency_id_seq'::regclass);


--
-- TOC entry 2721 (class 2604 OID 19342)
-- Name: order_graph_12h id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_12h ALTER COLUMN id SET DEFAULT nextval('order_graph_12h_id_seq'::regclass);


--
-- TOC entry 2717 (class 2604 OID 19250)
-- Name: order_graph_15m id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_15m ALTER COLUMN id SET DEFAULT nextval('order_graph_15m_id_seq'::regclass);


--
-- TOC entry 2722 (class 2604 OID 19365)
-- Name: order_graph_1d id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1d ALTER COLUMN id SET DEFAULT nextval('order_graph_1d_id_seq'::regclass);


--
-- TOC entry 2719 (class 2604 OID 19296)
-- Name: order_graph_1h id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1h ALTER COLUMN id SET DEFAULT nextval('order_graph_1h_id_seq'::regclass);


--
-- TOC entry 2716 (class 2604 OID 19227)
-- Name: order_graph_1m id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1m ALTER COLUMN id SET DEFAULT nextval('order_graph_1m_id_seq'::regclass);


--
-- TOC entry 2718 (class 2604 OID 19273)
-- Name: order_graph_30m id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_30m ALTER COLUMN id SET DEFAULT nextval('order_graph_30m_id_seq'::regclass);


--
-- TOC entry 2714 (class 2604 OID 19044)
-- Name: order_graph_5m id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_5m ALTER COLUMN id SET DEFAULT nextval('order_graph_id_seq'::regclass);


--
-- TOC entry 2720 (class 2604 OID 19319)
-- Name: order_graph_6h id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_6h ALTER COLUMN id SET DEFAULT nextval('order_graph_6h_id_seq'::regclass);


--
-- TOC entry 2723 (class 2604 OID 19388)
-- Name: order_graph_7d id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_7d ALTER COLUMN id SET DEFAULT nextval('order_graph_7d_id_seq'::regclass);


--
-- TOC entry 2708 (class 2604 OID 18815)
-- Name: order_sell id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_sell ALTER COLUMN id SET DEFAULT nextval('order_sell_id_seq'::regclass);


--
-- TOC entry 2713 (class 2604 OID 19012)
-- Name: order_sell_broker id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_sell_broker ALTER COLUMN id SET DEFAULT nextval('order_sell_broker_id_seq'::regclass);


--
-- TOC entry 2715 (class 2604 OID 19067)
-- Name: order_wallet id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_wallet ALTER COLUMN id SET DEFAULT nextval('order_wallet_id_seq'::regclass);


--
-- TOC entry 2677 (class 2604 OID 17019)
-- Name: password id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY password ALTER COLUMN id SET DEFAULT nextval('user_password_id_seq'::regclass);


--
-- TOC entry 2695 (class 2604 OID 18122)
-- Name: payee id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee ALTER COLUMN id SET DEFAULT nextval('payee_id_seq'::regclass);


--
-- TOC entry 2696 (class 2604 OID 18135)
-- Name: payee_crypto id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee_crypto ALTER COLUMN id SET DEFAULT nextval('payee_crypto_id_seq'::regclass);


--
-- TOC entry 2699 (class 2604 OID 18220)
-- Name: payee_fiat id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee_fiat ALTER COLUMN id SET DEFAULT nextval('payee_fiat_id_seq'::regclass);


--
-- TOC entry 2694 (class 2604 OID 18095)
-- Name: payee_master id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee_master ALTER COLUMN id SET DEFAULT nextval('payee_master_id_seq'::regclass);


--
-- TOC entry 2710 (class 2604 OID 18906)
-- Name: site_bank id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY site_bank ALTER COLUMN id SET DEFAULT nextval('site_bank_id_seq'::regclass);


--
-- TOC entry 2672 (class 2604 OID 16550)
-- Name: transaction id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transaction ALTER COLUMN id SET DEFAULT nextval('transaction_id_seq'::regclass);


--
-- TOC entry 2679 (class 2604 OID 17227)
-- Name: transfer id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer ALTER COLUMN id SET DEFAULT nextval('transfer_id_seq'::regclass);


--
-- TOC entry 2701 (class 2604 OID 18274)
-- Name: transfer_bank id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_bank ALTER COLUMN id SET DEFAULT nextval('transfer_withdraw_id_seq'::regclass);


--
-- TOC entry 2693 (class 2604 OID 18063)
-- Name: transfer_master id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_master ALTER COLUMN id SET DEFAULT nextval('transfer_master_id_seq'::regclass);


--
-- TOC entry 2698 (class 2604 OID 18197)
-- Name: transfer_payee id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_payee ALTER COLUMN id SET DEFAULT nextval('transfer_payee_id_seq'::regclass);


--
-- TOC entry 2697 (class 2604 OID 18174)
-- Name: transfer_wallet id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_wallet ALTER COLUMN id SET DEFAULT nextval('transfer_wallet_id_seq'::regclass);


--
-- TOC entry 2673 (class 2604 OID 16551)
-- Name: user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- TOC entry 2685 (class 2604 OID 17567)
-- Name: user_country id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY user_country ALTER COLUMN id SET DEFAULT nextval('user_country_id_seq'::regclass);


--
-- TOC entry 2678 (class 2604 OID 17030)
-- Name: user_setting id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY user_setting ALTER COLUMN id SET DEFAULT nextval('user_setting_id_seq'::regclass);


--
-- TOC entry 2675 (class 2604 OID 16552)
-- Name: wallet id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet ALTER COLUMN id SET DEFAULT nextval('wallet_id_seq'::regclass);


--
-- TOC entry 2668 (class 2604 OID 16544)
-- Name: wallet_crypto id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_crypto ALTER COLUMN id SET DEFAULT nextval('wallet_crypto_id_seq'::regclass);


--
-- TOC entry 2691 (class 2604 OID 17906)
-- Name: wallet_fiat id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_fiat ALTER COLUMN id SET DEFAULT nextval('wallet_fiat_id_seq'::regclass);


--
-- TOC entry 2692 (class 2604 OID 17990)
-- Name: wallet_master id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_master ALTER COLUMN id SET DEFAULT nextval('wallet_master_id_seq'::regclass);


--
-- TOC entry 2684 (class 2604 OID 17445)
-- Name: wallet_passphrase id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_passphrase ALTER COLUMN id SET DEFAULT nextval('wallet_passphrase_id_seq'::regclass);


--
-- TOC entry 3270 (class 0 OID 0)
-- Dependencies: 265
-- Name: AdminEmail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('"AdminEmail_id_seq"', 5, true);


--
-- TOC entry 3086 (class 0 OID 16440)
-- Dependencies: 187
-- Data for Name: admin; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY admin (id, name, adminname, admin_group_id, created_at, updated_at, deleted_at, edit_name_times) FROM stdin;
3	Suman Admin 1	a2	1	2017-11-28 14:56:24.186432-05	2017-11-28 14:56:24.186432-05	\N	0
4	Suman Admin 1	a3	1	2017-11-28 16:46:04.764382-05	2017-11-28 16:46:04.764382-05	\N	0
5	Suman Admin 1	a4	1	2017-11-28 16:50:10.596018-05	2017-11-28 16:50:10.596018-05	\N	0
2	S Admin 1	a1	1	2017-11-28 14:53:44.681291-05	2017-11-28 19:20:47.953714-05	\N	2
9	Subhendu Das	a5	3	2017-12-02 21:13:09.585009-05	2017-12-02 21:13:09.585009-05	\N	0
\.


--
-- TOC entry 3087 (class 0 OID 16446)
-- Dependencies: 188
-- Data for Name: admin_access; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY admin_access (id, admin_id, access_key, allow, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3271 (class 0 OID 0)
-- Dependencies: 189
-- Name: admin_access_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('admin_access_id_seq', 1, false);


--
-- TOC entry 3169 (class 0 OID 18424)
-- Dependencies: 270
-- Data for Name: admin_country; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY admin_country (id, admin_id, country_id, created_at, updated_at, deleted_at) FROM stdin;
1	2	1	2017-11-28 16:50:10.596018-05	2017-11-28 16:50:10.596018-05	\N
\.


--
-- TOC entry 3272 (class 0 OID 0)
-- Dependencies: 269
-- Name: admin_country_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('admin_country_id_seq', 1, true);


--
-- TOC entry 3165 (class 0 OID 18396)
-- Dependencies: 266
-- Data for Name: admin_email; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY admin_email (id, admin_id, email, "primary", created_at, updated_at, deleted_at) FROM stdin;
2	3	s@s.com	t	2017-11-28 14:56:24.193978-05	2017-11-28 14:56:24.193979-05	\N
3	4	s@s.com	t	2017-11-28 16:46:04.773715-05	2017-11-28 16:46:04.773716-05	\N
4	5	s@s.com	t	2017-11-28 16:50:10.604276-05	2017-11-28 16:50:10.604276-05	\N
5	9	sumaninster7@gmail.com	t	2017-12-02 21:13:09.615138-05	2017-12-02 21:13:09.615139-05	\N
\.


--
-- TOC entry 3167 (class 0 OID 18406)
-- Dependencies: 268
-- Data for Name: admin_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY admin_group (id, key, name, created_at, updated_at, deleted_at, description) FROM stdin;
1	SUPER_ADMIN	Super Admin	2017-08-18 18:13:05.72073-04	2017-08-18 18:13:05.72073-04	\N	\N
3	CUSTOMER_SUPPORT	Customer Support	2017-08-18 18:13:05.72073-04	2017-08-18 18:13:05.72073-04	\N	\N
2	ADMIN	Admin	2017-08-18 18:13:05.72073-04	2017-08-18 18:13:05.72073-04	\N	\N
\.


--
-- TOC entry 3273 (class 0 OID 0)
-- Dependencies: 267
-- Name: admin_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('admin_group_id_seq', 3, true);


--
-- TOC entry 3274 (class 0 OID 0)
-- Dependencies: 190
-- Name: admin_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('admin_id_seq', 9, true);


--
-- TOC entry 3161 (class 0 OID 18369)
-- Dependencies: 262
-- Data for Name: admin_password; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY admin_password (id, admin_id, password, created_at, updated_at, deleted_at) FROM stdin;
1	2	3132333c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2	2017-11-28 14:53:44.690325-05	2017-11-28 14:53:44.690326-05	\N
2	3	3132333c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2	2017-11-28 14:56:24.193266-05	2017-11-28 14:56:24.193267-05	\N
3	4	3132333c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2	2017-11-28 16:46:04.772815-05	2017-11-28 16:46:04.772815-05	\N
4	5	3132333c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2	2017-11-28 16:50:10.603336-05	2017-11-28 16:50:10.603337-05	\N
5	9	3132333c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2	2017-12-02 21:13:09.591539-05	2017-12-02 21:13:09.59154-05	\N
\.


--
-- TOC entry 3275 (class 0 OID 0)
-- Dependencies: 261
-- Name: admin_password_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('admin_password_id_seq', 5, true);


--
-- TOC entry 3163 (class 0 OID 18380)
-- Dependencies: 264
-- Data for Name: admin_token; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY admin_token (id, admin_id, token, expiration_time, created_at, updated_at, deleted_at) FROM stdin;
1	5	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTE5OTIyMTAsImlhdCI6MTUxMTkwNTgxMH0.D5kdp2SoANTAOOnRyrtpsx4SclHNQQe6oGFWHYfqKvfEF1Z6VE8omGTag0Sdi-96BqOTJcYddKQNkNiGXLyfiVn12zZtw36yPp58F7wtcZ7IiMgRhhgPD6j-ySgLoBoTD0SFhVk_hl7emOoLPsEaS0wKRX9xpYJskSvZImizGe5iwCEYDz78BdoP1yIMNtQ8RvTbN8plCqKkxxgcFbSZQrq1a9hd07NT7bLYOs6JgevbenfqV_vu1MCAK2O99zn6LAkm93r5UKxEgZS5BuPQsIT9bZWFx-odeiBx0ofVZo7wtDPPvNw-dk_WQ7MEy6c_mLVFy7Oso5KO_tn29gipNg	2017-11-29 16:50:10.605005-05	2017-11-28 16:50:10.610544-05	2017-11-28 16:50:10.611835-05	\N
3	9	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTI0NDM1MzIsImlhdCI6MTUxMjM1NzEzMn0.JL2e-ixkTPsGUc2ksQGHf5EjV0pgaQFVKQ8RJ0-xvDdtmczgDvO3Yh_AqDzqffQgltFv0vHkGYN6rO5Zc5usWQeyRP4C3jfrLt-ciRHV94Ti19CCMm0gHXIdAULna5Ox_-G9z88MY2mRuhL-u7TWYuGlMRmzw_NZo_n16x5GIGOkE2T1WBfTxCwC_TNs8zvvunnihqLXkZ7TKznph1uhiPDkI92bipDg1Lcq7xX9zfWfxjrRbOMMxBATYJ_5-Pt9zKEsrimUD7rk8Tb8_YaKCynkFcQ-Kn4jMcW8-FG6CrCHk_0ifh-CV4Gi-exKInjOzd3xQrKTlND8qWGeaa5JRQ	2017-12-04 22:12:12.846762-05	2017-12-02 21:14:06.007904-05	2017-12-03 22:12:12.85097-05	\N
2	2	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjQ1ODA4NDYsImlhdCI6MTUyNDQ5NDQ0Nn0.jbVAmR-kpZVxBUmO_JRcrs5OwfKHcOtQ33n6LhccjRWmoWVXZ1JI5s7QdMfRwxN9jBh1UagW6yBQ7OCqWYu8Wbr3dmW8TnIzzOfF57Ku_wviF9AGaDv01m9_xOKwCqLxRlEfMjRlIe3UfZkFxpoPVCV8Vd6Ft6tznrzoCCHpSLOVwTp7yFLVANNSMxsgjDezEImN6r1_PkBRNcVDanrvAa8Vz4Oul8c7rlF6oJQ73WhfP1Uu-oqac03sAJfkJAnoWMf_6TdETW_D5OeM60rlW4v_NvOyZAPcrnQ3XGKQ6cOc035nq49L1NZsLMFj6ih1-SzdsfwbYaJhVsDlDERobg	2018-04-24 10:40:46.50667-04	2017-11-28 17:46:51.0923-05	2018-04-23 10:40:46.535049-04	\N
\.


--
-- TOC entry 3276 (class 0 OID 0)
-- Dependencies: 263
-- Name: admin_token_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('admin_token_id_seq', 3, true);


--
-- TOC entry 3123 (class 0 OID 17431)
-- Dependencies: 224
-- Data for Name: api_key; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY api_key (id, app_name, description, token, created_at, updated_at, deleted_at, domain_name, ip_address) FROM stdin;
\.


--
-- TOC entry 3277 (class 0 OID 0)
-- Dependencies: 223
-- Name: api_key_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('api_key_id_seq', 1, false);


--
-- TOC entry 3090 (class 0 OID 16453)
-- Dependencies: 191
-- Data for Name: auth_token; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY auth_token (id, user_id, token, expiration_time, created_at, updated_at, deleted_at) FROM stdin;
21	19	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDM3MTkwMjAsImlhdCI6MTUwMzYzMjYyMH0.lx-kmYHl7Um666xbK0eGSbNImtapeuqPLR2061qyo8TCFF5sTK-m0oB7Ckoj0CxFmvaVR8LCK46k06hpBJJxXjkpg17Py_NLaWifExgGJCoStCmWRsS8cxmfv8ZqwW85Rspsq2KtNOtvGg7pNfqd0wS8nke8FpJd9yErVXT6H4GFbyeOFoUByNNxqP2jtz09wQl_NHlwjTB5Tx4c3NKWJuUVr1gekCkehHtz9M1DzOL7Qa79BAeWA4iNaSepOWYrp9dQYmP-DnNHnDFpSaZU9nW89rTjr7NnCgGJSZuh7gu5XsBWGhfoyqXgIhkdTqHRMGS1yJnUreFjMedaiWEOeg	2017-08-25 23:43:40.111451-04	2017-08-24 23:43:40.114608-04	2017-08-24 23:43:40.11569-04	\N
24	21	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDQ3ODgzNjksImlhdCI6MTUwNDcwMTk2OX0.V-rBwrrrlhvNpR3huJHQ0YAgl3aPSBICFyv6MM70nnJgpTuVM9I1J94Op5MHswt3lQmC7u0MXqwy30hd3ZJ8TS0xrFVzclIzuvGVluNPrYkEo2q45lJrexUOO0EUj0BfMbQTYwwyrYNy3CaZVo7LiJJ1yBXxlMMMWalfO9fmTlS2__tnjJMjnDWBUvAGdPtWPZm0VhJBijcteUv2JlXOWLuAZY1oXgN0vvwjmTbnyKram1Wg1jvaawxYKDxcUnhECPiCySFJg3ESIedVKEudSHLHYrlMJRusP1gzPu6PKnQxz92Z3umHLmx7M0e_i3hloI63G09Xt4d2QmRnYGYS8A	2017-09-07 08:46:09.0296-04	2017-09-06 08:45:59.188641-04	2017-09-06 08:46:09.033564-04	\N
20	11	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDM3MDk4NDMsImlhdCI6MTUwMzYyMzQ0M30.RIGv7mGCKcXaeJifiWOzwhdJD01p5UXwwXMkqXpVhcSMIjm9ZSnev3WYReZo-qMFbKGLnwWj1zYyeLJ3SV-tcWo6IKBgil_8c8NMXNwY1ClBjMkNFqTVVtPM0KxagF4-ohK69eHKoCzSQBlMTa9zw4Yh_JX0E-mOMynPDfk_jpMMzBm_FOWrinaG70mRUnSQqJry1Qayr8R3NMd4aHu7bJp9BLpVjLnAFYW3NBN_AT1YzMOoo0WbZSFg5M513psGLyz9w-mte0LWg-UGHiS2Hjllxv0xiIK-T4VodLZCbJVFxA8koszswPvbclkYWcoLa4OFwUm_VHqZ38wfPiLY-Q	2017-08-25 21:10:43.484661-04	2017-08-23 11:36:45.848771-04	2017-08-24 21:10:43.489458-04	\N
26	17	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDc1MTgwOTYsImlhdCI6MTUwNzQzMTY5Nn0.oz6zYyjFd5MH7RXCzNCHcxxM46Ax0cz-rNK6fJlYeTxIdyneFptnVNse-rpsere6Shd2-YL5OFHsuyXyDjcCYIVtTeOjylX8vJd7rZDWAvAREmnFr8OOXAcU61N95xdM8hFe4tyA6OYyRP8a4Guu-LP_ZQxYsryfoQ3lpkV_7AEOtFuqhOnNgig212MNegmEs5K5kNRG-eBVd2LPm7ssuOdyxWP3IeJKeIEac7-KcX9G-9qdqqHKyIWOyqsKEzXR_4-pr3JOJ7sbYrYiLsc_48sonc0xCCnZozp1LqWXrjmu1ZMPNlyQhtEYKZm3EbVVz1hfGobPNDXN23mBW7nG_Q	2017-10-08 23:01:36.595664-04	2017-10-07 20:22:12.694831-04	2017-10-07 23:01:36.599838-04	\N
19	9	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mjk4OTQ4MDcsImlhdCI6MTUyOTgwODQwN30.gSeaEpDvrJvr3UTg2xi5XFGDdKX_qnJYHStML7FSga0z_ZqjN3QsrJn-7ww0Bv8GLDqXFWvL5W8esLpDEVGJrwe9un7VofOt6_K4AqyU4iTsh5KyW9Y3VOExbWLexJwUj9Aom1Mw9b0Iv7keSPUL0-d6Zyl6uwoCkOEQv2vZ-MnaWFulEs4PuhzY7XIKrA8lbdJQB-4-uYXqEBIAT8qsLQtMvWspk6edttekLvs84qdQnYSYZKqGEtOr_bn2l5W5RBSUCEWzzgv96F5PWxNhxP1m72nFyNmMZqc_gZ9ytwJ0fmty08326Qv9gQEE3feC22Wp20VCrJ6HX4PxVLS9TA	2018-06-24 22:46:47.479488-04	2017-08-21 19:42:22.638501-04	2018-06-23 22:46:47.484015-04	\N
25	18	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDc1MDcyNzQsImlhdCI6MTUwNzQyMDg3NH0.h_sA8-A21FIk-9ErYhdfI7J6aBQkXtjK1hAVDcTWZWi_0kdfu8Wa9Z94DDsbTHhdlIZvRqxRHjqeUtjnZD8gRL8H_aj3LZ_njKmf9yl0P3uaCeVqKVTvVHLFi7LRDyQ1W-iTvvOW__0jhb_bQPgoU0VvvDs-nyIDIWhMnzrLy1wLDkrB3Q0iFS4mXu9Z8nPiUWEJGqZtgijRtxIxGpy9EtfQQZSbIxvH7okbA6z16qzn3xnWXRyadvwN3GkbU5LUdmcdyTo_l0HtPWqYcnPw_JaZdfIi7P9qZT0ghJvwD7w_Lzc-LxiQixrkDICS1YgI0GFPoHQ1Gm_AqRglHueecQ	2017-10-08 20:01:14.898418-04	2017-10-07 20:01:14.901755-04	2017-10-07 20:01:14.90925-04	\N
23	10	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDIzMTMxODMsImlhdCI6MTU0MjIyNjc4M30.twZMa9ezX_mgGQp0YyqL2gWH47IwvOyOOQ_qZh5ZnhIH6RZO0DkmoTbSls0VNp38LH7vCbFfmVGw4RxnNhSg3yVfyyMWsvbvNhg5ptHWj7Zc8dL0b_WrYcOdVq2bmW2wu4_dCmXbeLxvR_RoyD_Zh6-Du3nhivF9jPstyouFrgzEPFeT4YkHwj-jThmMKLk_qDbkLaA-9GBCGGzuhrfInn0iOdVS5dKSN0LL_KA9c9rgjpAXUcLbzJYYESwKysF5O-ueS8Ztst8s88imPsOsN6htvndqpMpSyi7LhS2z2P0Bo8_DfELXkUO_7nvji5mC5G01iEexA3SsvTITLWL27Q	2018-11-15 15:19:43.08006-05	2017-08-30 18:21:41.313194-04	2018-11-14 15:19:43.087231-05	\N
\.


--
-- TOC entry 3278 (class 0 OID 0)
-- Dependencies: 192
-- Name: auth_token_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('auth_token_id_seq', 26, true);


--
-- TOC entry 3092 (class 0 OID 16474)
-- Dependencies: 193
-- Data for Name: country; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY country (id, name, iso_code, dial_code, code, created_at, updated_at, deleted_at) FROM stdin;
3	India	IN	91	IN	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
2	United States	US	1	US	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
5	Japan	JP	81	JP	2017-11-29 20:01:01.387145-05	2017-11-29 20:01:01.387146-05	\N
4	Australia	AU	61	AU	2017-11-29 19:48:02.886552-05	2017-11-29 20:01:45.015936-05	\N
1	Global	GLOBAL	0	GLOBAL	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
6	Canada	CAN	1	CA	2017-11-29 20:01:01.387145-05	2017-11-29 20:01:01.387145-05	\N
\.


--
-- TOC entry 3279 (class 0 OID 0)
-- Dependencies: 194
-- Name: country_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('country_id_seq', 5, true);


--
-- TOC entry 3119 (class 0 OID 17341)
-- Dependencies: 220
-- Data for Name: currency; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY currency (id, code, name, description, type, country_id, created_at, updated_at, deleted_at) FROM stdin;
5	INR	Indian Rupees	Fiat Currency	FIAT	3	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N
6	USD	US Dollar	Fiat Currency	FIAT	2	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N
7	JPY	Yen	Fiat Currency	FIAT	5	2017-11-29 20:18:47.379844-05	2017-11-29 20:21:03.947467-05	\N
3	CAD	Canadian Dollar	Fiat Currency	FIAT	6	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N
1	BTC	Bitcoin	Digital Currency	CRYPTO	1	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N
2	ETH	Ethereum	Digital Currency	CRYPTO	1	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N
4	BCH	Bitcoin Cash	Digital Currency	CRYPTO	1	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N
8	LTC	Litecoin	Digital Currency	CRYPTO	1	2017-11-29 20:18:47.379844-05	2017-11-29 20:18:47.379844-05	\N
\.


--
-- TOC entry 3280 (class 0 OID 0)
-- Dependencies: 219
-- Name: currency_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('currency_id_seq', 7, true);


--
-- TOC entry 3121 (class 0 OID 17389)
-- Dependencies: 222
-- Data for Name: currency_pair; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY currency_pair (id, from_currency_id, to_currency_id, created_at, updated_at, deleted_at, pair_type, supported) FROM stdin;
1	1	2	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N	CRYPTO_TO_CRYPTO	YES
2	2	1	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N	CRYPTO_TO_CRYPTO	YES
\.


--
-- TOC entry 3281 (class 0 OID 0)
-- Dependencies: 221
-- Name: currency_pair_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('currency_pair_id_seq', 1, false);


--
-- TOC entry 3131 (class 0 OID 17607)
-- Dependencies: 232
-- Data for Name: data; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY data (id, user_id, country_id, nickname, "primary", active, created_at, updated_at, deleted_at, field_type) FROM stdin;
11	10	6	nun	t	t	2017-10-07 18:08:14.318323-04	2017-10-07 18:08:14.318325-04	\N	ADDRESS
7	10	6	fake address	f	f	2017-10-02 08:46:49.374386-04	2017-10-02 08:47:06.135237-04	\N	BANK
8	10	6	fake address	t	t	2017-10-02 08:47:06.142193-04	2017-10-02 08:47:06.142194-04	\N	BANK
9	10	6	xyz	f	t	2017-10-02 11:12:23.970599-04	2017-10-07 18:08:14.289577-04	\N	ADDRESS
1	10	6		f	t	2017-09-26 12:05:19.967862-04	2017-10-07 18:08:14.317496-04	\N	ADDRESS
2	10	6		f	t	2017-09-26 17:16:10.093746-04	2017-10-07 18:08:14.316656-04	\N	ADDRESS
10	10	6	file test	t	t	2017-10-02 11:12:23.970599-04	2017-10-02 11:12:23.970599-04	\N	FILE
6	10	6	good address	f	t	2017-10-02 08:45:21.551261-04	2017-10-07 18:08:14.297628-04	\N	ADDRESS
4	10	6	home address	f	t	2017-10-02 08:36:59.659358-04	2017-10-07 18:08:14.31006-04	\N	ADDRESS
3	10	6		f	t	2017-09-26 17:17:42.669484-04	2017-10-02 08:47:06.141201-04	\N	BANK
5	10	6	fake address	f	f	2017-10-02 08:45:03.736454-04	2017-10-07 18:08:14.303622-04	\N	ADDRESS
\.


--
-- TOC entry 3137 (class 0 OID 17664)
-- Dependencies: 238
-- Data for Name: data_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY data_category (id, data_id, field_id, field_category_id, created_at, updated_at, deleted_at) FROM stdin;
1	2	9	9	2017-09-26 17:16:10.121078-04	2017-09-26 17:16:10.121079-04	\N
2	4	9	19	2017-10-02 08:36:59.68447-04	2017-10-02 08:36:59.684471-04	\N
3	5	9	15	2017-10-02 08:45:03.746933-04	2017-10-02 08:45:03.746934-04	\N
4	6	9	15	2017-10-02 08:45:21.555029-04	2017-10-02 08:45:21.55503-04	\N
5	9	9	11	2017-10-02 11:12:23.979378-04	2017-10-02 11:12:23.979379-04	\N
6	11	9	20	2017-10-07 18:08:14.325216-04	2017-10-07 18:08:14.325217-04	\N
7	1	9	9	2017-10-07 18:08:14.325216-04	2017-10-07 18:08:14.325216-04	\N
\.


--
-- TOC entry 3282 (class 0 OID 0)
-- Dependencies: 237
-- Name: data_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_category_id_seq', 7, true);


--
-- TOC entry 3135 (class 0 OID 17646)
-- Dependencies: 236
-- Data for Name: data_file; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY data_file (id, data_id, field_id, name, extension, created_at, updated_at, deleted_at) FROM stdin;
4	10	1	4	png	2017-10-03 19:09:59.964346-04	2017-10-03 19:09:59.974425-04	\N
5	10	1	5	png	2017-10-03 19:12:30.974409-04	2017-10-03 19:12:30.980884-04	\N
6	10	1	6	png	2017-10-03 19:14:36.899216-04	2017-10-03 19:14:36.90205-04	\N
3	10	1	3	png	2017-10-03 18:58:00.347094-04	2017-10-03 18:58:00.359287-04	\N
7	10	1	7	png	2017-10-03 19:25:15.055038-04	2017-10-03 19:25:15.062547-04	\N
8	10	1	8	png	2017-10-03 19:27:56.173677-04	2017-10-03 19:27:56.18135-04	\N
9	10	1	9	png	2017-11-10 19:18:55.172919-05	2017-11-10 19:18:55.246782-05	\N
12	10	1	12	png	2017-11-20 18:02:23.644147-05	2017-11-20 18:02:23.656055-05	\N
11	10	1	11	png	2017-11-20 17:59:42.596463-05	2017-11-20 17:59:42.596467-05	\N
10	10	1	10	png	2017-11-20 17:59:01.103263-05	2017-11-20 17:59:01.103265-05	\N
13	10	1	13	png	2017-11-20 18:19:51.99056-05	2017-11-20 18:19:51.996513-05	\N
\.


--
-- TOC entry 3283 (class 0 OID 0)
-- Dependencies: 235
-- Name: data_file_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_file_id_seq', 13, true);


--
-- TOC entry 3284 (class 0 OID 0)
-- Dependencies: 231
-- Name: data_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_id_seq', 11, true);


--
-- TOC entry 3133 (class 0 OID 17615)
-- Dependencies: 234
-- Data for Name: data_text; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY data_text (id, data_id, field_id, text, created_at, updated_at, deleted_at) FROM stdin;
1	1	6	1614	2017-09-26 12:05:19.976345-04	2017-09-26 12:05:19.976347-04	\N
2	1	7	95 Thorncliffe Park Drive, East York	2017-09-26 12:05:19.978496-04	2017-09-26 12:05:19.978498-04	\N
3	1	8	M4H 1L7	2017-09-26 12:05:19.979768-04	2017-09-26 12:05:19.979769-04	\N
4	1	9	Ontario	2017-09-26 12:05:19.980564-04	2017-09-26 12:05:19.980565-04	\N
5	2	6	1614	2017-09-26 17:16:10.101266-04	2017-09-26 17:16:10.101267-04	\N
6	2	7	95 Thorncliffe Park Drive, East York	2017-09-26 17:16:10.108943-04	2017-09-26 17:16:10.108944-04	\N
7	2	8	M4H 1L7	2017-09-26 17:16:10.115539-04	2017-09-26 17:16:10.11554-04	\N
8	3	5	suman@gmail.com	2017-09-26 17:17:42.675688-04	2017-09-26 17:17:42.675689-04	\N
9	4	6	1614	2017-10-02 08:36:59.663917-04	2017-10-02 08:36:59.663918-04	\N
10	4	7	95 Thorncliffe Park Drive, East York	2017-10-02 08:36:59.671392-04	2017-10-02 08:36:59.671393-04	\N
11	4	8	M4H 1L7	2017-10-02 08:36:59.67785-04	2017-10-02 08:36:59.677851-04	\N
12	5	6	1614	2017-10-02 08:45:03.742851-04	2017-10-02 08:45:03.742852-04	\N
13	5	7	95 Thorncliffe Park Drive, East York	2017-10-02 08:45:03.744683-04	2017-10-02 08:45:03.744685-04	\N
14	5	8	M4H 1L7	2017-10-02 08:45:03.745901-04	2017-10-02 08:45:03.745902-04	\N
15	6	6	1614	2017-10-02 08:45:21.552214-04	2017-10-02 08:45:21.552216-04	\N
16	6	7	95 Thorncliffe Park Drive, East York	2017-10-02 08:45:21.553371-04	2017-10-02 08:45:21.553372-04	\N
17	6	8	M4H 1L7	2017-10-02 08:45:21.554242-04	2017-10-02 08:45:21.554243-04	\N
18	7	5	d@gmail	2017-10-02 08:46:49.375101-04	2017-10-02 08:46:49.375102-04	\N
19	8	5	d@gmail.com	2017-10-02 08:47:06.143053-04	2017-10-02 08:47:06.143055-04	\N
20	9	6	1614	2017-10-02 11:12:23.977246-04	2017-10-02 11:12:23.977247-04	\N
21	9	7	95 Thorncliffe Park Drive, East York	2017-10-02 11:12:23.978025-04	2017-10-02 11:12:23.978025-04	\N
22	9	8	M4H 1L7	2017-10-02 11:12:23.978736-04	2017-10-02 11:12:23.978736-04	\N
23	11	6	1614	2017-10-07 18:08:14.32101-04	2017-10-07 18:08:14.321011-04	\N
24	11	7	95 Thorncliffe Park Drive, East York	2017-10-07 18:08:14.32301-04	2017-10-07 18:08:14.323012-04	\N
25	11	8	M4H 1L7	2017-10-07 18:08:14.324087-04	2017-10-07 18:08:14.324088-04	\N
\.


--
-- TOC entry 3285 (class 0 OID 0)
-- Dependencies: 233
-- Name: data_text_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_text_id_seq', 25, true);


--
-- TOC entry 3108 (class 0 OID 16939)
-- Dependencies: 209
-- Data for Name: email; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY email (id, user_id, email, "primary", created_at, updated_at, deleted_at) FROM stdin;
9	9	s1@gmail.com	YES	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
10	10	s2@gmail.com	YES	2017-07-23 20:46:25.047371-04	2017-07-23 20:46:25.047371-04	\N
11	11	sukanya@gmail.com	YES	2017-08-10 17:30:45.223409-04	2017-08-10 17:30:45.223409-04	\N
13	13	abc@gmail.com	YES	2017-08-11 20:55:18.674922-04	2017-08-11 20:55:18.674922-04	\N
14	14	suman@gmail.com	YES	2017-08-11 22:53:32.265227-04	2017-08-11 22:53:32.265227-04	\N
15	15	chanda@gmail.com	YES	2017-08-11 23:34:12.943804-04	2017-08-11 23:34:12.943804-04	\N
16	16	sangita@gmail.com	YES	2017-08-11 23:53:20.985644-04	2017-08-11 23:53:20.985644-04	\N
17	17	suman.inster@gmail.com	YES	2017-08-17 21:00:11.868824-04	2017-08-17 21:00:11.868824-04	\N
18	18	s@gmail.com	YES	2017-08-18 18:13:05.749073-04	2017-08-18 18:13:05.749073-04	\N
19	19	s2@gmail.com	YES	2017-08-24 23:43:40.081956-04	2017-08-24 23:43:40.081956-04	\N
21	21	s@gmail.com	YES	2017-09-06 08:45:59.183185-04	2017-09-06 08:45:59.183185-04	\N
\.


--
-- TOC entry 3286 (class 0 OID 0)
-- Dependencies: 210
-- Name: email_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('email_id_seq', 21, true);


--
-- TOC entry 3096 (class 0 OID 16487)
-- Dependencies: 197
-- Data for Name: field; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY field (id, country_id, name, created_at, updated_at, deleted_at, description, "order", field_type, key, data_type, has_category, has_input_text, has_file) FROM stdin;
4	6	Photo	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please provide your Photo	1	FILE	\N	FILE	\N	\N	t
1	6	Photo ID	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please provide your Photo ID	2	FILE	\N	FILE	t	\N	t
6	6	Apartment / Unit	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please provide your apartment/unit number	1	ADDRESS	\N	TEXT	\N	t	\N
2	6	Selfie With Photo ID	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please provide a selfie with your Photo ID	3	FILE	\N	FILE	\N	\N	t
10	2	Apartment / Unit	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please provide your apartment/unit number	1	ADDRESS	\N	TEXT	\N	t	\N
11	2	Street / Address	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please provide your street name and number	2	ADDRESS	\N	TEXT	\N	t	\N
12	2	Zip Code	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please provide your Zip Code	3	ADDRESS	\N	TEXT	\N	t	\N
13	2	Province	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please select your province	4	ADDRESS	\N	TEXT	\N	t	\N
9	6	Province	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please select your province	4	ADDRESS	\N	CATEGORY	t	\N	\N
5	6	Email ID	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please provide your email id to recevie payment	1	BANK	PRIMARY	TEXT	\N	t	\N
7	6	Street / Address	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please provide your street name and number	2	ADDRESS	\N	TEXT	\N	t	\N
3	6	Utility Bill	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please provide your utility bill or bank statement	4	FILE	\N	FILE	t	\N	t
8	6	Zip Code	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N	Please provide your Zip Code	3	ADDRESS	\N	TEXT	\N	t	\N
\.


--
-- TOC entry 3129 (class 0 OID 17590)
-- Dependencies: 230
-- Data for Name: field_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY field_category (id, field_id, name, created_at, updated_at, deleted_at) FROM stdin;
1	1	Passport	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
2	1	Driver's Licence	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
3	1	National ID Card	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
4	1	Permanent Residance Card	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
5	3	Bank Statement	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
6	3	Phone Bill	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
7	3	Electricity Bill	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
8	3	Other	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
9	9	Ontario	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
16	9	New Brunswick	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
14	9	Saskatchewan	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
10	9	Quebec	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
12	9	Alberta	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
13	9	Manitoba	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
17	9	Newfoundland and Labrador	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
15	9	Nova Scotia	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
11	9	British Columbia	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
18	9	Prince Edward Island	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
19	9	Northwest Territories	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
20	9	Nunavut	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
21	9	Yukon	2017-07-23 20:44:38.249528-04	2017-07-23 20:44:38.249528-04	\N
\.


--
-- TOC entry 3287 (class 0 OID 0)
-- Dependencies: 229
-- Name: field_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('field_category_id_seq', 21, true);


--
-- TOC entry 3288 (class 0 OID 0)
-- Dependencies: 198
-- Name: field_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('field_id_seq', 13, true);


--
-- TOC entry 3157 (class 0 OID 18249)
-- Dependencies: 258
-- Data for Name: field_subcategory; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY field_subcategory (id, field_id, sub_field_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3289 (class 0 OID 0)
-- Dependencies: 257
-- Name: field_subcategory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('field_subcategory_id_seq', 1, false);


--
-- TOC entry 3117 (class 0 OID 17237)
-- Dependencies: 218
-- Data for Name: global_setting; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY global_setting (id, country_id, key, value, admin_id, created_at, updated_at, deleted_at) FROM stdin;
1	0	BTC_MIN_BUY	0.01	1	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N
\.


--
-- TOC entry 3290 (class 0 OID 0)
-- Dependencies: 217
-- Name: global_setting_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('global_setting_id_seq', 1, true);


--
-- TOC entry 3098 (class 0 OID 16508)
-- Dependencies: 199
-- Data for Name: mobile; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY mobile (id, user_id, number, "primary", created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3291 (class 0 OID 0)
-- Dependencies: 200
-- Name: mobile_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('mobile_id_seq', 1, false);


--
-- TOC entry 3100 (class 0 OID 16513)
-- Dependencies: 201
-- Data for Name: order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY "order" (id, seller_user_id, currency_id, rate_currency_id, amount, rate, total_amount, created_at, updated_at, deleted_at, buyer_user_id, seller_currency_wallet_id, seller_rate_currency_wallet_id, buyer_currency_wallet_id, buyer_rate_currency_wallet_id) FROM stdin;
283	10	1	8	1	10	10	2018-01-25 23:42:39.345552-05	2018-01-25 23:42:39.345554-05	\N	10	47	56	47	56
284	10	1	8	1	10	10	2018-01-25 23:47:04.91543-05	2018-01-25 23:47:04.915432-05	\N	10	47	56	47	56
285	10	1	8	1	5	5	2018-01-25 23:54:11.573264-05	2018-01-25 23:54:11.573267-05	\N	10	47	56	47	56
286	10	1	8	1	10	10	2018-01-26 14:38:22.022283-05	2018-01-26 14:38:22.022285-05	\N	10	46	56	46	56
287	10	1	8	3	4	12	2018-01-26 14:38:58.146027-05	2018-01-26 14:38:58.146029-05	\N	10	46	56	46	56
288	10	1	8	2	7	14	2018-01-26 14:39:46.378922-05	2018-01-26 14:39:46.378925-05	\N	10	46	56	46	56
289	10	1	8	1	5	5	2018-01-26 22:27:26.276294-05	2018-01-26 22:27:26.276297-05	\N	10	46	56	46	56
290	10	1	8	3	7	21	2018-01-26 22:27:29.300558-05	2018-01-26 22:27:29.300561-05	\N	10	46	56	46	56
291	10	1	8	1	7	7	2018-01-26 22:28:41.413291-05	2018-01-26 22:28:41.413294-05	\N	10	46	56	46	56
292	10	1	8	1	10	10	2018-01-26 22:28:44.437179-05	2018-01-26 22:28:44.437182-05	\N	10	46	56	46	56
293	10	1	8	2	5	10	2018-01-26 22:28:56.466979-05	2018-01-26 22:28:56.466981-05	\N	10	46	56	46	56
294	10	1	8	1	10	10	2018-01-26 22:29:02.494742-05	2018-01-26 22:29:02.494745-05	\N	10	46	56	46	56
\.


--
-- TOC entry 3179 (class 0 OID 18967)
-- Dependencies: 280
-- Data for Name: order_broker; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_broker (id, order_id, seller_broker_id, buyer_broker_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3292 (class 0 OID 0)
-- Dependencies: 279
-- Name: order_broker_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_broker_id_seq', 1, false);


--
-- TOC entry 3171 (class 0 OID 18804)
-- Dependencies: 272
-- Data for Name: order_buy; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_buy (id, user_id, currency_id, rate_currency_id, amount, rate, total_amount, created_at, currency_wallet_id, rate_currency_wallet_id, lock) FROM stdin;
1995	10	1	8	1	5	5	2018-01-26 22:28:29.485006-05	46	56	f
1991	10	1	8	4	5	20	2018-01-26 14:38:39.106855-05	46	56	f
1999	10	2	5	2	200	400	2018-07-25 23:07:50.204542-04	51	54	f
2000	10	1	2	1	10	10	2018-07-25 23:19:02.228189-04	50	52	f
\.


--
-- TOC entry 3181 (class 0 OID 18991)
-- Dependencies: 282
-- Data for Name: order_buy_broker; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_buy_broker (id, order_buy_id, broker_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3293 (class 0 OID 0)
-- Dependencies: 281
-- Name: order_buy_broker_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_buy_broker_id_seq', 1, false);


--
-- TOC entry 3294 (class 0 OID 0)
-- Dependencies: 271
-- Name: order_buy_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_buy_id_seq', 2000, true);


--
-- TOC entry 3175 (class 0 OID 18876)
-- Dependencies: 276
-- Data for Name: order_currency; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_currency (id, user_id, currency_id, rate_currency_id, created_at, updated_at, deleted_at) FROM stdin;
5	9	1	2	2017-12-26 18:25:09.168764-05	2018-04-23 11:00:36.350148-04	\N
3	10	1	2	2017-12-24 21:35:58.566283-05	2018-11-14 15:25:49.82656-05	\N
\.


--
-- TOC entry 3295 (class 0 OID 0)
-- Dependencies: 275
-- Name: order_currency_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_currency_id_seq', 5, true);


--
-- TOC entry 3199 (class 0 OID 19339)
-- Dependencies: 300
-- Data for Name: order_graph_12h; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_graph_12h (id, last_order_id, currency_id, rate_currency_id, open, high, low, close, volume, split, dividend, absolute_change, percent_change, date) FROM stdin;
\.


--
-- TOC entry 3296 (class 0 OID 0)
-- Dependencies: 299
-- Name: order_graph_12h_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_graph_12h_id_seq', 1, false);


--
-- TOC entry 3191 (class 0 OID 19247)
-- Dependencies: 292
-- Data for Name: order_graph_15m; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_graph_15m (id, last_order_id, currency_id, rate_currency_id, open, high, low, close, volume, split, dividend, absolute_change, percent_change, date) FROM stdin;
\.


--
-- TOC entry 3297 (class 0 OID 0)
-- Dependencies: 291
-- Name: order_graph_15m_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_graph_15m_id_seq', 4, true);


--
-- TOC entry 3201 (class 0 OID 19362)
-- Dependencies: 302
-- Data for Name: order_graph_1d; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_graph_1d (id, last_order_id, currency_id, rate_currency_id, open, high, low, close, volume, split, dividend, absolutue_change, percent_change, date) FROM stdin;
\.


--
-- TOC entry 3298 (class 0 OID 0)
-- Dependencies: 301
-- Name: order_graph_1d_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_graph_1d_id_seq', 1, false);


--
-- TOC entry 3195 (class 0 OID 19293)
-- Dependencies: 296
-- Data for Name: order_graph_1h; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_graph_1h (id, last_order_id, currency_id, rate_currency_id, open, high, low, close, volume, split, dividend, absolute_change, percent_change, date) FROM stdin;
\.


--
-- TOC entry 3299 (class 0 OID 0)
-- Dependencies: 295
-- Name: order_graph_1h_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_graph_1h_id_seq', 1, false);


--
-- TOC entry 3189 (class 0 OID 19224)
-- Dependencies: 290
-- Data for Name: order_graph_1m; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_graph_1m (id, last_order_id, currency_id, rate_currency_id, open, high, low, close, volume, split, dividend, absolute_change, percent_change, date) FROM stdin;
\.


--
-- TOC entry 3300 (class 0 OID 0)
-- Dependencies: 289
-- Name: order_graph_1m_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_graph_1m_id_seq', 4, true);


--
-- TOC entry 3193 (class 0 OID 19270)
-- Dependencies: 294
-- Data for Name: order_graph_30m; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_graph_30m (id, last_order_id, currency_id, rate_currency_id, open, high, low, close, volume, split, dividend, absolutue_change, percent_change, date) FROM stdin;
\.


--
-- TOC entry 3301 (class 0 OID 0)
-- Dependencies: 293
-- Name: order_graph_30m_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_graph_30m_id_seq', 1, false);


--
-- TOC entry 3185 (class 0 OID 19041)
-- Dependencies: 286
-- Data for Name: order_graph_5m; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_graph_5m (id, open, high, low, close, volume, split, dividend, absolute_change, percent_change, date, last_order_id, currency_id, rate_currency_id) FROM stdin;
1	5	10	5	10	10	0	0	5	100	2018-01-26 22:32:26.262154-05	294	1	8
2	10	10	10	10	0	0	0	0	0	2018-01-26 23:49:37.695087-05	294	1	8
3	10	10	10	10	0	0	0	0	0	2018-01-27 01:31:14.61717-05	294	1	8
4	10	10	10	10	0	0	0	0	0	2018-01-27 01:33:59.742793-05	294	1	8
5	10	10	10	10	0	0	0	0	0	2018-01-27 01:37:51.690963-05	294	1	8
\.


--
-- TOC entry 3197 (class 0 OID 19316)
-- Dependencies: 298
-- Data for Name: order_graph_6h; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_graph_6h (id, last_order_id, currency_id, rate_currency_id, open, high, low, close, volume, split, dividend, absolute_change, percent_change, date) FROM stdin;
\.


--
-- TOC entry 3302 (class 0 OID 0)
-- Dependencies: 297
-- Name: order_graph_6h_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_graph_6h_id_seq', 1, false);


--
-- TOC entry 3203 (class 0 OID 19385)
-- Dependencies: 304
-- Data for Name: order_graph_7d; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_graph_7d (id, last_order_id, currency_id, rate_currency_id, open, high, low, close, volume, split, dividend, absolute_change, percent_change, date) FROM stdin;
\.


--
-- TOC entry 3303 (class 0 OID 0)
-- Dependencies: 303
-- Name: order_graph_7d_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_graph_7d_id_seq', 1, false);


--
-- TOC entry 3304 (class 0 OID 0)
-- Dependencies: 285
-- Name: order_graph_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_graph_id_seq', 5, true);


--
-- TOC entry 3305 (class 0 OID 0)
-- Dependencies: 202
-- Name: order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_id_seq', 294, true);


--
-- TOC entry 3173 (class 0 OID 18812)
-- Dependencies: 274
-- Data for Name: order_sell; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_sell (id, user_id, currency_id, rate_currency_id, amount, rate, total_amount, created_at, currency_wallet_id, rate_currency_wallet_id, lock) FROM stdin;
49417	10	1	8	1	10	10	2018-01-26 22:28:25.223335-05	46	56	f
49413	10	1	8	1	10	10	2018-01-26 14:37:43.579372-05	46	56	f
49420	10	1	2	1	11	11	2018-07-25 23:19:15.520181-04	50	52	f
\.


--
-- TOC entry 3183 (class 0 OID 19009)
-- Dependencies: 284
-- Data for Name: order_sell_broker; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_sell_broker (id, order_sell_id, broker_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3306 (class 0 OID 0)
-- Dependencies: 283
-- Name: order_sell_broker_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_sell_broker_id_seq', 1, false);


--
-- TOC entry 3307 (class 0 OID 0)
-- Dependencies: 273
-- Name: order_sell_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_sell_id_seq', 49420, true);


--
-- TOC entry 3187 (class 0 OID 19064)
-- Dependencies: 288
-- Data for Name: order_wallet; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY order_wallet (id, user_id, currency_id, rate_currency_id, currency_wallet_id, rate_currency_wallet_id, created_at, updated_at, deleted_at) FROM stdin;
5	10	8	5	56	54	2018-01-24 21:08:48.396949-05	2018-01-24 21:08:54.378164-05	\N
4	10	1	8	47	56	2018-01-23 20:19:48.791516-05	2018-01-28 17:22:45.113237-05	\N
8	10	2	5	51	54	2018-07-06 11:20:53.979267-04	2018-07-06 11:20:54.005453-04	\N
6	10	1	2	50	52	2018-06-23 22:43:51.817965-04	2018-11-14 15:25:56.627389-05	\N
\.


--
-- TOC entry 3308 (class 0 OID 0)
-- Dependencies: 287
-- Name: order_wallet_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('order_wallet_id_seq', 8, true);


--
-- TOC entry 3111 (class 0 OID 17016)
-- Dependencies: 212
-- Data for Name: password; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY password (id, user_id, password, created_at, updated_at, deleted_at) FROM stdin;
7	9	3132333c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2	2017-07-23 20:44:38.248167-04	2017-07-23 20:44:38.248167-04	\N
8	10	3132333c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2	2017-07-23 20:46:25.046633-04	2017-07-23 20:46:25.046633-04	\N
9	11	73756b616e79617ed0cae5a6ac875a550273dab3f199e562d2b4c89b7849dfd7e967ddc61222bf138cf2da8ceb54faca7bdce6a3a544e0118b02bfc8afb31eec4e356db72c0db8	2017-08-10 17:30:45.220878-04	2017-08-10 17:30:45.220878-04	\N
11	13	616263ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f	2017-08-11 20:55:18.673278-04	2017-08-11 20:55:18.673278-04	\N
12	14	73756d616e6e340f9b8e7b05c61f76a4bbaae0ee14a2bc42c7aa74547c35d24e1d06f381af21e0de374f64d0f1daa4536e7155923a92748feb9f67d17307c579e61946d10e	2017-08-11 22:53:32.263862-04	2017-08-11 22:53:32.263862-04	\N
13	15	6368616e6461c6e8914793597b5d3a08aa5777a47b4f9af9fd3e0f3a05793ead46bca17bf201d00cf89a40dfb80cd2080a839d8ca3bac86ae23be1612df16ccf189dd51d86ff	2017-08-11 23:34:12.942345-04	2017-08-11 23:34:12.942345-04	\N
14	16	73616e67697461f8937749789192a4bb03987beef16089e49cd657e1f9e740f69759122a3fc4dd3af95278d5a58784ec8aeb8eed58a878a4d5b6fc72d96b4383d5c924d647b312	2017-08-11 23:53:20.98442-04	2017-08-11 23:53:20.98442-04	\N
15	17	3132333c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2	2017-08-17 21:00:11.864738-04	2017-08-17 21:00:11.864738-04	\N
16	18	3132333c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2	2017-08-18 18:13:05.737397-04	2017-08-18 18:13:05.737397-04	\N
17	19	3132333c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2	2017-08-24 23:43:40.077094-04	2017-08-24 23:43:40.077095-04	\N
19	21	3132333c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2	2017-09-06 08:45:59.181576-04	2017-09-06 08:45:59.181576-04	\N
\.


--
-- TOC entry 3147 (class 0 OID 18119)
-- Dependencies: 248
-- Data for Name: payee; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY payee (id, payee_master_id, name, nickname, email, created_at, updated_at, deleted_at) FROM stdin;
1	1	Rishimitra Nath	Rishi Kolkata	georgenath@gmail.com	2017-10-19 08:35:46.060941-04	2017-10-19 08:35:46.060942-04	\N
2	1	Rishimitra Nath	Rishi Kolkata 2	georgenath@gmail.com	2017-10-19 08:36:58.348876-04	2017-10-19 08:36:58.348877-04	\N
3	2	Rishimitra Nath	Rishi Ether	georgenath@gmail.com	2017-10-19 08:45:31.103004-04	2017-10-19 08:45:31.103005-04	\N
4	1	Rishimitra Nath	georgenath	georgenath@gmail.com	2017-10-19 11:34:57.116125-04	2017-10-19 11:34:57.116127-04	\N
\.


--
-- TOC entry 3149 (class 0 OID 18132)
-- Dependencies: 250
-- Data for Name: payee_crypto; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY payee_crypto (id, payee_id, address, created_at, updated_at, deleted_at) FROM stdin;
1	1	mhnr7DsacptCc26EkJDvmrZCjH4XZw2wEa	2017-10-19 08:35:46.065762-04	2017-10-19 08:35:46.065763-04	\N
2	2	mtoeMSgNnLUkvkji32WmvcjbfjtAZxwkBr	2017-10-19 08:36:58.358733-04	2017-10-19 08:36:58.358735-04	\N
3	3	0x734d523ee8d057d35586c885325e7a5fbb55c603	2017-10-19 08:45:31.104092-04	2017-10-19 08:45:31.104094-04	\N
4	4	mtoeMSgNnLUkvkji32WmvcjbfjtAZxwkBr	2017-10-19 11:34:57.122682-04	2017-10-19 11:34:57.122683-04	\N
\.


--
-- TOC entry 3309 (class 0 OID 0)
-- Dependencies: 249
-- Name: payee_crypto_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('payee_crypto_id_seq', 4, true);


--
-- TOC entry 3155 (class 0 OID 18217)
-- Dependencies: 256
-- Data for Name: payee_fiat; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY payee_fiat (id, payee_id, data_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3310 (class 0 OID 0)
-- Dependencies: 255
-- Name: payee_fiat_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('payee_fiat_id_seq', 1, false);


--
-- TOC entry 3311 (class 0 OID 0)
-- Dependencies: 247
-- Name: payee_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('payee_id_seq', 4, true);


--
-- TOC entry 3145 (class 0 OID 18092)
-- Dependencies: 246
-- Data for Name: payee_master; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY payee_master (id, user_id, currency_id, currency_type, created_at, updated_at, deleted_at) FROM stdin;
1	10	1	CRYPTO	2017-10-19 08:35:46.033742-04	2017-10-19 08:35:46.033744-04	\N
2	10	2	CRYPTO	2017-10-19 08:45:31.095899-04	2017-10-19 08:45:31.095901-04	\N
\.


--
-- TOC entry 3312 (class 0 OID 0)
-- Dependencies: 245
-- Name: payee_master_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('payee_master_id_seq', 2, true);


--
-- TOC entry 3177 (class 0 OID 18903)
-- Dependencies: 278
-- Data for Name: site_bank; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY site_bank (id, updated_admin_id, country_id, currency_id, beneficiary_name, beneficiary_address, beneficiary_city, beneficiary_province, beneficiary_postal_code, beneficiary_account_number, reference_text, bank_name, bank_address, bank_province, bank_postal_code, bank_phone_number, bank_institution_number, bank_transit_number, bank_swift_code, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3313 (class 0 OID 0)
-- Dependencies: 277
-- Name: site_bank_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('site_bank_id_seq', 1, false);


--
-- TOC entry 3102 (class 0 OID 16518)
-- Dependencies: 203
-- Data for Name: transaction; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY transaction (id, order_id, from_user, to_user, amount, created_at, updated_at, deleted_at, currency_id) FROM stdin;
\.


--
-- TOC entry 3314 (class 0 OID 0)
-- Dependencies: 204
-- Name: transaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('transaction_id_seq', 116, true);


--
-- TOC entry 3115 (class 0 OID 17224)
-- Dependencies: 216
-- Data for Name: transfer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY transfer (id, transfer_master_id, amount, transfer_type, from_user_id, to_user_id, processed_at, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3159 (class 0 OID 18271)
-- Dependencies: 260
-- Data for Name: transfer_bank; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY transfer_bank (id, transfer_id, from_wallet_id, to_data_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3315 (class 0 OID 0)
-- Dependencies: 215
-- Name: transfer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('transfer_id_seq', 6, true);


--
-- TOC entry 3143 (class 0 OID 18060)
-- Dependencies: 244
-- Data for Name: transfer_master; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY transfer_master (id, user_id, currency_id, currency_type, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3316 (class 0 OID 0)
-- Dependencies: 243
-- Name: transfer_master_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('transfer_master_id_seq', 2, true);


--
-- TOC entry 3153 (class 0 OID 18194)
-- Dependencies: 254
-- Data for Name: transfer_payee; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY transfer_payee (id, transfer_id, from_wallet_id, to_payee_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3317 (class 0 OID 0)
-- Dependencies: 253
-- Name: transfer_payee_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('transfer_payee_id_seq', 1, false);


--
-- TOC entry 3151 (class 0 OID 18171)
-- Dependencies: 252
-- Data for Name: transfer_wallet; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY transfer_wallet (id, transfer_id, from_wallet_id, to_wallet_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3318 (class 0 OID 0)
-- Dependencies: 251
-- Name: transfer_wallet_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('transfer_wallet_id_seq', 3, true);


--
-- TOC entry 3319 (class 0 OID 0)
-- Dependencies: 259
-- Name: transfer_withdraw_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('transfer_withdraw_id_seq', 1, false);


--
-- TOC entry 3104 (class 0 OID 16523)
-- Dependencies: 205
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY "user" (id, name, username, created_at, updated_at, deleted_at, edit_name_times) FROM stdin;
11	Sukanya Parashar	suku123	2017-08-10 17:30:45.191241-04	2017-08-10 17:30:45.191241-04	\N	0
13	abc	abc	2017-08-11 20:55:18.654488-04	2017-08-11 20:55:18.654488-04	\N	0
14	Suman Das	suman123	2017-08-11 22:53:32.256227-04	2017-08-11 22:53:32.256227-04	\N	0
15	Chanda Parashar	chanda123	2017-08-11 23:34:12.934144-04	2017-08-11 23:34:12.934144-04	\N	0
16	Sangita Goswami	sangita123	2017-08-11 23:53:20.977279-04	2017-08-11 23:53:20.977279-04	\N	0
17	Subhendu Das	s3	2017-08-17 21:00:11.833833-04	2017-08-17 21:00:11.833833-04	\N	0
18	Suman Das	s4	2017-08-18 18:13:05.72073-04	2017-08-18 18:13:05.72073-04	\N	0
19	sukanya	s123	2017-08-24 23:43:40.031227-04	2017-08-24 23:43:40.031227-04	\N	0
21	Suman Das	s	2017-09-06 08:45:59.173789-04	2017-09-06 08:45:59.173789-04	\N	0
9	Suman Das	s1	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N	0
10	Subhendu Das	s2	2017-07-23 20:46:25.039969-04	2017-07-23 20:46:25.039969-04	\N	0
\.


--
-- TOC entry 3127 (class 0 OID 17564)
-- Dependencies: 228
-- Data for Name: user_country; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY user_country (id, user_id, country_id, created_at, updated_at, deleted_at, eligible) FROM stdin;
13	9	2	2017-07-23 20:44:38.239669-04	2017-11-03 20:57:24.693702-04	\N	YES
14	10	3	2017-09-05 22:48:37.754537-04	2017-11-19 16:47:08.09809-05	\N	YES
15	10	2	2017-09-20 21:03:08.753053-04	2017-11-19 16:47:22.571789-05	\N	YES
7	16	6	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N	YES
5	14	6	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N	YES
3	11	6	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N	YES
6	15	6	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N	YES
10	19	6	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N	NO
4	13	6	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N	YES
9	18	6	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N	NO
8	17	6	2017-07-23 20:44:38.239669-04	2017-07-23 20:44:38.239669-04	\N	YES
1	9	6	2017-07-23 20:44:38.239669-04	2017-12-10 11:42:59.780755-05	2017-12-10 11:42:59.780747-05	NO
12	9	3	2017-07-23 20:44:38.239669-04	2017-12-10 11:43:00.995596-05	2017-12-10 11:43:00.995587-05	NO
19	9	5	2017-12-10 11:43:44.727613-05	2017-12-10 11:43:44.727615-05	\N	NO
17	10	4	2017-12-01 12:15:36.333062-05	2018-08-04 09:55:09.222773-04	2018-08-04 09:55:09.222763-04	NO
18	10	5	2017-12-01 12:15:46.052073-05	2018-08-04 09:55:10.182543-04	2018-08-04 09:55:10.182535-04	NO
2	10	6	2017-07-23 20:44:38.239669-04	2018-08-04 09:55:12.596035-04	\N	NO
16	17	3	2017-10-07 22:58:27.326395-04	2017-10-07 22:58:27.326396-04	\N	NO
\.


--
-- TOC entry 3320 (class 0 OID 0)
-- Dependencies: 227
-- Name: user_country_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_country_id_seq', 19, true);


--
-- TOC entry 3321 (class 0 OID 0)
-- Dependencies: 206
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_id_seq', 21, true);


--
-- TOC entry 3322 (class 0 OID 0)
-- Dependencies: 211
-- Name: user_password_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_password_id_seq', 19, true);


--
-- TOC entry 3113 (class 0 OID 17027)
-- Dependencies: 214
-- Data for Name: user_setting; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY user_setting (id, user_id, key, value, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3323 (class 0 OID 0)
-- Dependencies: 213
-- Name: user_setting_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_setting_id_seq', 1, false);


--
-- TOC entry 3106 (class 0 OID 16531)
-- Dependencies: 207
-- Data for Name: wallet; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY wallet (id, wallet_master_id, amount, amount_locked, nickname, "primary", created_at, updated_at, deleted_at) FROM stdin;
42	3	5600	0	No Name	f	2017-10-13 22:16:12.655889-04	2017-10-13 23:09:53.986046-04	\N
43	3	5600	0	new	f	2017-10-13 22:51:10.741477-04	2017-10-13 23:09:53.998772-04	\N
44	3	5600	0	la la	f	2017-10-13 23:07:51.78645-04	2017-10-13 23:09:53.979522-04	\N
41	3	5600	0	No Name	f	2017-10-13 22:14:13.023463-04	2017-10-13 23:09:53.992551-04	\N
48	3	5600	0	Rishi Kolkata	f	2017-10-19 08:32:56.756789-04	2017-10-19 08:32:56.75679-04	\N
51	4	5600	0	new eth acc	f	2017-11-03 13:39:56.029014-04	2017-11-03 13:40:11.453142-04	\N
45	3	5600	0	pa pa	f	2017-10-13 23:09:54.005197-04	2017-10-13 23:30:35.697993-04	\N
49	3	5600	0	checking	f	2017-10-19 10:05:04.6924-04	2017-10-19 10:19:02.495808-04	\N
54	5	5600	400	inr bank	t	2017-11-03 13:40:11.459391-04	2018-07-25 23:07:50.198257-04	\N
50	3	7000	1	new checking	t	2017-10-19 10:19:02.501468-04	2018-07-25 23:19:15.512513-04	\N
52	4	5600	10	new eth acc 1	t	2017-11-03 13:40:11.459391-04	2018-07-25 23:23:46.847705-04	\N
47	3	3000	0	sa sa	f	2017-10-13 23:31:56.698868-04	2018-01-25 23:54:11.572691-05	\N
55	6	7000	0	s1	t	2017-11-03 13:40:11.459391-04	2018-01-08 19:43:18.054172-05	\N
57	8	7000	0	s1	t	2017-11-03 13:40:11.459391-04	2018-01-25 00:45:59.205282-05	\N
56	7	7000	25	s2	t	2017-11-03 13:40:11.459391-04	2018-01-26 22:29:02.492468-05	\N
46	3	5600	2	ka ka	f	2017-10-13 23:30:35.704432-04	2018-01-26 22:29:02.494068-05	\N
\.


--
-- TOC entry 3094 (class 0 OID 16479)
-- Dependencies: 195
-- Data for Name: wallet_crypto; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY wallet_crypto (id, wallet_id, address, created_at, updated_at, deleted_at) FROM stdin;
18	41	n3yHWJpfttmYTyJLhe5yLGk7Biper8w5Kr	2017-10-13 22:14:13.237164-04	2017-10-13 22:14:13.237165-04	\N
19	42	mzJPZQ8VCqCR3bDHcoX1GDACa1PT4ifKXn	2017-10-13 22:16:12.868637-04	2017-10-13 22:16:12.868639-04	\N
20	43	n3MEmNnSeDDKg6jYQci174XPtGNdXvmzMt	2017-10-13 22:51:10.950462-04	2017-10-13 22:51:10.950463-04	\N
21	44	mqAarwmSLijEcna1dn4NH4qaLu8dSe6fS8	2017-10-13 23:07:51.97733-04	2017-10-13 23:07:51.977331-04	\N
22	45	myGM7d93sNF4idCcD5pnAQKZ3WgUgP6GNw	2017-10-13 23:09:54.175701-04	2017-10-13 23:09:54.175702-04	\N
23	46	n17njC4BuFDonz3KWT5BAiq4bR15MRFLAo	2017-10-13 23:30:35.889833-04	2017-10-13 23:30:35.889834-04	\N
24	47	myx3RheDuUWZZHvFjT8Cf9iZT2jmBUqZAV	2017-10-13 23:31:56.888423-04	2017-10-13 23:31:56.888424-04	\N
25	48	mqqou569hiC2a6BQGmDukNbgth7x5BXcWx	2017-10-19 08:32:56.994409-04	2017-10-19 08:32:56.99441-04	\N
26	49	mqM3fvTH6aEnhFT8JJnawj6y5akfZDJDDt	2017-10-19 10:05:04.901416-04	2017-10-19 10:05:04.901417-04	\N
27	50	mnkZZpZBuYHXtBMUyLeRiHE6BEjSw8ApzF	2017-10-19 10:19:02.674569-04	2017-10-19 10:19:02.67457-04	\N
28	51	0xf61b595563be571b084454e6231e401e7c6f07bb	2017-11-03 13:39:56.839174-04	2017-11-03 13:39:56.839175-04	\N
29	52	0x764350bdcc5a8a5985c7d260089de7c4469c07b3	2017-11-03 13:40:12.381597-04	2017-11-03 13:40:12.381597-04	\N
\.


--
-- TOC entry 3324 (class 0 OID 0)
-- Dependencies: 196
-- Name: wallet_crypto_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('wallet_crypto_id_seq', 29, true);


--
-- TOC entry 3139 (class 0 OID 17903)
-- Dependencies: 240
-- Data for Name: wallet_fiat; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY wallet_fiat (id, wallet_id, data_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- TOC entry 3325 (class 0 OID 0)
-- Dependencies: 239
-- Name: wallet_fiat_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('wallet_fiat_id_seq', 1, false);


--
-- TOC entry 3326 (class 0 OID 0)
-- Dependencies: 208
-- Name: wallet_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('wallet_id_seq', 55, true);


--
-- TOC entry 3141 (class 0 OID 17987)
-- Dependencies: 242
-- Data for Name: wallet_master; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY wallet_master (id, user_id, currency_id, currency_type, created_at, updated_at, deleted_at) FROM stdin;
3	10	1	CRYPTO	2017-10-13 22:14:13.016458-04	2017-10-13 22:14:13.016459-04	\N
4	10	2	CRYPTO	2017-11-03 13:39:55.995574-04	2017-11-03 13:39:55.995575-04	\N
5	10	5	FIAT	2017-11-03 13:39:55.995574-04	2017-11-03 13:39:55.995574-04	\N
6	9	8	CRYPTO	2017-11-03 13:39:55.995574-04	2017-11-03 13:39:55.995574-04	\N
7	10	8	CRYPTO	2017-11-03 13:39:55.995574-04	2017-11-03 13:39:55.995574-04	\N
8	9	1	CRYPTO	2017-11-03 13:39:55.995574-04	2017-11-03 13:39:55.995574-04	\N
\.


--
-- TOC entry 3327 (class 0 OID 0)
-- Dependencies: 241
-- Name: wallet_master_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('wallet_master_id_seq', 5, true);


--
-- TOC entry 3125 (class 0 OID 17442)
-- Dependencies: 226
-- Data for Name: wallet_passphrase; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY wallet_passphrase (id, wallet_id, passphrase, created_at, updated_at, deleted_at) FROM stdin;
12	51	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUxMDk3MzA3OTYsImlhdCI6MTUwOTczMDc5Nn0.H32Go7vpvm36M9w8_dRrzWjWyNgZyi0amkC-Nw4h4FmZ310ZPvD2MxkXT2CO7E7kgUH9G4xuwTk9KpMGmpJTzttu6VYXCTUNWyk54xygLX6UYioMGVyZsUX62tpwvnoMR5Hn8G1PaWFDTQZthoM75ImgdcD3V4teeatMBEpI1dSgOSTfHaQSrXcLoJlHdBpYGe7Fqp4bUkfo5Hp6lQAbsDeV8od9b1d0CezkBNAxgeqgA0zrDh9EQOizXw88g5Cx7HsANjpxWppqUm3UtCbxFpQXlJaxRp3P08At4VHX4q0KKeCqUeddJI2gz9jjqdjSnHWOgC_5Sh2L-9OfV80KNw	2017-11-03 13:39:56.035261-04	2017-11-03 13:39:56.035262-04	\N
13	52	eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUxMDk3MzA4MTEsImlhdCI6MTUwOTczMDgxMX0.dGH1n8Hzd4IlcGV3l0K1ZwfxsyQi2m3M-6LYik2YAgTZxuQJU0kqPJT9DyxG2-PHlKylHdhKa6gG9KAN_fknsi3Iw-5Z-2PQrWYaio60VSez9VnghfhfhSjURD1pqIbuwyK7IEnsK1J-D6tLIT13oY7I5a1Lsqf0aY2SdeQuXu1x110pfoA0V5v72R9TLuZcoaQ3eHJw2VrP18fcJMl0fmu0T2ZP4pkwFCGpHAM4pZyUAPPkmlLH7lfpPnqMAVNvfIaU1Xgu1H1KTVwNrb67qhMlhiOJBcSTWJXXx5RI8_8gZyB0AETQ7YvHZboKoiFUsziNt1QZzJf9KUeeO6xBtQ	2017-11-03 13:40:11.470667-04	2017-11-03 13:40:11.470669-04	\N
\.


--
-- TOC entry 3328 (class 0 OID 0)
-- Dependencies: 225
-- Name: wallet_passphrase_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('wallet_passphrase_id_seq', 13, true);


--
-- TOC entry 2805 (class 2606 OID 18401)
-- Name: admin_email AdminEmail_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_email
    ADD CONSTRAINT "AdminEmail_pkey" PRIMARY KEY (id);


--
-- TOC entry 2727 (class 2606 OID 16558)
-- Name: admin_access admin_access_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_access
    ADD CONSTRAINT admin_access_pkey PRIMARY KEY (id);


--
-- TOC entry 2811 (class 2606 OID 18429)
-- Name: admin_country admin_country_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_country
    ADD CONSTRAINT admin_country_pkey PRIMARY KEY (id);


--
-- TOC entry 2807 (class 2606 OID 18421)
-- Name: admin_group admin_group_key_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_group
    ADD CONSTRAINT admin_group_key_key UNIQUE (key);


--
-- TOC entry 2809 (class 2606 OID 18414)
-- Name: admin_group admin_group_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_group
    ADD CONSTRAINT admin_group_pkey PRIMARY KEY (id);


--
-- TOC entry 2801 (class 2606 OID 18377)
-- Name: admin_password admin_password_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_password
    ADD CONSTRAINT admin_password_pkey PRIMARY KEY (id);


--
-- TOC entry 2725 (class 2606 OID 16560)
-- Name: admin admin_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin
    ADD CONSTRAINT admin_pkey PRIMARY KEY (id);


--
-- TOC entry 2803 (class 2606 OID 18388)
-- Name: admin_token admin_token_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_token
    ADD CONSTRAINT admin_token_pkey PRIMARY KEY (id);


--
-- TOC entry 2763 (class 2606 OID 17439)
-- Name: api_key api_key_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY api_key
    ADD CONSTRAINT api_key_pkey PRIMARY KEY (id);


--
-- TOC entry 2729 (class 2606 OID 16562)
-- Name: auth_token auth_token_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY auth_token
    ADD CONSTRAINT auth_token_pkey PRIMARY KEY (id);


--
-- TOC entry 2731 (class 2606 OID 16568)
-- Name: country country_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY country
    ADD CONSTRAINT country_pkey PRIMARY KEY (id);


--
-- TOC entry 2733 (class 2606 OID 16570)
-- Name: wallet_crypto crypto_address_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_crypto
    ADD CONSTRAINT crypto_address_pkey PRIMARY KEY (id);


--
-- TOC entry 2761 (class 2606 OID 17397)
-- Name: currency_pair currency_pair_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY currency_pair
    ADD CONSTRAINT currency_pair_pkey PRIMARY KEY (id);


--
-- TOC entry 2759 (class 2606 OID 17346)
-- Name: currency currency_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY currency
    ADD CONSTRAINT currency_pkey PRIMARY KEY (id);


--
-- TOC entry 2777 (class 2606 OID 17669)
-- Name: data_category data_category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_category
    ADD CONSTRAINT data_category_pkey PRIMARY KEY (id);


--
-- TOC entry 2775 (class 2606 OID 17651)
-- Name: data_file data_file_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_file
    ADD CONSTRAINT data_file_pkey PRIMARY KEY (id);


--
-- TOC entry 2771 (class 2606 OID 17612)
-- Name: data data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data
    ADD CONSTRAINT data_pkey PRIMARY KEY (id);


--
-- TOC entry 2773 (class 2606 OID 17623)
-- Name: data_text data_text_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_text
    ADD CONSTRAINT data_text_pkey PRIMARY KEY (id);


--
-- TOC entry 2749 (class 2606 OID 16943)
-- Name: email email_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY email
    ADD CONSTRAINT email_pkey PRIMARY KEY (id);


--
-- TOC entry 2769 (class 2606 OID 17595)
-- Name: field_category field_category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY field_category
    ADD CONSTRAINT field_category_pkey PRIMARY KEY (id);


--
-- TOC entry 2735 (class 2606 OID 16572)
-- Name: field field_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY field
    ADD CONSTRAINT field_pkey PRIMARY KEY (id);


--
-- TOC entry 2797 (class 2606 OID 18254)
-- Name: field_subcategory field_subcategory_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY field_subcategory
    ADD CONSTRAINT field_subcategory_pkey PRIMARY KEY (id);


--
-- TOC entry 2757 (class 2606 OID 17245)
-- Name: global_setting global_setting_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY global_setting
    ADD CONSTRAINT global_setting_pkey PRIMARY KEY (id);


--
-- TOC entry 2737 (class 2606 OID 16578)
-- Name: mobile mobile_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY mobile
    ADD CONSTRAINT mobile_pkey PRIMARY KEY (id);


--
-- TOC entry 2823 (class 2606 OID 18972)
-- Name: order_broker order_broker_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_broker
    ADD CONSTRAINT order_broker_pkey PRIMARY KEY (id);


--
-- TOC entry 2825 (class 2606 OID 18996)
-- Name: order_buy_broker order_buy_broker_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_buy_broker
    ADD CONSTRAINT order_buy_broker_pkey PRIMARY KEY (id);


--
-- TOC entry 2813 (class 2606 OID 18809)
-- Name: order_buy order_buy_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_buy
    ADD CONSTRAINT order_buy_pkey PRIMARY KEY (id);


--
-- TOC entry 2817 (class 2606 OID 18881)
-- Name: order_currency order_currency_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_currency
    ADD CONSTRAINT order_currency_pkey PRIMARY KEY (id);


--
-- TOC entry 2819 (class 2606 OID 18898)
-- Name: order_currency order_currency_user_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_currency
    ADD CONSTRAINT order_currency_user_id_key UNIQUE (user_id);


--
-- TOC entry 2843 (class 2606 OID 19344)
-- Name: order_graph_12h order_graph_12h_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_12h
    ADD CONSTRAINT order_graph_12h_pkey PRIMARY KEY (id);


--
-- TOC entry 2835 (class 2606 OID 19252)
-- Name: order_graph_15m order_graph_15m_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_15m
    ADD CONSTRAINT order_graph_15m_pkey PRIMARY KEY (id);


--
-- TOC entry 2845 (class 2606 OID 19367)
-- Name: order_graph_1d order_graph_1d_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1d
    ADD CONSTRAINT order_graph_1d_pkey PRIMARY KEY (id);


--
-- TOC entry 2839 (class 2606 OID 19298)
-- Name: order_graph_1h order_graph_1h_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1h
    ADD CONSTRAINT order_graph_1h_pkey PRIMARY KEY (id);


--
-- TOC entry 2833 (class 2606 OID 19229)
-- Name: order_graph_1m order_graph_1m_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1m
    ADD CONSTRAINT order_graph_1m_pkey PRIMARY KEY (id);


--
-- TOC entry 2837 (class 2606 OID 19275)
-- Name: order_graph_30m order_graph_30m_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_30m
    ADD CONSTRAINT order_graph_30m_pkey PRIMARY KEY (id);


--
-- TOC entry 2841 (class 2606 OID 19321)
-- Name: order_graph_6h order_graph_6h_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_6h
    ADD CONSTRAINT order_graph_6h_pkey PRIMARY KEY (id);


--
-- TOC entry 2847 (class 2606 OID 19390)
-- Name: order_graph_7d order_graph_7d_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_7d
    ADD CONSTRAINT order_graph_7d_pkey PRIMARY KEY (id);


--
-- TOC entry 2829 (class 2606 OID 19049)
-- Name: order_graph_5m order_graph_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_5m
    ADD CONSTRAINT order_graph_pkey PRIMARY KEY (id);


--
-- TOC entry 2739 (class 2606 OID 16580)
-- Name: order order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "order"
    ADD CONSTRAINT order_pkey PRIMARY KEY (id);


--
-- TOC entry 2827 (class 2606 OID 19014)
-- Name: order_sell_broker order_sell_broker_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_sell_broker
    ADD CONSTRAINT order_sell_broker_pkey PRIMARY KEY (id);


--
-- TOC entry 2815 (class 2606 OID 18817)
-- Name: order_sell order_sell_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_sell
    ADD CONSTRAINT order_sell_pkey PRIMARY KEY (id);


--
-- TOC entry 2831 (class 2606 OID 19069)
-- Name: order_wallet order_wallet_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_wallet
    ADD CONSTRAINT order_wallet_pkey PRIMARY KEY (id);


--
-- TOC entry 2751 (class 2606 OID 17024)
-- Name: password password_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY password
    ADD CONSTRAINT password_pkey PRIMARY KEY (id);


--
-- TOC entry 2789 (class 2606 OID 18140)
-- Name: payee_crypto payee_crypto_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee_crypto
    ADD CONSTRAINT payee_crypto_pkey PRIMARY KEY (id);


--
-- TOC entry 2795 (class 2606 OID 18222)
-- Name: payee_fiat payee_fiat_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee_fiat
    ADD CONSTRAINT payee_fiat_pkey PRIMARY KEY (id);


--
-- TOC entry 2785 (class 2606 OID 18129)
-- Name: payee_master payee_master_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee_master
    ADD CONSTRAINT payee_master_pkey PRIMARY KEY (id);


--
-- TOC entry 2787 (class 2606 OID 18127)
-- Name: payee payee_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee
    ADD CONSTRAINT payee_pkey PRIMARY KEY (id);


--
-- TOC entry 2821 (class 2606 OID 18911)
-- Name: site_bank site_bank_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY site_bank
    ADD CONSTRAINT site_bank_pkey PRIMARY KEY (id);


--
-- TOC entry 2741 (class 2606 OID 16582)
-- Name: transaction transaction_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (id);


--
-- TOC entry 2783 (class 2606 OID 18065)
-- Name: transfer_master transfer_master_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_master
    ADD CONSTRAINT transfer_master_pkey PRIMARY KEY (id);


--
-- TOC entry 2793 (class 2606 OID 18199)
-- Name: transfer_payee transfer_payee_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_payee
    ADD CONSTRAINT transfer_payee_pkey PRIMARY KEY (id);


--
-- TOC entry 2755 (class 2606 OID 17229)
-- Name: transfer transfer_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer
    ADD CONSTRAINT transfer_pkey PRIMARY KEY (id);


--
-- TOC entry 2791 (class 2606 OID 18176)
-- Name: transfer_wallet transfer_wallet_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_wallet
    ADD CONSTRAINT transfer_wallet_pkey PRIMARY KEY (id);


--
-- TOC entry 2799 (class 2606 OID 18276)
-- Name: transfer_bank transfer_withdraw_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_bank
    ADD CONSTRAINT transfer_withdraw_pkey PRIMARY KEY (id);


--
-- TOC entry 2767 (class 2606 OID 17569)
-- Name: user_country user_country_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY user_country
    ADD CONSTRAINT user_country_pkey PRIMARY KEY (id);


--
-- TOC entry 2743 (class 2606 OID 16584)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 2753 (class 2606 OID 17035)
-- Name: user_setting user_setting_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY user_setting
    ADD CONSTRAINT user_setting_pkey PRIMARY KEY (id);


--
-- TOC entry 2745 (class 2606 OID 16946)
-- Name: user user_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_username_key UNIQUE (username);


--
-- TOC entry 2779 (class 2606 OID 17908)
-- Name: wallet_fiat wallet_fiat_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_fiat
    ADD CONSTRAINT wallet_fiat_pkey PRIMARY KEY (id);


--
-- TOC entry 2781 (class 2606 OID 17992)
-- Name: wallet_master wallet_master_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_master
    ADD CONSTRAINT wallet_master_pkey PRIMARY KEY (id);


--
-- TOC entry 2765 (class 2606 OID 17450)
-- Name: wallet_passphrase wallet_password_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_passphrase
    ADD CONSTRAINT wallet_password_pkey PRIMARY KEY (id);


--
-- TOC entry 2747 (class 2606 OID 16586)
-- Name: wallet wallet_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet
    ADD CONSTRAINT wallet_pkey PRIMARY KEY (id);


--
-- TOC entry 2848 (class 2606 OID 18415)
-- Name: admin admin_admin_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin
    ADD CONSTRAINT admin_admin_group_id_fkey FOREIGN KEY (admin_group_id) REFERENCES admin_group(id);


--
-- TOC entry 2913 (class 2606 OID 18435)
-- Name: admin_country admin_country_admin_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_country
    ADD CONSTRAINT admin_country_admin_id_fkey FOREIGN KEY (admin_id) REFERENCES admin(id);


--
-- TOC entry 2914 (class 2606 OID 18440)
-- Name: admin_country admin_country_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_country
    ADD CONSTRAINT admin_country_country_id_fkey FOREIGN KEY (country_id) REFERENCES country(id);


--
-- TOC entry 2849 (class 2606 OID 16592)
-- Name: admin_access admin_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_access
    ADD CONSTRAINT admin_id FOREIGN KEY (admin_id) REFERENCES admin(id);


--
-- TOC entry 2912 (class 2606 OID 18389)
-- Name: admin_token admin_token_admin_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY admin_token
    ADD CONSTRAINT admin_token_admin_id_fkey FOREIGN KEY (admin_id) REFERENCES admin(id);


--
-- TOC entry 2852 (class 2606 OID 16602)
-- Name: field country_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY field
    ADD CONSTRAINT country_id FOREIGN KEY (country_id) REFERENCES country(id);


--
-- TOC entry 2873 (class 2606 OID 18430)
-- Name: currency currency_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY currency
    ADD CONSTRAINT currency_country_id_fkey FOREIGN KEY (country_id) REFERENCES country(id);


--
-- TOC entry 2874 (class 2606 OID 17398)
-- Name: currency_pair currency_pair_from_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY currency_pair
    ADD CONSTRAINT currency_pair_from_currency_id_fkey FOREIGN KEY (from_currency_id) REFERENCES currency(id);


--
-- TOC entry 2875 (class 2606 OID 17403)
-- Name: currency_pair currency_pair_to_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY currency_pair
    ADD CONSTRAINT currency_pair_to_currency_id_fkey FOREIGN KEY (to_currency_id) REFERENCES currency(id);


--
-- TOC entry 2888 (class 2606 OID 17680)
-- Name: data_category data_category_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_category
    ADD CONSTRAINT data_category_category_id_fkey FOREIGN KEY (field_category_id) REFERENCES field_category(id);


--
-- TOC entry 2886 (class 2606 OID 17670)
-- Name: data_category data_category_data_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_category
    ADD CONSTRAINT data_category_data_id_fkey FOREIGN KEY (data_id) REFERENCES data(id);


--
-- TOC entry 2887 (class 2606 OID 17675)
-- Name: data_category data_category_field_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_category
    ADD CONSTRAINT data_category_field_id_fkey FOREIGN KEY (field_id) REFERENCES field(id);


--
-- TOC entry 2880 (class 2606 OID 17639)
-- Name: data data_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data
    ADD CONSTRAINT data_country_id_fkey FOREIGN KEY (country_id) REFERENCES country(id);


--
-- TOC entry 2884 (class 2606 OID 17652)
-- Name: data_file data_file_data_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_file
    ADD CONSTRAINT data_file_data_id_fkey FOREIGN KEY (data_id) REFERENCES data(id);


--
-- TOC entry 2885 (class 2606 OID 17657)
-- Name: data_file data_file_field_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_file
    ADD CONSTRAINT data_file_field_id_fkey FOREIGN KEY (field_id) REFERENCES field(id);


--
-- TOC entry 2882 (class 2606 OID 17624)
-- Name: data_text data_text_data_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_text
    ADD CONSTRAINT data_text_data_id_fkey FOREIGN KEY (data_id) REFERENCES data(id);


--
-- TOC entry 2883 (class 2606 OID 17629)
-- Name: data_text data_text_field_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_text
    ADD CONSTRAINT data_text_field_id_fkey FOREIGN KEY (field_id) REFERENCES field(id);


--
-- TOC entry 2881 (class 2606 OID 17634)
-- Name: data data_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data
    ADD CONSTRAINT data_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2867 (class 2606 OID 17139)
-- Name: email email_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY email
    ADD CONSTRAINT email_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2879 (class 2606 OID 17596)
-- Name: field_category field_category_field_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY field_category
    ADD CONSTRAINT field_category_field_id_fkey FOREIGN KEY (field_id) REFERENCES field(id);


--
-- TOC entry 2907 (class 2606 OID 18255)
-- Name: field_subcategory field_subcategory_field_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY field_subcategory
    ADD CONSTRAINT field_subcategory_field_id_fkey FOREIGN KEY (field_id) REFERENCES field(id);


--
-- TOC entry 2908 (class 2606 OID 18260)
-- Name: field_subcategory field_subcategory_sub_field_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY field_subcategory
    ADD CONSTRAINT field_subcategory_sub_field_id_fkey FOREIGN KEY (sub_field_id) REFERENCES field(id);


--
-- TOC entry 2862 (class 2606 OID 16632)
-- Name: transaction from_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transaction
    ADD CONSTRAINT from_user FOREIGN KEY (from_user) REFERENCES "user"(id);


--
-- TOC entry 2933 (class 2606 OID 18983)
-- Name: order_broker order_broker_buyer_broker_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_broker
    ADD CONSTRAINT order_broker_buyer_broker_id_fkey FOREIGN KEY (buyer_broker_id) REFERENCES "user"(id);


--
-- TOC entry 2931 (class 2606 OID 18973)
-- Name: order_broker order_broker_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_broker
    ADD CONSTRAINT order_broker_order_id_fkey FOREIGN KEY (order_id) REFERENCES "order"(id);


--
-- TOC entry 2932 (class 2606 OID 18978)
-- Name: order_broker order_broker_seller_broker_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_broker
    ADD CONSTRAINT order_broker_seller_broker_id_fkey FOREIGN KEY (seller_broker_id) REFERENCES "user"(id);


--
-- TOC entry 2935 (class 2606 OID 19002)
-- Name: order_buy_broker order_buy_broker_broker_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_buy_broker
    ADD CONSTRAINT order_buy_broker_broker_id_fkey FOREIGN KEY (broker_id) REFERENCES "user"(id);


--
-- TOC entry 2934 (class 2606 OID 18997)
-- Name: order_buy_broker order_buy_broker_order_buy_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_buy_broker
    ADD CONSTRAINT order_buy_broker_order_buy_id_fkey FOREIGN KEY (order_buy_id) REFERENCES order_buy(id);


--
-- TOC entry 2916 (class 2606 OID 18838)
-- Name: order_buy order_buy_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_buy
    ADD CONSTRAINT order_buy_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2918 (class 2606 OID 19170)
-- Name: order_buy order_buy_currency_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_buy
    ADD CONSTRAINT order_buy_currency_wallet_id_fkey FOREIGN KEY (currency_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2917 (class 2606 OID 18843)
-- Name: order_buy order_buy_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_buy
    ADD CONSTRAINT order_buy_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2919 (class 2606 OID 19175)
-- Name: order_buy order_buy_rate_currency_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_buy
    ADD CONSTRAINT order_buy_rate_currency_wallet_id_fkey FOREIGN KEY (rate_currency_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2915 (class 2606 OID 18833)
-- Name: order_buy order_buy_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_buy
    ADD CONSTRAINT order_buy_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2860 (class 2606 OID 19212)
-- Name: order order_buyer_currency_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "order"
    ADD CONSTRAINT order_buyer_currency_wallet_id_fkey FOREIGN KEY (buyer_currency_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2861 (class 2606 OID 19217)
-- Name: order order_buyer_rate_currency_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "order"
    ADD CONSTRAINT order_buyer_rate_currency_wallet_id_fkey FOREIGN KEY (buyer_rate_currency_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2857 (class 2606 OID 18960)
-- Name: order order_buyer_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "order"
    ADD CONSTRAINT order_buyer_user_id_fkey FOREIGN KEY (buyer_user_id) REFERENCES "user"(id);


--
-- TOC entry 2926 (class 2606 OID 18887)
-- Name: order_currency order_currency_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_currency
    ADD CONSTRAINT order_currency_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2854 (class 2606 OID 18861)
-- Name: order order_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "order"
    ADD CONSTRAINT order_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2927 (class 2606 OID 18892)
-- Name: order_currency order_currency_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_currency
    ADD CONSTRAINT order_currency_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2925 (class 2606 OID 18882)
-- Name: order_currency order_currency_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_currency
    ADD CONSTRAINT order_currency_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2961 (class 2606 OID 19350)
-- Name: order_graph_12h order_graph_12h_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_12h
    ADD CONSTRAINT order_graph_12h_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2962 (class 2606 OID 19345)
-- Name: order_graph_12h order_graph_12h_last_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_12h
    ADD CONSTRAINT order_graph_12h_last_order_id_fkey FOREIGN KEY (last_order_id) REFERENCES "order"(id);


--
-- TOC entry 2960 (class 2606 OID 19355)
-- Name: order_graph_12h order_graph_12h_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_12h
    ADD CONSTRAINT order_graph_12h_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2949 (class 2606 OID 19258)
-- Name: order_graph_15m order_graph_15m_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_15m
    ADD CONSTRAINT order_graph_15m_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2950 (class 2606 OID 19253)
-- Name: order_graph_15m order_graph_15m_last_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_15m
    ADD CONSTRAINT order_graph_15m_last_order_id_fkey FOREIGN KEY (last_order_id) REFERENCES "order"(id);


--
-- TOC entry 2948 (class 2606 OID 19263)
-- Name: order_graph_15m order_graph_15m_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_15m
    ADD CONSTRAINT order_graph_15m_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2964 (class 2606 OID 19373)
-- Name: order_graph_1d order_graph_1d_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1d
    ADD CONSTRAINT order_graph_1d_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2965 (class 2606 OID 19368)
-- Name: order_graph_1d order_graph_1d_last_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1d
    ADD CONSTRAINT order_graph_1d_last_order_id_fkey FOREIGN KEY (last_order_id) REFERENCES "order"(id);


--
-- TOC entry 2963 (class 2606 OID 19378)
-- Name: order_graph_1d order_graph_1d_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1d
    ADD CONSTRAINT order_graph_1d_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2955 (class 2606 OID 19304)
-- Name: order_graph_1h order_graph_1h_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1h
    ADD CONSTRAINT order_graph_1h_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2956 (class 2606 OID 19299)
-- Name: order_graph_1h order_graph_1h_last_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1h
    ADD CONSTRAINT order_graph_1h_last_order_id_fkey FOREIGN KEY (last_order_id) REFERENCES "order"(id);


--
-- TOC entry 2954 (class 2606 OID 19309)
-- Name: order_graph_1h order_graph_1h_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1h
    ADD CONSTRAINT order_graph_1h_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2946 (class 2606 OID 19235)
-- Name: order_graph_1m order_graph_1m_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1m
    ADD CONSTRAINT order_graph_1m_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2947 (class 2606 OID 19230)
-- Name: order_graph_1m order_graph_1m_last_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1m
    ADD CONSTRAINT order_graph_1m_last_order_id_fkey FOREIGN KEY (last_order_id) REFERENCES "order"(id);


--
-- TOC entry 2945 (class 2606 OID 19240)
-- Name: order_graph_1m order_graph_1m_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_1m
    ADD CONSTRAINT order_graph_1m_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2952 (class 2606 OID 19281)
-- Name: order_graph_30m order_graph_30m_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_30m
    ADD CONSTRAINT order_graph_30m_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2953 (class 2606 OID 19276)
-- Name: order_graph_30m order_graph_30m_last_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_30m
    ADD CONSTRAINT order_graph_30m_last_order_id_fkey FOREIGN KEY (last_order_id) REFERENCES "order"(id);


--
-- TOC entry 2951 (class 2606 OID 19286)
-- Name: order_graph_30m order_graph_30m_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_30m
    ADD CONSTRAINT order_graph_30m_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2958 (class 2606 OID 19327)
-- Name: order_graph_6h order_graph_6h_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_6h
    ADD CONSTRAINT order_graph_6h_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2959 (class 2606 OID 19322)
-- Name: order_graph_6h order_graph_6h_last_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_6h
    ADD CONSTRAINT order_graph_6h_last_order_id_fkey FOREIGN KEY (last_order_id) REFERENCES "order"(id);


--
-- TOC entry 2957 (class 2606 OID 19332)
-- Name: order_graph_6h order_graph_6h_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_6h
    ADD CONSTRAINT order_graph_6h_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2967 (class 2606 OID 19396)
-- Name: order_graph_7d order_graph_7d_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_7d
    ADD CONSTRAINT order_graph_7d_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2968 (class 2606 OID 19391)
-- Name: order_graph_7d order_graph_7d_last_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_7d
    ADD CONSTRAINT order_graph_7d_last_order_id_fkey FOREIGN KEY (last_order_id) REFERENCES "order"(id);


--
-- TOC entry 2966 (class 2606 OID 19401)
-- Name: order_graph_7d order_graph_7d_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_7d
    ADD CONSTRAINT order_graph_7d_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2938 (class 2606 OID 19051)
-- Name: order_graph_5m order_graph_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_5m
    ADD CONSTRAINT order_graph_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2939 (class 2606 OID 19056)
-- Name: order_graph_5m order_graph_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_graph_5m
    ADD CONSTRAINT order_graph_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2863 (class 2606 OID 16637)
-- Name: transaction order_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transaction
    ADD CONSTRAINT order_id FOREIGN KEY (order_id) REFERENCES "order"(id);


--
-- TOC entry 2855 (class 2606 OID 18866)
-- Name: order order_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "order"
    ADD CONSTRAINT order_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2937 (class 2606 OID 19020)
-- Name: order_sell_broker order_sell_broker_broker_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_sell_broker
    ADD CONSTRAINT order_sell_broker_broker_id_fkey FOREIGN KEY (broker_id) REFERENCES "user"(id);


--
-- TOC entry 2936 (class 2606 OID 19015)
-- Name: order_sell_broker order_sell_broker_order_sell_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_sell_broker
    ADD CONSTRAINT order_sell_broker_order_sell_id_fkey FOREIGN KEY (order_sell_id) REFERENCES order_sell(id);


--
-- TOC entry 2921 (class 2606 OID 18823)
-- Name: order_sell order_sell_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_sell
    ADD CONSTRAINT order_sell_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2923 (class 2606 OID 19180)
-- Name: order_sell order_sell_currency_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_sell
    ADD CONSTRAINT order_sell_currency_wallet_id_fkey FOREIGN KEY (currency_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2922 (class 2606 OID 18828)
-- Name: order_sell order_sell_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_sell
    ADD CONSTRAINT order_sell_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2924 (class 2606 OID 19185)
-- Name: order_sell order_sell_rate_currency_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_sell
    ADD CONSTRAINT order_sell_rate_currency_wallet_id_fkey FOREIGN KEY (rate_currency_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2920 (class 2606 OID 18818)
-- Name: order_sell order_sell_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_sell
    ADD CONSTRAINT order_sell_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2858 (class 2606 OID 19202)
-- Name: order order_seller_currency_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "order"
    ADD CONSTRAINT order_seller_currency_wallet_id_fkey FOREIGN KEY (seller_currency_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2859 (class 2606 OID 19207)
-- Name: order order_seller_rate_currency_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "order"
    ADD CONSTRAINT order_seller_rate_currency_wallet_id_fkey FOREIGN KEY (seller_rate_currency_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2856 (class 2606 OID 18955)
-- Name: order order_seller_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "order"
    ADD CONSTRAINT order_seller_user_id_fkey FOREIGN KEY (seller_user_id) REFERENCES "user"(id);


--
-- TOC entry 2941 (class 2606 OID 19105)
-- Name: order_wallet order_wallet_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_wallet
    ADD CONSTRAINT order_wallet_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2943 (class 2606 OID 19115)
-- Name: order_wallet order_wallet_currency_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_wallet
    ADD CONSTRAINT order_wallet_currency_wallet_id_fkey FOREIGN KEY (currency_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2942 (class 2606 OID 19110)
-- Name: order_wallet order_wallet_rate_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_wallet
    ADD CONSTRAINT order_wallet_rate_currency_id_fkey FOREIGN KEY (rate_currency_id) REFERENCES currency(id);


--
-- TOC entry 2944 (class 2606 OID 19120)
-- Name: order_wallet order_wallet_rate_currency_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_wallet
    ADD CONSTRAINT order_wallet_rate_currency_wallet_id_fkey FOREIGN KEY (rate_currency_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2940 (class 2606 OID 19100)
-- Name: order_wallet order_wallet_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY order_wallet
    ADD CONSTRAINT order_wallet_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2868 (class 2606 OID 17043)
-- Name: password password_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY password
    ADD CONSTRAINT password_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2898 (class 2606 OID 18141)
-- Name: payee_crypto payee_crypto_payee_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee_crypto
    ADD CONSTRAINT payee_crypto_payee_id_fkey FOREIGN KEY (payee_id) REFERENCES payee(id);


--
-- TOC entry 2906 (class 2606 OID 18228)
-- Name: payee_fiat payee_fiat_data_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee_fiat
    ADD CONSTRAINT payee_fiat_data_id_fkey FOREIGN KEY (data_id) REFERENCES data(id);


--
-- TOC entry 2905 (class 2606 OID 18223)
-- Name: payee_fiat payee_fiat_payee_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee_fiat
    ADD CONSTRAINT payee_fiat_payee_id_fkey FOREIGN KEY (payee_id) REFERENCES payee(id);


--
-- TOC entry 2896 (class 2606 OID 18101)
-- Name: payee_master payee_master_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee_master
    ADD CONSTRAINT payee_master_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2895 (class 2606 OID 18096)
-- Name: payee_master payee_master_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee_master
    ADD CONSTRAINT payee_master_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2897 (class 2606 OID 18146)
-- Name: payee payee_payee_master_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY payee
    ADD CONSTRAINT payee_payee_master_id_fkey FOREIGN KEY (payee_master_id) REFERENCES payee_master(id);


--
-- TOC entry 2929 (class 2606 OID 18917)
-- Name: site_bank site_bank_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY site_bank
    ADD CONSTRAINT site_bank_country_id_fkey FOREIGN KEY (country_id) REFERENCES country(id);


--
-- TOC entry 2930 (class 2606 OID 18922)
-- Name: site_bank site_bank_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY site_bank
    ADD CONSTRAINT site_bank_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2928 (class 2606 OID 18912)
-- Name: site_bank site_bank_updated_admin_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY site_bank
    ADD CONSTRAINT site_bank_updated_admin_id_fkey FOREIGN KEY (updated_admin_id) REFERENCES admin(id);


--
-- TOC entry 2864 (class 2606 OID 16642)
-- Name: transaction to_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transaction
    ADD CONSTRAINT to_user FOREIGN KEY (to_user) REFERENCES "user"(id);


--
-- TOC entry 2865 (class 2606 OID 17372)
-- Name: transaction transaction_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transaction
    ADD CONSTRAINT transaction_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2871 (class 2606 OID 18339)
-- Name: transfer transfer_from_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer
    ADD CONSTRAINT transfer_from_user_id_fkey FOREIGN KEY (from_user_id) REFERENCES "user"(id);


--
-- TOC entry 2894 (class 2606 OID 18071)
-- Name: transfer_master transfer_master_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_master
    ADD CONSTRAINT transfer_master_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2893 (class 2606 OID 18066)
-- Name: transfer_master transfer_master_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_master
    ADD CONSTRAINT transfer_master_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2903 (class 2606 OID 18205)
-- Name: transfer_payee transfer_payee_from_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_payee
    ADD CONSTRAINT transfer_payee_from_wallet_id_fkey FOREIGN KEY (from_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2904 (class 2606 OID 18210)
-- Name: transfer_payee transfer_payee_to_payee_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_payee
    ADD CONSTRAINT transfer_payee_to_payee_id_fkey FOREIGN KEY (to_payee_id) REFERENCES payee(id);


--
-- TOC entry 2902 (class 2606 OID 18200)
-- Name: transfer_payee transfer_payee_transfer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_payee
    ADD CONSTRAINT transfer_payee_transfer_id_fkey FOREIGN KEY (transfer_id) REFERENCES transfer(id);


--
-- TOC entry 2872 (class 2606 OID 18344)
-- Name: transfer transfer_to_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer
    ADD CONSTRAINT transfer_to_user_id_fkey FOREIGN KEY (to_user_id) REFERENCES "user"(id);


--
-- TOC entry 2870 (class 2606 OID 18076)
-- Name: transfer transfer_transfer_master_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer
    ADD CONSTRAINT transfer_transfer_master_id_fkey FOREIGN KEY (transfer_master_id) REFERENCES transfer_master(id);


--
-- TOC entry 2900 (class 2606 OID 18182)
-- Name: transfer_wallet transfer_wallet_from_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_wallet
    ADD CONSTRAINT transfer_wallet_from_wallet_id_fkey FOREIGN KEY (from_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2901 (class 2606 OID 18187)
-- Name: transfer_wallet transfer_wallet_to_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_wallet
    ADD CONSTRAINT transfer_wallet_to_wallet_id_fkey FOREIGN KEY (to_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2899 (class 2606 OID 18177)
-- Name: transfer_wallet transfer_wallet_transfer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_wallet
    ADD CONSTRAINT transfer_wallet_transfer_id_fkey FOREIGN KEY (transfer_id) REFERENCES transfer(id);


--
-- TOC entry 2911 (class 2606 OID 18287)
-- Name: transfer_bank transfer_withdraw_data_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_bank
    ADD CONSTRAINT transfer_withdraw_data_id_fkey FOREIGN KEY (to_data_id) REFERENCES data(id);


--
-- TOC entry 2909 (class 2606 OID 18277)
-- Name: transfer_bank transfer_withdraw_transfer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_bank
    ADD CONSTRAINT transfer_withdraw_transfer_id_fkey FOREIGN KEY (transfer_id) REFERENCES transfer(id);


--
-- TOC entry 2910 (class 2606 OID 18282)
-- Name: transfer_bank transfer_withdraw_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY transfer_bank
    ADD CONSTRAINT transfer_withdraw_wallet_id_fkey FOREIGN KEY (from_wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2878 (class 2606 OID 17575)
-- Name: user_country user_country_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY user_country
    ADD CONSTRAINT user_country_country_id_fkey FOREIGN KEY (country_id) REFERENCES country(id);


--
-- TOC entry 2877 (class 2606 OID 17570)
-- Name: user_country user_country_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY user_country
    ADD CONSTRAINT user_country_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2853 (class 2606 OID 16672)
-- Name: mobile user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY mobile
    ADD CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2850 (class 2606 OID 16677)
-- Name: auth_token user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY auth_token
    ADD CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2869 (class 2606 OID 17038)
-- Name: user_setting user_setting_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY user_setting
    ADD CONSTRAINT user_setting_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2851 (class 2606 OID 17979)
-- Name: wallet_crypto wallet_crypto_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_crypto
    ADD CONSTRAINT wallet_crypto_wallet_id_fkey FOREIGN KEY (wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2890 (class 2606 OID 17914)
-- Name: wallet_fiat wallet_fiat_data_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_fiat
    ADD CONSTRAINT wallet_fiat_data_id_fkey FOREIGN KEY (data_id) REFERENCES data(id);


--
-- TOC entry 2889 (class 2606 OID 17909)
-- Name: wallet_fiat wallet_fiat_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_fiat
    ADD CONSTRAINT wallet_fiat_wallet_id_fkey FOREIGN KEY (wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2892 (class 2606 OID 18033)
-- Name: wallet_master wallet_master_currency_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_master
    ADD CONSTRAINT wallet_master_currency_id_fkey FOREIGN KEY (currency_id) REFERENCES currency(id);


--
-- TOC entry 2891 (class 2606 OID 18028)
-- Name: wallet_master wallet_master_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_master
    ADD CONSTRAINT wallet_master_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id);


--
-- TOC entry 2876 (class 2606 OID 17451)
-- Name: wallet_passphrase wallet_password_wallet_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet_passphrase
    ADD CONSTRAINT wallet_password_wallet_id_fkey FOREIGN KEY (wallet_id) REFERENCES wallet(id);


--
-- TOC entry 2866 (class 2606 OID 18038)
-- Name: wallet wallet_wallet_master_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wallet
    ADD CONSTRAINT wallet_wallet_master_id_fkey FOREIGN KEY (wallet_master_id) REFERENCES wallet_master(id);


-- Completed on 2019-02-27 20:57:56 EST

--
-- PostgreSQL database dump complete
--

