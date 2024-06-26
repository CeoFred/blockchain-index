
CREATE TABLE faucets (
    id UUID PRIMARY KEY,
    asset VARCHAR(255) NOT NULL,
    address VARCHAR(42) NOT NULL,
    start_block INTEGER NOT NULL,
    end_block INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    UNIQUE (address)
);

CREATE TABLE faucet_transfers (
    id UUID PRIMARY KEY,
    faucet_id UUID NOT NULL REFERENCES faucets(id) ON DELETE CASCADE,
    recipient VARCHAR(42) NOT NULL,
    value NUMERIC NOT NULL,
    transaction_hash VARCHAR(66) NOT NULL,
    block_number INTEGER NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (transaction_hash, block_number)
);
