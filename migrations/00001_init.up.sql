CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE rewards
(
    id         UUID PRIMARY KEY   DEFAULT uuid_generate_v4(),
    reward     TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

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
    reward_id  UUID               REFERENCES rewards (id) ON DELETE SET NULL,
    created_at TIMESTAMP          NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_promocodes ON promocodes (code);

CREATE TABLE applied_rewards
(
    id         UUID PRIMARY KEY   DEFAULT uuid_generate_v4(),
    player_id  UUID REFERENCES players (id) ON DELETE CASCADE,
    promo_id   UUID REFERENCES promocodes (id) ON DELETE CASCADE,
    reward_id  UUID      REFERENCES rewards (id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_rewards_player_for_promo ON applied_rewards (player_id, promo_id, reward_id);

INSERT INTO rewards (reward)
VALUES ('reward111'),
       ('222reward'),
       ('reward3333'),
       ('reward4'),
       ('5');

INSERT INTO players (username, email)
VALUES ('raz', 'raz@raz.ru'),
       ('dva', 'dva@dva.ru'),
       ('tri', 'tri@tri.ru'),
       ('chetire', 'chetire@chetire.ru');
