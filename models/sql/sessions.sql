create table if not exists public.sessions (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT UNIQUE NOT NULL,
  token_hash TEXT UNIQUE NOT NULL
);

