

CREATE TABLE "branch" (
    id uuid primary key,
    name varchar not null,
    address varchar,
    phone_number varchar,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);

CREATE TABLE "category" (
    id uuid primary key,
    name varchar not null,
    parent_id uuid references category(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp
);

CREATE TABLE "product" (
    id uuid primary key,
    name varchar not null,
    price numeric not null,
    barcode varchar not null unique,
    category_id uuid references category(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp
);

create table "coming_table" (
    id uuid primary key,
    coming_id varchar not null,
    branch_id uuid references branch(id),
    date_time timestamp,
    status varchar default 'in_process',
    created_at timestamp default current_timestamp,
    updated_at timestamp
);

create table "coming_table_product" (
    id uuid primary key,
    category_id uuid references category(id),
    name varchar not null,
    price numeric not null,
    barcode varchar not null unique,
    count numeric not null default 0,
    total_price numeric default 0,
    coming_table_id uuid references coming_table(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp
);


create table "remaining" (
    id uuid primary key,
    branch_id uuid references branch(id),
    category_id uuid references category(id),
    name varchar not null,
    price numeric not null,
    barcode varchar not null unique,
    count numeric not null default 0,
    total_price numeric default 0,
    created_at timestamp default current_timestamp,
    updated_at timestamp
);

