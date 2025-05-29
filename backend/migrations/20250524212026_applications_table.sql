-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS applications (
    id SERIAL PRIMARY KEY,
    organization_name TEXT,
    contact_person TEXT,
    address TEXT,
    phone TEXT,
    requisites TEXT,
    email TEXT NOT NULL,
    premise_number INT NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    additional_info TEXT,
    is_processed BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS applications;
-- +goose StatementEnd
