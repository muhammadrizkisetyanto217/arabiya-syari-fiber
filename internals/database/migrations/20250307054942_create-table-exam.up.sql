CREATE TABLE IF NOT EXISTS exams (
    id SERIAL PRIMARY KEY,
    name_exams VARCHAR(50) NOT NULL,
    status VARCHAR(10) CHECK (status IN ('active', 'pending', 'archived')) DEFAULT 'pending',
    point INT NOT NULL DEFAULT 30,
    total_question INT,
    icon_url VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    unit_id INT REFERENCES units(id) ON DELETE CASCADE, -- ✅ Perbaiki "unit_Id" menjadi "unit_id"
    created_by INT REFERENCES users(id) ON DELETE CASCADE
);