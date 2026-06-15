# Changelog

Все значимые изменения проекта фиксируются в этом файле.

Формат основан на принципах Keep a Changelog.
Версионирование проекта следует Semantic Versioning.

---

## [Unreleased]

### Added

### Changed

### Removed

### Fixed

---

## [0.0.3] - 2026-06-15

### Added

#### Account
- Управление организациями и пользователями.
- Базовая модель ролей и доступа.

#### Agronomy
- Aggregate Crop.
- Aggregate Variety.
- Репозитории и проекции для культур и сортов.
- CRUD-операции для управления справочником культур.
- CRUD-операции для управления справочником сортов.

#### Spatial
- Aggregate ProductionUnit.
- Поддержка различных типов производственных единиц.
- Управление площадями производственных единиц.
- Проекции и команды для работы с участками.

#### Production
- Aggregate GrowingCycle.
- Aggregate Allocation.
- Aggregate Planting.
- Команда StartGrowingCycle.
- Команда AllocateProductionUnit.
- Команда ChangeAllocation.
- Команда ReleaseAllocation.
- Команда RegisterPlanting.
- Событие GrowingCycleCreated.
- Событие AllocationAllocated.
- Событие ProductionUnitOccupied.
- Проекции GrowingCycle.
- Поддержка множественного размещения одного цикла выращивания на нескольких участках.

### Changed

#### Production
- Площадь Allocation хранится в гектарах.
- Один GrowingCycle может содержать несколько Allocation.
- ProtocolID переведен в необязательное поле.
- ProductionMethod вынесен в справочные значения ProductionUnit.

#### Architecture
- В качестве идентификаторов используется value object `vo.ID`.
- Межмодульные связи реализуются через идентификаторы без внешних ключей.
- Внешние ключи используются только внутри bounded context.

### Fixed

#### Infrastructure
- Исправлена блокировка при публикации доменных событий через UnitOfWork.
- Исправлена обработка транзакций при выполнении вложенных сценариев.
- Исправлена регистрация агрегатов при обработке событий.

---

## [0.0.2] - 2026-05-XX

### Added

#### Foundation
- Первичная структура модульного монолита.
- Базовая реализация Domain Events.
- Event Bus.
- Unit Of Work.
- Repository Provider.
- Command Handler.
- Query Handler.
- PostgreSQL инфраструктура.

---

## [0.0.1] - 2026-05-XX

### Added

#### Foundation
- Инициализация проекта.
- Базовая структура каталогов.
- Настройка PostgreSQL.
- Настройка фронтенд-приложения.
- Базовые общие value objects.