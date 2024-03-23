package internal

import "errors"

var (
	// ErrSellerServiceNotFound is returned when the seller is not found
	ErrSellerServiceNotFound = errors.New("sellers service: seller not found")
	// ErrSellerServiceDuplicated is returned when the seller already exists
	ErrSellerServiceDuplicated = errors.New("sellers service: seller already exists")
	// ErrSellerServiceLocalityIdNotFound is returned when the locality id does not exist
	ErrSellerServiceLocalityIdNotFound = errors.New("sellers service: locality id does not exist")
	// ErrSellerServiceTransaction is returned when there is an error with the transaction
	ErrSellerServiceDB = errors.New("sellers service: database error")
	// ErrSellerServiceConn is returned when there is an error with the connection

	ErrSellerServiceForeignKey = errors.New("sellers service: foreign key error")
	// ErrSellerServiceUnknown is returned when there is an unknown error
	ErrSellerServiceUnknown = errors.New("sellers service: unknown error")
	// ErrSellerServiceNothingToUpdate is returned when there is nothing to update
	ErrSellerServiceNothingToUpdate = errors.New("sellers service: nothing to update")
)

// SellerService is an interface that contains the methods that the seller service should support
type SellerService interface {
	// GetAll returns all the sellers
	GetAll() ([]Seller, error)
	// Get returns the seller with the given ID
	Get(id int) (Seller, error)
	// Save saves the given seller
	Save(seller *Seller) (Seller, error)
	// Update updates the given seller
	Update(seller *Seller) error
	// Delete deletes the seller with the given ID
	Delete(id int) error
}
