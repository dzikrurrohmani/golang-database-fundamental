--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2
-- Dumped by pg_dump version 14.2

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
-- Name: restowmb; Type: SCHEMA; Schema: -; Owner: jutioncandrakirana
--

CREATE SCHEMA restowmb;


ALTER SCHEMA restowmb OWNER TO jutioncandrakirana;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: customer; Type: TABLE; Schema: public; Owner: jutioncandrakirana
--

CREATE TABLE public.customer (
    id character varying(60) NOT NULL,
    name character varying(100),
    address character varying(200),
    phone character varying(15),
    email character varying(100),
    saldo integer,
    is_status integer DEFAULT 1,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone
);


ALTER TABLE public.customer OWNER TO jutioncandrakirana;

--
-- Name: product; Type: TABLE; Schema: public; Owner: jutioncandrakirana
--

CREATE TABLE public.product (
    id character varying(60) NOT NULL,
    name character varying(100),
    price integer,
    description character varying(100),
    stock integer,
    store_id character varying(60),
    is_status integer DEFAULT 1,
    CONSTRAINT product_stock_check CHECK ((stock >= 0))
);


ALTER TABLE public.product OWNER TO jutioncandrakirana;

--
-- Name: shop; Type: TABLE; Schema: public; Owner: jutioncandrakirana
--

CREATE TABLE public.shop (
    id character varying(60) NOT NULL,
    no_siup character varying(50) NOT NULL,
    name character varying(100),
    address character varying(200),
    phone character varying(15),
    is_status integer DEFAULT 1
);


ALTER TABLE public.shop OWNER TO jutioncandrakirana;

--
-- Name: transaction; Type: TABLE; Schema: public; Owner: jutioncandrakirana
--

CREATE TABLE public.transaction (
    id character varying(60) NOT NULL,
    customer_id character varying(60),
    product_id character varying(60),
    store_id character varying(60),
    purchase_date date,
    quantity integer
);


ALTER TABLE public.transaction OWNER TO jutioncandrakirana;

--
-- Name: m_customer; Type: TABLE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE TABLE restowmb.m_customer (
    id bigint NOT NULL,
    customer_name character varying(50) NOT NULL,
    mobile_phone_no character varying(20) NOT NULL,
    is_member boolean
);


ALTER TABLE restowmb.m_customer OWNER TO jutioncandrakirana;

--
-- Name: m_customer_discount_id_seq; Type: SEQUENCE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE SEQUENCE restowmb.m_customer_discount_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE restowmb.m_customer_discount_id_seq OWNER TO jutioncandrakirana;

--
-- Name: m_customer_discount; Type: TABLE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE TABLE restowmb.m_customer_discount (
    id integer DEFAULT nextval('restowmb.m_customer_discount_id_seq'::regclass) NOT NULL,
    disc_id bigint,
    customer_id bigint
);


ALTER TABLE restowmb.m_customer_discount OWNER TO jutioncandrakirana;

--
-- Name: m_discount_id_seq; Type: SEQUENCE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE SEQUENCE restowmb.m_discount_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE restowmb.m_discount_id_seq OWNER TO jutioncandrakirana;

--
-- Name: m_discount; Type: TABLE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE TABLE restowmb.m_discount (
    id integer DEFAULT nextval('restowmb.m_discount_id_seq'::regclass) NOT NULL,
    disc_desciption character varying(50),
    pct integer
);


ALTER TABLE restowmb.m_discount OWNER TO jutioncandrakirana;

--
-- Name: m_menu; Type: TABLE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE TABLE restowmb.m_menu (
    id bigint NOT NULL,
    menu_name character varying(100) NOT NULL
);


ALTER TABLE restowmb.m_menu OWNER TO jutioncandrakirana;

--
-- Name: m_menu_price_id_seq; Type: SEQUENCE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE SEQUENCE restowmb.m_menu_price_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE restowmb.m_menu_price_id_seq OWNER TO jutioncandrakirana;

