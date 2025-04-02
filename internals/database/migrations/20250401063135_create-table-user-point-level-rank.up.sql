CREATE TABLE IF NOT EXISTS user_point_level_rank (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL UNIQUE,
    amount_total_quiz INTEGER DEFAULT 0,
    level_id INTEGER NOT NULL,
    min_point_level INTEGER NOT NULL,
    icon_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
