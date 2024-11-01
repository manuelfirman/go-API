package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/manuelfirman/go-API/platform/validate"
	"github.com/manuelfirman/go-API/platform/web/request"
	"github.com/manuelfirman/go-API/platform/web/response"

	"github.com/go-chi/chi/v5"
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
		// process
		// - get all products from the service
		products, err := h.sv.GetAll()
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductServiceNotFound):
				response.Error(w, http.StatusNotFound, "products not found")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// response
		// - serialize products
		var data []ProductJSON
		for _, p := range products {
			jsonData := serializeProduct(p)
			data = append(data, jsonData)
		}

		// response
		response.JSON(w, http.StatusOK, response.Res{
			Message: "success",
			Data:    data,
		})
	}
}

// GetByID returns a product
func (h *ProductDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - validate the product
		product, err := h.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductServiceNotFound):
				response.Error(w, http.StatusNotFound, "product not found")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		data := serializeProduct(product)

		//response
		response.JSON(w, http.StatusOK, response.Res{
			Message: "success",
			Data:    data,
		})
	}
}

// Create creates a new product
func (h *ProductDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - read the body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - unmarshal the body
		bodyMap := map[string]any{}
		err = json.Unmarshal(body, &bodyMap)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - validate the body
		var bodyRequest ProductJSON
		err = validate.CheckFieldExistance(bodyRequest, bodyMap)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - unmarshal the body
		err = json.Unmarshal(body, &bodyRequest)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - map the body to a product
		product := deserializeProduct(bodyRequest)
		err = validateProductZeroValues(product)
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		// process
		// - create a new product
		err = h.sv.Save(&product)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductServiceDuplicated):
				response.Error(w, http.StatusConflict, "duplicated product code")
			case errors.Is(err, internal.ErrSellerServiceNotFound):
				response.Error(w, http.StatusConflict, "seller not found")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// - deserialize the product
		data := serializeProduct(product)

		//response
		response.JSON(w, http.StatusOK, response.Res{
			Message: "success",
			Data:    data,
		})
	}
}

// Update updates a product
func (h *ProductDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - get the product from the service by ID
		p, err := h.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductServiceNotFound):
				response.Error(w, http.StatusNotFound, "product not found")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// - deserialize to ProductJSON
		productJSONData := serializeProduct(p)

		// - map JSON to productJSON
		if err := request.JSON(r, &productJSONData); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - serialize to internal product
		updatedProduct := deserializeProduct(productJSONData)
		updatedProduct.ID = id
		// - validate required fields
		err = validateProductZeroValues(updatedProduct)
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		// - update the product
		err = h.sv.Update(&updatedProduct)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductServiceNotFound):
				response.Error(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrProductServiceDuplicated):
				response.Error(w, http.StatusConflict, "duplicated product code")
			case errors.Is(err, internal.ErrProductServiceNothingToUpdate):
				response.Error(w, http.StatusConflict, "nothing to update")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// serialize to JSON
		data := serializeProduct(updatedProduct)

		// response
		response.JSON(w, http.StatusOK, response.Res{
			Message: "success",
			Data:    data,
		})
	}
}

// Delete deletes a product
func (h *ProductDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		err = h.sv.Delete(id)

		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductServiceNotFound):
				response.Error(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrProductServiceForeignKey):
				response.Error(w, http.StatusConflict, "product has dependencies")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, response.Res{
			Message: "success",
			Data:    nil,
		})
	}
}

// // GetReport returns the information of the product record report
// func (h *ProductDefault) GetReport() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// request
// 		id := r.URL.Query().Get("id")
// 		if id == "" {
// 			id = "0"
// 		}

// 		idInt, err := strconv.Atoi(id)

// 		if err != nil {
// 			response.Error(w, http.StatusUnprocessableEntity, "invalid product id")
// 			return
// 		}

// 		// process
// 		reportData, err := h.sv.GetReport(idInt)

// 		if err != nil {
// 			switch {
// 			case errors.Is(err, internal.ErrProductServiceNotFound):
// 				response.Error(w, http.StatusNotFound, "product not found")
// 			default:
// 				response.Error(w, http.StatusInternalServerError, "unknown error")
// 			}
// 			return
// 		}

// 		// response
// 		// - serialize product records
// 		var productsRecordResponseJSON []ProductRecordReportJSON
// 		for _, re := range reportData {
// 			jsonData := serializeProductRecordReport(re)
// 			productsRecordResponseJSON = append(productsRecordResponseJSON, jsonData)
// 		}

// 		response.JSON(w, http.StatusOK, map[string]any{
// 			"message": "product records report",
// 			"data":    productsRecordResponseJSON,
// 		})
// 	}
// }

// serializeProduct converts a internal Product to a ProductJSON
func serializeProduct(p internal.Product) ProductJSON {
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

// deserializeProduct converts a ProductJSON to a internal Product
func deserializeProduct(p ProductJSON) internal.Product {
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

// func serializeProductRecordReport(r internal.ProductRecordReport) ProductRecordReportJSON {
// 	return ProductRecordReportJSON{
// 		ID:          r.ID,
// 		Description: r.Description,
// 		RecordCount: r.RecordCount,
// 	}
// }

// validateProductZeroValues validates if the product has fields in zero value
func validateProductZeroValues(product internal.Product) error {
	if product.ID != 0 {
		return validate.ErrHandlerIdInRequest
	}
	// Validate required fields
	if product.ProductCode == "" {
		return fmt.Errorf("%w: product_code", validate.ErrHandlerMissingField)
	}
	if product.Description == "" {
		return fmt.Errorf("%w: description", validate.ErrHandlerMissingField)
	}
	if product.Height == 0 {
		return fmt.Errorf("%w: height", validate.ErrHandlerMissingField)
	}
	if product.Length == 0 {
		return fmt.Errorf("%w: length", validate.ErrHandlerMissingField)
	}
	if product.Width == 0 {
		return fmt.Errorf("%w: width", validate.ErrHandlerMissingField)
	}
	if product.Weight == 0 {
		return fmt.Errorf("%w: netweight", validate.ErrHandlerMissingField)
	}
	if product.ExpirationRate == 0 {
		return fmt.Errorf("%w: expiration_rate", validate.ErrHandlerMissingField)
	}
	if product.FreezingRate == 0 {
		return fmt.Errorf("%w: freezing_rate", validate.ErrHandlerMissingField)
	}
	if product.RecomFreezTemp == 0 {
		return fmt.Errorf("%w: recommended_freezing_temperature", validate.ErrHandlerMissingField)
	}
	if product.SellerID == 0 {
		return fmt.Errorf("%w: seller_id", validate.ErrHandlerMissingField)
	}

	return nil
}
