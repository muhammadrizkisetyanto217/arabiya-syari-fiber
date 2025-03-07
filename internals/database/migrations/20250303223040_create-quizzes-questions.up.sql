CREATE TABLE IF NOT EXISTS quizzes_questions (
    id SERIAL PRIMARY KEY,
    question_text VARCHAR(200) NOT NULL,
    question_answer TEXT[] NOT NULL,
    status VARCHAR(10) CHECK (status IN ('active', 'pending', 'archived')) DEFAULT 'pending',
    question_correct VARCHAR(50) NOT NULL,
    paragraph_help TEXT NOT NULL,
    explain_question TEXT NOT NULL,
    answer_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    quizzes_id INT REFERENCES quizzes(id) ON DELETE CASCADE
);