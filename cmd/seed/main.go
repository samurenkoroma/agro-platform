package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/crop"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/variety"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	crop2 "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
	agronomy "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/messaging/inmemory"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/postgres"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/providers"
	configs "github.com/samurenkoroma/agro-platform/internal/shared/config"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
	"github.com/samurenkoroma/agro-platform/pkg/db"
)

var (
	dataDir = flag.String("data", "./data", "Path to seed data directory")
	dryRun  = flag.Bool("dry-run", false, "Dry run mode - only validate, don't insert")
	module  = flag.String("module", "all", "Module to seed (crop, farm, growing, all)")
)

type seedData struct {
	CropTypes []struct {
		Name           string `json:"name"`
		Category       string `json:"category"`
		Description    string `json:"description"`
		IsPerennial    bool   `json:"is_perennial"`
		ScientificName string `json:"scientific_name"`
		Family         string `json:"family"`
		Icon           string `json:"lucideIcon"`
		ImageUrl       string `json:"imageUrl"`
	} `json:"crop_types"`
	Varieties []struct {
		Name               string   `json:"name"`
		Croptype           string   `json:"crop"`
		Description        string   `json:"description"`
		VegetationDays     string   `json:"vegetation_days"`
		YieldPotential     string   `json:"yield_potential"`
		RecommendedRegions []string `json:"recommended_regions"`
	}
	CultivationPlan []struct {
		Version   int     `json:"version"`
		Name      string  `json:"name"`
		CropKey   string  `json:"crop_key"`
		VarietyId *string `json:"varietyId"`
		Steps     []struct {
			Type    string `json:"type" validate:"required"`
			Title   string `json:"title" validate:"required"`
			Trigger struct {
				Type  string         `json:"type" validate:"required"`
				Value map[string]any `json:"value" validate:"required,min=1,dive,keys,min=3,endkeys,required"`
			} `json:"trigger" validate:"required"`
		} `json:"steps" validate:"required"`
	} `json:"cultivation_plans"`
}

func main() {
	flag.Parse()

	// Загружаем конфигурацию
	cfg := configs.LoadConfig()

	// Подключаемся к БД
	pool, err := db.NewPool(cfg.Db)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	bus := inmemory.NewInMemoryEventBus()

	uow := postgres.NewUnitOfWork(pool, bus)
	data := readFile()

	var parseData seedData

	if err := json.Unmarshal(data, &parseData); err != nil {
		fmt.Errorf("failed to parse crop_types.json: %w", err)
	}
	log.Printf("Found %d crop types to seed", len(parseData.CropTypes))
	switch *module {
	case "all":
		if err := seedCropTypes(uow, parseData, *dryRun); err != nil {
			log.Fatalf("Failed to seed crop types: %v", err)
		}
		if err := seedVarieties(uow, parseData, *dryRun); err != nil {
			log.Fatalf("Failed to seed varieties: %v", err)
		}

	case "crop":
		if err := seedCropTypes(uow, parseData, *dryRun); err != nil {
			log.Fatalf("Failed to seed crop types: %v", err)
		}
		if err := seedVarieties(uow, parseData, *dryRun); err != nil {
			log.Fatalf("Failed to seed varieties: %v", err)
		}
	}

	log.Println("Seeding completed successfully!")
}

func seedVarieties(u uow.UnitOfWork, data seedData, dryRun bool) error {
	var crops []*crop2.Crop

	u.Execute(context.Background(), providers.NewAgronomyProvider, func(provider repository.RepositoryProvider, exec uow.Execution) (any, error) {
		agronomyProvider, ok := provider.(agronomy.AgronomyProvider)
		if !ok {
			return nil, repository.ErrInvalidProviderType
		}
		crops, _ = agronomyProvider.Crops().GetAll(context.Background())
		return nil, nil
	})

	mapCrops := make(map[string]vo.ID)
	for _, c := range crops {
		mapCrops[c.Name] = c.ID
	}

	for _, v := range data.Varieties {
		if dryRun {
			log.Printf("[DRY RUN] Would create: %s (category: %s)", v.Name, v.Croptype)
			continue
		}

		cmd := &variety.CreateVarietyCommand{
			Name:   v.Name,
			CropID: mapCrops[v.Croptype],
		}

		// Создаём обработчик
		handler := variety.NewHandler(u)

		// Выполняем команду
		if _, err := handler.Create(context.Background(), cmd); err != nil {
			if errors.Is(err, variety.ErrVarietyAlreadyExists) {
				log.Printf("Skipping. Variety '%s' already exists", v.Name)
				continue
			}
			fmt.Errorf("failed to create variety %s: %w", v.Name, err)
			continue
		}

		log.Printf("Created varieties: %s", v.Name)
	}
	return nil
}

// // seedCropTypes заливает типы культур из JSON
func seedCropTypes(uow uow.UnitOfWork, data seedData, dryRun bool) error {
	for _, ct := range data.CropTypes {
		if dryRun {
			log.Printf("[DRY RUN] Would create: %s (category: %s, perennial: %v)",
				ct.Name, ct.Category, ct.IsPerennial)
			continue
		}

		re := regexp.MustCompile(`([а-яА-Я]*)\s\((\w+)\)`)
		family := re.FindStringSubmatch(ct.Family)
		category := re.FindStringSubmatch(ct.Category)
		cmd := &crop.CreateCropCommand{
			ScientificName: ct.ScientificName,
			Name:           ct.Name,
			Category:       strings.ToLower(category[1]),
			Family:         strings.ToLower(family[1]),
		}

		// Создаём обработчик
		handler := crop.NewHandler(uow)

		// Выполняем команду
		if _, err := handler.Create(context.Background(), cmd); err != nil {
			// Если уже существует — пропускаем
			if errors.Is(err, crop.ErrCropAlreadyExist) {
				log.Printf("Skipping. CropType '%s' already exists", ct.Name)
				continue
			}
			return fmt.Errorf("failed to create crop type %s: %w", ct.Name, err)
		}

		log.Printf("Created crop type: %s", ct.Name)
	}

	return nil
}

func readFile() []byte {
	dataPath := filepath.Join(*dataDir, "seeds.json")
	data, err := os.ReadFile(dataPath)
	if err != nil {
		fmt.Errorf("failed to read crop_types.json: %w", err)
		return nil
	}
	return data
}
