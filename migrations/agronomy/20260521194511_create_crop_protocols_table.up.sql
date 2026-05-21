CREATE TABLE crop_protocols
(
    id             UUID PRIMARY KEY,
    crop_id        UUID        NOT NULL,
    name           TEXT        NOT NULL,
    growing_method VARCHAR(50) NOT NULL,
    description    TEXT        NULL,
    metadata       JSONB       NOT NULL,
    created_at     TIMESTAMPTZ NOT NULL,
    updated_at     TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_crop_protocol_crop FOREIGN KEY (crop_id) REFERENCES crops (id)
);

CREATE TABLE crop_protocol_stage_profiles
(
    id                 UUID PRIMARY KEY,
    protocol_id        UUID        NOT NULL,
    stage_code         VARCHAR(50) NOT NULL,
    bbch               INTEGER     NULL,
    order_index        INTEGER     NOT NULL,
    duration_days      INTEGER     NULL,
    target_temperature JSONB       NULL,
    target_humidity    JSONB       NULL,
    target_ph          JSONB       NULL,
    target_ec          JSONB       NULL,
    target_light       JSONB       NULL,
    target_co2         JSONB       NULL,
    metadata           JSONB       NOT NULL,

    CONSTRAINT fk_protocol_stage FOREIGN KEY (protocol_id) REFERENCES crop_protocols (id) ON DELETE CASCADE
);

CREATE INDEX idx_crop_protocol_crop ON crop_protocols (crop_id);
CREATE INDEX idx_crop_protocol_method ON crop_protocols (growing_method);
CREATE INDEX idx_protocol_stage_protocol ON crop_protocol_stage_profiles (protocol_id);
CREATE INDEX idx_protocol_stage_order ON crop_protocol_stage_profiles (order_index);