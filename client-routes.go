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

	// atvs - brute force
	mux.Get("/atv-brute-force", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/atv-brute-force-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/atv-brute-force-inventory", standardMiddleWare.ThenFunc(GetAllBruteForce))
	mux.Get("/inventory/atv-brute-force-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllBruteForce))

	// atvs - teryx
	mux.Get("/atv-teryx", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/atv-teryx-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/atv-teryx-inventory", standardMiddleWare.ThenFunc(GetAllTeryx))
	mux.Get("/inventory/atv-teryx-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllTeryx))

	// atvs - mule
	mux.Get("/atv-mule", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/atv-mule-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/atv-mule-inventory", standardMiddleWare.ThenFunc(GetAllMule))
	mux.Get("/inventory/atv-mule-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllMule))

	// jetski
	mux.Get("/watercraft-jetski", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/watercraft-jetski-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/watercraft-jetski-inventory", standardMiddleWare.ThenFunc(GetAllJetSki))
	mux.Get("/inventory/watercraft-jetski-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllJetSki))

	// jetski
	mux.Get("/outboard-motors-mercury-outboards", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/outboard-motors-mercury-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/outboard-motors-mercury-inventory", standardMiddleWare.ThenFunc(GetAllMercuryOutboard))
	mux.Get("/inventory/outboard-motors-mercury-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllMercuryOutboard))

	// electric bikes
	mux.Get("/electric-bikes", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/outboard-motors-mercury-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/electric-bikes-inventory", standardMiddleWare.ThenFunc(GetAllElectricBikes))
	mux.Get("/inventory/electric-bikes-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllElectricBikes))

	// scooters
	mux.Get("/vespa-piaggio-gas-and-electric-scooters-mopeds", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/vespa-piaggio-gas-and-electric-scooters-mopeds-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/vespa-piaggio-gas-and-electric-scooters-mopeds-inventory", standardMiddleWare.ThenFunc(GetAllScooters))
	mux.Get("/inventory/vespa-piaggio-gas-and-electric-scooters-mopeds-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllScooters))

	// pontoon boats
	mux.Get("/pontoon-boats-bennington-and-crestliner", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/bennington-crestliner-pontoon-boats-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/bennington-crestliner-pontoon-boats-inventory", standardMiddleWare.ThenFunc(GetAllPontoonBoats))
	mux.Get("/inventory/bennington-crestliner-pontoon-boats-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllPontoonBoats))

	// power boats
	mux.Get("/speed-boats-aluminum-boats-power-boats", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/speed-boats-aluminum-boats-power-boats", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/speed-boats-aluminum-boats-power-boats-inventory", standardMiddleWare.ThenFunc(GetAllPowerBoats))
	mux.Get("/inventory/speed-boats-aluminum-boats-power-boats-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllPowerBoats))

	// used
	mux.Get("/used-motorcycles-used-atv-used-boats-used-pontoons", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/used-motorcycles-used-atv-used-boats-used-pontoons-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/used-motorcycles-used-atv-used-boats-used-pontoons-inventory", standardMiddleWare.ThenFunc(GetAllUsedPowerSports))
	mux.Get("/inventory/used-motorcycles-used-atv-used-boats-used-pontoons-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllUsedPowerSports))

	// traiers
	mux.Get("/atv-trailers-boattrailers-utility-trailerss", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/atv-trailers-boattrailers-utility-trailers-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/atv-trailers-boattrailers-utility-trailers-inventory", standardMiddleWare.ThenFunc(GetAllTrailers))
	mux.Get("/inventory/atv-trailers-boattrailers-utility-trailers-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllTrailers))

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
