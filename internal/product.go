package internal

// Product is a struct that contains the product's information
type Product struct {
	// ID is the unique identifier of the product
	ID int
	// ProductCode is the unique code of the product
	ProductCode string
	// Description is the description of the product
	Description string
	// Height is the height of the product
	Height float64
	// Length is the length of the product
	Length float64
	// Width is the width of the product
	Width float64
	// Weight is the weight of the product
	Weight float64
	// ExpirationRate is the rate at which the product expires
	ExpirationRate float64
	// FreezingRate is the rate at which the product should be frozen
	FreezingRate float64
	// RecomFreezTemp is the recommended freezing temperature for the product
	RecomFreezTemp float64
	// ProductTypeID is the unique identifier of the product type
	ProductTypeID int
	// SellerID is the unique identifier of the seller
	SellerID int
}
