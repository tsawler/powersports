package clienthandlers

import (
	"database/sql"
	"fmt"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/pkg/config"
	"github.com/tsawler/goblender/pkg/helpers"
	"net/http"
)

var vehicleModel *clientdb.VehicleModel

func NewClientVehicleHandlers(conn *sql.DB, app config.AppConfig) {
	vehicleModel = &clientdb.VehicleModel{DB: conn}
}

func GetAllMotorcycles(w http.ResponseWriter, r *http.Request) {
	vehicles, err := vehicleModel.GetAllMotorcycles()
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}
	for _, x := range vehicles {
		infoLog.Println(x.ID)
	}
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func TestProtectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Worked!")
}
