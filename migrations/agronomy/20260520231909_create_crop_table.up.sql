CREATE TABLE agronomy_crops
(
    id         UUID PRIMARY KEY,
    name       TEXT        NOT NULL,
    category   TEXT        NOT NULL,
    family     TEXT        NOT NULL,

    image      TEXT        NULL,
    agronomy   JSONB       NOT NULL DEFAULT '{}',
    metadata   JSONB       NOT NULL DEFAULT '{"description": ""}',
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    UNIQUE (name, category)
);

CREATE INDEX idx_crops_name ON agronomy_crops (name);

CREATE INDEX idx_crops_category ON agronomy_crops (category);