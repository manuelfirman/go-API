package service

import (
	"fmt"

	"github.com/manuelfirman/go-API/internal"
)

// NewBuyerDefault creates a new instance of the buyer service
func NewBuyerDefault(rp internal.BuyerRepository) *BuyerDefault {
	return &BuyerDefault{
		rp: rp,
	}
}

// BuyerDefault is the default implementation of the buyer service
type BuyerDefault struct {
	// rp is the buyer repository
	rp internal.BuyerRepository
}

// GetAll returns all buyers. Returns an error if the operation fails.
func (s *BuyerDefault) GetAll() (buyers []internal.Buyer, err error) {
	buyers, err = s.rp.GetAll()
	if err != nil {
		switch err {
		case internal.ErrBuyerRepository:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerService, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerServiceUnknown, err)
		}

		return
	}

	return
}

// Get returns a buyer by ID. Returns an error if the buyer is not found.
func (s *BuyerDefault) Get(id int) (buyer internal.Buyer, err error) {
	buyer, err = s.rp.Get(id)
	if err != nil {
		switch err {
		case internal.ErrBuyerRepositoryNotFound:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerServiceNotFound, err)
		case internal.ErrBuyerRepository:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerService, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerServiceUnknown, err)
		}

		return
	}

	return
}

// Save saves the given buyer. Returns an error if the operation fails.
func (s *BuyerDefault) Save(buyer *internal.Buyer) (err error) {
	// validate buyer
	if err = ValidateBuyer(buyer); err != nil {
		return
	}

	// save buyer
	err = s.rp.Save(buyer)
	if err != nil {
		switch err {
		case internal.ErrBuyerRepositoryDuplicated:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerServiceDuplicated, err)
		case internal.ErrBuyerRepository:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerService, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerServiceUnknown, err)
		}

		return
	}

	return
}

// Update updates the given buyer. Returns an error if the operation fails.
func (s *BuyerDefault) Update(buyer *internal.Buyer) (err error) {
	// validate buyer
	if err = ValidateBuyer(buyer); err != nil {
		return
	}

	// update buyer
	err = s.rp.Update(buyer)
	if err != nil {
		switch err {
		case internal.ErrBuyerRepositoryDuplicated:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerServiceDuplicated, err)
		case internal.ErrBuyerRepository:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerService, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerServiceUnknown, err)
		}
		return
	}

	return
}

// Delete deletes the buyer with the given ID. Returns an error if the operation fails.
func (s *BuyerDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrBuyerRepositoryNotFound:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerServiceNotFound, err)
		case internal.ErrBuyerRepositoryFK:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerServiceFK, err)
		case internal.ErrBuyerRepository:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerService, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrBuyerServiceUnknown, err)
		}

		return
	}

	return
}

// ValidateBuyer validates a buyer
func ValidateBuyer(buyer *internal.Buyer) (err error) {
	// - validate required fields
	if (*buyer).CardNumberID <= 0 {
		return fmt.Errorf("%w: card_number_id", internal.ErrBuyerServiceFieldRequired)
	}
	if (*buyer).FirstName == "" || len((*buyer).FirstName) < 3 {
		return fmt.Errorf("%w: first_name", internal.ErrBuyerServiceFieldRequired)
	}
	if (*buyer).LastName == "" || len((*buyer).LastName) < 3 {
		return fmt.Errorf("%w: last_name", internal.ErrBuyerServiceFieldRequired)
	}

	return
}
