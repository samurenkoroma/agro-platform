CREATE TABLE inventory_warehouses
(
    id          UUID         PRIMARY KEY,
    farm_id     UUID         NOT NULL,
    name        VARCHAR(255) NOT NULL,
    code        VARCHAR(100) NULL,
    created_at  TIMESTAMPTZ  NOT NULL,
    updated_at  TIMESTAMPTZ  NOT NULL,
    archived_at TIMESTAMPTZ  NULL
);

CREATE INDEX idx_inventory_warehouses_farm_id ON inventory_warehouses (farm_id);
