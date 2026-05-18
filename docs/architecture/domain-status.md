Agro Platform
DDD + Event Driven

Shared
- ID
- Metadata
- Quantity
- Area
- Range
- DomainEvent
- AggregateRoot

Spatial
- ProductionUnit
- hierarchy
- capabilities
- layout snapshots

Production
- GrowingCycle
- Plant
- Slot
- Substrate
- HarvestBatch
- YieldBatch

Agronomy
Crop
Variety
CropStage
CropProtocol

Variety:
- Maturity
- Growth
- Spacing
- HarvestProfile
- YieldPotential
- EnvironmentTolerance

CropProtocol:
- Climate
- Lighting
- Irrigation
- Nutrition
- WaterDemand
- VPD
- StageProfiles

Principle:
Biology != Technology

Variety -> biology
CropProtocol -> cultivation technology