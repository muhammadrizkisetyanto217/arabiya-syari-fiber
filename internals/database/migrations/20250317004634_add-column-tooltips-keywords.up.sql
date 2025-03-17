DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns 
                   WHERE table_name='readings' AND column_name='tooltips_keyword') THEN
        ALTER TABLE readings ADD COLUMN tooltips_keyword TEXT[];
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns 
                   WHERE table_name='exams_questions' AND column_name='tooltips_keyword') THEN
        ALTER TABLE exams_questions ADD COLUMN tooltips_keyword TEXT[];
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns 
                   WHERE table_name='quizzes_questions' AND column_name='tooltips_keyword') THEN
        ALTER TABLE quizzes_questions ADD COLUMN tooltips_keyword TEXT[];
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns 
                   WHERE table_name='evaluations_questions' AND column_name='tooltips_keyword') THEN
        ALTER TABLE evaluations_questions ADD COLUMN tooltips_keyword TEXT[];
    END IF;
END $$;
