-- +goose Up
-- +goose StatementBegin
ALTER TABLE author
    ADD COLUMN avatar VARCHAR(128) DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE author
    DROP COLUMN avatar;
-- +goose StatementEnd
