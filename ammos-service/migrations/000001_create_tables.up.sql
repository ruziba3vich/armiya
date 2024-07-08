BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS ammos (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    caliber VARCHAR(255) NOT NULL,
    description TEXT,
    type VARCHAR(255) NOT NULL,
    quantity BIGINT NOT NULL,
    last_update TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS ammo_history (
    id UUID PRIMARY KEY,
    ammo_id UUID NOT NULL,
    action VARCHAR(255) NOT NULL,
    actor_id UUID NOT NULL,
    action_timestamp TIMESTAMP NOT NULL,
    FOREIGN KEY (ammo_id) REFERENCES ammos(id)
);

COMMIT;
