package clientmodels

import "time"

// Vehicle holds a vehicle
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

// Option holds vehicle options
type Option struct {
	ID         int
	OptionName string
	Active     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// VehicleOption holds option for a given vehicle
type VehicleOption struct {
	ID         int
	VehicleID  int
	OptionID   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	OptionName string
}

// Make is vehicle make (i.e. Volvo)
type Make struct {
	ID        int
	Make      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Model is vehicle model (i.e. Camry)
type Model struct {
	ID        int
	Model     string
	MakeID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Image is a vehicle image
type Image struct {
	ID        int
	VehicleID int
	Image     string
	SortOrder int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Video holds a video
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

// SalesStaff holds sales people
type SalesStaff struct {
	ID    int
	Name  string
	Slug  string
	Email string
	Phone string
	Image string
}

// CreditApp holds a credit application
type CreditApp struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   string
	City      string
	Province  string
	Zip       string
	Vehicle   string
	Processed int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TestDrive holds a test drive
type TestDrive struct {
	ID            int
	UsersName     string
	Email         string
	Phone         string
	PreferredDate string
	PreferredTime string
	VehicleID     int
	Processed     int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
