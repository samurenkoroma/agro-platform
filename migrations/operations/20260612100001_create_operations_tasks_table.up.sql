CREATE TABLE operations_tasks
(
    id                  UUID         PRIMARY KEY,

    farm_id             UUID         NOT NULL,
    title               VARCHAR(255) NOT NULL,
    description         TEXT         NULL,

    operation_type      VARCHAR(100) NULL,

    production_unit_id  UUID         NULL,
    growing_cycle_id    UUID         NULL,
    plant_id            UUID         NULL,

    assigned_to         UUID         NULL,

    status              VARCHAR(50)  NOT NULL DEFAULT 'TODO',
    priority            VARCHAR(50)  NOT NULL DEFAULT 'MEDIUM',

    due_date            TIMESTAMPTZ  NULL,
    completed_at        TIMESTAMPTZ  NULL,
    archived_at         TIMESTAMPTZ  NULL,

    created_at          TIMESTAMPTZ  NOT NULL,
    updated_at          TIMESTAMPTZ  NOT NULL
);

CREATE INDEX idx_operations_tasks_farm_id          ON operations_tasks (farm_id);
CREATE INDEX idx_operations_tasks_growing_cycle_id ON operations_tasks (growing_cycle_id);
CREATE INDEX idx_operations_tasks_status           ON operations_tasks (status);
CREATE INDEX idx_operations_tasks_assigned_to      ON operations_tasks (assigned_to);
CREATE INDEX idx_operations_tasks_due_date         ON operations_tasks (due_date);
