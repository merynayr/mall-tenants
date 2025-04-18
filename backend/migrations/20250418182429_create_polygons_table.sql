-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS polygons (
    id SERIAL PRIMARY KEY,             
    premise_code INT NOT NULL,           
    points TEXT NOT NULL,                
    label VARCHAR(255) NOT NULL,         
    CONSTRAINT fk_premise_code FOREIGN KEY (premise_code) REFERENCES premises(code)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS polygons;
-- +goose StatementEnd
