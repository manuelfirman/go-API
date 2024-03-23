package internal

import "errors"

var (
	// ErrSellerRepositoryNotFound is returned when the seller is not found
	ErrSellerServiceNotFound = errors.New("repository: seller not found")
	// ErrSellerRepositoryDuplicated is returned when the seller already exists
	ErrSellerServiceDuplicated = errors.New("repository: seller already exists")
	// ErrSellerRepositoryLocalityIdNotFound is returned when the locality id does not exist
	ErrSellerServiceLocalityIdNotFound = errors.New("repository: locality id does not exist")
	// ErrSellerServiceMissingFields is returned when there are invalid fields in the request
	ErrSellerServiceInvalidFields = errors.New("service: seller has invalid fields")
)

// SellerService is an interface that contains the methods that the seller service should support
type SellerService interface {
	// GetAll returns all the sellers
	GetAll() ([]Seller, error)
	// Get returns the seller with the given ID
	Get(id int) (Seller, error)
	// Save saves the given seller
	Save(cid int, seller *Seller) error
	// Update updates the given seller
	Update(id int, seller *Seller) error
	// Delete deletes the seller with the given ID
	Delete(id int) error
}