--
-- Name: m_menu_price; Type: TABLE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE TABLE restowmb.m_menu_price (
    id integer DEFAULT nextval('restowmb.m_menu_price_id_seq'::regclass) NOT NULL,
    menu_id bigint NOT NULL,
    price real NOT NULL
);


ALTER TABLE restowmb.m_menu_price OWNER TO jutioncandrakirana;

--
-- Name: m_table; Type: TABLE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE TABLE restowmb.m_table (
    id integer NOT NULL,
    table_name character(3) NOT NULL
);


ALTER TABLE restowmb.m_table OWNER TO jutioncandrakirana;

--
-- Name: m_trans_type; Type: TABLE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE TABLE restowmb.m_trans_type (
    trans_type_id character varying(3) NOT NULL,
    description character varying(50)
);


ALTER TABLE restowmb.m_trans_type OWNER TO jutioncandrakirana;

--
-- Name: t_bill; Type: TABLE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE TABLE restowmb.t_bill (
    id bigint NOT NULL,
    trans_date date NOT NULL,
    customer_id bigint NOT NULL,
    table_id integer,
    trans_type character varying(3) NOT NULL
);


ALTER TABLE restowmb.t_bill OWNER TO jutioncandrakirana;

--
-- Name: t_bill_detail; Type: TABLE; Schema: restowmb; Owner: jutioncandrakirana
--

CREATE TABLE restowmb.t_bill_detail (
    id bigint NOT NULL,
    bill_id bigint NOT NULL,
    menu_price_id integer NOT NULL,
    qty real NOT NULL
);


ALTER TABLE restowmb.t_bill_detail OWNER TO jutioncandrakirana;

--
-- Data for Name: customer; Type: TABLE DATA; Schema: public; Owner: jutioncandrakirana
--

COPY public.customer (id, name, address, phone, email, saldo, is_status, created_at, updated_at, deleted_at) FROM stdin;
cd34aac6-071e-424f-976a-c33124bdd6e4	Jution Kirana	Ragunan	08292929	jutionck@gmail.com	150000	1	2022-06-27 03:38:56.002	\N	\N
cd34aac6-071e-424f-976a-c33124bdd6ex	Megalodon	Surakarta	0829292	megalodon.tok@gmail.com	1000000	0	2022-06-27 03:38:57.002	\N	\N
cd34aac6-071e-424f-976a-c33124bdd6e9	Rifqy Puasa	Depok	0823339889	rifqy.puasa@gmail.com	1000	1	2022-06-27 03:38:51.002	\N	\N
cd34aac6-071e-424f-976a-c33124bdc6e4	Mamang Resing	Bogor	028373892	mamang.resing@gmail.com	1000000	0	2022-06-27 03:38:54.002	\N	\N
cd34aac6-071e-424f-976a-c33124bd36e4	Tika Teki	Lampung	028222292	tika.teki@gmail.com	1000	1	2022-06-27 03:38:46.002	\N	\N
cd34aac6-071e-424f-97da-c33124bdd6e4	Ojan Riders	Bekasi	028282992	ojan.riders@gmail.com	1000000	0	2022-06-27 03:38:53.002	\N	\N
cd34aac6-071e-424f-976a-c33124bdd6es	Puan	Jakarta	0829298339	puan.aja@gmail.com	1000	1	2022-06-27 03:38:49.002	\N	\N
1c63f19d-e896-4116-a847-2dbb75d7eae9	Mega Mendung	Jakarta	082929889	mega.mendung@gmail.com	1000	1	2022-06-27 03:38:50.002	\N	\N
cd34aac6-071e-422f-976a-c33124bdd6e4	Atta Mini	Jakarta	08292929	atta.mini@gmail.com	900000	1	2022-06-27 03:38:47.002	\N	\N
cd34aac6-071e-424f-976a-c33124bdd6et	Dona Dinihari	\N	028239992	dona.dini@gmail.com	1000000	1	2022-06-27 03:38:45.002	\N	\N
cd34aac6-071e-424f-976a-c33124bdd614	Indo Dana	\N	82920200	indo.dana@gmail.com	900000	1	2022-06-27 03:38:48.002	\N	\N
5b6c5af5-e139-435f-aafc-771084ad7ea4	Rara Mendung	Pati	0823338282	rara.mendung@gmail.com	89000	1	2022-06-27 03:38:52.002	\N	\N
1c63f19d-e896-4116-a847-2dbb75d7eae8	Jution Mendung	Ragunan	08292929	jutionck@gmail.com	150000	1	2022-06-27 03:38:55.002	\N	\N
10001	Bulan	\N	\N	\N	4400	1	2022-06-27 16:44:50.689602	\N	\N
10002	Bintang	\N	\N	\N	5200	1	2022-06-27 16:44:50.689602	\N	\N
\.


