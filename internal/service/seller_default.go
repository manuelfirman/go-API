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
	products, err = s.rp.GetAll()
	if err != nil {
		switch err {
		case internal.ErrSellerRepositoryNotFound:
			err = internal.ErrSellerServiceNotFound
		case internal.ErrSellerRepositoryTransaction, internal.ErrSellerRepositoryConn:
			err = internal.ErrSellerServiceDB
		default:
			err = internal.ErrSellerServiceUnknown
		}
		return
	}

	return
}

// Get returns a product by ID. Returns an error if the product is not found.
func (s *SellerDefault) Get(id int) (p internal.Seller, err error) {
	p, err = s.rp.Get(id)
	if err != nil {
		switch err {
		case internal.ErrSellerRepositoryNotFound:
			err = internal.ErrSellerServiceNotFound
		case internal.ErrSellerRepositoryTransaction, internal.ErrSellerRepositoryConn:
			err = internal.ErrSellerServiceDB
		default:
			err = internal.ErrSellerServiceUnknown
		}
		return
	}

	return
}

// Save receives a product and saves it. It returns the ID of the product saved.
func (s *SellerDefault) Save(sell *internal.Seller) (seller internal.Seller, err error) {
	id, err := s.rp.Save(sell)
	if err != nil {
		switch err {
		case internal.ErrSellerRepositoryDuplicated:
			err = internal.ErrSellerServiceDuplicated
		case internal.ErrSellerRepositoryLocalityIdNotFound:
			err = internal.ErrSellerServiceForeignKey
		case internal.ErrSellerRepositoryTransaction, internal.ErrSellerRepositoryConn:
			err = internal.ErrSellerServiceDB
		default:
			err = internal.ErrSellerServiceUnknown
		}
		return
	}

	seller = *sell
	seller.ID = id
	return
}

// Update receives a product and updates it. Returns an error if the product is not found.
func (s *SellerDefault) Update(p *internal.Seller) (err error) {
	err = s.rp.Update(p)
	if err != nil {
		switch err {
		case internal.ErrSellerRepositoryNotFound:
			err = internal.ErrSellerServiceNotFound
		case internal.ErrSellerRepositoryNothingToUpdate:
			err = internal.ErrSellerServiceNothingToUpdate
		case internal.ErrSellerRepositoryTransaction, internal.ErrSellerRepositoryConn:
			err = internal.ErrSellerServiceDB
		default:
			err = internal.ErrSellerServiceUnknown
		}
		return
	}

	return
}

// Delete receives a product ID and deletes it. Returns an error if the product is not found.
func (s *SellerDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrSellerRepositoryNotFound:
			err = internal.ErrSellerServiceNotFound
		case internal.ErrSellerRepositoryTransaction, internal.ErrSellerRepositoryConn:
			err = internal.ErrSellerServiceDB
		default:
			err = internal.ErrSellerServiceUnknown
		}
		return
	}

	return
}

// GetRecordsByProductReport returns the product records.
// func (s *SellerDefault) GetRecordsByProductReport(id int) (products []internal.Product, err error) {
// 	products, err = s.rp.GetRecordsByProductReport(id)
// 	if err != nil {
// 		switch err {
// 		case internal.ErrSellerRepositoryNotFound:
// 			err = internal.ErrSellerServiceNotFound
// 		case internal.ErrSellerRepositoryTransaction, internal.ErrSellerRepositoryConn:
// 			err = internal.ErrSellerServiceDB
// 		default:
// 			err = internal.ErrSellerServiceUnknown
// 		}
// 		return
// 	}

// 	return
// }
