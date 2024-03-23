package handler

import (
	"net/http"

	"github.com/manuelfirman/go-API/internal"
)

// SellerJSON is a struct that contains the seller's information as JSON
type SellerJSON struct {
	// ID is the unique identifier of the seller
	ID int `json:"id"`
	// CID is the unique identifier of the company
	CID int `json:"cid"`
	// CompanyName is the name of the company
	CompanyName string `json:"company_name"`
	// Address is the address of the company
	Address string `json:"address"`
	// Telephone is the telephone number of the company
	Telephone string `json:"telephone"`
	// LocalityID is the seller's locality id
	LocalityID string `json:"locality_id"`
}

// NewProductDefault creates a new instance of the product handler
func NewSellerDefault(sv internal.SellerService) *SellerDefault {
	return &SellerDefault{
		sv: sv,
	}
}

// ProductDefault is the default implementation of the product handler
type SellerDefault struct {
	// sv is the service used by the handler
	sv internal.SellerService
}

// GetAll returns all products
func (h *SellerDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a product
func (h *SellerDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new product
func (h *SellerDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a product
func (h *SellerDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a product
func (h *SellerDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetReport returns the information of the product record report
func (h *SellerDefault) GetRecordsByProductReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// serializeSellerToJSON takes a seller struct and converts it to a json
func deserializeSellerToJSON(seller internal.Seller) (sellerJSON SellerJSON) {
	sellerJSON = SellerJSON{
		ID:          seller.ID,
		CID:         seller.CID,
		CompanyName: seller.CompanyName,
		Address:     seller.Address,
		Telephone:   seller.Telephone,
		LocalityID:  seller.LocalityID,
	}
	return
}

// serializeSellerFromJSON takes a json and converts it to a seller struct
func serializeSellerFromJSON(sellerJSON SellerJSON) (seller internal.Seller) {
	seller = internal.Seller{
		ID:          sellerJSON.ID,
		CID:         sellerJSON.CID,
		CompanyName: sellerJSON.CompanyName,
		Address:     sellerJSON.Address,
		Telephone:   sellerJSON.Telephone,
		LocalityID:  sellerJSON.LocalityID,
	}
	return
}
