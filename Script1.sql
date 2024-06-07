CREATE TABLE "Users" (
	id SERIAL PRIMARY KEY,
	name VARCHAR(200),
	email VARCHAR(200),
	password VARCHAR(200),
	address VARCHAR(255),
	phone VARCHAR(200),
	bird_date DATE,
	created_at timestamp default now()
);

CREATE TABLE "Genres" (
	id SERIAL PRIMARY KEY,
	genre_name VARCHAR(200) NOT NULL,
	created_at timestamp default now()
);

CREATE TABLE "Books" (
	id SERIAL PRIMARY KEY,
	genre_id INT,
	title VARCHAR(200) NOT NULL,
	author VARCHAR(200),
	publisher VARCHAR(200),
	publish_year INT,
	created_at timestamp default now(),
	FOREIGN KEY (genre_id) REFERENCES "Genres"(id)
);

CREATE TABLE "Loans" (
	id SERIAL PRIMARY KEY,
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
	id SERIAL PRIMARY KEY,
	name VARCHAR(200),
	email VARCHAR(200),
	password VARCHAR(200),
	phone VARCHAR(200),
	created_at timestamp default now()
);

CREATE TABLE "book_request" (
	id SERIAL PRIMARY KEY,
	user_id INT,
	approved_admin_id INT,
	title VARCHAR(200),
	author VARCHAR(200),
	publisher VARCHAR(200),
	publish_year INT,
	status_request BOOL DEFAULT TRUE,
	created_at timestamp default now(),
	FOREIGN KEY (user_id) REFERENCES "Users"(id),
	FOREIGN KEY (approved_admin_id) REFERENCES "Admins"(id)
);