CREATE TABLE accounts
(
    id          VARCHAR(20) PRIMARY KEY,
    private_key VARCHAR(255) NOT NULL UNIQUE,
    chain       VARCHAR(128) NOT NULL,
    is_in_use   BOOLEAN      NOT NULL
);

CREATE TYPE transaction_status AS ENUM (
    'in_queue',
    'processing',
    'completed',
    'failed'
    );

CREATE TABLE transactions
(
    id               BIGSERIAL PRIMARY KEY,
    chain            VARCHAR(255)       NOT NULL,
    "to"             VARCHAR(255)       NOT NULL,
    data             TEXT,
    gas_limit        VARCHAR(255),
    status           transaction_status NOT NULL,
    transaction_hash VARCHAR(255),
    create_timestamp BIGINT             NOT NULL
);
