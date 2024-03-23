package service

import "github.com/manuelfirman/go-API/internal"

// NewProductDefault creates a new instance of the product service
func NewSellerDefault(rp internal.SellerRepository) *SellerDefault {
	return &SellerDefault{
		rp: rp,
	}
}

// ProductDefault is the default implementation of the product service
type SellerDefault struct {
	// rp is the repository used by the service
	rp internal.SellerRepository
}

// GetAll returns all products. Returns an error if the operation fails.
func (s *SellerDefault) GetAll() (products []internal.Seller, err error) {
	return
}

// Get returns a product by ID. Returns an error if the product is not found.
func (s *SellerDefault) Get(id int) (p internal.Seller, err error) {
	return
}

// Save receives a product and saves it. It returns the ID of the product saved.
func (s *SellerDefault) Save(p *internal.Seller) (prod internal.Product, err error) {
	return
}

// Update receives a product and updates it. Returns an error if the product is not found.
func (s *SellerDefault) Update(p *internal.Seller) (err error) {
	return
}

// Delete receives a product ID and deletes it. Returns an error if the product is not found.
func (s *SellerDefault) Delete(id int) (err error) {
	return
}

// GetRecordsByProductReport returns the product records.
func (s *SellerDefault) GetRecordsByProductReport(id int) (products []internal.Product, err error) {
	return
}
