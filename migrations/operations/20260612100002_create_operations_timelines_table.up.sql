CREATE TABLE operations_timelines
(
    id                  UUID        PRIMARY KEY,

    farm_id             UUID        NOT NULL,
    production_unit_id  UUID        NULL,
    growing_cycle_id    UUID        NULL,

    items               JSONB       NOT NULL DEFAULT '[]',

    created_at          TIMESTAMPTZ NOT NULL,
    updated_at          TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_operations_timelines_farm_id          ON operations_timelines (farm_id);
CREATE INDEX idx_operations_timelines_growing_cycle_id ON operations_timelines (growing_cycle_id);
-- one timeline per (farm, cycle) pair
CREATE UNIQUE INDEX idx_operations_timelines_owner
    ON operations_timelines (farm_id, growing_cycle_id)
    WHERE growing_cycle_id IS NOT NULL;
