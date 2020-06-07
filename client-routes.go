package clienthandlers

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/pkg/config"
	"github.com/tsawler/goblender/pkg/driver"
	"github.com/tsawler/goblender/pkg/handlers"
	"log"
	"net/http"
)

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger
var parentDB *driver.DB

var repo *handlers.DBRepo

// ClientRoutes holds all app routes for the custom code
func ClientRoutes(mux *pat.PatternServeMux, standardMiddleWare, dynamicMiddleware alice.Chain) (*pat.PatternServeMux, error) {

	// public folder
	fileServer := http.FileServer(http.Dir("./client/clienthandlers/public/"))
	mux.Get("/client/static/", http.StripPrefix("/client/static", fileServer))
	fileServer = http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// api
	mux.Get("/api/facebook-marketplace", dynamicMiddleware.ThenFunc(FacebookMarketplaceFeed))

	mux.Get("/blog", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/motorsportsnews", http.StatusMovedPermanently)
	}))

	// blog
	mux.Get("/motorsportsnews", standardMiddleWare.ThenFunc(repo.ShowBlogPage(app)))

	// public buttons
	mux.Post("/inventory/compare-vehicles", standardMiddleWare.ThenFunc(CompareVehicles))
	mux.Post("/power-sports/quick-quote", standardMiddleWare.ThenFunc(QuickQuote))
	mux.Post("/power-sports/test-drive", standardMiddleWare.ThenFunc(TestDrive))
	mux.Post("/power-sports/send-to-friend", standardMiddleWare.ThenFunc(SendFriend))

	// credit app
	mux.Get("/credit-application", standardMiddleWare.ThenFunc(CreditApp))
	mux.Post("/credit-application", standardMiddleWare.ThenFunc(PostCreditApp))

	// motorcycles
	mux.Get("/motorcycle-inventory", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/motorcyclesforsale/fredericton", http.StatusMovedPermanently)
	}))
	mux.Get("/motorcyclesforsale/fredericton", standardMiddleWare.ThenFunc(GetAllMotorcycles))
	mux.Get("/motorcyclesforsale/fredericton/:pageIndex", standardMiddleWare.ThenFunc(GetAllMotorcycles))

	// atvs - brute force
	mux.Get("/atv-brute-force", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/atvs/brute-force-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/atvs/brute-force-inventory", standardMiddleWare.ThenFunc(GetAllBruteForce))
	mux.Get("/atvs/brute-force-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllBruteForce))

	// atvs - teryx
	mux.Get("/atv-teryx", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/sidexsides/teryx-for-sale", http.StatusMovedPermanently)
	}))
	mux.Get("/sidexsides/teryx-for-sale", standardMiddleWare.ThenFunc(GetAllTeryx))
	mux.Get("/sidexsides/teryx-for-sale/:pageIndex", standardMiddleWare.ThenFunc(GetAllTeryx))

	// atvs - mule
	mux.Get("/atv-mule", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/utvs/kawasaki mules-for-sale", http.StatusMovedPermanently)
	}))
	mux.Get("/utvs/kawasaki-mules-for-sale", standardMiddleWare.ThenFunc(GetAllMule))
	mux.Get("/utvs/kawasaki-mules-for-sale/:pageIndex", standardMiddleWare.ThenFunc(GetAllMule))

	// jetski
	mux.Get("/watercraft-jetski", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/personalwatercraft/new-jetskis-forsale", http.StatusMovedPermanently)
	}))
	mux.Get("/personalwatercraft/new-jetskis-forsale", standardMiddleWare.ThenFunc(GetAllJetSki))
	mux.Get("/personalwatercraft/new-jetskis-forsale/:pageIndex", standardMiddleWare.ThenFunc(GetAllJetSki))

	// mercury outboards
	mux.Get("/outboard-motors-mercury-outboards", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/inventory/outboard-motors-mercury", http.StatusMovedPermanently)
	}))
	mux.Get("/inventory/outboard-motors-mercury", standardMiddleWare.ThenFunc(GetAllMercuryOutboards))
	mux.Get("/inventory/outboard-motors-mercury/:pageIndex", standardMiddleWare.ThenFunc(GetAllMercuryOutboards))

	// electric bikes
	mux.Get("/electric-bikes", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/ebikes/pedego-inventory", http.StatusMovedPermanently)
	}))
	mux.Get("/ebikes/pedego-inventory", standardMiddleWare.ThenFunc(GetAllElectricBikes))
	mux.Get("/ebikes/pedego-inventory/:pageIndex", standardMiddleWare.ThenFunc(GetAllElectricBikes))

	// scooters
	mux.Get("/vespa-piaggio-gas-and-electric-scooters-mopeds", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/scooters-mopeds/vespa-piaggio-gas-electric", http.StatusMovedPermanently)
	}))
	mux.Get("/scooters-mopeds/vespa-piaggio-gas-electric", standardMiddleWare.ThenFunc(GetAllScooters))
	mux.Get("/scooters-mopeds/vespa-piaggio-gas-electric/:pageIndex", standardMiddleWare.ThenFunc(GetAllScooters))

	// pontoon boats
	mux.Get("/pontoon-boats-bennington-and-crestliner", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/pontoon-boats/bennington-crestliner-dealer", http.StatusMovedPermanently)
	}))
	mux.Get("/pontoon-boats/bennington-crestliner-dealer", standardMiddleWare.ThenFunc(GetAllPontoonBoats))
	mux.Get("/pontoon-boats/bennington-crestliner-dealer/:pageIndex", standardMiddleWare.ThenFunc(GetAllPontoonBoats))

	// power boats
	mux.Get("/speed-boats-aluminum-boats-power-boats", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/boatsforsale/fishing-speed-jon-aluminum-power-glastron", http.StatusMovedPermanently)
	}))
	mux.Get("/boatsforsale/fishing-speed-jon-aluminum-power-glastron", standardMiddleWare.ThenFunc(GetAllPowerBoats))
	mux.Get("/boatsforsale/fishing-speed-jon-aluminum-power-glastron/:pageIndex", standardMiddleWare.ThenFunc(GetAllPowerBoats))

	// used
	mux.Get("/used-motorcycles-used-atv-used-boats-used-pontoons", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/used/used-motorbikes-atvs-wheelers-trailers", http.StatusMovedPermanently)
	}))
	mux.Get("/used/used-motorbikes-atvs-wheelers-trailers", standardMiddleWare.ThenFunc(GetAllUsedPowerSports))
	mux.Get("/used/used-motorbikes-atvs-wheelers-trailers/:pageIndex", standardMiddleWare.ThenFunc(GetAllUsedPowerSports))

	// used boats
	mux.Get("/used-boats-for-sale/pontoons-pwc-jetski-fishboats-boattrailers", standardMiddleWare.ThenFunc(GetAllUsedBoats))
	mux.Get("/used-boats-for-sale/pontoons-pwc-jetski-fishboats-boattrailers/:pageIndex", standardMiddleWare.ThenFunc(GetAllUsedBoats))

	// trailers
	mux.Get("/atv-trailers-boattrailers-utility-trailers", standardMiddleWare.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/newtrailers/boat-stv-utility", http.StatusMovedPermanently)
	}))
	mux.Get("/newtrailers/boat-stv-utility", standardMiddleWare.ThenFunc(GetAllTrailers))
	mux.Get("/newtrailers/boat-stv-utility/:pageIndex", standardMiddleWare.ThenFunc(GetAllTrailers))

	mux.Get("/:item/:productType/:prefix/:ID/:description", standardMiddleWare.ThenFunc(ShowItem))

	return mux, nil
}

// ClientInit gives us access to site values for client code.
func ClientInit(c config.AppConfig, p *driver.DB, r *handlers.DBRepo) {
	app = c
	repo = r

	vehicleModel = &clientdb.VehicleModel{DB: p.SQL}
	infoLog = app.InfoLog
	errorLog = app.ErrorLog
	parentDB = p
}
