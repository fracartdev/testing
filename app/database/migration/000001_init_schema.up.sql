CREATE TABLE order_table (
    id text NOT NULL PRIMARY KEY,
    item text,
    "user" text,
    quantity bigint,
    city text,
    department text,
    price bigint
);

CREATE TABLE report_table (
    id text NOT NULL PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
