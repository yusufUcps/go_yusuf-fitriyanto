-- Membuat database alta_online_shop
CREATE DATABASE alta_online_shop;

-- Membuat tabel product_types
CREATE TABLE product_types (
    id serial PRIMARY KEY,
    name varchar(255),
    created_at timestamp,
    updated_at timestamp
);

-- Membuat tabel operator
CREATE TABLE operator (
    id serial PRIMARY KEY,
    name varchar(255),
    created_at timestamp,
    updated_at timestamp
);

-- Membuat tabel product_descriptions
CREATE TABLE product_descriptions (
    id serial PRIMARY KEY,
    description text,
    created_at timestamp,
    updated_at timestamp
);

-- Membuat tabel product
CREATE TABLE product (
    id serial PRIMARY KEY,
    product_type_id int,
    operator_id int,
    name varchar(50),
    status smallint,
    created_at timestamp,
    updated_at timestamp,
	FOREIGN KEY (product_type_id) REFERENCES product_types(id),
	FOREIGN KEY (operator_id) REFERENCES operator(id)
);

-- Membuat tabel transactions_details
CREATE TABLE transactions_details (
    transactions_id int,
    product_id int,
    status varchar(10),
    qty int,
    price numeric(25, 2),
    created_at timestamp,
    updated_at timestamp,
	FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

-- Membuat tabel transactions
CREATE TABLE transactions (
    id serial PRIMARY KEY,
    user_id int,
    payment_method_id int,
    status varchar(10),
    total_qty int,
    total_price numeric(25, 2),
    created_at timestamp,
    updated_at timestamp
);

-- Membuat tabel payment_methods
CREATE TABLE payment_methods (
    id serial PRIMARY KEY,
    name varchar(11),
    status smallint,
    created_at timestamp,
    updated_at timestamp
);

-- Membuat tabel users
CREATE TABLE users (
    id serial PRIMARY KEY,
    status smallint,
    dob date,
    gender char(1),
    created_at timestamp,
    updated_at timestamp
);

-- Membuat tabel kurir dengan field id, name, created_at, updated_at.
CREATE TABLE kurir (
    id serial PRIMARY KEY,
    name varchar(255),
    created_at timestamp,
    updated_at timestamp
);

-- Menambahkan kolom ongkos_dasar di tabel kurir.
ALTER TABLE kurir
ADD COLUMN ongkos_dasar numeric(25, 2);

-- Rename tabel kurir menjadi shipping.
ALTER TABLE kurir
RENAME TO shipping;

-- Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
DROP TABLE shipping;

-- a. 1-to-1: Tabel payment_method_description (Untuk menyimpan deskripsi metode pembayaran):
CREATE TABLE payment_method_description (
    id serial PRIMARY KEY,
    payment_method_id int,
    description text,
    created_at timestamp,
    updated_at timestamp,
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

-- b. 1-to-many: Tabel alamat (Untuk menyimpan alamat pengguna):
CREATE TABLE alamat (
    id serial PRIMARY KEY,
    user_id int,
    alamat text,
    kota varchar(255),
    created_at timestamp,
    updated_at timestamp,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- c. many-to-many: Tabel user_payment_method_detail (Untuk menghubungkan pengguna dengan metode pembayaran:
CREATE TABLE user_payment_method_detail (
    user_id int,
    payment_method_id int,
    created_at timestamp,
    updated_at timestamp,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);
