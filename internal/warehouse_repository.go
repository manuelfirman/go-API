package internal

import "errors"

var (
	// ErrWarehouseRepositoryNotFound is returned when a warehouse is not found.
	ErrWarehouseRepositoryNotFound = errors.New("warehouse repository: warehouse not found")
	// ErrWarehouseRepositoryDuplicated is returned when a warehouse already exists.
	ErrWarehouseRepositoryDuplicated = errors.New("warehouse repository: warehouse code already exists")
	// ErrWarehouseRepositoryUnknown is returned when an unknown error occurs.
	ErrWarehouseRepositoryUnknown = errors.New("warehouse repository: unknown error")
	// ErrWarehouseRepositoryForeignKey is returned when a warehouse couldn't be deleted because of a foreign key constraint.
	ErrWarehouseRepositoryForeignKey = errors.New("warehouse repository: warehouse foreign key constraint")
	// ErrWarehouseRepositoryNothingToUpdate is returned when there is nothing to update.
	ErrWarehouseRepositoryNothingToUpdate = errors.New("warehouse repository: nothing to update")
)

// WarehouseRepository is an interface that contains the methods that the warehouse repository should support
type WarehouseRepository interface {
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
