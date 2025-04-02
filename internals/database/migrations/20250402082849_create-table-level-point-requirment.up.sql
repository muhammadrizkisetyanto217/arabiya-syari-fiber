-- Buat tabel level_point_requirement
CREATE TABLE IF NOT EXISTS level_point_requirement (
    id SERIAL PRIMARY KEY,
    name_level VARCHAR(100) NOT NULL,
    min_point_level INT NOT NULL,
    icon_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);