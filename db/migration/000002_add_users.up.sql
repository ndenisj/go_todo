CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "phone" varchar UNIQUE,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "is_admin" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "todos" ADD "user_id" bigint NOT NULL;

ALTER TABLE "todos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "users" ("full_name");

CREATE INDEX ON "users" ("phone");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");