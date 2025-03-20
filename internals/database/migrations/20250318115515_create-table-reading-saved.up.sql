CREATE TABLE reading_saved (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    reading_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_reading FOREIGN KEY (reading_id) REFERENCES readings(id) ON DELETE CASCADE,
    UNIQUE (user_id, reading_id)
);
