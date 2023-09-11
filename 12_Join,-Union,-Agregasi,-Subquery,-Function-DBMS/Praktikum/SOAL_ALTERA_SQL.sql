-- Insert 5 operators pada tabel operators:
INSERT INTO operator (name, created_at) VALUES
    ('Operator 1', NOW()),
    ('Operator 2', NOW()),
    ('Operator 3', NOW()),
    ('Operator 4', NOW()),
    ('Operator 5', NOW());

-- Insert 3 product types:
INSERT INTO product_types (name, created_at) VALUES
    ('baju', NOW()),
    ('celana', NOW()),
    ('sepatu', NOW());

-- Insert 2 product dengan product type id = 1, dan operators id = 3:
INSERT INTO product (product_type_id, operator_id, name, status, created_at) VALUES
    (1, 3, 'baju merah', 1, NOW()),
    (1, 3, 'baju putih', 1, NOW());

-- Insert 3 product dengan product type id = 2, dan operators id = 1:
INSERT INTO product (product_type_id, operator_id, name, status, created_at) VALUES
    (2, 1, 'celana merah', 1, NOW()),
    (2, 1, 'celana putih', 1, NOW()),
    (2, 1, 'celana biru', 1, NOW());

-- Insert 3 product dengan product type id = 3, dan operators id = 4:
INSERT INTO product (product_type_id, operator_id, name, status, created_at) VALUES
    (3, 4, 'tas merah', 1, NOW()),
    (3, 4, 'tas putih', 1, NOW()),
    (3, 4, 'tas biru', 1, NOW());

-- Insert product description pada setiap product:
INSERT INTO product_descriptions (description, created_at) VALUES
    ('baju berwarna merah', NOW()),
    ('baju berwarna putih', NOW()),
    ('celana berwarna merah', NOW()),
    ('celana berwarna putih', NOW()),
    ('celana berwarna biru', NOW()),
    ('tas berwarna merah', NOW()),
    ('tas berwarna putih', NOW()),
    ('tas berwarna biru', NOW());

-- Insert 3 payment methods:
INSERT INTO payment_methods (name, status, created_at) VALUES
    ('dana', 1, NOW()),
    ('link aja', 1, NOW()),
    ('bank', 1, NOW());

-- Insert 5 user pada tabel user:
INSERT INTO users (nama, alamat, status, dob, gender, created_at) VALUES
    ('tono', 'bogor', 1, '2000-01-01', 'M', NOW()),
    ('andini', 'bandung', 1, '1995-05-15', 'F', NOW()),
    ('topan', 'jakarta', 1, '1988-09-20', 'M', NOW()),
    ('sri', 'tangerang', 1, '2002-03-10', 'F', NOW()),
    ('budi', 'depok', 1, '1990-11-30', 'M', NOW());

-- Insert 3 transaksi di masing-masing user:
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price, created_at) VALUES
    -- Transaksi untuk User 1
    (1, 1, 'Completed', 3, 150.00, NOW()),
    (1, 2, 'Pending', 3, 150000.00, NOW()),
    (1, 3, 'Pending', 3, 150000.00, NOW()),

    -- Transaksi untuk User 2
    (2, 1, 'Pending', 3, 150000.00, NOW()),
    (2, 2, 'Completed', 3, 150000.00, NOW()),
    (2, 3, 'Pending', 3, 150000.00, NOW()),

    -- Transaksi untuk User 3
    (3, 1, 'Pending', 3, 150000.00, NOW()),
    (3, 2, 'Pending', 3, 150000.00, NOW()),
    (3, 3, 'Completed', 3, 150000.00, NOW());

-- Insert 3 produk di masing-masing transaksi:
INSERT INTO transactions_details (transactions_id, product_id, status, qty, price, created_at) VALUES
    -- Produk untuk Transaksi User 1, Transaksi 1
    (1, 1, 'Completed', 1, 50000.00, NOW()),
    (1, 2, 'Completed', 1, 50000.00, NOW()),
    (1, 3, 'Completed', 1, 50000.00, NOW()),

    -- Produk untuk Transaksi User 1, Transaksi 2
    (2, 1, 'Pending', 1, 50000.00, NOW()),
    (2, 2, 'Pending', 1, 50000.00, NOW()),
    (2, 3, 'Pending', 1, 50000.00, NOW()),

    -- Produk untuk Transaksi User 1, Transaksi 3
    (3, 1, 'Pending', 1, 50000.00, NOW()),
    (3, 2, 'Completed', 1, 50000.00, NOW()),
    (3, 3, 'Completed', 1, 50000.00, NOW()),

    -- Produk untuk Transaksi User 2, Transaksi 1
    (4, 4, 'Completed', 1, 100.00, NOW()),
    (4, 5, 'Pending', 1, 50.00, NOW()),
    (4, 6, 'Pending', 1, 150.00, NOW()),

    -- Produk untuk Transaksi User 2, Transaksi 2
    (5, 1, 'Completed', 1, 50000.00, NOW()),
    (5, 2, 'Completed', 1, 50000.00, NOW()),
    (5, 3, 'Completed', 1, 50000.00, NOW()),

    -- Produk untuk Transaksi User 2, Transaksi 3
    (6, 4, 'Pending', 1, 50000.00, NOW()),
    (6, 5, 'Pending', 1, 50000.00, NOW()),
    (6, 6, 'Pending', 1, 50000.00, NOW()),

    -- Produk untuk Transaksi User 3, Transaksi 1
    (7, 7, 'Pending', 1, 50000.00, NOW()),
    (7, 8, 'Pending', 1, 50000.00, NOW()),
    (7, 9, 'Completed', 1, 50000.00, NOW()),

