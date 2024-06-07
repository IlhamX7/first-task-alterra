CREATE TABLE "Users" (
	id INT PRIMARY KEY,
	name VARCHAR(100),
	email VARCHAR(100),
	password VARCHAR(100),
	address VARCHAR(255),
	phone VARCHAR(20),
	bird_date DATE,
	created_at timestamp default now()
);

CREATE TABLE "Genres" (
	id INT PRIMARY KEY,
	genre_name VARCHAR(10) NOT NULL,
	created_at timestamp default now()
);

CREATE TABLE "Books" (
	id INT PRIMARY KEY,
	genre_id INT,
	title VARCHAR(100) NOT NULL,
	author VARCHAR(20),
	publisher VARCHAR(20),
	publish_year INT,
	created_at timestamp default now(),
	FOREIGN KEY (genre_id) REFERENCES "Genres"(id)
);

CREATE TABLE "Loans" (
	id INT PRIMARY KEY,
	user_id INT,
	book_id INT,
	deadline_date DATE NOT NULL,
	date_loan DATE NOT NULL,
	date_return DATE NOT NULL,
	status_loan BOOL DEFAULT TRUE,
	created_at timestamp default now(),
	FOREIGN KEY (user_id) REFERENCES "Users"(id),
	FOREIGN KEY (book_id) REFERENCES "Books"(id)
);

CREATE TABLE "Admins" (
	id INT PRIMARY KEY,
	name VARCHAR(50),
	email VARCHAR(50),
	password VARCHAR(50),
	phone VARCHAR(20),
	created_at timestamp default now()
);

CREATE TABLE "book_request" (
	id INT PRIMARY KEY,
	user_id INT,
	approved_admin_id INT,
	title VARCHAR(100),
	author VARCHAR(20),
	publisher VARCHAR(100),
	publish_year INT,
	status_request BOOL DEFAULT TRUE,
	created_at timestamp default now(),
	FOREIGN KEY (user_id) REFERENCES "Users"(id),
	FOREIGN KEY (approved_admin_id) REFERENCES "Admins"(id)
);