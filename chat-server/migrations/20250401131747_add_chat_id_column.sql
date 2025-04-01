-- +goose Up
-- +goose StatementBegin
ALTER TABLE message
ADD COLUMN chat_id INTEGER REFERENCES chat(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE message
DROP COLUMN IF EXISTS chat_id,
-- +goose StatementEnd
