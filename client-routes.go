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

// ClientRoutes holds all app routes for the custom code
func ClientRoutes(mux *pat.PatternServeMux, standardMiddleWare, dynamicMiddleware alice.Chain) (*pat.PatternServeMux, error) {

	mux.Post("/inventory/compare-vehicles", standardMiddleWare.ThenFunc(CompareVehicles))

	// motorcycles
	mux.Get("/motorcycle-inventory", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/motorcycle", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/motorcycle", standardMiddleWare.ThenFunc(GetAllMotorcycles))
	mux.Get("/inventory/motorcycle/:pageIndex", standardMiddleWare.ThenFunc(GetAllMotorcycles))

	// atvs - brute force
	mux.Get("/atv-brute-force", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/atv-brute-force", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/atv-brute-force", standardMiddleWare.ThenFunc(GetAllBruteForce))
	mux.Get("/inventory/atv-brute-force/:pageIndex", standardMiddleWare.ThenFunc(GetAllBruteForce))

	// atvs - teryx
	mux.Get("/atv-teryx", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/atv-teryx", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/atv-teryx", standardMiddleWare.ThenFunc(GetAllTeryx))
	mux.Get("/inventory/atv-teryx/:pageIndex", standardMiddleWare.ThenFunc(GetAllTeryx))

	// atvs - mule
	mux.Get("/atv-mule", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/atv-mule", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/atv-mule", standardMiddleWare.ThenFunc(GetAllMule))
	mux.Get("/inventory/atv-mule/:pageIndex", standardMiddleWare.ThenFunc(GetAllMule))

	// jetski
	mux.Get("/watercraft-jetski", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/watercraft-jetski", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/watercraft-jetski", standardMiddleWare.ThenFunc(GetAllJetSki))
	mux.Get("/inventory/watercraft-jetski/:pageIndex", standardMiddleWare.ThenFunc(GetAllJetSki))

	// mercury outboards
	mux.Get("/outboard-motors-mercury-outboards", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/outboard-motors-mercury", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/outboard-motors-mercury", standardMiddleWare.ThenFunc(GetAllMercuryOutboards))
	mux.Get("/inventory/outboard-motors-mercury/:pageIndex", standardMiddleWare.ThenFunc(GetAllMercuryOutboards))

	// electric bikes
	mux.Get("/electric-bikes", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/outboard-motors-mercury", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/electric-bikes", standardMiddleWare.ThenFunc(GetAllElectricBikes))
	mux.Get("/inventory/electric-bikes/:pageIndex", standardMiddleWare.ThenFunc(GetAllElectricBikes))

	// scooters
	mux.Get("/vespa-piaggio-gas-and-electric-scooters-mopeds", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/vespa-piaggio-gas-and-electric-scooters-mopeds", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/vespa-piaggio-gas-and-electric-scooters-mopeds", standardMiddleWare.ThenFunc(GetAllScooters))
	mux.Get("/inventory/vespa-piaggio-gas-and-electric-scooters-mopeds/:pageIndex", standardMiddleWare.ThenFunc(GetAllScooters))

	// pontoon boats
	mux.Get("/pontoon-boats-bennington-and-crestliner", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/bennington-crestliner-pontoon-boats", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/bennington-crestliner-pontoon-boats", standardMiddleWare.ThenFunc(GetAllPontoonBoats))
	mux.Get("/inventory/bennington-crestliner-pontoon-boats/:pageIndex", standardMiddleWare.ThenFunc(GetAllPontoonBoats))

	// power boats
	mux.Get("/speed-boats-aluminum-boats-power-boats", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/speed-boats-aluminum-boats-power-boats", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/speed-boats-aluminum-boats-power-boats", standardMiddleWare.ThenFunc(GetAllPowerBoats))
	mux.Get("/inventory/speed-boats-aluminum-boats-power-boats/:pageIndex", standardMiddleWare.ThenFunc(GetAllPowerBoats))

	// used
	mux.Get("/used-motorcycles-used-atv-used-boats-used-pontoons", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/used-motorcycles-used-atv-used-boats-used-pontoons", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/used-motorcycles-used-atv-used-boats-used-pontoons", standardMiddleWare.ThenFunc(GetAllUsedPowerSports))
	mux.Get("/inventory/used-motorcycles-used-atv-used-boats-used-pontoons/:pageIndex", standardMiddleWare.ThenFunc(GetAllUsedPowerSports))

	// trailers
	mux.Get("/atv-trailers-boattrailers-utility-trailerss", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/atv-trailers-boat-trailers-utility-trailers", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/atv-trailers-boatt-railers-utility-trailers", standardMiddleWare.ThenFunc(GetAllTrailers))
	mux.Get("/inventory/atv-trailers-boat-trailers-utility-trailers/:pageIndex", standardMiddleWare.ThenFunc(GetAllTrailers))

	mux.Get("/inventory/:productType/:prefix/:ID", standardMiddleWare.ThenFunc(ShowItem))
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
