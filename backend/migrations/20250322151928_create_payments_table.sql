-- +goose Up
CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    rent_id INTEGER NOT NULL,
    payment_date TIMESTAMP DEFAULT NOW(),
    amount INTEGER NOT NULL,
    CONSTRAINT fk_rent FOREIGN KEY (rent_id) REFERENCES rentals(rental_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS payments;
