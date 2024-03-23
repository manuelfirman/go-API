package service

import "github.com/manuelfirman/go-API/internal"

// NewWarehouseDefault creates a new instance of the warehouse service
type WarehouseDefault struct {
	// rp is the repository used by the service
	rp internal.WarehouseRepository
}

// NewWarehouseDefault creates a new instance of the warehouse service
func NewWarehouseDefault(rp internal.WarehouseRepository) *WarehouseDefault {
	return &WarehouseDefault{
		rp: rp,
	}
}

// GetAll returns all products. Returns an error if the operation fails.
func (w *WarehouseDefault) GetAll() (warehouses []internal.Warehouse, err error) {
	warehouses, err = w.rp.GetAll()
	if err != nil {
		switch err {
		case internal.ErrWarehouseRepositoryNotFound:
			err = internal.ErrWarehouseServiceNotFound
		default:
			err = internal.ErrWarehouseServiceUnknown
		}
		return
	}

	return
}

// Get returns a product by ID. Returns an error if the product is not found.
func (w *WarehouseDefault) Get(id int) (p internal.Warehouse, err error) {
	p, err = w.rp.Get(id)
	if err != nil {
		switch err {
		case internal.ErrWarehouseRepositoryNotFound:
			err = internal.ErrWarehouseServiceNotFound
		default:
			err = internal.ErrWarehouseServiceUnknown
		}
		return
	}

	return
}

// Save receives a product and saves it. It returns the ID of the product saved.
func (w *WarehouseDefault) Save(wh *internal.Warehouse) (warehouse internal.Warehouse, err error) {
	id, err := w.rp.Save(wh)
	if err != nil {
		switch err {
		case internal.ErrWarehouseRepositoryDuplicated:
			err = internal.ErrWarehouseServiceDuplicated
		case internal.ErrWarehouseRepositoryForeignKey:
			err = internal.ErrWarehouseServiceForeignKey
		default:
			err = internal.ErrWarehouseServiceUnknown
		}
		return
	}

	wh.ID = id
	warehouse = *wh

	return
}

// Update receives a product and updates it. Returns an error if the product is not found.
func (w *WarehouseDefault) Update(p *internal.Warehouse) (err error) {
	err = w.rp.Update(p)
	if err != nil {
		switch err {
		case internal.ErrWarehouseRepositoryNotFound:
			err = internal.ErrWarehouseServiceNotFound
		case internal.ErrWarehouseRepositoryNothingToUpdate:
			err = internal.ErrWarehouseServiceNothingToUpdate
		default:
			err = internal.ErrWarehouseServiceUnknown
		}
	}

	return
}

// Delete deletes a product by ID. Returns an error if the product is not found.
func (w *WarehouseDefault) Delete(id int) (err error) {
	err = w.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrWarehouseRepositoryNotFound:
			err = internal.ErrWarehouseServiceNotFound
		default:
			err = internal.ErrWarehouseServiceUnknown
		}
	}

	return
}
