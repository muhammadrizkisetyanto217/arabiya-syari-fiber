-- **🔽 DOWN Migration: Kembalikan ke Struktur Sebelumnya**
ALTER TABLE users
    RENAME COLUMN user_name TO name;

ALTER TABLE users
    DROP COLUMN IF EXISTS donation_name,
    DROP COLUMN IF EXISTS original_name;
