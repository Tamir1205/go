CREATE TABLE item
(
    id          SERIAL PRIMARY KEY,
    name        varchar(255)   NOT NULL,
    price       decimal(10, 2) NOT NULL,
    description text           NOT NULL,
    created_at  TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP
);