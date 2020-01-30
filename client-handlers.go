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
	var offset int

	pageIndex, err := strconv.Atoi(r.URL.Query().Get(":pageIndex"))
	if err != nil {
		pageIndex = 1
	}

	perPage := 10
	offset = (pageIndex - 1) * perPage

	vehicles, num, err := vehicleModel.AllVehiclesPaginated(7, perPage, offset)
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	infoLog.Println(num)

	rowSets := make(map[string]interface{})
	rowSets["vehicles"] = vehicles

	intMap := make(map[string]int)
	intMap["num-vehicles"] = num
	intMap["current-page"] = pageIndex

	pg, err := pageModel.GetBySlug("motorcycle-inventory")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	stringMap := make(map[string]string)
	stringMap["pager-url"] = "/motorcycle-inventory"

	helpers.Render(w, r, "motorcycles.page.tmpl", &templates.TemplateData{
		Page:      pg,
		RowSets:   rowSets,
		IntMap:    intMap,
		StringMap: stringMap,
	})
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func TestProtectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Worked!")
}
