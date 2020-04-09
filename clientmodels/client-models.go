package clientmodels

import "time"

// Vehicle holds a vehicle
type Vehicle struct {
	ID              int              `xml:"ID"`
	StockNo         string           `xml:"StockNo"`
	Cost            float32          `xml:"Cost"`
	Vin             string           `xml:"-"`
	Odometer        int              `xml:"-"`
	Year            int              `xml:"Year"`
	Trim            string           `xml:"-"`
	VehicleType     int              `xml:"-"`
	Body            string           `xml:"-"`
	SeatingCapacity string           `xml:"-"`
	DriveTrain      string           `xml:"-"`
	Engine          string           `xml:"-"`
	ExteriorColour  string           `xml:"-"`
	InteriorColour  string           `xml:"-"`
	Transmission    string           `xml:"-"`
	Options         string           `xml:"-"`
	ModelNumber     string           `xml:"-"`
	TotalMSR        float32          `xml:"-"`
	Status          int              `xml:"-"`
	Description     string           `xml:"Description"`
	VehicleMakesID  int              `xml:"-"`
	VehicleModelsID int              `xml:"-"`
	HandPicked      int              `xml:"-"`
	Used            int              `xml:"-"`
	PriceForDisplay string           `xml:"-"`
	CreatedAt       time.Time        `xml:"-"`
	UpdatedAt       time.Time        `xml:"-"`
	Make            Make             `xml:"-"`
	Model           Model            `xml:"-"`
	Video           Video            `xml:"-"`
	Images          []*Image         `xml:"-"`
	VehicleOptions  []*VehicleOption `xml:"-"`
	VehicleMake     string           `xml:"make"`
	VehicleModel    string           `xml:"model"`
	Photo1          string           `xml:"Photo1"`
	Photo2          string           `xml:"Photo2"`
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
	Image     string `xml:"image"`
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

// QuickQuote holds a quick quote
type QuickQuote struct {
	ID        int
	UsersName string
	Email     string
	Phone     string
	VehicleID int
	Processed int
	CreatedAt time.Time
	UpdatedAt time.Time
}
