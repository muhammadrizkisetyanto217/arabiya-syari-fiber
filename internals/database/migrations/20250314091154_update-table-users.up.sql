-- **🆙 UP Migration: Ubah dan Tambah Kolom**
ALTER TABLE users
    RENAME COLUMN name TO user_name;

ALTER TABLE users
    ADD COLUMN donation_name VARCHAR(100),
    ADD COLUMN original_name VARCHAR(100);
