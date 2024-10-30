package internal

// Section is a struct that contains the section's information
type Section struct {
	// ID is the unique identifier of the section
	ID int
	// SectionNumber is the number of the section
	SectionNumber int
	// CurrentTemperature is the current temperature of the section
	CurrentTemperature float64
	// MinimumTemperature is the minimum temperature that can be maintained in the section
	MinimumTemperature float64
	// CurrentCapacity is the current capacity of the section
	CurrentCapacity int
	// MinimumCapacity is the minimum capacity of the section
	MinimumCapacity int
	// MaximumCapacity is the maximum capacity of the section
	MaximumCapacity int
	// WarehouseID is the unique identifier of the warehouse to which the section belongs
	WarehouseID int
	// ProductTypeID is the unique identifier of the type of product stored in the section
	ProductTypeID int
}
