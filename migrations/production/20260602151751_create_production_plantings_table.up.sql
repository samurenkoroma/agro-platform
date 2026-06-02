CREATE TABLE production_plantings
(
    id         UUID PRIMARY KEY,
    cycle_id   UUID           NOT NULL,
    planted_at TIMESTAMPTZ    NOT NULL,
    quantity   NUMERIC(18, 4) NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ    NOT NULL,
    updated_at TIMESTAMPTZ    NOT NULL
);

CREATE INDEX idx_production_plantings_cycle_id ON production_plantings (cycle_id);