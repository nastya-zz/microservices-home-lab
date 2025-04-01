-- +goose Up
-- +goose StatementBegin
ALTER TABLE message
DROP COLUMN IF EXISTS "to";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE message
ADD COLUMN "to" TEXT;
-- +goose StatementEnd
