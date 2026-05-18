Уровень 1 — официальные агрономические базы

Для базовой биологии культуры.

Источники:

FAO Crop Water Information
FAOSTAT
USDA PLANTS Database
USDA Agricultural Research Service
CGIAR Research Platform

Отсюда брать:

дни созревания
GDD
рост
корневая система
водопотребление
температуры
Уровень 2 — селекционеры и производители семян

Самый ценный слой для Variety.

Например:

Rijk Zwaan
Enza Zaden
Syngenta Vegetable Seeds
Bejo Seeds
Hazera Seeds

Там уже есть:

Tomato Cherry F1

65 дней

индет

250 см

2.5 растения/м²

18 кг/м²

multi harvest

Это почти готовый Variety.

Уровень 3 — тепличные исследования

Для CropProtocol.

Источники:

Wageningen University & Research
Cornell Controlled Environment Agriculture
University of Arizona CEAC

Отсюда:

PPFD
DLI
VPD
CO2
EC
pH
досвет
NFT
DWC
вертикалки
Уровень 4 — fertigation / hydro

Источники:

Yara Crop Nutrition
Haifa Group Agronomy
Grodan Knowledge Center

Отсюда:

EC
pH
NPK
Ca
Mg
дренаж
фертигация
кокос
минвата
Уровень 5 — климат и болезни

Источники:

EPPO Global Database
CABI Crop Protection Compendium

Для:

болезни
вредители
симптомы
стрессы
патогены

Я бы вообще ввёл отдельный bounded context:

knowledge/

Например:

internal/domain/knowledge

crop_dataset
variety_dataset
protocol_dataset
disease_dataset
research_reference
source_reference

Потому что:

Agronomy

— это наши нормализованные сущности.

А:

FAO
Rijk Zwaan
Grodan
WUR
Yara

— это сырьё.