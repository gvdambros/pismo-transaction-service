-- Creates the accounts table
CREATE TABLE accounts
(
    id                serial    primary key,
    document_number   text    	not null
);

-- Creates the operation types table
CREATE TABLE operation_types
(
    id                serial    primary key,
    description       text      not null
);

-- Insert values to operation types table
INSERT INTO operation_types (description) VALUES ('COMPRA A VISTA');
INSERT INTO operation_types (description) VALUES ('COMPRA PARCELADA');
INSERT INTO operation_types (description) VALUES ('COMPRA SAQUE');
INSERT INTO operation_types (description) VALUES ('COMPRA PAGAMENTO');

-- Creates the transactions table
CREATE TABLE transactions
(
    id                serial        primary key,
    account_id        integer       references accounts(id),
    operation_type_id integer       references operation_types(id),
    amount            numeric(10,2) not null,
    event_date        timestamp     not null default now()
);