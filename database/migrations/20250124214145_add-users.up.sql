CREATE TABLE IF NOT EXISTS "users" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "username" VARCHAR(64) NOT NULL UNIQUE,
    "password" VARCHAR(64) NOT NULL,
    "is_admin" BOOLEAN NOT NULL DEFAULT FALSE
);

DO $$
DECLARE
    "admin_id" UUID;
BEGIN
    INSERT INTO "users" ("id", "username", "password", "is_admin")
    VALUES (gen_random_uuid(), 'adminas', '$2y$14$gFMwRWYfl9v8igp7ekso7eXt085PirlkRXY6Qd6lb6MDT/lAAFusG', TRUE)
    RETURNING "id" INTO "admin_id";

    ALTER TABLE "tasks"
    ADD COLUMN "user_id" UUID;

    ALTER TABLE "tasks"
    ADD CONSTRAINT "fk_tasks_user_id"
    FOREIGN KEY ("user_id")
    REFERENCES "users"("id")
    ON DELETE SET NULL;

    UPDATE "tasks"
    SET "user_id" = "admin_id"
    WHERE "user_id" IS NULL;
END $$;

ALTER TABLE "tasks"
ALTER COLUMN "user_id" SET NOT NULL;