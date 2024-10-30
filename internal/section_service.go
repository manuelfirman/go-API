package internal

import "errors"

var (
	//ErrSectionFieldRequired is returned when the Section field is required
	ErrSectionServiceFieldRequired = errors.New("service: section field is required")
	// ErrSectionServiceNotFound is returned when the Section is not found
	ErrSectionServiceNotFound = errors.New("service: section not found")
	// ErrSectionServiceDuplicated is returned when the Section already exists
	ErrSectionServiceDuplicated = errors.New("service: section already exists")
	//Generic error for service
	ErrSectionService = errors.New("service: internal error")
	//ErrSectionServiceFK is returned when the Section
	ErrSectionServiceFK = errors.New("service: section fk error")
	// ErrSectionServiceUnkown is returned when the repository returns an unknown error (not defined in repository errors)
	ErrSectionServiceUnkown = errors.New("service: unknown error")
	// ErrSectionServiceInvalidField is returned when the field is invalid
	ErrSectionServiceInvalidField = errors.New("service: invalid field")
)

// SectionService is an interface that contains the methods that the section service should support
type SectionService interface {
	// FindAll returns all the sections
	GetAll() ([]Section, error)
	// FindByID returns the section with the given ID
	Get(id int) (Section, error)
	// Save saves the given section
	Save(section *Section) error
	// Update updates the given section
	Update(section *Section) error
	// Delete deletes the section with the given ID
	Delete(id int) error
	// GetAllProducts returns all the products
	// GetAllProducts(id int) ([]map[string]interface{}, error)
}
