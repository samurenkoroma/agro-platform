CREATE TABLE diseases
(
    id            UUID PRIMARY KEY,
    crop_id       UUID        NULL,
    name          TEXT        NOT NULL,
    pathogen_type VARCHAR(50) NOT NULL,
    description   TEXT        NULL,
    symptoms      JSONB       NOT NULL,
    conditions    JSONB       NOT NULL,
    prevention    JSONB       NOT NULL,
    treatment     JSONB       NOT NULL,
    metadata      JSONB       NOT NULL,

    created_at    TIMESTAMPTZ NOT NULL,
    updated_at    TIMESTAMPTZ NOT NULL,
    archived_at   TIMESTAMPTZ NULL
);

CREATE INDEX idx_disease_crop ON diseases (crop_id);

CREATE INDEX idx_disease_pathogen ON diseases (pathogen_type);