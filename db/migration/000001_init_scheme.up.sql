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
    "id"            uuid,
    "user_id"       bigserial    NOT NULL,
    "refresh_token" varchar(512) NOT NULL,
    "expires_at"    timestamptz  NOT NULL,
    "created_at"    timestamptz  NOT NULL DEFAULT (now()),
    "ip_address"    varchar(45),
    "user_agent"    varchar(2048),
    "revoked"       boolean      NOT NULL DEFAULT (false),
    "revoked_at"    timestamptz,
    PRIMARY KEY ("id", "created_at")
) PARTITION BY RANGE (created_at);

-- Create sessions partitions (2025/1 - 2026/12)
CREATE TABLE "sessions_y2025m01" PARTITION OF "sessions" FOR VALUES FROM ('2025-01-01') TO ('2025-02-01');
CREATE TABLE "sessions_y2025m02" PARTITION OF "sessions" FOR VALUES FROM ('2025-02-01') TO ('2025-03-01');
CREATE TABLE "sessions_y2025m03" PARTITION OF "sessions" FOR VALUES FROM ('2025-03-01') TO ('2025-04-01');
CREATE TABLE "sessions_y2025m04" PARTITION OF "sessions" FOR VALUES FROM ('2025-04-01') TO ('2025-05-01');
CREATE TABLE "sessions_y2025m05" PARTITION OF "sessions" FOR VALUES FROM ('2025-05-01') TO ('2025-06-01');
CREATE TABLE "sessions_y2025m06" PARTITION OF "sessions" FOR VALUES FROM ('2025-06-01') TO ('2025-07-01');
CREATE TABLE "sessions_y2025m07" PARTITION OF "sessions" FOR VALUES FROM ('2025-07-01') TO ('2025-08-01');
CREATE TABLE "sessions_y2025m08" PARTITION OF "sessions" FOR VALUES FROM ('2025-08-01') TO ('2025-09-01');
CREATE TABLE "sessions_y2025m09" PARTITION OF "sessions" FOR VALUES FROM ('2025-09-01') TO ('2025-10-01');
CREATE TABLE "sessions_y2025m10" PARTITION OF "sessions" FOR VALUES FROM ('2025-10-01') TO ('2025-11-01');
CREATE TABLE "sessions_y2025m11" PARTITION OF "sessions" FOR VALUES FROM ('2025-11-01') TO ('2025-12-01');
CREATE TABLE "sessions_y2025m12" PARTITION OF "sessions" FOR VALUES FROM ('2025-12-01') TO ('2026-01-01');

CREATE TABLE "sessions_y2026m01" PARTITION OF "sessions" FOR VALUES FROM ('2026-01-01') TO ('2026-02-01');
CREATE TABLE "sessions_y2026m02" PARTITION OF "sessions" FOR VALUES FROM ('2026-02-01') TO ('2026-03-01');
CREATE TABLE "sessions_y2026m03" PARTITION OF "sessions" FOR VALUES FROM ('2026-03-01') TO ('2026-04-01');
CREATE TABLE "sessions_y2026m04" PARTITION OF "sessions" FOR VALUES FROM ('2026-04-01') TO ('2026-05-01');
CREATE TABLE "sessions_y2026m05" PARTITION OF "sessions" FOR VALUES FROM ('2026-05-01') TO ('2026-06-01');
CREATE TABLE "sessions_y2026m06" PARTITION OF "sessions" FOR VALUES FROM ('2026-06-01') TO ('2026-07-01');
CREATE TABLE "sessions_y2026m07" PARTITION OF "sessions" FOR VALUES FROM ('2026-07-01') TO ('2026-08-01');
CREATE TABLE "sessions_y2026m08" PARTITION OF "sessions" FOR VALUES FROM ('2026-08-01') TO ('2026-09-01');
CREATE TABLE "sessions_y2026m09" PARTITION OF "sessions" FOR VALUES FROM ('2026-09-01') TO ('2026-10-01');
CREATE TABLE "sessions_y2026m10" PARTITION OF "sessions" FOR VALUES FROM ('2026-10-01') TO ('2026-11-01');
CREATE TABLE "sessions_y2026m11" PARTITION OF "sessions" FOR VALUES FROM ('2026-11-01') TO ('2026-12-01');
CREATE TABLE "sessions_y2026m12" PARTITION OF "sessions" FOR VALUES FROM ('2026-12-01') TO ('2027-01-01');

-- Create default partition for sessions
CREATE TABLE "sessions_default" PARTITION OF "sessions" DEFAULT;

