CREATE TABLE production_allocations
(
    id                 UUID PRIMARY KEY,

    cycle_id           UUID           NOT NULL,
    production_unit_id UUID           NOT NULL,
    area               NUMERIC(12, 6) NOT NULL DEFAULT 0,

    started_at         TIMESTAMPTZ    NULL,
    ended_at           TIMESTAMPTZ    NULL,

    created_at         TIMESTAMPTZ    NOT NULL,
    updated_at         TIMESTAMPTZ    NOT NULL,

    CONSTRAINT fk_growing_cycle_allocation
        FOREIGN KEY (cycle_id) REFERENCES production_growing_cycles (id) ON DELETE CASCADE
);

CREATE INDEX idx_production_allocations_cycle_id ON production_allocations (cycle_id);
CREATE INDEX idx_production_allocations_production_unit_id ON production_allocations (production_unit_id);