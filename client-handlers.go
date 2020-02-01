package clienthandlers

import (
	"encoding/json"
	"fmt"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/pkg/helpers"
	"github.com/tsawler/goblender/pkg/templates"
	"net/http"
	"strconv"
)

var vehicleModel *clientdb.VehicleModel

// JSONResponse is a generic struct to hold json responses
type JSONResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// CompareVehicles Show 2 or 3 vehicles in table TODO
func CompareVehicles(w http.ResponseWriter, r *http.Request) {
	ids := r.Form.Get("ids")
	infoLog.Println("Ids:", ids)
	helpers.Render(w, r, "compare.page.tmpl", &templates.TemplateData{})
}

// GetAllMotorcycles gets all motorcycles
func GetAllMotorcycles(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/motorcycle"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 7
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "motorcycle-inventory")
}

// GetAllBruteForce gets all brute force
func GetAllBruteForce(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/atv-brute-force"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 8
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-brute-force")
}

// GetAllTeryx gets all teryx
func GetAllTeryx(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/atv-teryx"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 12
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-teryx")
}

// GetAllMule gets all mules
func GetAllMule(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/atv-mule"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 11
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-mule")
}

// GetAllJetSki gets all new jetskis
func GetAllJetSki(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/watercraft-jetski"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 13
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "watercraft-jetski-kawaski")
}

// GetAllMercuryOutboards gets all mercury outboards
func GetAllMercuryOutboards(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/outboard-motors-mercury"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 10
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "outboard-motors-mercury-outboards")
}

// GetAllElectricBikes gets all Pedegogo
func GetAllElectricBikes(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/electric-bikes"
	intMap := make(map[string]int)
	intMap["show-makes"] = 0
	vehicleType := 16
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "electric-bikes")
}

// GetAllScooters gets all scooters
func GetAllScooters(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/vespa-piaggio-gas-and-electric-scooters-mopeds"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 17
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "vespa-piaggio-gas-and-electric-scooters-mopeds")
}

// GetAllPontoonBoats gets all pontoon boats (bennington & crestliner
func GetAllPontoonBoats(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/bennington-crestliner-pontoon-boats"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 9
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "pontoon-boats-bennington-and-crestliner")
}

// GetAllPowerBoats gets all power boats
func GetAllPowerBoats(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/speed-boats-aluminum-boats-power-boats"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 15
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "speed-boats-aluminum-boats-power-boats")
}

// GetAllUsedPowerSports gets all used powersports
func GetAllUsedPowerSports(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/used-motorcycles-used-atv-used-boats-used-pontoons"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 1000
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "used-motorcycles-used-atv-used-boats-used-pontoons")
}

// GetAllTrailers gets all used powersports
func GetAllTrailers(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/inventory/atv-trailers-boat-trailers-utility-trailers"
	stringMap["item-link-prefix"] = "view"
	intMap := make(map[string]int)
	intMap["show-makes"] = 1
	vehicleType := 14
	templateName := "inventory.page.tmpl"

	renderInventory(r, stringMap, vehicleType, w, intMap, templateName, "atv-trailers-boattrailers-utility-trailers")
}

// renderInventory renders inventory for a product type
func renderInventory(r *http.Request, stringMap map[string]string, vehicleType int, w http.ResponseWriter, intMap map[string]int, templateName, slug string) {
	var offset int
	var selectedYear, selectedMake, selectedModel, selectedPrice int
	pagerSuffix := ""
	stringMap["item-link-prefix"] = "view"
	stringMap["pager-prefix"] = "powersports-inventory"

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

// ShowItem shows a product (i.e. ATV, whatever)
func ShowItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get(":ID"))

	pg, err := pageModel.GetBySlug("power-sports-item")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	pg.PageNotEditable = 1

	item, err := vehicleModel.GetPowerSportItem(id)
	if err != nil {
		fmt.Fprint(w, "custom 404")
		return
	}

	rowSets := make(map[string]interface{})
	rowSets["item"] = item

	staff, err := vehicleModel.GetSales()
	if err != nil {
		errorLog.Println(err)
	}

	rowSets["sales"] = staff

	helpers.Render(w, r, "item.page.tmpl", &templates.TemplateData{
		Page:    pg,
		RowSets: rowSets,
	})
}

func QuickQuote(w http.ResponseWriter, r *http.Request) {
	infoLog.Println("HIt")
	infoLog.Println(r.Form.Get("name"))

	theData := JSONResponse{
		OK: true,
	}

	// build the json response from the struct
	out, err := json.MarshalIndent(theData, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// send json to client
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		errorLog.Println(err)
	}
}
