CREATE TABLE harvest_batches
(
    id                 UUID PRIMARY KEY,
    growing_cycle_id   UUID        NOT NULL,
    production_unit_id UUID        NOT NULL,
    quantity           JSONB       NOT NULL,
    harvested_area     JSONB       NULL,
    grade              VARCHAR(50) NOT NULL,
    marketable         BOOLEAN     NOT NULL,
    notes              TEXT        NULL,
    harvested_at       TIMESTAMPTZ NOT NULL,
    metadata           JSONB       NOT NULL,
    created_at         TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_harvest_cycle FOREIGN KEY (growing_cycle_id) REFERENCES growing_cycles (id),
    CONSTRAINT fk_harvest_unit FOREIGN KEY (production_unit_id) REFERENCES production_units (id)
);

CREATE INDEX idx_harvest_cycle ON harvest_batches (growing_cycle_id);
CREATE INDEX idx_harvest_unit ON harvest_batches (production_unit_id);
CREATE INDEX idx_harvest_grade ON harvest_batches (grade);
CREATE INDEX idx_harvest_date ON harvest_batches (harvested_at);