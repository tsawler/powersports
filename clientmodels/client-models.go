package clientmodels

import "time"

type Vehicle struct {
	ID              int
	StockNo         string
	Cost            float32
	Vin             string
	Odometer        int
	Year            int
	Trim            string
	VehicleType     int
	Body            string
	SeatingCapacity string
	DriveTrain      string
	Engine          string
	ExteriorColour  string
	InteriorColour  string
	Transmission    string
	Options         string
	ModelNumber     string
	TotalMSR        float32
	Status          int
	Description     string
	VehicleMakesID  int
	VehicleModelsID int
	HandPicked      int
	Used            int
	PriceForDisplay string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
