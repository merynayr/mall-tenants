-- +goose Up
CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    rental_id INTEGER NOT NULL REFERENCES rentals(rental_id) ON DELETE CASCADE,
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    amount INTEGER NOT NULL,
    is_paid BOOLEAN NOT NULL DEFAULT FALSE,
    payment_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS payments;
