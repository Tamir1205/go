CREATE TABLE comment
(
    id         SERIAL PRIMARY KEY,
    user_id    int       NOT NULL,
    item_id    int       NOT NULL,
    content    text      NOT NULL,
    parent_id  int,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);