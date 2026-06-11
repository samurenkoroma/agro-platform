# Agro Platform

**Unified Agricultural Production Platform (Agro ERP)**
DDD + Clean Architecture на Go по ТЗ.md

🌱 В разработке

* Agronomy     = справочник знаний
* Spatial      = где выращиваем
* Production   = что выращиваем сейчас
* Operations   = что делаем
* Inventory    = чем обеспечиваем
* Environment  = что происходит вокруг
* Automation   = автоматизация
* Analytics    = анализ результатов


commands
StartGrowingCycle
AdvanceCycleStage
AllocateProductionUnit
RegisterPlanting
RecordObservation
RecordDisease
RecordStress
RegisterHarvest
CompleteCycle
FailCycle
ArchiveCycle



CreateGrowingCycle
↓
GrowingCycleCreated

AllocateProductionUnit
↓
AllocationAllocated

RegisterPlanting
↓
PlantingRegistered

RegisterHarvestBatch
↓
HarvestBatchRegistered

ReleaseAllocation
↓
AllocationReleased

CompleteGrowingCycle
↓
GrowingCycleCompleted