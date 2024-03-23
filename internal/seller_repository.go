package internal

import "errors"

var (
	// ErrSellerRepositoryNotFound is returned when the seller is not found
	ErrSellerRepositoryNotFound = errors.New("sellers repository: seller not found")
	// ErrSellerRepositoryDuplicated is returned when the seller already exists
	ErrSellerRepositoryDuplicated = errors.New("sellers repository: seller already exists")
	// ErrSellerRepositoryLocalityIdNotFound is returned when the locality id does not exist
	ErrSellerRepositoryLocalityIdNotFound = errors.New("sellers repository: locality id does not exist")
	// ErrSellerRepositoryTransaction is returned when there is an error with the transaction
	ErrSellerRepositoryTransaction = errors.New("sellers repository: transaction error")
	// ErrSellerRepositoryConn is returned when there is an error with the connection
	ErrSellerRepositoryConn = errors.New("sellers repository: connection error")
	// ErrSellerRepositoryForeignKey is returned when there is an error with the foreign key
	ErrSellerRepositoryForeignKey = errors.New("sellers repository: foreign key error")
	// ErrSellerRepositoryUnknown is returned when there is an unknown error
	ErrSellerRepositoryUnknown = errors.New("sellers repository: unknown error")
	// ErrSellerRepositoryNothingToUpdate is returned when there is nothing to update
	ErrSellerRepositoryNothingToUpdate = errors.New("sellers repository: nothing to update")
)

// SellerRepository is an interface that contains the methods that the seller repository should support
type SellerRepository interface {
	// GetAll returns all the sellers
	GetAll() ([]Seller, error)
	// Get returns the seller with the given ID
	Get(id int) (Seller, error)
	// Save saves the given seller
	Save(seller *Seller) (int, error)
	// Update updates the given seller
	Update(seller *Seller) error
	// Delete deletes the seller with the given ID
	Delete(id int) error
}
