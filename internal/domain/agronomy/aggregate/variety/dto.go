package variety

// ========== ТИПЫ ДЛЯ ФЕНОЛОГИИ (BBCH + GDD) ==========

// PhenophaseGDD - фаза развития с требованиями по GDD
type PhenophaseGDD struct {
	Code        string  // "BBCH-10"
	Name        string  // "Первый настоящий лист"
	GDDRequired float64 // накопленное GDD для достижения
	Description string  // описание фазы
	IsCritical  bool    // критическая фаза?
}

// ========== ТИПЫ ДЛЯ НОРМ ВЫСЕВА ==========

// SeedingRate - норма высева для одного способа выращивания
type SeedingRate struct {
	GrowingType     string  // "open_ground", "greenhouse"
	RowSpacing      float64 // расстояние между рядами (м)
	PlantSpacing    float64 // расстояние между растениями (м)
	SowingDepth     float64 // глубина посева (см)
	GerminationRate float64 // всхожесть (%)
	SafetyFactor    float64 // страховой коэффициент (1.1-1.3)
}

func (r SeedingRate) CalculateSeedsNeeded(areaM2 float64) (seeds int, weightKg float64) {
	// Количество растений на 1 м²
	plantsPerM2 := (1 / r.RowSpacing) * (1 / r.PlantSpacing)

	// Количество семян с учетом всхожести и страховки
	seedsNeeded := plantsPerM2 * areaM2 / (r.GerminationRate / 100) * r.SafetyFactor

	// Примерный вес семян (1 семя ≈ 0.005 г)
	weightG := float64(int(seedsNeeded)) * 0.005

	return int(seedsNeeded), weightG / 1000
}

type WaterRequirement struct {
	DailyNeedMin   float64  // л/м² в день (минимально)
	DailyNeedOpt   float64  // л/м² в день (оптимально)
	CriticalPhases []string // критические BBCH коды
}

// LightRequirement потребность в освещении
type LightRequirement struct {
	PPFDMin         int      // μmol/m²/s (минимальный фотосинтетический поток)
	PPFDOpt         int      // μmol/m²/s (оптимальный)
	DayLengthMin    float64  // часов (минимальный световой день)
	DayLengthOpt    float64  // часов (оптимальный световой день)
	PhotoperiodType string   // "short_day", "long_day", "day_neutral"
	CriticalPhases  []string // критические BBCH коды для света
}
