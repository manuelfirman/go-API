package internal

import "errors"

// Errors
var (
	// ErrRepositoryNotFound is returned when a section is not found.
	ErrProductRepositoryNotFound = errors.New("products repository: product not found")
	// ErrRepositoryDuplicated is returned when a section already exists.
	ErrProductRepositoryDuplicated = errors.New("products repository: product already exists")
	// ErrRepositoryTransaction is returned when an operation is performed on a commited transaction
	ErrProductRepositoryTransaction = errors.New("products repository: transaction error")
	// ErrRepositoryConn is returned when the DB connection is done.
	ErrProductRepositoryConn = errors.New("products repository: connection is done")
	// ErrRepositoryUnknown is returned when an unknown error occurs.
	ErrProductRepositoryUnknown = errors.New("products repository: unknown error")
	// ErrProductRepositoryNothingToUpdate is returned when there is nothing to update.
	ErrProductRepositoryNothingToUpdate = errors.New("products repository: nothing to update")
	// ErrProductRepositoryForeignKey is returned when a product couldn't be deleted because of a foreign key constraint.
	ErrProductRepositoryForeignKey = errors.New("products repository: product couldn't be deleted because foreign key constraint")
)

// Repository encapsulates the storage of a Product.
type ProductRepository interface {
	// GetAll returns all the products.
	GetAll() ([]Product, error)
	// Get returns the product with the given id.
	Get(id int) (Product, error)
	// Save saves the product in the storage.
	Save(p *Product) (int64, error)
	// Update updates the product in the storage.
	Update(p *Product) error
	// Delete deletes the product with the given id from the storage.
	Delete(id int) error
	// GetRecordsByProductReport returns the product records.
	GetRecordsByProductReport(id int) ([]Product, error)
}
