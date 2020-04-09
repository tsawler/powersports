package clienthandlers

import (
	"encoding/xml"
	"github.com/tsawler/goblender/client/clienthandlers/clientmodels"
	"github.com/tsawler/goblender/pkg/helpers"
	"net/http"
)

type Xml struct {
	XMLName  xml.Name               `xml:"vehicles"`
	Version  string                 `xml:"version,attr"`
	Vehicles []clientmodels.Vehicle `xml:"Vehicle"`
}

func FacebookMarketplaceFeed(w http.ResponseWriter, r *http.Request) {

	v := &Xml{Version: "1"}
	vehicles, err := vehicleModel.GetAllVehiclesForSale()
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	v.Vehicles = vehicles
	feed, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(xml.Header))
	w.Write(feed)
}
