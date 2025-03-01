CREATE TABLE quizzes (
    id SERIAL PRIMARY KEY,
    name_quizzes VARCHAR(50) UNIQUE NOT NULL,
    status VARCHAR(10) CHECK (status IN ('active', 'pending', 'archived')) DEFAULT 'pending',
    point INT NOT NULL DEFAULT 30,
    total_question INT,
    icon_url VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    section_quizzes_Id INT REFERENCES section_quizzes(id) ON DELETE CASCADE,
    unit_Id INT REFERENCES units(id) ON DELETE CASCADE,
    created_by INT REFERENCES users(id) ON DELETE CASCADE
);
