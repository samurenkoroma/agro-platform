CREATE TABLE growing_cycles
(
    id                 UUID PRIMARY KEY,
    farm_id            UUID        NOT NULL,
    crop_id            UUID        NOT NULL,
    production_unit_id UUID        NOT NULL,
    method             VARCHAR(50) NOT NULL,
    status             VARCHAR(50) NOT NULL,
    metadata           JSONB       NOT NULL,
    created_at         TIMESTAMPTZ NOT NULL,
    updated_at         TIMESTAMPTZ NOT NULL,
    archived_at        TIMESTAMPTZ NULL,

    CONSTRAINT fk_cycle_crop FOREIGN KEY (crop_id) REFERENCES crops (id),
    CONSTRAINT fk_cycle_unit FOREIGN KEY (production_unit_id) REFERENCES production_units (id)
);

CREATE INDEX idx_cycle_farm ON growing_cycles (farm_id);
CREATE INDEX idx_cycle_crop ON growing_cycles (crop_id);
CREATE INDEX idx_cycle_unit ON growing_cycles (production_unit_id);
CREATE INDEX idx_cycle_status ON growing_cycles (status);