-- +goose Up
-- +goose StatementBegin
CREATE TABLE premises (
    code INTEGER PRIMARY KEY,
    floor INTEGER NOT NULL,
    area INTEGER NOT NULL,
    type TEXT NOT NULL CHECK (type IN ('retail', 'service')),
    rent_per_month INTEGER CHECK (
        (type = 'service' AND (rent_per_month IS NULL OR rent_per_month = 0)) OR 
        (type != 'service' AND rent_per_month > 0)
    ),
    status TEXT CHECK (
        (type = 'service' AND (status IS NULL OR status = '')) OR 
        (type != 'service' AND status IN ('available', 'occupied', 'maintenance'))
    ),
    security_system TEXT,
    air_conditioning TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE premises;
-- +goose StatementEnd
