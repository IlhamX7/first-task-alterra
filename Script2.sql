INSERT INTO "Genres" (genre_name) VALUES 
('horor'),
('comedy'),
('drama'),
('thriller'),
('action'),
('romance'),
('mystery');

INSERT INTO "Books"(genre_id, title, author, publisher, publish_year)
VALUES (7, 'Harry Potter', 'JK. Rowling', 'Gramedia', 2010);
INSERT INTO "Books"(genre_id, title, author, publisher, publish_year)
VALUES (6, 'Laskar Pelangi', 'Andrea Hirata', 'Gramedia', 2005);
INSERT INTO "Books"(genre_id, title, author, publisher, publish_year)
VALUES (4, 'The Chronicles of Narnia', 'Clive Staples Lewis', 'Gramedia', 1950);
INSERT INTO "Books"(genre_id, title, author, publisher, publish_year)
VALUES (3, 'Ayat-Ayat Cinta', 'Habiburrahman El Shirazy', 'Gramedia', 2003);
INSERT INTO "Books"(genre_id, title, author, publisher, publish_year)
VALUES (1, 'Kisah Tanah Jawa', 'Bonaventura D. Genta', 'GagasMedia', 2018);
INSERT INTO "Books"(genre_id, title, author, publisher, publish_year)
VALUES (2, 'Koala Kumal', 'Raditya Dika', 'GagasMedia', 2017);

INSERT INTO "Users"(name, email, "password", address, phone, bird_date) VALUES
('Jerry', 'jerry@alterra.id', 'p12345', 'Jl. Hayati Mahim no 11', '0819432144', '1991-09-21');
INSERT INTO "Users"(name, email, "password", address, phone, bird_date) VALUES
('Ilham', 'ilham@gmail.com', 'p6789', 'Jl. Jend. Ahmad Yani no 25', '081949565427', '1997-07-20');
INSERT INTO "Users"(name, email, "password", address, phone, bird_date) VALUES
('Hafiz', 'hafiz@gmail.co.id', 'p101112', 'Jl. Kebon Jeruk no 09', '08195566321', '1999-01-11');
INSERT INTO "Users"(name, email, "password", address, phone, bird_date) VALUES
('Farhan', 'farhan@gmail.co.id', 'p131415', 'Jl. Dukuh Atas no 15', '081969778811', '2000-11-02');
INSERT INTO "Users"(name, email, "password", address, phone, bird_date) VALUES
('Edi Saputra', 'edi.saputra@gmail.co.id', 'p161718', 'Jl. Sriwijaya Kasih no 66', '081944321112', '1994-05-15');

SELECT * FROM "Books" WHERE title like 'Harry Potter%';

SELECT * FROM "Books" WHERE CAST(genre_id AS TEXT) LIKE '1%';

INSERT INTO "Loans"(user_id, book_id, deadline_date, date_loan, date_return) VALUES
(1, 1, '2024-07-12', '2023-05-11', '2024-01-02');
INSERT INTO "Loans"(user_id, book_id, deadline_date, date_loan, date_return) VALUES
(1, 2, '2024-06-20', '2023-08-19', '2024-02-03');
INSERT INTO "Loans"(user_id, book_id, deadline_date, date_loan, date_return) VALUES
(1, 6, '2024-05-19', '2023-10-15', '2024-03-04');
INSERT INTO "Loans"(user_id, book_id, deadline_date, date_loan, date_return) VALUES
(2, 3, '2024-07-08', '2023-01-02', '2024-04-11');
INSERT INTO "Loans"(user_id, book_id, deadline_date, date_loan, date_return) VALUES
(3, 1, '2024-08-22', '2023-02-03', '2024-06-07');

UPDATE "Loans" SET status_loan = FALSE WHERE user_id = 1 AND book_id = 1;

SELECT u."name" , b.title , l."status_loan" 
FROM "Users" u
JOIN "Loans" l on l.user_id = u.id 
JOIN "Books" b on b.id = l.book_id
WHERE l.status_loan = TRUE;

SELECT u."name" , b.title , l."status_loan" 
FROM "Users" u
JOIN "Loans" l on l.user_id = u.id 
JOIN "Books" b on b.id = l.book_id
WHERE l.status_loan = FALSE;
