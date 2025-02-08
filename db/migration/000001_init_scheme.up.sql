CREATE TYPE "reading_status" AS ENUM (
  'unread',
  'reading',
  'done',
  'unknown'
);

CREATE TABLE "users"
(
    "id"              bigserial PRIMARY KEY,
    "name"            varchar(30)         NOT NULL,
    "email"           varchar(320) UNIQUE NOT NULL,
    "hashed_password" varchar             NOT NULL,
    "created_at"      timestamptz         NOT NULL DEFAULT (now()),
    "updated_at"      timestamptz         NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions"
(
    "id"            uuid PRIMARY KEY,
    "user_id"       bigserial    NOT NULL,
    "refresh_token" varchar(512) NOT NULL,
    "expires_at"    timestamptz  NOT NULL,
    "created_at"    timestamptz  NOT NULL DEFAULT (now()),
    "ip_address"    varchar(45),
    "user_agent"    varchar(2048),
    "revoked"       boolean      NOT NULL DEFAULT (false),
    "revoked_at"    timestamptz
);

CREATE TABLE "books"
(
    "id"              bigserial PRIMARY KEY,
    "title"           varchar(255) NOT NULL,
    "description"     varchar(500),
    "cover_image_url" varchar(2048),
    "url"             varchar(2048),
    "author_name"     varchar(255),
    "publisher_name"  varchar(255),
    "published_date"  timestamptz,
    "isbn"            char(13),
    "created_at"      timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"      timestamptz  NOT NULL DEFAULT (now())
);

CREATE TABLE "authors"
(
    "name"       varchar(255) PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "publishers"
(
    "name"       varchar(255) PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "genres"
(
    "name"       varchar(255) PRIMARY KEY,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "book_genres"
(
    "book_id"    bigserial    NOT NULL,
    "genre_name" varchar(255) NOT NULL,
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

CREATE INDEX ON "books" ("author_name");

CREATE INDEX ON "books" ("publisher_name");

CREATE INDEX ON "books" ("published_date");

CREATE INDEX ON "books" ("isbn");

CREATE INDEX ON "reading_histories" ("user_id", "status");

COMMENT
ON TABLE "users" IS 'Stores user data.';

COMMENT
ON TABLE "sessions" IS 'Stores session data.';

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

ALTER TABLE "sessions"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "books"
    ADD FOREIGN KEY ("author_name") REFERENCES "authors" ("name");

ALTER TABLE "books"
    ADD FOREIGN KEY ("publisher_name") REFERENCES "publishers" ("name");

ALTER TABLE "book_genres"
    ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "book_genres"
    ADD FOREIGN KEY ("genre_name") REFERENCES "genres" ("name");

ALTER TABLE "reading_histories"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reading_histories"
    ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");
