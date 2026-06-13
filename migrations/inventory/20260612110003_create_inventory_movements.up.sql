CREATE TABLE inventory_movements
(
    id               UUID           PRIMARY KEY,
    farm_id          UUID           NOT NULL,
    item_id          UUID           NOT NULL REFERENCES inventory_items (id) ON DELETE CASCADE,

    type             VARCHAR(50)    NOT NULL,
    quantity         NUMERIC(14, 4) NOT NULL,

    from_warehouse_id UUID          NULL REFERENCES inventory_warehouses (id) ON DELETE SET NULL,
    to_warehouse_id   UUID          NULL REFERENCES inventory_warehouses (id) ON DELETE SET NULL,

    reference_type   VARCHAR(100)   NULL,
    reference_id     TEXT           NULL,

    note             TEXT           NULL,
    timestamp        TIMESTAMPTZ    NOT NULL,

    CONSTRAINT fk_item_from_warehouse FOREIGN KEY (from_warehouse_id) REFERENCES inventory_warehouses (id) ON DELETE SET NULL,
    CONSTRAINT fk_item_to_warehouse FOREIGN KEY (to_warehouse_id) REFERENCES inventory_warehouses (id) ON DELETE SET NULL,
    CONSTRAINT fk_inventory_item FOREIGN KEY (item_id) REFERENCES inventory_items (id) ON DELETE SET NULL
);

CREATE INDEX idx_inventory_movements_farm_id   ON inventory_movements (farm_id);
CREATE INDEX idx_inventory_movements_item_id   ON inventory_movements (item_id);
CREATE INDEX idx_inventory_movements_type      ON inventory_movements (type);
CREATE INDEX idx_inventory_movements_timestamp ON inventory_movements (timestamp DESC);
