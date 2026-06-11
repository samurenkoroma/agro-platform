CREATE TABLE crops
(
    id              UUID PRIMARY KEY,
    name            TEXT        NOT NULL,
    scientific_name TEXT        NOT NULL,
    category        TEXT        NOT NULL,
    family          TEXT        NOT NULL,
    ImageUrl        TEXT        NULL,
    metadata        JSONB       NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL,
    updated_at      TIMESTAMPTZ NOT NULL,
    UNIQUE (scientific_name, category)
);

CREATE INDEX idx_crops_name ON crops (name);

CREATE INDEX idx_crops_category ON crops (category);