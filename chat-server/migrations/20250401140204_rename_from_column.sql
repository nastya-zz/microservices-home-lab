-- +goose Up
-- +goose StatementBegin
ALTER TABLE message
    RENAME COLUMN "from" TO from_user;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
