package internal

import "errors"

// Errors
var (
	// ErrServiceNotFound is returned when a product is not found.
	ErrProductServiceNotFound = errors.New("products service: product not found")
	// ErrServiceDBError is returned when a database connection or transaction error occurs.
	ErrProductServiceDBError = errors.New("products service: database error")
	// ErrServiceUnkown is returned when an unknown error occurs.
	ErrProductServiceUnkown = errors.New("products service: unknown error")
	// ErrServiceProductCodeInUse is returned when a product's code is already in use.
	ErrProductServiceDuplicated = errors.New("products service: product code already in use")
	// ErrServiceInconsistentData is returned when a map's type is not consistent with the domain
	ErrProductServiceInconsistentData = errors.New("products service: inconsistent data")
	// ErrInvalidContent is returned when the content is invalid
	ErrProductServuceInvalidContent = errors.New("products service: invalid content")
	// ErrProductServiceNothingToUpdate is returned when there is nothing to update.
	ErrProductServiceNothingToUpdate = errors.New("products service: nothing to update")

	ErrProductServiceForeignKey = errors.New("products service: product couldn't be deleted because foreign key constraint")
)

type ProductService interface {
	// GetAll returns all products.
	GetAll() ([]Product, error)
	// Get returns a product by ID.
	Get(id int) (Product, error)
	// Save saves a new product.
	Save(p *Product) (Product, error)
	// Update updates a product by ID.
	Update(p *Product) error
	// Delete deletes a product by ID.
	Delete(id int) error
	// GetRecordsByProductReport returns a report of the product records.
	GetRecordsByProductReport(id int) ([]Product, error)
}
