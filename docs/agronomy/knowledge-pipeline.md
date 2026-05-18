# Agronomy Knowledge Pipeline

## Назначение

Этот документ описывает внешние источники данных для наполнения
Agronomy Context.

Agronomy НЕ хранит сырой импорт.

Agronomy хранит нормализованные модели:

- Crop
- Variety
- CropStage
- CropProtocol
- Disease
- Stress
- Pest
- NutrientDeficiency

---

# Layer 1 — Crop Biology

Используется для:

- Crop
- CropStage
- GDD
- Water demand
- Climate ranges

Sources:

FAO Crop Water Information
https://www.fao.org/land-water/databases-and-software/crop-information/en/

FAOSTAT
https://www.fao.org/faostat/

USDA Plants
https://plants.usda.gov/

USDA ARS
https://www.ars.usda.gov/

CGIAR
https://www.cgiar.org/

---

# Layer 2 — Variety Genetics

Используется для:

- Variety
- maturity
- height
- spacing
- yield potential
- harvest profile

Sources:

Rijk Zwaan
https://www.rijkzwaan.com/

Enza Zaden
https://www.enzazaden.com/

Syngenta Vegetable Seeds
https://www.syngentavegetables.com/

Bejo Seeds
https://www.bejo.com/

Hazera
https://www.hazera.com/

---

# Layer 3 — Controlled Environment Agriculture

Используется для:

- CropProtocol
- PPFD
- DLI
- VPD
- CO₂
- hydroponics
- vertical farming

Sources:

Wageningen University
https://www.wur.nl/

Cornell CEA
https://cea.cals.cornell.edu/

University of Arizona CEAC
https://ceac.arizona.edu/

---

# Layer 4 — Fertigation

Используется для:

- EC
- pH
- nutrients
- fertigation
- drainage

Sources:

Yara
https://www.yara.com/crop-nutrition/

Haifa Group
https://www.haifa-group.com/

Grodan
https://www.grodan.com/

---

# Layer 5 — Diseases / Pests

Используется для:

- Disease
- Pest
- Stress
- Pathogens

Sources:

EPPO Global Database
https://gd.eppo.int/

CABI Crop Protection
https://www.cabi.org/

---

# Rule

External datasets != Domain Models

External data:

FAO
Rijk Zwaan
Grodan
WUR
Yara

↓

Normalization

↓

Domain:

Crop
Variety
CropProtocol
Disease
Stress
Pest