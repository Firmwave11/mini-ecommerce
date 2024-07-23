-- public.order definition

-- Drop table

-- DROP TABLE public.order;

CREATE TABLE public.order (
	id serial4 NOT NULL,
	customer_id int not null,
	total_quantity int NOT NULL,
	total_weight decimal not null,
	total_price decimal not null,
	is_paymented bool not null,
	updated_at timestamp NOT NULL DEFAULT now(),
	created_at timestamp NOT NULL DEFAULT now(),
	is_deleted bool NOT NULL DEFAULT false,
	CONSTRAINT order_pk PRIMARY KEY (id)
);


CREATE TABLE public.order_product (
	id serial4 NOT NULL,
	product_id int not null,
	order_id int not null,
	warehouse_id int not null,
	quantity int NOT NULL,
	weight decimal not null,
	price decimal not null,
	updated_at timestamp NOT NULL DEFAULT now(),
	created_at timestamp NOT NULL DEFAULT now(),
	is_deleted bool NOT NULL DEFAULT false,
	CONSTRAINT order_product_pk PRIMARY KEY (id)
);

ALTER TABLE public.order_product ADD CONSTRAINT order_product_fk FOREIGN KEY (order_id) REFERENCES public.order(id) ON UPDATE CASCADE;