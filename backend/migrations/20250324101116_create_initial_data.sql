-- +goose Up
-- +goose StatementBegin

-- Создание начальных данных

-- Вставка пользователей
INSERT INTO users (email, password_hash, role) VALUES
('john.doe@example.com', '$2a$10$x3eW4dqNo8cMqsatIVIBH.ev6SxEOZlktcUoZJjKgNC6Qz5w0eaF.', 'client'),
('jane.smith@example.com', '$2a$10$HdBVBZvsv9UaxuXh9Rb9oOpY119DmkZ8ojmVF37Cl8EZXN/f7vXcW', 'client'),
('alice.brown@example.com', '$2a$10$sXEo6pqPra1UTAOEtxSV4uQqXKIJvrFJY.N2RmYRsj4CYafcnBvpa', 'client'),
('1@mail.ru', '$2a$10$0TFVc4ubbYMUYOo.OpDnk.dcVDc.CdEZm5JhL.z.RvzsA0N.UlbpK', 'moderator'),
('2@mail.ru', '$2a$10$0TFVc4ubbYMUYOo.OpDnk.dcVDc.CdEZm5JhL.z.RvzsA0N.UlbpK', 'client'),
('director@mail.ru', '$2a$10$OOPYJr4wm6N2kewOz9r8s.b2M76rQVDxJmRm7OG8DjL4DSU..itLK', 'director'),
('moderator@mail.ru', '$2a$10$UM52Ej3oGkWCiLLy.A45TOuHnG4Q95eKmO0DQvR0ZCw28oO7RXzvW', 'moderator');

-- Вставка клиентов с использованием последних вставленных пользователей
INSERT INTO clients (organization_name, contact_person, address, phone, requisites, client_id) VALUES
('ABC Ltd.', 'John Doe', '123 Main St', '+11234567890', 'Bank details', (SELECT user_id FROM users WHERE email = 'john.doe@example.com')),
('XYZ Corp.', 'Jane Smith', '456 Elm St', '+19876543210', 'Payment info', (SELECT user_id FROM users WHERE email = 'jane.smith@example.com')),
('Tech Solutions', 'Alice Brown', '789 Oak St', '+12125550123', 'Tax ID: 123456789', (SELECT user_id FROM users WHERE email = 'alice.brown@example.com'));


-- Вставка помещений
INSERT INTO premises (code, floor, area, type, rent_per_month, status, security_system, air_conditioning) VALUES
(1, 1, 50, 'service', 0, '', 'YES', 'NO'),
(2, 1, 40, 'service', 0, '', 'NO', 'YES'),
(3, 1, 100, 'retail', 50000, 'available', 'YES', 'YES'),
(4, 1, 120, 'retail', 60000, 'available', 'NO', 'YES'),
(5, 1, 80, 'retail', 45000, 'available', 'YES', 'NO'),
(6, 1, 90, 'retail', 47000, 'available', 'NO', 'YES'),
(7, 1, 110, 'retail', 55000, 'maintenance', 'YES', 'YES');


-- +goose StatementEnd
