-- +goose Up
-- +goose StatementBegin
ALTER TABLE posts
ALTER COLUMN id SET DATA TYPE BIGSERIAL PRIMARY KEY;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE posts
ALTER COLUMN id SET DATA TYPE SERIAL PRIMARY KEY;
-- +goose StatementEnd
