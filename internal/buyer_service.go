package internal

import "errors"

var (
	//ErrBuyerFieldRequired is returned when the buyer field is required
	ErrBuyerServiceFieldRequired = errors.New("service: buyer field is required")
	// ErrBuyerServiceNotFound is returned when the buyer is not found
	ErrBuyerServiceNotFound = errors.New("service: buyer not found")
	// ErrBuyerServiceDuplicated is returned when the buyer already exists
	ErrBuyerServiceDuplicated = errors.New("service: buyer already exists")
	//Generic error for service
	ErrBuyerService = errors.New("service: internal error")
	//ErrBuyerServiceFK is returned when the buyer has purchase orders
	ErrBuyerServiceFK = errors.New("service: buyer has purchase orders")
)

// BuyerService is an interface that contains the methods that the buyer service should support
type BuyerService interface {
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
