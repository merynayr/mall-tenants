-- +goose Up
-- +goose StatementBegin

-- Создание таблицы пользователей (авторизация)
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role VARCHAR(9) NOT NULL
);

-- Создание таблицы клиентов
CREATE TABLE clients (
    client_id INTEGER PRIMARY KEY, 
    organization_name VARCHAR(255),
    contact_person VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    phone BIGINT NOT NULL,
    requisites TEXT NOT NULL,
    CONSTRAINT fk_clients_user FOREIGN KEY (client_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Индексы для ускорения поиска
CREATE INDEX idx_clients_organization_name ON clients (organization_name);
CREATE INDEX idx_clients_phone ON clients (phone);
CREATE INDEX idx_users_email ON users (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clients;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
