CREATE TABLE varieties
(
    id            UUID PRIMARY KEY,
    crop_id       UUID        NOT NULL,

    name          TEXT        NOT NULL,
    breeder       TEXT        NULL,

    maturity      JSONB       NOT NULL,
    growth        JSONB       NOT NULL,
    spacing       JSONB       NOT NULL,

    harvest       JSONB       NOT NULL,
    yield_profile JSONB       NOT NULL,

    tolerance     JSONB       NOT NULL,

    metadata      JSONB       NOT NULL,

    created_at    TIMESTAMPTZ NOT NULL,
    updated_at    TIMESTAMPTZ NOT NULL,

    archived_at   TIMESTAMPTZ NULL,

    CONSTRAINT fk_variety_crop FOREIGN KEY (crop_id) REFERENCES crops (id)
);