BEGIN;

CREATE TABLE IF NOT EXISTS fuel_management (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    quantity DECIMAL(10, 2) NOT NULL,
    last_update TIMESTAMP NOT NULL 
);

CREATE TABLE IF NOT EXISTS fuel_history (
    id UUID PRIMARY KEY,
    fuel_id UUID NOT NULL,
    action VARCHAR(255) NOT NULL,
    actior_id UUID,
    action_timestamp TIMESTAMP NOT NULL,
    FOREIGN KEY (fuel_id) REFERENCES fuel_management(id)
);

COMMIT;