-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA products
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA products
-- +goose StatementEnd
