package service

import (
	"fmt"

	"github.com/manuelfirman/go-API/internal"
)

// NewSectionDefault creates a new instance of the section service
func NewSectionDefault(rp internal.SectionRepository) *SectionDefault {
	return &SectionDefault{
		rp: rp,
	}
}

// SectionDefault is the default implementation of the section service
type SectionDefault struct {
	rp internal.SectionRepository
}

// GetAll returns all sections. Returns an error if the operation fails.
func (s *SectionDefault) GetAll() (sections []internal.Section, err error) {
	sections, err = s.rp.GetAll()
	if err != nil {
		switch err {
		case internal.ErrSectionRepository:
			err = fmt.Errorf("%w: %v", internal.ErrSectionService, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrSectionServiceUnkown, err)
		}

		return
	}

	return
}

// Get returns a section by ID. Returns an error if the section is not found.
func (s *SectionDefault) Get(id int) (section internal.Section, err error) {
	section, err = s.rp.Get(id)
	if err != nil {
		switch err {
		case internal.ErrSectionRepositoryNotFound:
			err = fmt.Errorf("%w: %v", internal.ErrSectionServiceNotFound, err)
		case internal.ErrSectionRepository:
			err = fmt.Errorf("%w: %v", internal.ErrSectionService, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrSectionServiceUnkown, err)
		}

		return
	}

	return
}

// Save saves the given section. Returns an error if the operation fails.
func (s *SectionDefault) Save(section *internal.Section) (err error) {
	if err = validateSection(section); err != nil {
		return
	}

	err = s.rp.Save(section)
	if err != nil {
		switch err {
		case internal.ErrSectionRepositoryDuplicated:
			err = fmt.Errorf("%w: %v", internal.ErrSectionServiceDuplicated, err)
		case internal.ErrSectionRepositoryFK:
			err = fmt.Errorf("%w: %v", internal.ErrSectionServiceFK, err)
		case internal.ErrSectionRepository:
			err = fmt.Errorf("%w: %v", internal.ErrSectionService, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrSectionServiceUnkown, err)
		}

		return
	}

	return
}

// Update updates the given section. Returns an error if the operation fails.
func (s *SectionDefault) Update(section *internal.Section) (err error) {
	if err = validateSection(section); err != nil {
		return
	}

	err = s.rp.Update(section)
	if err != nil {
		switch err {
		case internal.ErrSectionRepositoryFK:
			err = fmt.Errorf("%w: %v", internal.ErrSectionServiceFK, err)
		case internal.ErrSectionRepository:
			err = fmt.Errorf("%w: %v", internal.ErrSectionService, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrSectionServiceUnkown, err)
		}

		return
	}

	return
}

// Delete deletes a section by ID. Returns an error if the operation fails.
func (s *SectionDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrSectionRepositoryNotFound:
			err = fmt.Errorf("%w: %v", internal.ErrSectionServiceNotFound, err)
		case internal.ErrSectionRepository:
			err = fmt.Errorf("%w: %v", internal.ErrSectionService, err)
		case internal.ErrSectionRepositoryFK:
			err = fmt.Errorf("%w: %v", internal.ErrSectionServiceFK, err)
		default:
			err = fmt.Errorf("%w: %v", internal.ErrSectionServiceUnkown, err)
		}

		return
	}

	return
}

// validateSection validates the section fields
func validateSection(section *internal.Section) (err error) {
	if section.SectionNumber == 0 {
		err = fmt.Errorf("%w: %v", internal.ErrSectionServiceInvalidField, "section_number")
	}

	if section.CurrentTemperature == 0 {
		err = fmt.Errorf("%w: %v", internal.ErrSectionServiceInvalidField, "current_temperature")
	}

	if section.MinimumTemperature < -30 {
		err = fmt.Errorf("%w: %v", internal.ErrSectionServiceInvalidField, "minimum_temperature")
	}

	if section.CurrentCapacity < 0 {
		err = fmt.Errorf("%w: %v", internal.ErrSectionServiceInvalidField, "current_capacity")
	}

	if section.MinimumCapacity < 0 {
		err = fmt.Errorf("%w: %v", internal.ErrSectionServiceInvalidField, "minimum_capacity")
	}

	if section.MaximumCapacity < 0 {
		err = fmt.Errorf("%w: %v", internal.ErrSectionServiceInvalidField, "maximum_capacity")
	}

	if section.WarehouseID <= 0 {
		err = fmt.Errorf("%w: %v", internal.ErrSectionServiceInvalidField, "warehouse_id")
	}

	if section.ProductTypeID <= 0 {
		err = fmt.Errorf("%w: %v", internal.ErrSectionServiceInvalidField, "product_type_id")
	}

	return
}
