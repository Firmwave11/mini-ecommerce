-- public.product definition

-- Drop table

-- DROP TABLE public.product;

CREATE TABLE public.product (
	id serial4 NOT NULL,
	name varchar not null,
	price decimal NOT NULL,
	weight decimal not null,
	enabled bool not null,
	updated_at timestamp NOT NULL DEFAULT now(),
	created_at timestamp NOT NULL DEFAULT now(),
	is_deleted bool NOT NULL DEFAULT false,
	CONSTRAINT product_pk PRIMARY KEY (id)
);