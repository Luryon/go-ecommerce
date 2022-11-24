CREATE TABLE users (
    id UUID NOT NULL ,
    email VARCHAR(254) NOT NULL ,
    password VARCHAR(72) not null ,
    is_admin BOOL NOT NULL DEFAULT FALSE,
    details JSONB NOT NULL ,
    created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
    updated_at INTEGER,
    CONSTRAINT users_id_pk PRIMARY KEY (id),
    CONSTRAINT users_email_uk UNIQUE (email)
);

COMMENT ON TABLE users IS 'Staorage the admins and costumers for the e-commerce';