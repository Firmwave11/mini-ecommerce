
-- public.user_account definition

-- Drop table

-- DROP TABLE public.user_account;

CREATE TABLE public.user_account (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    email varchar NOT NULL,
    "password" varchar NOT NULL,
    salt varchar NOT NULL,
    "name" varchar NOT NULL,
    phone varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT false,
    CONSTRAINT user_pkey PRIMARY KEY (id)
);

-- public.authentication_tokens definition

-- Drop table

-- DROP TABLE public.authentication_tokens;

CREATE TABLE public.authentication_tokens (
    token_id uuid NOT NULL DEFAULT gen_random_uuid(),
    user_account_id uuid NOT NULL,
    auth_token varchar NOT NULL,
    generated_at timestamp NULL,
    expires_at timestamp NULL,
    CONSTRAINT authentication_tokens_pkey PRIMARY KEY (token_id)
);


-- public.authentication_tokens foreign keys

ALTER TABLE public.authentication_tokens ADD CONSTRAINT authentication_tokens_fk FOREIGN KEY (user_account_id) REFERENCES public.user_account(id) ON UPDATE CASCADE;

-- public.customer definition

-- Drop table

-- DROP TABLE public.customer;

CREATE TABLE public.customer (
    id serial4 NOT NULL,
    first_name varchar NOT NULL,
    last_name varchar NOT NULL,
    gender varchar NOT NULL,
    birth_date timestamp NOT NULL,
    phone_number varchar NOT NULL,
    nickname varchar NOT NULL,
    description varchar NULL,
    photo varchar NULL,
    updated_at timestamp NOT NULL DEFAULT now(),
    created_at timestamp NOT NULL DEFAULT now(),
    is_deleted bool NOT NULL DEFAULT false,
    user_account_id uuid NOT NULL,
    is_verified bool NOT NULL DEFAULT false,
    CONSTRAINT profile_pk PRIMARY KEY (id)
);


-- public.customer foreign keys

ALTER TABLE public.customer ADD CONSTRAINT customer_fk FOREIGN KEY (user_account_id) REFERENCES public.user_account(id) ON UPDATE CASCADE;

