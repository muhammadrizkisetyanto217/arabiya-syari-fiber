CREATE TABLE themes_or_levels (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    status VARCHAR(10) CHECK (status IN ('active', 'pending', 'archived')) DEFAULT 'pending',
    description_short VARCHAR(100),
    description_long VARCHAR(2000),
    total_unit INT,
    update_news JSONB,
    image_url VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    subcategories_id INT REFERENCES subcategories(id)
)