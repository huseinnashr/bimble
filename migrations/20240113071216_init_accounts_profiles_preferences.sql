-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS public.accounts (
    id bigserial PRIMARY KEY,
    email text NOT NULL UNIQUE,
    password text NOT NULL,
    is_verified boolean NOT NULL DEFAULT false
);

CREATE type profile_type AS ENUM ('DATING', 'BFF');
CREATE TABLE IF NOT EXISTS public.profiles (
    id bigserial PRIMARY KEY,
    account_id bigint NOT NULL,
    type profile_type NOT NULL,
    name text,
    hobbies text[],
    about text,
    CONSTRAINT fk_accounts FOREIGN KEY(account_id) REFERENCES accounts(id)
);

CREATE TABLE IF NOT EXISTS public.preferences (
    id bigserial PRIMARY KEY,
    profile_id bigint NOT NULL,
    max_distance int,
    min_age int,
    max_age int,
    CONSTRAINT fk_profiles FOREIGN KEY(profile_id) REFERENCES profiles(id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS public.preferences;

DROP TABLE IF EXISTS public.profiles;
DROP TYPE IF EXISTS profile_type;

DROP TABLE IF EXISTS public.accounts;