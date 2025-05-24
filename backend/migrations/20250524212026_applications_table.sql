-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS applications (
    id SERIAL PRIMARY KEY,
    organization_name TEXT NOT NULL,
    contact_person TEXT NOT NULL,
    address TEXT NOT NULL,
    phone TEXT NOT NULL,
    requisites TEXT NOT NULL,
    email TEXT NOT NULL,
    premise_number TEXT NOT NULL,
    additional_info TEXT,
    is_processed BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS applications;
-- +goose StatementEnd
