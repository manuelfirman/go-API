package internal

import "errors"

var (
	// ErrWarehouseServiceNotFound is returned when a warehouse is not found.
	ErrWarehouseServiceNotFound = errors.New("warehouse service: warehouse not found")
	// ErrWarehouseServiceDuplicated is returned when a warehouse already exists.
	ErrWarehouseServiceDuplicated = errors.New("warehouse service: warehouse code already exists")
	// ErrWarehouseServiceUnknown is returned when an unknown error occurs.
	ErrWarehouseServiceUnknown = errors.New("warehouse service: unknown error")
	// ErrWarehouseServiceForeignKey is returned when a warehouse couldn't be deleted because of a foreign key constraint.
	ErrWarehouseServiceForeignKey = errors.New("warehouse service: warehouse foreign key constraint")
	// ErrWarehouseServiceNothingToUpdate is returned when there is nothing to update.
	ErrWarehouseServiceNothingToUpdate = errors.New("warehouse service: nothing to update")
)

// WarehouseService is an interface that contains the methods that the warehouse service should support
type WarehouseService interface {
	// GetAll returns all the warehouses
	GetAll() ([]Warehouse, error)
	// Get returns the warehouse with the given ID
	Get(id int) (Warehouse, error)
	// Save saves the given warehouse
	Save(warehouse *Warehouse) (Warehouse, error)
	// Update updates the given warehouse
	Update(warehouse *Warehouse) error
	// Delete deletes the warehouse with the given ID
	Delete(id int) error
}
