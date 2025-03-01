DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_name = 'units' 
        AND column_name = 'total_quizzes_section'
    ) THEN
        ALTER TABLE units 
        ADD COLUMN total_quizzes_section INT NOT NULL DEFAULT 0;
    END IF;
END $$;
