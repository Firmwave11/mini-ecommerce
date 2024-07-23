-- public.shop definition

-- Drop table

-- DROP TABLE public.shop;

CREATE TABLE public.shop (
	id serial4 NOT NULL,
	customer_id int NOT NULL,
	address varchar NOT NULL,
	name varchar NOT NULL,
	phone varchar not null,
	updated_at timestamp NOT NULL DEFAULT now(),
	created_at timestamp NOT NULL DEFAULT now(),
	is_deleted bool NOT NULL DEFAULT false,
	CONSTRAINT shop_pk PRIMARY KEY (id)
);