CREATE TABLE agronomy_varieties
(
    id                UUID PRIMARY KEY,
    crop_id           UUID          NOT NULL,
    name              TEXT          NOT NULL,
    breeder           TEXT          NULL,
    image             TEXT,

    profile JSONB NOT NULL DEFAULT '{}',
    created_at        TIMESTAMPTZ   NOT NULL,
    updated_at        TIMESTAMPTZ   NOT NULL,

    CONSTRAINT fk_variety_crop FOREIGN KEY (crop_id) REFERENCES agronomy_crops (id),
    UNIQUE (name, crop_id)
);