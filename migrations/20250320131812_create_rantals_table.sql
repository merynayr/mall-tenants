-- +goose Up
CREATE TABLE rentals (
    rental_id SERIAL PRIMARY KEY,
    space_code INTEGER NOT NULL,
    client_id INTEGER NOT NULL,
    start_date INTEGER NOT NULL,
    end_date INTEGER NOT NULL,
    paid_months INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_space_code FOREIGN KEY (space_code) REFERENCES premises(code) ON DELETE CASCADE,
    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES clients(client_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE rentals;
