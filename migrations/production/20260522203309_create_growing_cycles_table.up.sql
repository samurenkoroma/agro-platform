CREATE TABLE production_growing_cycles
(
    id                  UUID PRIMARY KEY,

    farm_id             UUID         NOT NULL,
    crop_id             UUID         NOT NULL,

    variety_id          UUID         NULL,
    protocol_id         UUID         NULL,

    name                VARCHAR(255) NOT NULL,
    code                VARCHAR(100) NOT NULL UNIQUE,

    method              VARCHAR(50)  NOT NULL,

    status              VARCHAR(50)  NOT NULL,
    stage               VARCHAR(50)  NOT NULL,

    created_at          TIMESTAMPTZ  NOT NULL,
    updated_at          TIMESTAMPTZ  NOT NULL
);

CREATE INDEX idx_production_growing_cycles_farm_id ON production_growing_cycles (farm_id);

CREATE INDEX idx_production_growing_cycles_crop_id ON production_growing_cycles (crop_id);

CREATE INDEX idx_production_growing_cycles_status ON production_growing_cycles (status);

CREATE INDEX idx_production_growing_cycles_code ON production_growing_cycles (code);