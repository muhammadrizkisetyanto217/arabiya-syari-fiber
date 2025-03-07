CREATE TABLE IF NOT EXISTS evaluations_questions (
    id SERIAL PRIMARY KEY,
    question_text VARCHAR(200) NOT NULL,
    question_answer TEXT[] NOT NULL,
    question_correct VARCHAR(50) NOT NULL,
    status VARCHAR(10) CHECK (status IN ('active', 'pending', 'archived')) DEFAULT 'pending',
    paragraph_help TEXT NOT NULL,
    explain_question TEXT NOT NULL,
    answer_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    evaluation_id INT NOT NULL REFERENCES evaluations(id) ON DELETE CASCADE
);
