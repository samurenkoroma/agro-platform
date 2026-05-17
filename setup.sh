#!/usr/bin/env bash

set -e

BASE="internal/domain"

mkdir -p \
$BASE/shared/{event,valueobject,errors,types,specification} \
\
$BASE/spatial/{aggregate,entity,service,event,repository,projection} \
$BASE/spatial/aggregate/production_unit \
$BASE/spatial/entity/{geometry,layout_snapshot,production_unit_snapshot} \
\
$BASE/production/{aggregate,entity,service,event,repository,projection} \
$BASE/production/aggregate/{growing_cycle,plant,reservoir} \
$BASE/production/entity/{slot,substrate} \
\
$BASE/agronomy/{aggregate,entity,service,event,repository,projection} \
$BASE/agronomy/aggregate/{crop,protocol} \
$BASE/agronomy/entity/{variety,stage,disease} \
\
$BASE/operations/{aggregate,entity,service,event,repository,projection} \
$BASE/operations/aggregate/{task,timeline} \
$BASE/operations/entity/{operation_event} \
\
internal/application/{command,query,dto} \
\
internal/application/spatial/{command,query,dto} \
internal/application/production/{command,query,dto} \
internal/application/agronomy/{command,query,dto} \
internal/application/operations/{command,query,dto} \
\
internal/interfaces/http \
internal/interfaces/grpc \
internal/interfaces/ws \
\
internal/infrastructure/postgres \
internal/infrastructure/eventbus \
internal/infrastructure/cache \
internal/infrastructure/clock \
\
tests/{integration,unit,e2e}

touch \
$BASE/shared/event/domain_event.go \
$BASE/shared/errors/errors.go \
$BASE/shared/valueobject/range.go \
$BASE/shared/valueobject/quantity.go \
$BASE/shared/valueobject/id.go \
\
$BASE/spatial/repository/production_unit_repository.go \
$BASE/spatial/service/topology_rules.go \
\
$BASE/production/repository/growing_cycle_repository.go \
$BASE/production/repository/plant_repository.go \
\
$BASE/agronomy/repository/crop_repository.go \
$BASE/agronomy/repository/protocol_repository.go \
\
$BASE/operations/repository/task_repository.go \
$BASE/operations/repository/timeline_repository.go

echo "DDD structure created"