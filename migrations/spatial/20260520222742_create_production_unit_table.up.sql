CREATE TABLE production_units
(
    id          UUID PRIMARY KEY,
    owner_id    UUID REFERENCES auth_organizations (id) ON DELETE CASCADE,
    parent_id   UUID                     NULL REFERENCES production_units (id),
    type        VARCHAR(50)              NOT NULL,
    sequence    INTEGER                  NOT NULL,
    status      VARCHAR(30)              NOT NULL,
    code        TEXT                     NOT NULL,
    area        NUMERIC(12, 6)           NOT NULL DEFAULT 0,
    geometry    geometry(Geometry, 4326) NULL,
    properties  JSONB                    NOT NULL DEFAULT '{}'::jsonb,
    created_at  TIMESTAMPTZ              NOT NULL,
    updated_at  TIMESTAMPTZ              NOT NULL,
    archived_at TIMESTAMPTZ              NULL
);

CREATE INDEX idx_production_units_parent_id ON production_units (parent_id);
CREATE INDEX idx_production_units_type ON production_units (type);
CREATE INDEX idx_production_units_status ON production_units (status);
CREATE INDEX idx_production_units_archived_at ON production_units (archived_at);
CREATE INDEX idx_production_units_properties ON production_units USING GIN (properties);

CREATE UNIQUE INDEX ux_production_units_code ON production_units (owner_id, code);
CREATE UNIQUE INDEX ux_production_units_sequence ON production_units (owner_id,parent_id,type,sequence);
CREATE INDEX idx_production_units_parent ON production_units (parent_id);
CREATE INDEX idx_production_units_org    ON production_units (owner_id);