-- Produk untuk Transaksi User 3, Transaksi 2
    (8, 7, 'Completed', 1, 50000.00, NOW()),
    (8, 8, 'Completed', 1, 50000.00, NOW()),
    (8, 9, 'Pending', 1, 50000.00, NOW()),

-- Produk untuk Transaksi User 3, Transaksi 3
    (9, 1, 'Completed', 1, 50000.00, NOW()),
    (9, 2, 'Completed', 1, 50000.00, NOW()),
    (9, 3, 'Completed', 1, 50000.00, NOW());



-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT nama, gender
FROM users
WHERE gender = 'M';

-- Tampilkan product dengan id = 3.
SELECT *
FROM product
WHERE id = 3;

-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata 'a'.
SELECT *
FROM users
WHERE created_at >= NOW() - INTERVAL '7 days' AND nama LIKE '%a%';

-- Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*)
FROM users
WHERE gender = 'F';

-- Tampilkan data pelanggan dengan urutan sesuai nama abjad.
SELECT *
FROM users
ORDER BY nama;

-- Tampilkan 5 data pada tabel product.
SELECT *
FROM product
LIMIT 5;
	
-- Ubah data product id 1 dengan nama 'product dummy':
UPDATE product
SET name = 'product dummy'
WHERE id = 1;

-- Update qty = 3 pada transaction detail dengan product id = 1:
UPDATE transactions_details
SET qty = 3
WHERE product_id = 1;

-- Delete data pada tabel product dengan id = 1:
DELETE FROM product
WHERE id = 1;

-- Delete pada tabel product dengan product type id = 1:
DELETE FROM product
WHERE product_type_id = 1;

-- 1. Gabungkan data transaksi dari user id 1 dan user id 2:
SELECT *
FROM transactions
WHERE user_id IN (1, 2);

-- 2. Tampilkan jumlah harga transaksi user id 1:
SELECT user_id, SUM(total_price) AS total_harga
FROM transactions
WHERE user_id = 1
GROUP BY user_id;

-- 3. Tampilkan total transaksi dengan product type 2:
SELECT pt.name AS product_type, SUM(t.price) AS total_harga
FROM transactions_details t
JOIN product p ON t.product_id = p.id
JOIN product_types pt ON p.product_type_id = pt.id
WHERE pt.id = 2
GROUP BY pt.name;

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan:
SELECT p.*, pt.name AS product_type_name
FROM product p
JOIN product_types pt ON p.product_type_id = pt.id;

CREATE OR REPLACE VIEW transaction_with_product_id AS
SELECT t.*, td.product_id
FROM transactions t
JOIN transactions_details td ON t.id = td.transactions_id;

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user:
SELECT t.*, p.name AS product_name, u.nama AS user_name
FROM transaction_with_product_id t
JOIN product p ON t.product_id = p.id
JOIN users u ON t.user_id = u.id;

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud:
CREATE OR REPLACE FUNCTION delete_transaction_with_details(transaction_id INT) RETURNS VOID AS $$
BEGIN
    DELETE FROM transactions_details
    WHERE transactions_id = transaction_id;

    DELETE FROM transactions
    WHERE id = transaction_id;
END;
$$ LANGUAGE plpgsql;

-- 7.Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus:
CREATE OR REPLACE FUNCTION update_total_qty(transaction_id INT) RETURNS VOID AS $$
BEGIN
    UPDATE transactions
    SET total_qty = (
        SELECT SUM(qty)
        FROM transactions_details
        WHERE transactions_id = transaction_id
    )
    WHERE id = transaction_id;
END;
$$ LANGUAGE plpgsql;


-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query:
SELECT *
FROM product
WHERE id NOT IN (
    SELECT DISTINCT product_id
    FROM transactions_details
);
