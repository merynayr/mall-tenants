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
(1, 1, 100, 'retail', 51000, 'available', 'NO', 'YES'),
(2, 1, 65, 'retail', 41000, 'available', 'NO', 'NO'),

(3, 1, 80, 'retail', 42000, 'available', 'YES', 'YES'),
(4, 1, 95, 'retail', 49000, 'available', 'NO', 'YES'),
(5, 1, 70, 'retail', 41000, 'available', 'YES', 'NO'),
(6, 1, 100, 'retail', 53000, 'available', 'YES', 'YES'),
(7, 1, 85, 'retail', 46000, 'available', 'NO', 'NO'),
(8, 1, 110, 'retail', 58000, 'available', 'YES', 'YES'),
(9, 1, 75, 'retail', 44000, 'available', 'NO', 'YES'),
(10, 1, 90, 'retail', 47000, 'available', 'YES', 'NO'),

(11, 1, 65, 'retail', 43000, 'available', 'YES', 'NO'),
(12, 1, 105, 'retail', 55000, 'available', 'NO', 'YES'),
(13, 1, 120, 'retail', 60000, 'available', 'YES', 'YES'),
(14, 1, 95, 'retail', 49000, 'available', 'NO', 'YES'),
(15, 1, 70, 'retail', 42000, 'available', 'YES', 'NO'),
(16, 1, 100, 'retail', 52000, 'available', 'YES', 'YES'),
(17, 1, 85, 'retail', 47000, 'available', 'NO', 'NO'),
(18, 1, 90, 'retail', 48000, 'available', 'YES', 'YES'),
(19, 1, 75, 'retail', 43000, 'available', 'NO', 'YES'),
(20, 1, 110, 'retail', 57000, 'available', 'YES', 'NO'),

(21, 1, 45, 'service', 0, '', 'YES', 'NO'),
(22, 1, 60, 'retail', 40000, 'available', 'YES', 'NO'),
(23, 1, 80, 'retail', 44000, 'available', 'YES', 'YES'),
(24, 1, 95, 'retail', 49000, 'available', 'NO', 'YES'),
(25, 1, 70, 'retail', 42000, 'available', 'YES', 'NO'),
(26, 1, 85, 'retail', 46000, 'available', 'YES', 'YES'),
(27, 1, 120, 'retail', 61000, 'maintenance', 'NO', 'YES'),
(28, 1, 105, 'retail', 56000, 'maintenance', 'YES', 'YES'),
(29, 1, 90, 'retail', 48000, 'available', 'NO', 'NO'),
(30, 1, 75, 'retail', 43000, 'available', 'YES', 'YES'),
(31, 1, 55, 'service', 0, '', 'NO', 'YES');


-- Полигональные данные
INSERT INTO polygons (premise_code, floor, points, label) VALUES
(1, 1, '52,270 93,305 151,305 151,364 137,364 52,342 17,311', '1'),
(2, 1, '92,224 123,252 123,302 93,302 53,268', '2'),
(3, 1, '137,169 168,197 124,249 93,223', '3'),
(4, 1, '179,122 208,148 208,178 185,178 169,195 139,169', '4'),
(5, 1, '218,76 254,109 254,178 210,178 210,147 180,120', '5'),
(6, 1, '256,111 307,156 307,178 256,178', '6'),
(7, 1, '388,117 393,117 431,108 431,178 388,178', '7'),
(8, 1, '433,108 477,98 477,178 433,178', '8'),
(9, 1, '479,97 519,89 519,178 479,178', '9'),
(10, 1, '522,88 560,79 560,178 522,178', '10'),
(11, 1, '562,79 622,66 662,178 562,178', '11'),
(12, 1, '154,257 200,207 211,206 211,275 154,275', '12'),
(13, 1, '213,206 254,206 254,275 213,275', '13'),
(14, 1, '256,206 297,206 297,275 256,275', '14'),
(15, 1, '299,206 362,206 362,275 299,275', '15'),
(16, 1, '388,206 431,206 431,262 388,262', '16'),
(17, 1, '433,206 477,206 477,262 433,262', '17'),
(18, 1, '479,206 519,206 519,262 479,262', '18'),
(19, 1, '521,206 562,206 562,262 522,262', '19'),
(20, 1, '564,206 604,206 604,262 564,262', '20'),
(21, 1, '171,305 211,305 211,337 171,337', '21'),
(22, 1, '213,305 254,305 254,364 213,305', '22'),
(23, 1, '256,305 297,305 297,364 257,364', '23'),
(24, 1, '300,305 341,305 341,364 300,364', '24'),
(25, 1, '343,305 392,305 392,364 343,364', '25'),
(26, 1, '394,292 431,292 431,364 394,364', '26'),
(27, 1, '433,292 477,292 477,364 433,364', '27'),
(28, 1, '479,292 520,292 520,364 479,364', '28'),
(29, 1, '522,292 562,292 562,364 522,364', '29'),
(30, 1, '564,292 604,292 604,364 564,364', '30'),
(31, 1, '606,292 647,292 647,364 606,364', '31');

-- +goose StatementEnd
