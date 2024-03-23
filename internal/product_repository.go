package internal

import "errors"

// Errors
var (
	// ErrRepositoryNotFound is returned when a section is not found.
	ErrRepositoryNotFound = errors.New("repository: product not found")
	// ErrRepositoryTransaction is returned when an operation is performed on a commited transaction
	ErrRepositoryTransaction = errors.New("repository: transaction error")
	// ErrRepositoryConn is returned when the DB connection is done.
	ErrRepositoryConn = errors.New("repository: connection is done")
	// ErrRepositoryUnknown is returned when an unknown error occurs.
	ErrRepositoryUnknown = errors.New("repository: unknown error")
)

// Repository encapsulates the storage of a Product.
type ProductRepository interface {
	// GetAll returns all the products.
	GetAll() ([]Product, error)
	// Get returns the product with the given id.
	Get(id int) (Product, error)
	// Exists returns true if the product with the given product code exists.
	Exists(productCode string) bool
	// Save saves the product in the storage.
	Save(p *Product) (int, error)
	// Update updates the product in the storage.
	Update(p *Product) error
	// Delete deletes the product with the given id from the storage.
	Delete(id int) error
	// GetRecordsByProductReport returns the product records.
	GetRecordsByProductReport(id int) ([]Product, error)
}
