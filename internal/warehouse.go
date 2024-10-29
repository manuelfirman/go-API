package internal

// Warehouse is a struct that contains the warehouse's information
type Warehouse struct {
	// ID is the unique identifier of the warehouse
	ID int
	// WarehouseCode is the unique code of the warehouse
	WarehouseCode string
	// Address is the address of the warehouse
	Address string
	// Telephone is the telephone number of the warehouse
	Telephone string
	// MinimumCapacity is the minimum capacity of the warehouse
	MinimumCapacity int
	// MinimumTemperature is the minimum temperature that can be maintained in the warehouse
	MinimumTemperature float64
	// LocalityId is the id of the locality where the warehouse is located
	LocalityId string
}