--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: jutioncandrakirana
--

COPY public.product (id, name, price, description, stock, store_id, is_status) FROM stdin;
fd97dbee-f48d-4d8b-9900-45bae3b736f8	Gula Batu	15000	Untuk yang kurang manis	20	7dc92174-82ad-41a0-98d8-44aa44662c13	1
fd97dbee-f48d-4d8b-9900-45bae3b736f9	Minyak Telon	14500	Untuk yang digigit nyamuk	11	7dc92174-82ad-41a0-98d8-44aa44662c12	1
fd97dbee-f48d-4d8b-9900-45bae3b836f9	Beras Rojo Lele	190000	Untuk yang dimakan	100	7dc92174-82ad-41a0-98d8-44aa44662c12	1
fd97dbee-f48d-4d8b-9900-25bae3b736f9	Baking Soda	10000	Untuk yang berkembang	5	7dc92174-82ad-41a0-98d8-44aa44662c16	1
\.


--
-- Data for Name: shop; Type: TABLE DATA; Schema: public; Owner: jutioncandrakirana
--

COPY public.shop (id, no_siup, name, address, phone, is_status) FROM stdin;
7dc92174-82ad-41a0-98d8-44aa44662c13	SIPU/XII/2/2222	Toko Serba Guna	Jakarta	20020202	1
7dc92174-82ad-41a0-98d8-44aa44662c12	SIPU/XII/2/1111	Toko Ada Apa Aja	Banten	29299292	1
7dc92174-82ad-41a0-98d8-44aa44662c16	SIPU/XII/2/3333	Toko Pak Udun	Surakarta	27288282	1
\.


--
-- Data for Name: transaction; Type: TABLE DATA; Schema: public; Owner: jutioncandrakirana
--

COPY public.transaction (id, customer_id, product_id, store_id, purchase_date, quantity) FROM stdin;
\.


--
-- Data for Name: m_customer; Type: TABLE DATA; Schema: restowmb; Owner: jutioncandrakirana
--

COPY restowmb.m_customer (id, customer_name, mobile_phone_no, is_member) FROM stdin;
1	Kadir	0877123333	f
2	Basuki	0812399123	f
3	Munaroh	0888901920	t
4	Adinda	0815343411	f
5	Rozak	0857129823	f
6	Devi	0877745983	t
7	Qibil	0821834583	f
8	Nurman	0877567211	f
\.


--
-- Data for Name: m_customer_discount; Type: TABLE DATA; Schema: restowmb; Owner: jutioncandrakirana
--

COPY restowmb.m_customer_discount (id, disc_id, customer_id) FROM stdin;
1	1	3
2	1	6
\.


--
-- Data for Name: m_discount; Type: TABLE DATA; Schema: restowmb; Owner: jutioncandrakirana
--

COPY restowmb.m_discount (id, disc_desciption, pct) FROM stdin;
1	Discount 10%	10
\.


--
-- Data for Name: m_menu; Type: TABLE DATA; Schema: restowmb; Owner: jutioncandrakirana
--

COPY restowmb.m_menu (id, menu_name) FROM stdin;
1	Nasi Putih
2	Sayur Sop
3	Tahu
4	Es Teh Tawar
5	Sayur Lodeh
6	Tempe
7	Nasi Goreng
8	Telor Ceplok
9	Nasi Goreng Spesial
10	Sayur Kangkung
11	Telor Dadar
12	Kopi Kapal Api
13	Aneka Gorengan
14	Indomie Goreng Telor
15	Es Ovaltine
16	Telor Balado
17	Sayur Buncis
18	Es Teh Manis
19	Tempe Orek
20	Sayur Tahu
21	Indomie Kari Ayam Telor
\.


