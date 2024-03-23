package internal

import "errors"

var (
	// ErrSellerRepositoryNotFound is returned when the seller is not found
	ErrSellerRepositoryNotFound = errors.New("repository: seller not found")
	// ErrSellerRepositoryDuplicated is returned when the seller already exists
	ErrSellerRepositoryDuplicated = errors.New("repository: seller already exists")
	// ErrSellerRepositoryLocalityIdNotFound is returned when the locality id does not exist
	ErrSellerRepositoryLocalityIdNotFound = errors.New("repository: locality id does not exist")
)

// SellerRepository is an interface that contains the methods that the seller repository should support
type SellerRepository interface {
	// GetAll returns all the sellers
	GetAll() ([]Seller, error)
	// Get returns the seller with the given ID
	Get(id int) (Seller, error)
	// Save saves the given seller
	Save(seller *Seller) error
	// Update updates the given seller
	Update(id int, seller *Seller) error
	// Delete deletes the seller with the given ID
	Delete(id int) error
}
