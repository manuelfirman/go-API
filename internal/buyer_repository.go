package internal

import "errors"

var (
	// ErrBuyerRepositoryNotFound is returned when the buyer is not found
	ErrBuyerRepositoryNotFound = errors.New("repository: buyer not found")
	// ErrBuyerRepositoryDuplicated is returned when the buyer already exists
	ErrBuyerRepositoryDuplicated = errors.New("repository: buyer already exists")
	//Generic error for repository
	ErrBuyerRepository = errors.New("repository: internal error")
	//ErrBuyerRepositoryFK is returned when the buyer has purchase orders
	ErrBuyerRepositoryFK = errors.New("repository: buyer has purchase orders")
	// ErrBuyerRepositoryNoData is returned when the buyer has no data
	ErrBuyerRepositoryNoData = errors.New("repository: buyer table has no data")
)

// BuyerRepository is an interface that contains the methods that the buyer repository should support
type BuyerRepository interface {
	// FindAll returns all the buyers
	FindAll() ([]Buyer, error)
	// FindByID returns the buyer with the given ID
	FindByID(id int) (Buyer, error)
	// Save saves the given buyer
	Save(buyer *Buyer) error
	// Update updates the given buyer
	Update(buyer *Buyer) error
	// Delete deletes the buyer with the given ID
	Delete(id int) error
	// ReportPurchaseOrders returns the report of the purchase orders of the buyer with the given ID or all the buyers if the ID is 0
	// ReportPurchaseOrders(id int) (report []PurchaseOrderReport, err error)
}
