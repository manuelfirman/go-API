package handler

import (
	"net/http"

	"github.com/manuelfirman/go-API/internal"
)

// ProductJSON is a struct that contains the product's information as JSON
type ProductJSON struct {
	// ID is the unique identifier of the product
	ID int `json:"id"`
	// ProductCode is the unique code of the product
	ProductCode string `json:"product_code"`
	// Description is the description of the product
	Description string `json:"description"`
	// Height is the height of the product
	Height float64 `json:"height"`
	// Length is the length of the product
	Length float64 `json:"length"`
	// Width is the width of the product
	Width float64 `json:"width"`
	// Weight is the weight of the product
	Weight float64 `json:"netweight"`
	// ExpirationRate is the rate at which the product expires
	ExpirationRate float64 `json:"expiration_rate"`
	// FreezingRate is the rate at which the product should be frozen
	FreezingRate float64 `json:"freezing_rate"`
	// RecomFreezTemp is the recommended freezing temperature for the product
	RecomFreezTemp float64 `json:"recommended_freezing_temperature"`
	// ProductTypeID is the unique identifier of the product type
	ProductTypeID int `json:"product_type_id"`
	// SellerID is the unique identifier of the seller
	SellerID int `json:"seller_id"`
}

// ProductRecordReportJSON is a struct that contains the product record report information as JSON
type ProductRecordReportJSON struct {
	// ID is the unique identifier of the product record
	ID int `json:"product_id"`
	// Description is the name of the product
	Description string `json:"description"`
	// RecordCount is the amount of records of the product
	RecordCount int `json:"records_count"`
}

// NewProductDefault creates a new instance of the product handler
func NewProductDefault(sv internal.ProductService) *ProductDefault {
	return &ProductDefault{
		sv: sv,
	}
}

// ProductDefault is the default implementation of the product handler
type ProductDefault struct {
	// sv is the service used by the handler
	sv internal.ProductService
}

// GetAll returns all products
func (h *ProductDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a product
func (h *ProductDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new product
func (h *ProductDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a product
func (h *ProductDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a product
func (h *ProductDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetReport returns the information of the product record report
func (h *ProductDefault) GetRecordsByProductReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// deserializeProduct converts a internal Product to a ProductJSON
func deserializeProduct(p internal.Product) ProductJSON {
	return ProductJSON{
		ID:             p.ID,
		ProductCode:    p.ProductCode,
		Description:    p.Description,
		Height:         p.Height,
		Length:         p.Length,
		Width:          p.Width,
		Weight:         p.Weight,
		ExpirationRate: p.ExpirationRate,
		FreezingRate:   p.FreezingRate,
		RecomFreezTemp: p.RecomFreezTemp,
		ProductTypeID:  p.ProductTypeID,
		SellerID:       p.SellerID,
	}
}

// serializeProduct converts a ProductJSON to a internal Product
func serializeProduct(p ProductJSON) internal.Product {
	return internal.Product{
		ID:             p.ID,
		ProductCode:    p.ProductCode,
		Description:    p.Description,
		Height:         p.Height,
		Length:         p.Length,
		Width:          p.Width,
		Weight:         p.Weight,
		ExpirationRate: p.ExpirationRate,
		FreezingRate:   p.FreezingRate,
		RecomFreezTemp: p.RecomFreezTemp,
		ProductTypeID:  p.ProductTypeID,
		SellerID:       p.SellerID,
	}
}
