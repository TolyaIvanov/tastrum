CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE players
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username   VARCHAR(100) UNIQUE NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP        DEFAULT NOW()
);

CREATE INDEX idx_players_email ON players (email);

CREATE TABLE promocodes
(
    id         UUID PRIMARY KEY            DEFAULT uuid_generate_v4(),
    code       VARCHAR(50) UNIQUE NOT NULL,
    max_uses   INT                NOT NULL,
    uses_count INT                         DEFAULT 0,
    created_at TIMESTAMP          NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_promocodes ON promocodes (code);

CREATE TABLE rewards
(
    id         UUID PRIMARY KEY   DEFAULT uuid_generate_v4(),
    player_id  UUID      NOT NULL,
    promo_id   UUID REFERENCES promocodes (id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_rewards_player_for_promo ON rewards (player_id, promo_id);
