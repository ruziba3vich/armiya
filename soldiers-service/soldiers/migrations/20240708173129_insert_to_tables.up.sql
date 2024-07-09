CREATE TABLE IF NOT EXISTS objects (
    object_id UUID PRIMARY KEY,
    object_name VARCHAR(64),
    object_surname VARCHAR(64),
    position VARCHAR(20),
    created_at TIMESTAMP,
    deleted_at TIMESTAMP,
    deleted BOOLEAN,
    deleted_by UUID
);

CREATE TABLE IF NOT EXISTS groups (
    group_id UUID PRIMARY KEY,
    group_name VARCHAR(100),
    created_at TIMESTAMP,
    created_by UUID REFERENCES objects(object_id),
    deleted BOOLEAN,
    deleted_at TIMESTAMP,
    deleted_by UUID REFERENCES objects(object_id)
);


CREATE TABLE IF NOT EXISTS soldiers (
    soldier_id UUID PRIMARY KEY,
    soldier_name VARCHAR(64),
    soldier_surname VARCHAR(64),
    birth_date TIMESTAMP,
    join_date DATE,
    leave_date DATE,
    group_id UUID REFERENCES groups(group_id),
    completed BOOLEAN,
    deleted BOOLEAN,
    created_by UUID REFERENCES objects(object_id)
);

CREATE TABLE IF NOT EXISTS trainings (
    training_id UUID PRIMARY KEY,
    trainer_name VARCHAR(64),
    training_description TEXT,
    created_by UUID REFERENCES objects(object_id),
    trainer_id UUID REFERENCES objects(object_id)
);

CREATE TABLE IF NOT EXISTS attendances (
    attendance_id UUID PRIMARY KEY,
    soldier_id UUID REFERENCES soldiers(soldier_id),
    training_id UUID REFERENCES trainings(training_id),
    event_time TIMESTAMP,
    used_ammos INTEGER,
    used_ammo_type VARCHAR(64),
    used_fuels FLOAT,
    deleted BOOLEAN,
    deleted_by UUID REFERENCES objects(object_id),
    created_by UUID REFERENCES objects(object_id)
);
