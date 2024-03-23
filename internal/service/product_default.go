package service

import "github.com/manuelfirman/go-API/internal"

// NewProductDefault creates a new instance of the product service
func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

// ProductDefault is the default implementation of the product service
type ProductDefault struct {
	// rp is the repository used by the service
	rp internal.ProductRepository
}

// GetAll returns all products. Returns an error if the operation fails.
func (s *ProductDefault) GetAll() (products []internal.Product, err error) {
	products, err = s.rp.GetAll()
	if err != nil {
		switch err {
		case internal.ErrProductRepositoryNotFound:
			err = internal.ErrProductServiceNotFound
		case internal.ErrProductRepositoryTransaction, internal.ErrProductRepositoryConn:
			err = internal.ErrProductServiceDBError
		default:
			err = internal.ErrProductServiceUnkown
		}

		return
	}

	return

}

// Get returns a product by ID. Returns an error if the product is not found.
func (s *ProductDefault) Get(id int) (p internal.Product, err error) {
	p, err = s.rp.Get(id)
	if err != nil {
		switch err {
		case internal.ErrProductRepositoryNotFound:
			err = internal.ErrProductServiceNotFound
		case internal.ErrProductRepositoryTransaction, internal.ErrProductRepositoryConn:
			err = internal.ErrProductServiceDBError
		default:
			err = internal.ErrProductServiceUnkown
		}

		return
	}

	return
}

// Save receives a product and saves it. It returns the ID of the product saved.
func (s *ProductDefault) Save(p *internal.Product) (prod internal.Product, err error) {
	id, err := s.rp.Save(p)
	if err != nil {
		switch err {
		case internal.ErrProductRepositoryDuplicated:
			err = internal.ErrProductServiceDuplicated
		case internal.ErrProductRepositoryTransaction, internal.ErrProductRepositoryConn:
			err = internal.ErrProductServiceDBError
		case internal.ErrSellerRepositoryNotFound:
			err = internal.ErrSellerServiceNotFound
		default:
			err = internal.ErrProductServiceUnkown
		}

		return
	}

	p.ID = id
	prod = *p

	return
}

// Update receives a product and updates it. Returns an error if the product is not found.
func (s *ProductDefault) Update(p *internal.Product) (err error) {
	err = s.rp.Update(p)
	if err != nil {
		switch err {
		case internal.ErrProductRepositoryNotFound:
			err = internal.ErrProductServiceNotFound
		case internal.ErrProductRepositoryTransaction, internal.ErrProductRepositoryConn:
			err = internal.ErrProductServiceDBError
		case internal.ErrProductRepositoryNothingToUpdate:
			err = internal.ErrProductServiceNothingToUpdate
		default:
			err = internal.ErrProductServiceUnkown
		}

		return
	}

	return
}

// Delete receives a product ID and deletes it. Returns an error if the product is not found.
func (s *ProductDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrProductRepositoryNotFound:
			err = internal.ErrProductServiceNotFound
		case internal.ErrProductRepositoryTransaction, internal.ErrProductRepositoryConn:
			err = internal.ErrProductServiceDBError
		case internal.ErrProductRepositoryForeignKey:
			err = internal.ErrProductServiceForeignKey
		default:
			err = internal.ErrProductServiceUnkown
		}

		return
	}

	return
}

// GetRecordsByProductReport returns the product records.
func (s *ProductDefault) GetRecordsByProductReport(id int) (products []internal.Product, err error) {
	return
}
