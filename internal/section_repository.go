package internal

import "errors"

var (
	// ErrSectionRepositoryNotFound is returned when the Section is not found
	ErrSectionRepositoryNotFound = errors.New("repository: section not found")
	// ErrSectionRepositoryDuplicated is returned when the Section already exists
	ErrSectionRepositoryDuplicated = errors.New("repository: section already exists")
	//Generic error for repository
	ErrSectionRepository = errors.New("repository: internal error")
	//ErrSectionRepositoryFK is returned when the Section has purchase orders
	ErrSectionRepositoryFK = errors.New("repository: section has purchase orders")
	// ErrSectionRepositoryNoData is returned when the Section has no data
	ErrSectionRepositoryNoData = errors.New("repository: section table has no data")
)

// SectionRepository is an interface that contains the methods that the section repository should support
type SectionRepository interface {
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
	// GetAllProducts
	// GetAllProducts(id int) ([]map[string]interface{}, error)
}
