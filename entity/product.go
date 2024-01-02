package entity

type Product struct {
	ID                      uint
	Name                    string
	Description             string
	Presentation            string
	HeightMM                *int
	WidthMM                 *int
	LengthMM                *int
	Quantity                int
	StorageCondition        string
	GrossWeightG            *int
	NetWeightG              *int
	Brand                   string
	Ncm                     string
	Gtins                   []string
	ManufacturerID          uint
	GroupID                 uint
	AssociatedConditionID   *uint
	UnitOfMeasurementSymbol string
	UmbrellaItemID          *uint
	Availabilities          []Availability
}

type Manufacturer struct {
	ID   uint
	Name string
}

type ManufacturerProducts struct {
	Manufacturer
	Products []Product
}

type Measurement struct {
	Symbol string
	Name   string
}

type Group struct {
	ID            uint
	Name          string
	Products      []Product
	ParentGroupID *uint
	ParentGroup   []Group
}

type AvailabilityStatus string

const (
	Available    AvailabilityStatus = "AVAILABLE"
	Unavailable  AvailabilityStatus = "UNAVAILABLE"
	Discontinued AvailabilityStatus = "DISCONTINUED"
)

type Availability struct {
	ID            uint
	CompanyUnitID uint
	ProductID     uint
	Availability  AvailabilityStatus
}
