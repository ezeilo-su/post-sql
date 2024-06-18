-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS posts ALTER COLUMN id DROP DEFAULT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS posts ALTER COLUMN id SET DEFAULT gen_random_uuid();
-- +goose StatementEnd
