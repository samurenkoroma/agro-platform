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

	_ "github.com/lib/pq"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/crop"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	crop2 "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/messaging/inmemory"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/postgres"
	configs "github.com/samurenkoroma/agro-platform/internal/shared/config"
	"github.com/samurenkoroma/agro-platform/pkg/db"
)

var (
	dataDir = flag.String("data", "./data", "Path to seed data directory")
	dryRun  = flag.Bool("dry-run", false, "Dry run mode - only validate, don't insert")
	module  = flag.String("module", "all", "Module to seed (crop, farm, growing, all)")
)

type seedData struct {
	Crops []struct {
		Name     string          `json:"name"`
		Category string          `json:"category"`
		Family   string          `json:"family"`
		Agronomy json.RawMessage `json:"agronomy"`
	} `json:"crops"`
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
	log.Printf("Found %d crop types to seed", len(parseData.Crops))
	switch *module {
	case "all":
		if err := seedCropTypes(uow, parseData, *dryRun); err != nil {
			log.Fatalf("Failed to seed crop types: %v", err)
		}
	}

	log.Println("Seeding completed successfully!")
}

// // seedCropTypes заливает типы культур из JSON
func seedCropTypes(uow uow.UnitOfWork, data seedData, dryRun bool) error {
	for _, ct := range data.Crops {
		if dryRun {
			log.Printf("[DRY RUN] Would create: %s (category: %s, perennial: %v)",
				ct.Name, ct.Category, ct.Family)
			continue
		}

		var agronomy crop2.AgronomyProfile
		json.Unmarshal(ct.Agronomy, &agronomy)
		cmd := &crop.CreateCropCommand{
			Agronomy: agronomy,
			Name:     ct.Name,
			Category: ct.Category,
			Family:   ct.Family,
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
