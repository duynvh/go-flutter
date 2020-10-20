-- -------------------------------------------------------------
-- TablePlus 3.5.0(308)
--
-- https://tableplus.com/
--
-- Database: go-flutter
-- Generation Time: 2020-10-20 20:08:48.7430
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."bookmarks" (
    "bid" text NOT NULL,
    "user_id" text,
    "repo_name" text,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    PRIMARY KEY ("bid")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."migrations" (
    "id" text NOT NULL,
    "applied_at" timestamptz,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."repos" (
    "name" text NOT NULL,
    "description" text,
    "url" text,
    "color" text,
    "lang" text,
    "fork" text,
    "stars" text,
    "stars_today" text,
    "build_by" text,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    PRIMARY KEY ("name")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."users" (
    "user_id" text NOT NULL,
    "full_name" text,
    "email" text,
    "password" text,
    "role" text,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL,
    PRIMARY KEY ("user_id")
);

INSERT INTO "public"."migrations" ("id", "applied_at") VALUES
('1_init.sql', '2020-10-17 00:23:52.464013+07');

INSERT INTO "public"."users" ("user_id", "full_name", "email", "password", "role", "created_at", "updated_at") VALUES
('162ed222-122c-11eb-aa5a-f40f24211069', 'Duy Nguyen', 'admin@gmail.com', '$2a$04$X1iGv3JZmqxTJwj0f6eN4efrq4PmLNivO8DgmmKdi5dPrx8Lhum8m', 'MEMBER', '2020-10-19 23:56:52.496134+07', '2020-10-20 00:11:07.103564+07');

ALTER TABLE "public"."bookmarks" ADD FOREIGN KEY ("repo_name") REFERENCES "public"."repos"("name");
