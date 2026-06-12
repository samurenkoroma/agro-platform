CREATE TABLE operations_events
(
    id                  UUID         PRIMARY KEY,

    farm_id             UUID         NOT NULL,
    type                VARCHAR(100) NOT NULL,

    production_unit_id  UUID         NULL,
    growing_cycle_id    UUID         NULL,
    plant_id            UUID         NULL,
    harvest_batch_id    UUID         NULL,

    performed_by        UUID         NULL,

    payload             JSONB        NOT NULL DEFAULT '{}',

    timestamp           TIMESTAMPTZ  NOT NULL
);

CREATE INDEX idx_operations_events_farm_id          ON operations_events (farm_id);
CREATE INDEX idx_operations_events_growing_cycle_id ON operations_events (growing_cycle_id);
CREATE INDEX idx_operations_events_type             ON operations_events (type);
CREATE INDEX idx_operations_events_timestamp        ON operations_events (timestamp DESC);