CREATE TABLE "books"
(
    "id"              bigserial,
    "title"           varchar(255) NOT NULL,
    "description"     varchar(500),
    "cover_image_url" varchar(2048),
    "url"             varchar(2048),
    "author_name"     varchar(255),
    "publisher_name"  varchar(255),
    "published_date"  date,
    "isbn"            char(13),
    "created_at"      timestamptz  NOT NULL DEFAULT (now()),
    "updated_at"      timestamptz  NOT NULL DEFAULT (now()),
    PRIMARY KEY ("id")
) PARTITION BY HASH (id);

-- Create books hash partitions (4 partitions)
CREATE TABLE "books_0" PARTITION OF "books" FOR VALUES WITH (modulus 4, remainder 0);
CREATE TABLE "books_1" PARTITION OF "books" FOR VALUES WITH (modulus 4, remainder 1);
CREATE TABLE "books_2" PARTITION OF "books" FOR VALUES WITH (modulus 4, remainder 2);
CREATE TABLE "books_3" PARTITION OF "books" FOR VALUES WITH (modulus 4, remainder 3);

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
    PRIMARY KEY ("user_id", "book_id", "created_at")
) PARTITION BY RANGE (created_at);

-- Create reading_histories partitions (2025/1 - 2026/12)
CREATE TABLE "reading_histories_y2025m01" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-01-01') TO ('2025-02-01');
CREATE TABLE "reading_histories_y2025m02" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-02-01') TO ('2025-03-01');
CREATE TABLE "reading_histories_y2025m03" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-03-01') TO ('2025-04-01');
CREATE TABLE "reading_histories_y2025m04" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-04-01') TO ('2025-05-01');
CREATE TABLE "reading_histories_y2025m05" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-05-01') TO ('2025-06-01');
CREATE TABLE "reading_histories_y2025m06" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-06-01') TO ('2025-07-01');
CREATE TABLE "reading_histories_y2025m07" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-07-01') TO ('2025-08-01');
CREATE TABLE "reading_histories_y2025m08" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-08-01') TO ('2025-09-01');
CREATE TABLE "reading_histories_y2025m09" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-09-01') TO ('2025-10-01');
CREATE TABLE "reading_histories_y2025m10" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-10-01') TO ('2025-11-01');
CREATE TABLE "reading_histories_y2025m11" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-11-01') TO ('2025-12-01');
CREATE TABLE "reading_histories_y2025m12" PARTITION OF "reading_histories" FOR VALUES FROM ('2025-12-01') TO ('2026-01-01');

CREATE TABLE "reading_histories_y2026m01" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-01-01') TO ('2026-02-01');
CREATE TABLE "reading_histories_y2026m02" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-02-01') TO ('2026-03-01');
CREATE TABLE "reading_histories_y2026m03" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-03-01') TO ('2026-04-01');
CREATE TABLE "reading_histories_y2026m04" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-04-01') TO ('2026-05-01');
CREATE TABLE "reading_histories_y2026m05" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-05-01') TO ('2026-06-01');
CREATE TABLE "reading_histories_y2026m06" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-06-01') TO ('2026-07-01');
CREATE TABLE "reading_histories_y2026m07" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-07-01') TO ('2026-08-01');
CREATE TABLE "reading_histories_y2026m08" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-08-01') TO ('2026-09-01');
CREATE TABLE "reading_histories_y2026m09" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-09-01') TO ('2026-10-01');
CREATE TABLE "reading_histories_y2026m10" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-10-01') TO ('2026-11-01');
CREATE TABLE "reading_histories_y2026m11" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-11-01') TO ('2026-12-01');
CREATE TABLE "reading_histories_y2026m12" PARTITION OF "reading_histories" FOR VALUES FROM ('2026-12-01') TO ('2027-01-01');

-- Create default partition for reading_histories
CREATE TABLE "reading_histories_default" PARTITION OF "reading_histories" DEFAULT;

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

-- Foreign key constraints for partitioned tables
-- Note: Foreign keys on partitioned tables require the partition key to be included

ALTER TABLE "sessions"
    ADD CONSTRAINT "fk_sessions_user_id"
    FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "books"
    ADD CONSTRAINT "fk_books_author_name"
    FOREIGN KEY ("author_name") REFERENCES "authors" ("name");

ALTER TABLE "books"
    ADD CONSTRAINT "fk_books_publisher_name"
    FOREIGN KEY ("publisher_name") REFERENCES "publishers" ("name");

ALTER TABLE "book_genres"
    ADD CONSTRAINT "fk_book_genres_book_id"
    FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "book_genres"
    ADD CONSTRAINT "fk_book_genres_genre_name"
    FOREIGN KEY ("genre_name") REFERENCES "genres" ("name");

ALTER TABLE "reading_histories"
    ADD CONSTRAINT "fk_reading_histories_user_id"
    FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "reading_histories"
    ADD CONSTRAINT "fk_reading_histories_book_id"
    FOREIGN KEY ("book_id") REFERENCES "books" ("id");