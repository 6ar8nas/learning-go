ALTER TABLE "tasks"
DROP CONSTRAINT "fk_tasks_user_id";

ALTER TABLE "tasks"
DROP COLUMN "user_id";

DROP TABLE IF EXISTS "users";