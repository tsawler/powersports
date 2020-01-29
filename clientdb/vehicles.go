package clientdb

import (
	"database/sql"
	"github.com/tsawler/goblender/client/clienthandlers/clientmodels"
)

type VehicleModel struct {
	DB *sql.DB
}

func (m *VehicleModel) GetAllMotorcycles() ([]clientmodels.Vehicle, error) {
	var v []clientmodels.Vehicle

	return v, nil
}
