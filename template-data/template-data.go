package template_data

import (
	"database/sql"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/pkg/templates"
)

var vehicleModel *clientdb.VehicleModel

func NewTemplateData(p *sql.DB) {
	vehicleModel = &clientdb.VehicleModel{DB: p}
}

// AddDefaultData
func AddDefaultData(td *templates.TemplateData) *templates.TemplateData {
	return td
}
