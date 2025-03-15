-- Create enum types
CREATE TYPE level AS ENUM ('beginner', 'intermediate', 'advanced', 'expert');
CREATE TYPE plan AS ENUM ('free', 'pro');
CREATE TYPE provider AS ENUM ('extension', 'app', 'google');

-- Create words table
CREATE TABLE words (
                       id VARCHAR(255) PRIMARY KEY,
                       word VARCHAR(255) NOT NULL,
                       level level NOT NULL,
                       parts_of_speech TEXT[] NOT NULL,
                       ipa VARCHAR(255),
                       audio VARCHAR(255),
                       created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

-- Create an index on the word column for faster lookups
CREATE INDEX idx_words_word ON words(word);

-- Create word_definitions table
CREATE TABLE word_definitions (
                                  id SERIAL PRIMARY KEY,
                                  word_id VARCHAR(255) NOT NULL REFERENCES words(id) ON DELETE CASCADE,
                                  pos VARCHAR(50) NOT NULL,  -- part of speech
                                  definition TEXT NOT NULL
);

-- Create index for faster word_id lookups
CREATE INDEX idx_word_definitions_word_id ON word_definitions(word_id);

-- Create word_examples table
CREATE TABLE word_examples (
                               id VARCHAR(255) PRIMARY KEY,
                               word_id VARCHAR(255) NOT NULL REFERENCES words(id) ON DELETE CASCADE,
                               example TEXT NOT NULL,
                               audio VARCHAR(255),
                               main_word VARCHAR(255) NOT NULL,
                               level level NOT NULL,
                               created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

-- Create index for faster word_id lookups
CREATE INDEX idx_word_examples_word_id ON word_examples(word_id);

-- Create word_news table
CREATE TABLE word_news (
                           id VARCHAR(255) PRIMARY KEY,
                           word_id VARCHAR(255) NOT NULL REFERENCES words(id) ON DELETE CASCADE,
                           categories TEXT[] NOT NULL,
                           source_url TEXT NOT NULL,
                           title TEXT NOT NULL,
                           summary TEXT NOT NULL,
                           created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster word_id lookups
CREATE INDEX idx_word_news_word_id ON word_news(word_id);

-- Create users table
CREATE TABLE users (
                       id VARCHAR(255) PRIMARY KEY,
                       belongs_to VARCHAR(255) REFERENCES users(id) ON DELETE SET NULL,
                       name VARCHAR(255) NOT NULL,
                       subscription JSONB NOT NULL, -- JSON Format: {"plan": "free|pro", "expiry": "ISO8601 date string", "customerId": "string"}
                       auth_providers JSONB NOT NULL, -- JSON Format: [{"provider": "extension|app|google", "id": "string", "name": "string", "email": "string"}]
                       created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
                       updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

-- Create index for faster belongs_to lookups
CREATE INDEX idx_users_belongs_to ON users(belongs_to);

-- Create index for subscription plan queries using JSONB operators
CREATE INDEX idx_users_subscription_plan ON users((subscription->>'plan'));