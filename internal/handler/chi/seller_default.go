package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/manuelfirman/go-API/internal"
	"github.com/manuelfirman/go-API/platform/validate"
	"github.com/manuelfirman/go-API/platform/web/request"
	"github.com/manuelfirman/go-API/platform/web/response"
)

var (
	// ErrProductHandlerMissingField is the error returned when a required field is missing
	ErrSellerHandlerMissingField = errors.New("missing field")
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

// SellerJSON is a struct that contains the seller's information as JSON
type SellerRequestJSON struct {
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

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
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
		// Get all the sellers
		sellers, err := h.sv.GetAll()
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrSellerServiceNotFound):
				response.Error(w, http.StatusNotFound, "sellers not found")
			case errors.Is(err, internal.ErrSellerServiceDB):
				response.Error(w, http.StatusInternalServerError, "database error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// Serialize the sellers to JSON
		var sellersJSON []SellerJSON
		for _, seller := range sellers {
			sellersJSON = append(sellersJSON, deserializeSellerToJSON(seller))
		}

		// Return the sellers as JSON
		response.JSON(w, http.StatusOK, Response{
			Message: "success",
			Data:    sellersJSON,
		})

	}
}

// GetByID returns a product
func (h *SellerDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// Get the seller
		seller, err := h.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrSellerServiceNotFound):
				response.Error(w, http.StatusNotFound, "seller not found")
			case errors.Is(err, internal.ErrSellerServiceDB):
				response.Error(w, http.StatusInternalServerError, "database error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// Serialize the seller to JSON
		sellerJSON := deserializeSellerToJSON(seller)

		// Return the seller as JSON
		response.JSON(w, http.StatusOK, Response{
			Message: "success",
			Data:    sellerJSON,
		})
	}
}

// Create creates a new product
func (h *SellerDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - unmarshal the body to a map for validation
		bodyMap := map[string]any{}
		err = json.Unmarshal(body, &bodyMap)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - validate the body
		sellerRequest := SellerRequestJSON{}
		err = validate.CheckFieldExistance(sellerRequest, bodyMap)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - unmarshal the body to a sellerRequest struct for further processing
		err = json.Unmarshal(body, &sellerRequest)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// process
		// - create a sellerJSON
		sellerJSON := SellerJSON{
			CID:         sellerRequest.CID,
			CompanyName: sellerRequest.CompanyName,
			Address:     sellerRequest.Address,
			Telephone:   sellerRequest.Telephone,
			LocalityID:  sellerRequest.LocalityID,
		}

		// - deserialize the sellerJSON
		seller := serializeSellerFromJSON(sellerJSON)
		err = validateSellerFields(&seller)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// - save the seller
		s, err := h.sv.Save(&seller)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrSellerServiceDuplicated):
				response.Error(w, http.StatusConflict, "seller already exists")
			case errors.Is(err, internal.ErrSellerServiceForeignKey):
				response.Error(w, http.StatusConflict, "foreign key error")
			case errors.Is(err, internal.ErrSellerServiceDB):
				response.Error(w, http.StatusInternalServerError, "database error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// response
		// - deserialize the seller to JSON
		data := deserializeSellerToJSON(s)

		// - return the seller as JSON
		response.JSON(w, http.StatusOK, Response{
			Message: "success",
			Data:    data,
		})
	}
}

// Update updates a product
func (h *SellerDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - get the seller
		s, err := h.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrSellerServiceNotFound):
				response.Error(w, http.StatusNotFound, "seller not found")
			case errors.Is(err, internal.ErrSellerServiceDB):
				response.Error(w, http.StatusInternalServerError, "database error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// - deserialize the seller to JSON
		sellerJSONData := deserializeSellerToJSON(s)
		sellerJSONData.ID = id

		// - map JSON to the seller
		if err := request.JSON(r, &sellerJSONData); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - serialize to internal seller
		s = serializeSellerFromJSON(sellerJSONData)

		err = validateSellerFields(&s)
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		// - update the seller
		err = h.sv.Update(&s)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrSellerServiceDuplicated):
				response.Error(w, http.StatusConflict, "seller already exists")
			case errors.Is(err, internal.ErrSellerServiceNotFound):
				response.Error(w, http.StatusNotFound, "seller not found")
			case errors.Is(err, internal.ErrSellerServiceDB):
				response.Error(w, http.StatusInternalServerError, "database error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// response
		// - return the updated seller
		response.JSON(w, http.StatusOK, Response{
			Message: "success",
			Data:    sellerJSONData,
		})

	}
}

// Delete deletes a product
func (h *SellerDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - delete the seller
		err = h.sv.Delete(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrSellerServiceNotFound):
				response.Error(w, http.StatusNotFound, "seller not found")
			case errors.Is(err, internal.ErrSellerServiceDB):
				response.Error(w, http.StatusInternalServerError, "database error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// response
		// - return a success message
		response.JSON(w, http.StatusNoContent, Response{
			Message: "success",
		})
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

func validateSellerFields(seller *internal.Seller) (err error) {
	if seller.CID == 0 {
		return fmt.Errorf("%w: CID", ErrSellerHandlerMissingField)
	}
	if seller.CompanyName == "" {
		return fmt.Errorf("%w: company name", ErrSellerHandlerMissingField)
	}
	if seller.Address == "" {
		return fmt.Errorf("%w: address", ErrSellerHandlerMissingField)
	}
	if seller.Telephone == "" {
		return fmt.Errorf("%w: telephone", ErrSellerHandlerMissingField)
	}
	if seller.LocalityID == "" {
		return fmt.Errorf("%w: locality id", ErrSellerHandlerMissingField)
	}
	return

}
