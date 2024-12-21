CREATE TYPE "reading_status" AS ENUM (
  'unread',
  'reading',
  'done'
);

CREATE TABLE "users"
(
    "id"              bigserial PRIMARY KEY,
    "name"            varchar        NOT NULL,
    "email"           varchar UNIQUE NOT NULL,
    "hashed_password" varchar        NOT NULL,
    "created_at"      timestamptz    NOT NULL DEFAULT (now()),
    "updated_at"      timestamptz    NOT NULL DEFAULT (now())
);

CREATE TABLE "books"
(
    "id"              bigserial PRIMARY KEY,
    "title"           varchar,
    "description"     varchar,
    "cover_image_url" varchar,
    "url"             varchar,
    "author_id"       bigserial   NOT NULL,
    "publisher_id"    bigserial   NOT NULL,
    "published_date"  timestamptz,
    "isbn"            char(13),
    "created_at"      timestamptz NOT NULL DEFAULT (now()),
    "updated_at"      timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "authors"
(
    "id"         bigserial PRIMARY KEY,
    "name"       varchar,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "publishers"
(
    "id"         bigserial PRIMARY KEY,
    "name"       varchar,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "genres"
(
    "name"       varchar PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "book_genres"
(
    "book_id"    bigserial NOT NULL,
    "genre_name" varchar   NOT NULL,
    PRIMARY KEY ("book_id", "genre_name")
);

CREATE TABLE "reading_histories"
(
    "user_id"    bigserial      NOT NULL,
    "book_id"    bigserial      NOT NULL,
    "status"     reading_status NOT NULL,
    "start_date" date,
    "end_date"   date,
    "created_at" timestamptz    NOT NULL DEFAULT (now()),
    "updated_at" timestamptz    NOT NULL DEFAULT (now()),
    PRIMARY KEY ("user_id", "book_id")
);

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "books" ("title");

CREATE INDEX ON "books" ("author_id");

CREATE INDEX ON "books" ("publisher_id");

CREATE INDEX ON "books" ("published_date");

CREATE INDEX ON "books" ("isbn");

CREATE INDEX ON "authors" ("name");

CREATE INDEX ON "publishers" ("name");

CREATE INDEX ON "reading_histories" ("user_id", "status");

COMMENT
ON TABLE "users" IS 'Stores user data.';

COMMENT
ON TABLE "books" IS 'Stores book data.';

COMMENT
ON TABLE "authors" IS 'Stores author data.';

COMMENT
ON TABLE "publishers" IS 'Stores publisher data.';

COMMENT
ON TABLE "genres" IS 'Stores genre data.';

COMMENT
ON TABLE "book_genres" IS 'Stores book and genre. Normalize using intermediate tables because of the many-to-many relationship between books and genres.';

COMMENT
ON TABLE "reading_histories" IS 'Stores reading history.';

ALTER TABLE "books"
    ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "books"
    ADD FOREIGN KEY ("publisher_id") REFERENCES "publishers" ("id");

ALTER TABLE "book_genres"
    ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "book_genres"
    ADD FOREIGN KEY ("genre_name") REFERENCES "genres" ("name");

ALTER TABLE "reading_histories"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reading_histories"
    ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");
