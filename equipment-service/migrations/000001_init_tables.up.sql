CREATE TABLE "equipment_history"(
    "id" UUID,
    "equipment_id" UUID,
    "action" VARCHAR(255),
    "actor_id" BIGINT,
    "action_timestamp" TIMESTAMP(0) WITHOUT TIME ZONE
);
ALTER TABLE
    "equipment_history" ADD PRIMARY KEY("id");

    
CREATE TABLE "equipments"(
    "id" UUID,
    "name" VARCHAR(255),
    "description" VARCHAR(255),
    "origin_country" VARCHAR(100),
    "classification" VARCHAR(100),
    "quantity" BIGINT,
    "main_armament" BIGINT,
    "crew_size" INTEGER,
    "weight_kg" DECIMAL(10, 2),
    "length_cm" DECIMAL(10, 2),
    "width_cm" DECIMAL(10, 2),
    "height_cm" DECIMAL(10, 2),
    "max_speed_kmh" INTEGER,
    "operational_range_km" DECIMAL(10, 2),
    "year_of_introduction" INTEGER,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE
    "equipments" ADD PRIMARY KEY("id");
ALTER TABLE
    "equipment_history" ADD CONSTRAINT "equipment_history_equipment_id_foreign" FOREIGN KEY("equipment_id") REFERENCES "equipments"("id");
