package internal

import "errors"

var (
	// ErrWarehouseServiceNotFound is returned when a warehouse is not found.
	ErrWarehouseServiceNotFound = errors.New("warehouse service: warehouse not found")
	// ErrWarehouseServiceDuplicated is returned when a warehouse already exists.
	ErrWarehouseServiceDuplicated = errors.New("warehouse service: warehouse code already exists")
	// ErrWarehouseServiceUnknown is returned when an unknown error occurs.
	ErrWarehouseServiceUnknown = errors.New("warehouse service: unknown error")
)

// WarehouseService is an interface that contains the methods that the warehouse service should support
type WarehouseService interface {
	// FindAll returns all the warehouses
	FindAll() ([]Warehouse, error)
	// FindByID returns the warehouse with the given ID
	FindByID(id int) (Warehouse, error)
	// Save saves the given warehouse
	Save(warehouse *Warehouse) error
	// Update updates the given warehouse
	Update(warehouse *Warehouse) error
	// Delete deletes the warehouse with the given ID
	Delete(id int) error
}
