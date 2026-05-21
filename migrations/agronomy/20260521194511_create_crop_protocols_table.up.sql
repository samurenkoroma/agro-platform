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
    archived_at    TIMESTAMPTZ NULL,

    CONSTRAINT fk_crop_protocol_crop FOREIGN KEY (crop_id) REFERENCES crops (id)
);

CREATE INDEX idx_crop_protocol_crop ON crop_protocols (crop_id);
CREATE INDEX idx_crop_protocol_method ON crop_protocols (growing_method);

CREATE TABLE crop_protocol_stage_profiles
(
    id            UUID PRIMARY KEY,
    protocol_id   UUID    NOT NULL,
    crop_stage_id UUID    NOT NULL,
    climate       JSONB   NOT NULL,
    lighting      JSONB   NOT NULL,
    irrigation    JSONB   NOT NULL,
    nutrition     JSONB   NOT NULL,
    water         JSONB   NOT NULL,
    vpd           JSONB   NOT NULL,
    order_index   INTEGER NOT NULL,
    metadata      JSONB   NOT NULL,

    CONSTRAINT fk_protocol_stage_protocol FOREIGN KEY (protocol_id) REFERENCES crop_protocols (id) ON DELETE CASCADE
);

CREATE INDEX idx_protocol_stage_protocol ON crop_protocol_stage_profiles (protocol_id);
CREATE INDEX idx_protocol_stage_crop ON crop_protocol_stage_profiles (crop_stage_id);

