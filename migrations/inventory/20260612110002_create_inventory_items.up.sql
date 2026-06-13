CREATE TABLE inventory_items
(
    id              UUID PRIMARY KEY,
    farm_id         UUID           NOT NULL,
    name            VARCHAR(255)   NOT NULL,
    sku             VARCHAR(100)   NULL,

    type            VARCHAR(100)   NOT NULL,
    unit            VARCHAR(50)    NOT NULL,

    warehouse_id    UUID           NULL,

    stock_available NUMERIC(14, 4) NOT NULL DEFAULT 0,
    stock_reserved  NUMERIC(14, 4) NOT NULL DEFAULT 0,
    stock_consumed  NUMERIC(14, 4) NOT NULL DEFAULT 0,
    stock_lost      NUMERIC(14, 4) NOT NULL DEFAULT 0,

    created_at      TIMESTAMPTZ    NOT NULL,
    updated_at      TIMESTAMPTZ    NOT NULL,
    archived_at     TIMESTAMPTZ    NULL,

    CONSTRAINT fk_item_belong_warehouse FOREIGN KEY (warehouse_id) REFERENCES inventory_warehouses (id) ON DELETE SET NULL
);

CREATE INDEX idx_inventory_items_farm_id ON inventory_items (farm_id);
CREATE INDEX idx_inventory_items_warehouse_id ON inventory_items (warehouse_id);
CREATE INDEX idx_inventory_items_type ON inventory_items (type);
