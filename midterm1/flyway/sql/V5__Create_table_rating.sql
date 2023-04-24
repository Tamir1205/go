CREATE TABLE rating
(
    id         SERIAL PRIMARY KEY,
    user_id    int       NOT NULL,
    item_id    int       NOT NULL,
    rating     int       NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);