-- +goose Up
-- +goose StatementBegin
ALTER TABLE book
    ADD COLUMN cover VARCHAR(128) DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE book
    DROP COLUMN cover;
-- +goose StatementEnd
