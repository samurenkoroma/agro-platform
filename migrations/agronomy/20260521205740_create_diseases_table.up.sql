CREATE TABLE diseases
(
    id              UUID PRIMARY KEY,
    name            TEXT        NOT NULL,
    scientific_name TEXT        NULL,
    pathogen_type   VARCHAR(50) NOT NULL,
    hosts           JSONB       NOT NULL,
    symptoms        JSONB       NOT NULL,
    description     TEXT        NULL,
    metadata        JSONB       NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL,
    updated_at      TIMESTAMPTZ NOT NULL,
    archived_at     TIMESTAMPTZ NULL
);

CREATE INDEX idx_disease_name ON diseases (name);

CREATE INDEX idx_disease_pathogen ON diseases (pathogen_type);