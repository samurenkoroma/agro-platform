package service

import pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"

type TopologyRules interface {
	CanAttach(
		parent pu.ProductionUnitType,

		child pu.ProductionUnitType,
	) bool
}

type DefaultTopology struct{}

func (
	DefaultTopology,
) CanAttach(
	parent pu.ProductionUnitType,

	child pu.ProductionUnitType,
) bool {

	allowed := map[pu.ProductionUnitType][]pu.ProductionUnitType{

		pu.Field: {
			pu.Block,
		},

		pu.Block: {
			pu.Bed,
		},

		pu.Bed: {
			pu.Row,
		},

		pu.Greenhouse: {
			pu.GreenhouseZone,
		},

		pu.Rack: {
			pu.Shelf,
		},

		pu.Shelf: {
			pu.Slot,
		},
	}

	for _, v := range allowed[parent] {

		if v == child {
			return true
		}
	}

	return false
}
