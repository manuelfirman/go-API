package repository

import (
	"database/sql"

	"github.com/manuelfirman/go-API/internal"
)

type repository struct {
	db *sql.DB
}

// NewProductMySQL creates a new instance of the product repository
func NewProductMySQL(db *sql.DB) internal.ProductRepository {
	return &repository{
		db: db,
	}
}

// GetAll returns all products. Returns an error if the operation fails.
func (r *repository) GetAll() (products []internal.Product, err error) {
	return
}

// Get returns a product by ID. Returns an error if the product is not found.
func (r *repository) Get(id int) (p internal.Product, err error) {
	return
}

// Exists receives a product code and returns true if a product exists.
func (r *repository) Exists(productCode string) (exist bool) {
	return
}

// Save receives a product and saves it. It returns the ID of the product saved.
func (r *repository) Save(p *internal.Product) (id int, err error) {
	return
}

// Update receives a product and updates it. Returns an error if the product is not found.
func (r *repository) Update(p *internal.Product) (err error) {
	return
}

// Delete receives a product ID and deletes it. Returns an error if the product is not found.
func (r *repository) Delete(id int) (err error) {
	return
}

// GetRecordsByProductReport returns the product records.
func (r *repository) GetRecordsByProductReport(id int) (products []internal.Product, err error) {
	return
}
