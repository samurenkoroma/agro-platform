CREATE TABLE production_plantings
(
    id         UUID PRIMARY KEY,
    cycle_id   UUID           NOT NULL,
    planted_at TIMESTAMPTZ    NOT NULL,
    quantity   NUMERIC(18, 4) NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ    NOT NULL,
    updated_at TIMESTAMPTZ    NOT NULL,

    CONSTRAINT fk_growing_cycle_planting
        FOREIGN KEY (cycle_id) REFERENCES production_growing_cycles (id) ON DELETE CASCADE
);

CREATE INDEX idx_production_plantings_cycle_id ON production_plantings (cycle_id);