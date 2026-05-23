CREATE TABLE plants
(
    id                 UUID PRIMARY KEY,
    growing_cycle_id   UUID        NOT NULL,
    crop_id            UUID        NOT NULL,
    variety_id         UUID        NULL,
    production_unit_id UUID        NOT NULL,
    slot_id            UUID        NULL,
    substrate_id       UUID        NULL,
    status             VARCHAR(50) NOT NULL,
    health             JSONB       NOT NULL,
    current_stage_id   UUID        NULL,
    planted_at         TIMESTAMPTZ NOT NULL,
    transplanted_at    TIMESTAMPTZ NULL,
    harvested_at       TIMESTAMPTZ NULL,
    discarded_at       TIMESTAMPTZ NULL,
    metadata           JSONB       NOT NULL,
    created_at         TIMESTAMPTZ NOT NULL,
    updated_at         TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_plant_cycle FOREIGN KEY (growing_cycle_id) REFERENCES growing_cycles (id),
    CONSTRAINT fk_plant_crop FOREIGN KEY (crop_id) REFERENCES crops (id),
    CONSTRAINT fk_plant_variety FOREIGN KEY (variety_id) REFERENCES varieties (id),
    CONSTRAINT fk_plant_stage FOREIGN KEY (current_stage_id) REFERENCES crop_stages (id),
    CONSTRAINT fk_plant_unit FOREIGN KEY (production_unit_id) REFERENCES production_units (id)
);

CREATE INDEX idx_plant_cycle ON plants (growing_cycle_id);
CREATE INDEX idx_plant_crop ON plants (crop_id);
CREATE INDEX idx_plant_variety ON plants (variety_id);
CREATE INDEX idx_plant_unit ON plants (production_unit_id);
CREATE INDEX idx_plant_stage ON plants (current_stage_id);
CREATE INDEX idx_plant_status ON plants (status);