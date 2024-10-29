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
	"github.com/manuelfirman/go-API/platform/web/request"
	"github.com/manuelfirman/go-API/platform/web/response"
)

var (
	// ErrHandlerMissingKey is the error returned when a required key is missing
	ErrHandlerMissingKey = errors.New("missing key")
	// ErrHandlerMissingField is the error returned when a required field is missing
	ErrHandlerMissingField = errors.New("missing field")
)

// Response is a struct that contains the response message and data
type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// BuyerJSON is a struct that contains the buyer's information as JSON
type BuyerJSON struct {
	// ID is the unique identifier of the buyer
	ID int `json:"id"`
	// CardNumberID is the unique identifier of the card number
	CardNumberID int `json:"card_number_id"`
	// FirstName is the first name of the buyer
	FirstName string `json:"first_name"`
	// LastName is the last name of the buyer
	LastName string `json:"last_name"`
}

// NewBuyerDefault creates a new instance of the buyer handler
func NewBuyerDefault(sv internal.BuyerService) *BuyerDefault {
	return &BuyerDefault{
		sv: sv,
	}
}

// NewBuyerDefault creates a new instance of the buyer handler
type BuyerDefault struct {
	// sv is the service used by the handler
	sv internal.BuyerService
}

// GetAll returns all buyers
func (h *BuyerDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buyers, err := h.sv.GetAll()
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrBuyerService):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrBuyerServiceUnkown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// serialize buyers
		data := make([]BuyerJSON, len(buyers))
		for i, b := range buyers {
			data[i] = serializeBuyer(b)
		}

		response.JSON(w, http.StatusOK, Response{
			Message: "success",
			Data:    data,
		})
	}
}

// Get returns a buyer by ID
func (h *BuyerDefault) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - get buyer by id
		buyer, err := h.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrBuyerServiceNotFound):
				response.Error(w, http.StatusNotFound, "buyer not found")
			case errors.Is(err, internal.ErrBuyerService):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrBuyerServiceUnkown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}
		// - serialize buyer
		data := serializeBuyer(buyer)

		// response
		response.JSON(w, http.StatusOK, Response{
			Message: "success",
			Data:    data,
		})
	}
}

// Save saves the given buyer
func (h *BuyerDefault) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body: cannot read")
			return
		}

		// - unmarshal the body to a map for validation
		var bodyMap map[string]any
		if err = json.Unmarshal(body, &bodyMap); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body: cannot unmarshal to map")
			return
		}
		// - validate the body keys
		if err = validateKeyExistance(bodyMap, "card_number_id", "first_name", "last_name"); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// - unmarshal the body to a BuyerJSON
		var buyerJSON BuyerJSON
		if err = json.Unmarshal(body, &buyerJSON); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body: cannot unmarshal to struct")
			return
		}
		// - validate the body values
		if err = validateBuyerZeroValues(buyerJSON); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// process
		// - deserialize the BuyerJSON to an internal Buyer
		buyer := deserializeBuyer(buyerJSON)
		// - save buyer
		err = h.sv.Save(&buyer)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrBuyerServiceDuplicated):
				response.Error(w, http.StatusConflict, "buyer already exists")
			case errors.Is(err, internal.ErrBuyerService):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrBuyerServiceUnkown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// - serialize buyer to response
		data := serializeBuyer(buyer)

		// response
		response.JSON(w, http.StatusCreated, Response{
			Message: "success",
			Data:    data,
		})
	}
}

// Update updates the given buyer
func (h *BuyerDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - find buyer by id
		buyer, err := h.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrBuyerServiceNotFound):
				response.Error(w, http.StatusNotFound, "buyer not found")
			case errors.Is(err, internal.ErrBuyerService):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrBuyerServiceUnkown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// - serialize internal buyer
		buyerJSON := serializeBuyer(buyer)
		// - map the body json to the buyer
		if err := request.JSON(r, &buyerJSON); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - validate the buyer
		if err := validateBuyerZeroValues(buyerJSON); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// - deserialize buyerJSON to internal buyer
		buyer = deserializeBuyer(buyerJSON)

		// - update buyer
		err = h.sv.Update(&buyer)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrBuyerServiceNotFound):
				response.Error(w, http.StatusNotFound, "buyer not found")
			case errors.Is(err, internal.ErrBuyerService):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrBuyerServiceUnkown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// - serialize buyer to response
		data := serializeBuyer(buyer)

		// response
		response.JSON(w, http.StatusOK, Response{
			Message: "success",
			Data:    data,
		})
	}
}

// Delete deletes the buyer by ID
func (h *BuyerDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - delete buyer by id
		err = h.sv.Delete(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrBuyerServiceNotFound):
				response.Error(w, http.StatusNotFound, "buyer not found")
			case errors.Is(err, internal.ErrBuyerServiceFK):
				response.Error(w, http.StatusConflict, "buyer has purchase orders")
			case errors.Is(err, internal.ErrBuyerService):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrBuyerServiceUnkown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusNoContent, Response{
			Message: "success",
			Data:    nil,
		})
	}
}

// serializeBuyer converts an internal Buyer to a BuyerJSON
func serializeBuyer(b internal.Buyer) BuyerJSON {
	return BuyerJSON{
		ID:           b.ID,
		CardNumberID: b.CardNumberID,
		FirstName:    b.FirstName,
		LastName:     b.LastName,
	}
}

// deserializeBuyer converts a BuyerJSON to an internal Buyer
func deserializeBuyer(b BuyerJSON) internal.Buyer {
	return internal.Buyer{
		ID:           b.ID,
		CardNumberID: b.CardNumberID,
		FirstName:    b.FirstName,
		LastName:     b.LastName,
	}
}

// TODO: move these functions to a separate file
// validateKeyExistance validates if the key exists in the map
func validateKeyExistance(m map[string]any, keys ...string) error {
	for _, k := range keys {
		if _, ok := m[k]; !ok {
			return fmt.Errorf("%w: %s not found", ErrHandlerMissingKey, k)
		}
	}

	return nil
}

// validateBuyerZeroValues validates if the buyer has zero values
func validateBuyerZeroValues(b BuyerJSON) error {
	if b.CardNumberID == 0 {
		return fmt.Errorf("%w: card_number_id", ErrHandlerMissingKey)
	}

	if b.FirstName == "" {
		return fmt.Errorf("%w: first_name", ErrHandlerMissingKey)
	}

	if b.LastName == "" {
		return fmt.Errorf("%w: last_name", ErrHandlerMissingKey)
	}

	return nil
}
