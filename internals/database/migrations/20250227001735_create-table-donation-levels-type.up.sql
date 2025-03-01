CREATE TABLE donation_levels_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description VARCHAR(250) NOT NULL, -- Menghapus UNIQUE agar lebih fleksibel
    min_amount INT NOT NULL, -- Menggunakan DECIMAL jika mendukung pecahan
    prize_diamond INT DEFAULT 10 NOT NULL,
    icon_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL -- Agar mendukung soft delete
);
