CREATE TABLE oauth_access_token
(
    id      SERIAL PRIMARY KEY not null unique,
    user_id BIGINT             not null,
    revoked BOOLEAN DEFAULT false
);

CREATE TABLE oauth_refresh_tokens
(
    id            SERIAL PRIMARY KEY       not null unique,
    user_id       bigint                   not null,
--     access_token_id BIGINT             null,
    refresh_token VARCHAR(255)             NOT NULL,
    revoked       BOOLEAN                           DEFAULT false,
    expires_in    timestamp with time zone NOT NULL,
    created_at    timestamp with time zone NOT NULL DEFAULT now()
);
