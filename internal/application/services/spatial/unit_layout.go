package spatial

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pu "github.com/samurenkoroma/agro-platform/internal/domain/spatial/aggregate/production_unit"
	"github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
)

type LayoutGenerator interface {
	Generate(ctx context.Context, parent *pu.ProductionUnit, schema pu.LayoutSchema) error
}

type generator struct {
	repo repository.ProductionUnitRepository
	exec uow.Execution
}

func (g generator) Generate(ctx context.Context, parent *pu.ProductionUnit, schema pu.LayoutSchema) error {
	seq, err := g.repo.GetNextSequence(ctx, parent.OwnerID, parent.ParentID, pu.Bed)
	for _, element := range schema.Beds {

		unitType := element.Type.ToUpper()

		if err != nil {
			return err
		}

		code := pu.BuildCode(parent.Code, unitType, seq)

		unit := pu.New(parent.OwnerID, &parent.ID, unitType, code, &element.Name, seq)
		unit.AddDimensions(&pu.Dimensions{Length: &element.Length, Width: &element.Width})
		unit.Properties.Position = &valueobject.Position{
			X: element.X,
			Y: element.Y,
		}
		if err := g.repo.Save(ctx, unit); err != nil {
			return err
		}
		g.exec.RegisterAggregate(unit)
		seq++
	}
	return nil
}

func NewUnitLayoutGenerator(repo repository.ProductionUnitRepository, exec uow.Execution) LayoutGenerator {
	return generator{
		repo: repo,
		exec: exec,
	}
}
