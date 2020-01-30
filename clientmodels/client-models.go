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
	Make            Make
	Model           Model
	Video           Video
	Images          []*Image
	VehicleOptions  []*VehicleOption
}

type Option struct {
	ID         int
	OptionName string
	Active     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type VehicleOption struct {
	ID         int
	VehicleID  int
	OptionID   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	OptionName string
}

type Make struct {
	ID        int
	Make      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Model struct {
	ID        int
	Model     string
	MakeID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Image struct {
	ID        int
	VehicleID int
	Image     string
	SortOrder int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Video struct {
	ID                      int
	VideoName               string
	FileName                string
	Public                  int
	Description             string
	CategoryID              int
	SortOrder               int
	Thumb                   string
	ConvertedForStreamingAt time.Time
	Duration                int
	Is360                   int
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
