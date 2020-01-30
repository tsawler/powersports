package clienthandlers

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/pkg/config"
	"github.com/tsawler/goblender/pkg/driver"
	"github.com/tsawler/goblender/pkg/repository"
	"github.com/tsawler/goblender/pkg/repository/page"
	"log"
	"net/http"
)

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger
var pageModel repository.PageRepo
var parentDB *driver.DB

func ClientRoutes(mux *pat.PatternServeMux, standardMiddleWare, dynamicMiddleware alice.Chain) (*pat.PatternServeMux, error) {

	mux.Post("/inventory/compare-vehicles", standardMiddleWare.ThenFunc(CompareVehicles))

	// motorcycles
	mux.Get("/motorcycle-inventory", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/motorcycle-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/motorcycle-inventory", standardMiddleWare.ThenFunc(GetAllMotorcycles))
	mux.Get("/inventory/motorcycle-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllMotorcycles))

	// atvs
	mux.Get("/atv-brute-force-inventory", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/atv-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/atv-brute-force-inventory", standardMiddleWare.ThenFunc(GetAllBruteForce))
	mux.Get("/inventory/atv-brute-force-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllBruteForce))

	return mux, nil
}

// ClientInit gives us access to site values for client code.
func ClientInit(c config.AppConfig, p *driver.DB) {
	app = c
	conn := app.Connections["wheels"]
	vehicleModel = &clientdb.VehicleModel{DB: conn}
	infoLog = app.InfoLog
	errorLog = app.ErrorLog
	pageModel = page.NewSQLPageRepo(p.SQL)
	parentDB = p
}
