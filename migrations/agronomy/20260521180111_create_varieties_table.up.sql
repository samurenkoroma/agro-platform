CREATE TABLE varieties
(
    id                UUID PRIMARY KEY,
    crop_id           UUID          NOT NULL,

    name              TEXT          NOT NULL,
    breeder           TEXT          NULL,

    maturity          JSONB         NOT NULL,
    growth            JSONB         NOT NULL,
    spacing           JSONB         NOT NULL,

    harvest           JSONB         NOT NULL,
    yield_profile     JSONB         NOT NULL,

    tolerance         JSONB         NOT NULL,
    -- Температурные параметры
    base_temperature  DECIMAL(5, 2) NOT NULL DEFAULT 10.0,
    max_temperature   DECIMAL(5, 2) NOT NULL DEFAULT 30.0,

    -- Характеристики (JSONB для гибкости)
    characteristics   JSONB         NOT NULL,

    -- Описание
    description       TEXT,
    image             TEXT,

    -- Водные требования (JSONB)
    water_requirement JSONB         NOT NULL,
    -- Световые требования (JSONB)
    light_requirement JSONB         NOT NULL,
    -- Фазы развития (JSONB)
    phenophase_gdd    JSONB         NOT NULL,

    metadata          JSONB         NOT NULL,

    created_at        TIMESTAMPTZ   NOT NULL,
    updated_at        TIMESTAMPTZ   NOT NULL,

    archived_at       TIMESTAMPTZ   NULL,

    CONSTRAINT fk_variety_crop FOREIGN KEY (crop_id) REFERENCES crops (id),
    UNIQUE (name, crop_id)
);