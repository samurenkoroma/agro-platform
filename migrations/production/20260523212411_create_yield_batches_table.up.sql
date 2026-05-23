CREATE TABLE yield_batches
(
    id               UUID PRIMARY KEY,
    growing_cycle_id UUID        NOT NULL,
    plant_id         UUID        NOT NULL,
    quantity         JSONB       NOT NULL,
    fruit_count      INTEGER     NULL,
    grade            VARCHAR(50) NOT NULL,
    marketable       BOOLEAN     NOT NULL,
    notes            TEXT        NULL,
    harvested_at     TIMESTAMPTZ NOT NULL,
    metadata         JSONB       NOT NULL,
    created_at       TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_yield_cycle FOREIGN KEY (growing_cycle_id) REFERENCES growing_cycles (id),
    CONSTRAINT fk_yield_plant FOREIGN KEY (plant_id) REFERENCES plants (id)
);

CREATE INDEX idx_yield_cycle ON yield_batches (growing_cycle_id);
CREATE INDEX idx_yield_plant ON yield_batches (plant_id);
CREATE INDEX idx_yield_grade ON yield_batches (grade);
CREATE INDEX idx_yield_harvested ON yield_batches (harvested_at);