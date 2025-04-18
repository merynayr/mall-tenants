-- +goose Up
-- +goose StatementBegin
CREATE TABLE floor_plans (
    id SERIAL PRIMARY KEY,
    floor INTEGER NOT NULL UNIQUE,
    filename TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS floor_plans;
-- +goose StatementEnd
