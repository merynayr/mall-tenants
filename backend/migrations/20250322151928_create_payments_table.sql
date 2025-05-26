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

CREATE INDEX idx_payments_rental_id ON payments(rental_id);
CREATE INDEX idx_payments_period_start ON payments(period_start);
CREATE INDEX idx_payments_is_paid ON payments(is_paid);
CREATE INDEX idx_payments_is_paid_period_start ON payments(is_paid, period_start DESC);

-- +goose Down
DROP TABLE IF EXISTS payments;
