# Operations: полная пирамида тестов

## Зависимость

Integration и e2e тесты требуют Docker и `testcontainers-go`:

```bash
go get github.com/testcontainers/testcontainers-go@latest
go mod tidy
```

## Запуск

```bash
# Unit + application (без Docker, быстро)
go test ./internal/domain/operations/... \
        ./internal/application/commands/operations/... \
        ./internal/infrastructure/repository/inmemory/operations/... -v

# Integration (требует Docker, тег integration)
go test -tags=integration \
        ./internal/infrastructure/repository/postgres/operations/... \
        ./internal/infrastructure/postgres/... -v

# E2E (требует Docker)
go test -tags=integration ./e2e/operations/... -v

# Всё
go test -tags=integration ./... -v
```

## Что покрывает каждый уровень

**Domain** — переходы статусов Task, доменные события, Timeline.AddItem, OperationEvent.Payload.

**Application** — команды через FakeUoW + in-memory репозитории: happy path,
отсутствие organization_id, неверный payload, авто-создание и накопление Timeline.

**Postgres integration** — реальный SQL: UPSERT, ON CONFLICT, JSONB round-trip,
уникальный индекс (farm_id, growing_cycle_id), NULL-семантика, сортировка DESC.

**UoW integration** — атомарность (commit/rollback), регрессия на per-call state
(один uow.UnitOfWork обрабатывает N последовательных Execute без утечки состояния).

**E2E** — register → login → create_organization → switch_organization → реальные
CQRS-запросы через httptest.Server. Task: полный цикл TODO→IN_PROGRESS→DONE,
401 без токена, изоляция между организациями. Operations + Timeline: record_operation
создаёт запись и автоматически кладёт её в Timeline, несколько операций накапливаются,
данные разных организаций не пересекаются.
