package clienthandlers

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"github.com/tsawler/goblender/client/clienthandlers/clientdb"
	"github.com/tsawler/goblender/pkg/config"
	"github.com/tsawler/goblender/pkg/middleware"
	"log"
)

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func ClientRoutes(mux *pat.PatternServeMux, standardMiddleWare, dynamicMiddleware alice.Chain) (*pat.PatternServeMux, error) {

	mux.Get("/motorcycle-inventory", standardMiddleWare.ThenFunc(GetAllMotorcycles))

	mux.Get("/custom/my-test-route", standardMiddleWare.ThenFunc(TestHandler))
	mux.Get("/custom/protected-route", standardMiddleWare.Append(middleware.Auth).ThenFunc(TestProtectedHandler))

	return mux, nil
}

// ClientInit gives us access to site values for client code.
func ClientInit(c config.AppConfig) {
	app = c
	conn := app.Connections["wheels"]
	vehicleModel = &clientdb.VehicleModel{DB: conn}
	infoLog = app.InfoLog
	errorLog = app.ErrorLog
}
