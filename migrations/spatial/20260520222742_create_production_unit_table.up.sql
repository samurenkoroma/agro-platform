CREATE TABLE production_units
(
    id          UUID PRIMARY KEY,
    owner_id    UUID REFERENCES auth_organizations (id) ON DELETE CASCADE ,
    parent_id   UUID                     NULL REFERENCES production_units (id),
    type        VARCHAR(50)              NOT NULL,
    status      VARCHAR(30)              NOT NULL,
    name        TEXT                     NOT NULL,
    code        TEXT                     NULL,
    description TEXT                     NULL,
    geometry    geometry(Geometry, 4326) NULL,
    position    JSONB                    NULL,
    capacity    JSONB                    NULL,
    climate     JSONB                    NULL,
    properties  JSONB                    NOT NULL DEFAULT '{}'::jsonb,
    metadata    JSONB                    NOT NULL,
    created_at  TIMESTAMPTZ              NOT NULL,
    updated_at  TIMESTAMPTZ              NOT NULL,
    archived_at TIMESTAMPTZ              NULL
);

CREATE INDEX idx_production_units_parent_id ON production_units (parent_id);
CREATE INDEX idx_production_units_type ON production_units (type);
CREATE INDEX idx_production_units_status ON production_units (status);
CREATE INDEX idx_production_units_archived_at ON production_units (archived_at);
CREATE INDEX idx_production_units_properties ON production_units USING GIN (properties);