CREATE TABLE "Users" (
	id INT PRIMARY KEY,
	name VARCHAR(100),
	email VARCHAR(100),
	password VARCHAR(100),
	address VARCHAR(255),
	phone VARCHAR(20),
	bird_date VARCHAR(100)
);

CREATE TABLE "Genres" (
	id INT PRIMARY KEY,
	genre_name VARCHAR(10)
);

CREATE TABLE "Books" (
	id INT PRIMARY KEY,
	genre_id INT,
	title VARCHAR(100),
	author VARCHAR(20),
	publisher VARCHAR(20),
	publish_year INT,
	FOREIGN KEY (genre_id) REFERENCES "Genres"(id)
);

CREATE TABLE "Loans" (
	id INT PRIMARY KEY,
	user_id INT,
	book_id INT,
	deadline_date TIMESTAMP,
	date_loan TIMESTAMP,
	date_return TIMESTAMP,
	status_loan VARCHAR(20),
	FOREIGN KEY (user_id) REFERENCES "Users"(id),
	FOREIGN KEY (book_id) REFERENCES "Books"(id)
);

CREATE TABLE "Admins" (
	id INT PRIMARY KEY,
	name VARCHAR(50),
	email VARCHAR(50),
	password VARCHAR(50),
	phone VARCHAR(20)
);

CREATE TABLE "book_request" (
	id INT PRIMARY KEY,
	user_id INT,
	approved_admin_id INT,
	title VARCHAR(100),
	author VARCHAR(20),
	publisher VARCHAR(100),
	publish_year INT,
	status_request VARCHAR(20),
	FOREIGN KEY (user_id) REFERENCES "Users"(id),
	FOREIGN KEY (approved_admin_id) REFERENCES "Admins"(id)
);