--
-- Data for Name: m_menu_price; Type: TABLE DATA; Schema: restowmb; Owner: jutioncandrakirana
--

COPY restowmb.m_menu_price (id, menu_id, price) FROM stdin;
1	1	3000
2	2	2000
3	3	2000
4	4	1500
5	5	2500
6	6	2000
7	7	12000
8	8	5000
9	9	25000
10	10	1500
11	11	5000
12	12	4000
13	13	2000
14	14	10000
15	15	6000
16	16	4000
17	17	3000
18	18	1500
19	19	3000
20	20	3000
21	21	10000
22	1	4000
\.


--
-- Data for Name: m_table; Type: TABLE DATA; Schema: restowmb; Owner: jutioncandrakirana
--

COPY restowmb.m_table (id, table_name) FROM stdin;
1	T01
2	T02
3	T03
4	T04
\.


--
-- Data for Name: m_trans_type; Type: TABLE DATA; Schema: restowmb; Owner: jutioncandrakirana
--

COPY restowmb.m_trans_type (trans_type_id, description) FROM stdin;
EI	Eat In
TA	Take Away
\.


--
-- Data for Name: t_bill; Type: TABLE DATA; Schema: restowmb; Owner: jutioncandrakirana
--

COPY restowmb.t_bill (id, trans_date, customer_id, table_id, trans_type) FROM stdin;
1	2022-06-01	1	1	EI
2	2022-06-01	2	2	EI
3	2022-06-01	3	1	EI
4	2022-06-02	4	\N	TA
5	2022-06-03	2	1	EI
6	2022-06-03	5	2	EI
7	2022-06-03	6	3	EI
8	2022-06-03	7	2	EI
9	2022-06-03	8	4	EI
10	2022-06-04	2	1	EI
11	2022-06-04	5	4	EI
12	2022-06-05	5	\N	TA
\.


--
-- Data for Name: t_bill_detail; Type: TABLE DATA; Schema: restowmb; Owner: jutioncandrakirana
--

COPY restowmb.t_bill_detail (id, bill_id, menu_price_id, qty) FROM stdin;
1	1	1	1
2	1	2	1
3	1	3	2
4	1	4	1
5	2	1	1
6	2	5	1
7	2	6	2
8	2	4	1
9	3	7	1
10	3	8	1
11	4	9	1
12	5	1	1
13	5	10	1
14	5	11	1
15	5	4	1
16	6	12	1
17	6	13	5
18	7	14	1
19	7	15	1
20	8	1	0.5
21	8	16	1
22	8	17	1
23	8	13	2
24	8	18	1
25	9	1	1
26	9	19	1
27	9	20	1
28	9	4	1
29	10	22	1
30	10	5	1
31	10	6	1
32	10	18	1
33	11	21	1
34	11	18	1
35	12	13	15
\.


--
-- Name: m_customer_discount_id_seq; Type: SEQUENCE SET; Schema: restowmb; Owner: jutioncandrakirana
--

SELECT pg_catalog.setval('restowmb.m_customer_discount_id_seq', 1, false);


--
-- Name: m_discount_id_seq; Type: SEQUENCE SET; Schema: restowmb; Owner: jutioncandrakirana
--

SELECT pg_catalog.setval('restowmb.m_discount_id_seq', 1, false);


--
-- Name: m_menu_price_id_seq; Type: SEQUENCE SET; Schema: restowmb; Owner: jutioncandrakirana
--

SELECT pg_catalog.setval('restowmb.m_menu_price_id_seq', 1, false);


--
-- Name: customer customer_pkey; Type: CONSTRAINT; Schema: public; Owner: jutioncandrakirana
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (id);


--
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: jutioncandrakirana
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- Name: shop shop_pkey; Type: CONSTRAINT; Schema: public; Owner: jutioncandrakirana
--

ALTER TABLE ONLY public.shop
    ADD CONSTRAINT shop_pkey PRIMARY KEY (id);


