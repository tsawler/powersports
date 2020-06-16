package template_data

import (
	"database/sql"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/pkg/templates"
)

var vehicleModel *clientdb.DBModel

func NewTemplateData(p *sql.DB) {
	vehicleModel = &clientdb.DBModel{DB: p}
}

// AddDefaultData
func AddDefaultData(td *templates.TemplateData) *templates.TemplateData {

	return td
}
