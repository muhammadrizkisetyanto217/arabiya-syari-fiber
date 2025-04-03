-- Buat tabel rank_level_requirement
CREATE TABLE IF NOT EXISTS rank_level_requirement (
    id SERIAL PRIMARY KEY,
    name_rank VARCHAR(100) NOT NULL,
    max_level INT NOT NULL,
    icon_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
