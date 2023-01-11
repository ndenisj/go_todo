
ALTER TABLE IF EXISTS "todos" DROP CONSTRAINT IF EXISTS "todos_user_id_fkey";

DROP TABLE IF EXISTS users cascade;