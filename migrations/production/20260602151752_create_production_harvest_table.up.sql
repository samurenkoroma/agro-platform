CREATE TABLE production_harvest_batch
(
    id           UUID PRIMARY KEY,
    cycle_id     UUID           NOT NULL,
    harvested_at TIMESTAMPTZ    NOT NULL,
    quantity     NUMERIC(18, 4) NOT NULL DEFAULT 0,
    created_at   TIMESTAMPTZ    NOT NULL,
    updated_at   TIMESTAMPTZ    NOT NULL,

    CONSTRAINT fk_growing_cycle_harvest
        FOREIGN KEY (cycle_id) REFERENCES production_growing_cycles (id) ON DELETE CASCADE
);

CREATE INDEX idx_production_harvest_batch_cycle_id ON production_harvest_batch (cycle_id);