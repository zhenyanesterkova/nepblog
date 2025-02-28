BEGIN TRANSACTION;

CREATE TABLE posts
(
    created_at          timestamptz NOT NULL,
    id                  uuid PRIMARY KEY,
    user_id             uuid        NOT NULL,
    title               text        NOT NULL,
    content             text NOT NULL
    allowed_comments    boolean NOT NULL,
);

CREATE TABLE comments
(
    post_id     uuid NOT NULL,
    user_id     uuid NOT NULL,
    created_at  timestamptz NOT NULL,
    id          uuid PRIMARY KEY,
    comment     text NOT NULL

    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE
);

CREATE INDEX post_id ON posts (id);
CREATE INDEX comment_id ON comments (id);

COMMIT;