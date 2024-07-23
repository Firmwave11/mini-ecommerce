-- public.warehouse

-- Drop table

-- DROP TABLE public.warehouse;

CREATE TABLE public.warehouse (
	id serial4 NOT NULL,
	name varchar NOT NULL,
	shop_id int NOT NULL,
	enabled bool NOT NULL,
	area varchar NOT NULL,
	address varchar NOT NULL,
	updated_at timestamp NOT NULL DEFAULT now(),
	created_at timestamp NOT NULL DEFAULT now(),
	is_deleted bool NOT NULL DEFAULT false,
	CONSTRAINT warehouse_pk PRIMARY KEY (id)
);

-- public.stock definition

-- Drop table

-- DROP TABLE public.stock;

CREATE TABLE public.stock (
	id serial4 NOT NULL,
	warehouse_id int NOT NULL,
	product_id int NOT NULL,
	quantity int NOT NULL,
	updated_at timestamp NOT NULL DEFAULT now(),
	created_at timestamp NOT NULL DEFAULT now(),
	is_deleted bool NOT NULL DEFAULT false,
	CONSTRAINT stock_pk PRIMARY KEY (id)
);

ALTER TABLE public.stock ADD CONSTRAINT stock_fk FOREIGN KEY (warehouse_id) REFERENCES public.warehouse(id) ON UPDATE CASCADE;

-- public.stock_histories definition

-- Drop table

-- DROP TABLE public.stock_histories;

CREATE TABLE public.stock_histories (
	id serial4 NOT NULL,
	stock_id int4 NOT NULL,
	product_id int NOT NULL,
	quantity int NOT NULL,
	transaction_type varchar not null,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	is_deleted bool NOT NULL DEFAULT false,
	CONSTRAINT stock_histories_pk PRIMARY KEY (id)
);


-- public.stock_histories foreign keys

ALTER TABLE public.stock_histories ADD CONSTRAINT stock_histories_fk FOREIGN KEY (stock_id) REFERENCES public.stock(id) ON UPDATE CASCADE;
