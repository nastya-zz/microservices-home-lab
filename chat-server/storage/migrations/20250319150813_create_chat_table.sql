-- +goose Up
-- +goose StatementBegin
CREATE TABLE message (
    "from" VARCHAR(50) not null,
    "to" VARCHAR(50) not null,
    timestamp timestamp
);
CREATE TABLE chat (
    id serial primary key,
    usernames  VARCHAR(50)[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table message;
DROP table chat;
-- +goose StatementEnd
