CREATE TABLE crop_stages
(
    id            UUID PRIMARY KEY,
    crop_id       UUID        NOT NULL,
    code          VARCHAR(50) NOT NULL,
    name          TEXT        NOT NULL,
    bbch          INTEGER     NULL,
    order_index   INTEGER     NOT NULL,
    duration_days INTEGER     NULL,
    metadata      JSONB       NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL,
    updated_at    TIMESTAMPTZ NOT NULL,
    archived_at   TIMESTAMPTZ NULL,

    CONSTRAINT fk_crop_stage_crop FOREIGN KEY (crop_id) REFERENCES crops (id)
);

CREATE INDEX idx_crop_stage_crop ON crop_stages (crop_id);

CREATE INDEX idx_crop_stage_bbch ON crop_stages (bbch);

CREATE INDEX idx_crop_stage_order ON crop_stages (order_index);