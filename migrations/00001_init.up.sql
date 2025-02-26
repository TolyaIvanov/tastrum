CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE promocodes
(
    id           UUID PRIMARY KEY     DEFAULT uuid_generate_v4(),
    code         TEXT UNIQUE NOT NULL,
    max_uses     INT         NOT NULL,
    current_uses INT                  DEFAULT 0,
    reward       TEXT        NOT NULL,
    created_at   TIMESTAMP   NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_promocodes ON promocodes (code);

CREATE TABLE rewards
(
    id          UUID PRIMARY KEY   DEFAULT uuid_generate_v4(),
    player_id   UUID      NOT NULL,
    promo_id    UUID REFERENCES promocodes (id),
    received_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_rewards_player_for_promo ON rewards (player_id, promo_id);
