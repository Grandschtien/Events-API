CREATE TABLE refresh_tokens (
  id SERIAL PRIMARY KEY,
  user_id      INTEGER     NOT NULL REFERENCES users(id),
  token_hash   TEXT        NOT NULL UNIQUE,
  issued_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
  expires_at   TIMESTAMPTZ NOT NULL,
  revoked      BOOLEAN     NOT NULL DEFAULT FALSE
);
CREATE INDEX ON refresh_tokens(user_id);
