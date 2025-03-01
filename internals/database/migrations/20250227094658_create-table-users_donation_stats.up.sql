CREATE TABLE IF NOT EXISTS users_donation_stats (
    id SERIAL PRIMARY KEY,
    amount INT NOT NULL,
    donatable_type VARCHAR(10) CHECK (donatable_type IN ('quiz', 'article', 'exam', 'other', 'pending')) DEFAULT 'pending',
    prayer VARCHAR(100),
    total_amount INT,
    total_question INT, -- Sesuai dengan total_amount
    diamond INT NOT NULL DEFAULT 30,
    donatable_id TEXT, -- Bisa menyimpan array ID dalam format string dipisahkan koma
    donated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INT REFERENCES users(id) ON DELETE CASCADE
);
