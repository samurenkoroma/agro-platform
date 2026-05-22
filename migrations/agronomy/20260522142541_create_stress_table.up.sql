CREATE TABLE stresses
(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    type VARCHAR(50) NOT NULL,
    triggers JSONB NOT NULL,
    symptoms JSONB NOT NULL,
    description TEXT NULL,
    metadata JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    archived_at TIMESTAMPTZ NULL
);

CREATE INDEX idx_stress_name ON stresses(name);

CREATE INDEX idx_stress_type ON stresses(type);