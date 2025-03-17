-- DROP tooltips_keyword jika ada
ALTER TABLE readings DROP COLUMN IF EXISTS tooltips_keyword;
ALTER TABLE exams_questions DROP COLUMN IF EXISTS tooltips_keyword;
ALTER TABLE quizzes_questions DROP COLUMN IF EXISTS tooltips_keyword;
ALTER TABLE evaluations_questions DROP COLUMN IF EXISTS tooltips_keyword;