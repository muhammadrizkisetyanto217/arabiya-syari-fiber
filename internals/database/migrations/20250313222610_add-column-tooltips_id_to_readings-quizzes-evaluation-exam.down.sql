ALTER TABLE readings DROP COLUMN IF EXISTS tooltips_id;
ALTER TABLE exams_questions DROP COLUMN tooltips_id;
ALTER TABLE quizzes_questions DROP COLUMN tooltips_id;
ALTER TABLE evaluations_questions DROP COLUMN tooltips_id;