--
-- Name: transaction transaction_pkey; Type: CONSTRAINT; Schema: public; Owner: jutioncandrakirana
--

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (id);


--
-- Name: m_customer_discount m_customer_discount_pkey; Type: CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.m_customer_discount
    ADD CONSTRAINT m_customer_discount_pkey PRIMARY KEY (id);


--
-- Name: m_customer m_customer_pkey; Type: CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.m_customer
    ADD CONSTRAINT m_customer_pkey PRIMARY KEY (id);


--
-- Name: m_discount m_discount_pkey; Type: CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.m_discount
    ADD CONSTRAINT m_discount_pkey PRIMARY KEY (id);


--
-- Name: m_menu m_menu_pkey; Type: CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.m_menu
    ADD CONSTRAINT m_menu_pkey PRIMARY KEY (id);


--
-- Name: m_menu_price m_menu_price_pkey; Type: CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.m_menu_price
    ADD CONSTRAINT m_menu_price_pkey PRIMARY KEY (id);


--
-- Name: m_table m_table_pkey; Type: CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.m_table
    ADD CONSTRAINT m_table_pkey PRIMARY KEY (id);


--
-- Name: m_trans_type m_trans_type_pkey; Type: CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.m_trans_type
    ADD CONSTRAINT m_trans_type_pkey PRIMARY KEY (trans_type_id);


--
-- Name: t_bill_detail t_bill_detail_pkey; Type: CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.t_bill_detail
    ADD CONSTRAINT t_bill_detail_pkey PRIMARY KEY (id);


--
-- Name: t_bill t_bill_pkey; Type: CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.t_bill
    ADD CONSTRAINT t_bill_pkey PRIMARY KEY (id);


--
-- Name: product product_shop_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: jutioncandrakirana
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_shop_id_fk FOREIGN KEY (store_id) REFERENCES public.shop(id);


--
-- Name: m_customer_discount m_customer_discount_customer_id_fkey; Type: FK CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.m_customer_discount
    ADD CONSTRAINT m_customer_discount_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES restowmb.m_customer(id);


--
-- Name: m_customer_discount m_customer_discount_disc_id_fkey; Type: FK CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.m_customer_discount
    ADD CONSTRAINT m_customer_discount_disc_id_fkey FOREIGN KEY (disc_id) REFERENCES restowmb.m_discount(id);


--
-- Name: m_menu_price m_menu_price_menu_id_fkey; Type: FK CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.m_menu_price
    ADD CONSTRAINT m_menu_price_menu_id_fkey FOREIGN KEY (menu_id) REFERENCES restowmb.m_menu(id);


--
-- Name: t_bill t_bill_customer_id_fkey; Type: FK CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.t_bill
    ADD CONSTRAINT t_bill_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES restowmb.m_customer(id);


--
-- Name: t_bill_detail t_bill_detail_bill_id_fkey; Type: FK CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.t_bill_detail
    ADD CONSTRAINT t_bill_detail_bill_id_fkey FOREIGN KEY (bill_id) REFERENCES restowmb.t_bill(id);


--
-- Name: t_bill_detail t_bill_detail_menu_price_id_fkey; Type: FK CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.t_bill_detail
    ADD CONSTRAINT t_bill_detail_menu_price_id_fkey FOREIGN KEY (menu_price_id) REFERENCES restowmb.m_menu_price(id);


--
-- Name: t_bill t_bill_table_id_fkey; Type: FK CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.t_bill
    ADD CONSTRAINT t_bill_table_id_fkey FOREIGN KEY (table_id) REFERENCES restowmb.m_table(id);


--
-- Name: t_bill t_bill_trans_type_fkey; Type: FK CONSTRAINT; Schema: restowmb; Owner: jutioncandrakirana
--

ALTER TABLE ONLY restowmb.t_bill
    ADD CONSTRAINT t_bill_trans_type_fkey FOREIGN KEY (trans_type) REFERENCES restowmb.m_trans_type(trans_type_id);


--
-- PostgreSQL database dump complete
--

