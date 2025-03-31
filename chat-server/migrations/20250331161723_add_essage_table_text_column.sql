-- +goose Up
-- +goose StatementBegin
ALTER TABLE message
ADD COLUMN id SERIAL PRIMARY KEY,
ADD COLUMN text TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE message
DROP COLUMN IF EXISTS id,
DROP COLUMN IF EXISTS text;
-- +goose StatementEnd
