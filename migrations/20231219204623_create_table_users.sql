-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.users;
-- +goose StatementEnd
