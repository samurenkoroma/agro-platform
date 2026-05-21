CREATE TABLE production_units
(
    id         UUID PRIMARY KEY,
    farm_id    UUID        NOT NULL,
    parent_id  UUID NULL,
    type       VARCHAR(50) NOT NULL,
    name       TEXT        NOT NULL,
    code       TEXT NULL,
    geometry   JSONB NULL,
    capacity   JSONB NULL,
    metadata   JSONB       NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_production_units_parent ON production_units (parent_id);
CREATE INDEX idx_production_units_farm ON production_units (farm_id);
CREATE INDEX idx_production_units_type ON production_units (type);