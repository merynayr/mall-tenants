-- +goose Up
CREATE TABLE contracts (
    contract_id SERIAL PRIMARY KEY,
    rental_id INTEGER NOT NULL,
    file_path TEXT NOT NULL,
    signature BYTEA,
    public_key TEXT,
    is_active BOOLEAN DEFAULT FALSE,
    is_signed BOOLEAN DEFAULT FALSE,
    signed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_rental FOREIGN KEY (rental_id) REFERENCES rentals(rental_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS contracts;
