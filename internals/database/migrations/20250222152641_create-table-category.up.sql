CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description_short VARCHAR(100),
    description_long VARCHAR(2000),
    total_subcategories INT,
    status VARCHAR(10) CHECK (status IN ('active', 'pending', 'archived')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    difficulty_id INT REFERENCES difficulties(id)
)
