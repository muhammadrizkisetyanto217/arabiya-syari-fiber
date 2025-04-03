CREATE TABLE IF NOT EXISTS user_point_level_rank (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL UNIQUE,
    amount_total_quiz INTEGER DEFAULT 0 NOT NULL,
    level_id INTEGER NOT NULL,
    max_point_level INTEGER NOT NULL,
    icon_url_level VARCHAR(255) NOT NULL,
    rank_id INTEGER NOT NULL,
    max_level_rank INTEGER NOT NULL,
    icon_url_rank VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
