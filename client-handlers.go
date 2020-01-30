package clienthandlers

import (
	"database/sql"
	"fmt"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/pkg/config"
	"github.com/tsawler/goblender/pkg/helpers"
	"github.com/tsawler/goblender/pkg/templates"
	"net/http"
	"strconv"
)

var vehicleModel *clientdb.VehicleModel

func NewClientVehicleHandlers(conn *sql.DB, app config.AppConfig) {
	vehicleModel = &clientdb.VehicleModel{DB: conn}
}

func CompareVehicles(w http.ResponseWriter, r *http.Request) {
	ids := r.Form.Get("ids")
	infoLog.Println("Ids:", ids)
	helpers.Render(w, r, "compare.page.tmpl", &templates.TemplateData{})
}

func GetAllMotorcycles(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/motorcycle-inventory"
	stringMap["item-link-prefix"] = "motorcycle"
	stringMap["pager-prefix"] = "motorcycles"

	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 7

	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "motorcycle-inventory")
}

func GetAllBruteForce(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/atv-brute-force-inventory"
	stringMap["item-link-prefix"] = "atv"
	stringMap["pager-prefix"] = "atvs"

	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 8

	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-brute-force")
}

func GetAllTeryx(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/atv-teryx-inventory"
	stringMap["item-link-prefix"] = "atv"
	stringMap["pager-prefix"] = "teryx"

	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 12

	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-teryx")
}

func GetAllMule(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/atv-mule-inventory"
	stringMap["item-link-prefix"] = "atv"
	stringMap["pager-prefix"] = "mule"

	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 11

	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-mule")
}

func GetAllJetSki(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/watercraft-jetski-inventory"
	stringMap["item-link-prefix"] = "atv"
	stringMap["pager-prefix"] = "jetski"

	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 13

	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "watercraft-jetski-kawaski")
}

func GetAllMercuryOutboard(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/outboard-motors-mercury-inventory"
	stringMap["item-link-prefix"] = "atv"
	stringMap["pager-prefix"] = "mercury"

	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 10

	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "outboard-motors-mercury-outboards")
}

func GetAllElectricBikes(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/electric-bikes-inventory"
	stringMap["item-link-prefix"] = "electric-bikes"
	stringMap["pager-prefix"] = "pedegogo"

	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 16

	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "electric-bikes")
}

func GetAllScooters(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/vespa-piaggio-gas-and-electric-scooters-mopeds-inventory"
	stringMap["item-link-prefix"] = "scooters"
	stringMap["pager-prefix"] = "scooters"

	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 16

	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "vespa-piaggio-gas-and-electric-scooters-mopeds")
}

func renderInventory(r *http.Request, stringMap map[string]string, vehicleType int, w http.ResponseWriter, intMap map[string]int, templateName, slug string) {
	var offset int
	var selectedYear, selectedMake, selectedModel, selectedPrice int
	pagerSuffix := ""

	pageIndex, err := strconv.Atoi(r.URL.Query().Get(":pageIndex"))
	if err != nil {
		pageIndex = 1
	}

	searching, ok := r.URL.Query()["year"]
	if !ok || len(searching[0]) < 1 {
		selectedYear = 0
		selectedMake = 0
		selectedModel = 0
		selectedPrice = 0
	} else {
		selectedYear, _ = strconv.Atoi(r.URL.Query()["year"][0])
		selectedMake, _ = strconv.Atoi(r.URL.Query()["make"][0])
		selectedModel, _ = strconv.Atoi(r.URL.Query()["model"][0])
		selectedPrice, _ = strconv.Atoi(r.URL.Query()["price"][0])
		pagerSuffix = fmt.Sprintf("?year=%d&make=%d&model=%d&price=%d", selectedYear, selectedMake, selectedModel, selectedPrice)
		stringMap["pager-suffix"] = pagerSuffix
	}

	perPage := 10
	offset = (pageIndex - 1) * perPage

	vehicles, num, err := vehicleModel.AllVehiclesPaginated(vehicleType, perPage, offset, selectedYear, selectedMake, selectedModel, selectedPrice)
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	rowSets := make(map[string]interface{})
	rowSets["vehicles"] = vehicles

	intMap["num-vehicles"] = num
	intMap["current-page"] = pageIndex
	intMap["year"] = selectedYear
	intMap["make"] = selectedMake
	intMap["model"] = selectedModel
	intMap["price"] = selectedPrice

	pg, err := pageModel.GetBySlug(slug)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// get makes
	makes, err := vehicleModel.GetMakesForVehicleType(vehicleType)
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	// get models
	models, err := vehicleModel.GetModelsForVehicleType(vehicleType)
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	// get years
	years, err := vehicleModel.GetYearsForVehicleType(vehicleType)
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	rowSets["years"] = years
	rowSets["models"] = models
	rowSets["makes"] = makes

	helpers.Render(w, r, templateName, &templates.TemplateData{
		Page:      pg,
		RowSets:   rowSets,
		IntMap:    intMap,
		StringMap: stringMap,
	})
}
