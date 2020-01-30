package clienthandlers

import (
	"database/sql"
	"fmt"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/pkg/config"
	"github.com/tsawler/goblender/pkg/helpers"
	"github.com/tsawler/goblender/pkg/templates"
	"net/http"
)

var vehicleModel *clientdb.VehicleModel

func NewClientVehicleHandlers(conn *sql.DB, app config.AppConfig) {
	vehicleModel = &clientdb.VehicleModel{DB: conn}
}

func GetAllMotorcycles(w http.ResponseWriter, r *http.Request) {
	vehicles, err := vehicleModel.GetVehiclesForSaleByType(7)
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	rowSets := make(map[string]interface{})
	rowSets["vehicles"] = vehicles
	//for _, x := range vehicles {
	//	infoLog.Println(x.ID)
	//	infoLog.Println(x.Make.Make, x.Model.Model, x.Trim, x.Cost)
	//	infoLog.Println(x.Images)
	//	for _, i := range x.Images {
	//		infoLog.Println("     ", i.Image)
	//	}
	//}

	pg, err := pageModel.GetBySlug("motorcycle-inventory")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	helpers.Render(w, r, "motorcycles.page.tmpl", &templates.TemplateData{
		Page:    pg,
		RowSets: rowSets,
	})
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func TestProtectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Worked!")
}
