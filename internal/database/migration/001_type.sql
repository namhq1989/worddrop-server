DO $do$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'level') THEN
CREATE TYPE level AS ENUM ('beginner', 'intermediate', 'advanced');
END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'plan') THEN
CREATE TYPE plan AS ENUM ('free', 'pro');
END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'provider') THEN
CREATE TYPE provider AS ENUM ('extension', 'app', 'google');
END IF;
END
$do$;