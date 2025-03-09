ALTER TABLE users
DROP COLUMN IF EXISTS security_question,
DROP COLUMN IF EXISTS security_answer;
