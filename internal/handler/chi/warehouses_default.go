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

// WarehouseJSON is the JSON representation of a warehouse
type WarehouseJSON struct {
	// Id is the identifier of the warehouse
	ID int `json:"id"`
	// WarehouseCode is the code of the warehouse
	WarehouseCode string `json:"warehouse_code"`
	// Address is the address of the warehouse
	Address string `json:"address"`
	// Telephone is the telephone number of the warehouse
	Telephone string `json:"telephone"`
	// MinimumCapacity is the minimum capacity of the warehouse
	MinimumCapacity int `json:"minimum_capacity"`
	// MinimumTemperature is the minimum temperature that can be maintained in the warehouse
	MinimumTemperature float64 `json:"minimum_temperature"`
	// LocalityID is the id of the locality where the warehouse is located
	LocalityId string `json:"locality_id"`
}

// WarehouseRequestJSON is the JSON representation of a warehouse request
type WarehouseRequestJSON struct {
	// WarehouseCode is the code of the warehouse
	WarehouseCode string `json:"warehouse_code"`
	// Address is the address of the warehouse
	Address string `json:"address"`
	// Telephone is the telephone number of the warehouse
	Telephone string `json:"telephone"`
	// MinimumCapacity is the minimum capacity of the warehouse
	MinimumCapacity int `json:"minimum_capacity"`
	// MinimumTemperature is the minimum temperature that can be maintained in the warehouse
	MinimumTemperature float64 `json:"minimum_temperature"`
	// LocalityID is the id of the locality where the warehouse is located
	LocalityId string `json:"locality_id"`
}

type WarehouseDefault struct {
	// rp is the repository used by the service
	sv internal.WarehouseService
}

// NewWarehouseDefault creates a new instance of the warehouse service
func NewWarehouseDefault(sv internal.WarehouseService) *WarehouseDefault {
	return &WarehouseDefault{
		sv: sv,
	}
}

// GetAll returns all products. Returns an error if the operation fails.
func (wd *WarehouseDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// process
		wh, err := wd.sv.GetAll()
		if err != nil {
			switch err {
			case internal.ErrWarehouseServiceNotFound:
				response.Error(w, http.StatusNotFound, "warehouse not found")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// response
		// - serialize the response
		var data []WarehouseJSON
		for _, w := range wh {
			data = append(data, deserializeWarehouse(w))
		}

		// - write the response
		response.JSON(w, http.StatusOK, response.Res{
			Message: "success",
			Data:    data,
		})
	}
}

// Get returns a product by ID. Returns an error if the product is not found.
func (wd *WarehouseDefault) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}
		// process
		wh, err := wd.sv.Get(id)
		if err != nil {
			switch err {
			case internal.ErrWarehouseServiceNotFound:
				response.Error(w, http.StatusNotFound, "warehouse not found")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// response
		// - serialize the response
		data := deserializeWarehouse(wh)
		// - write the response
		response.JSON(w, http.StatusOK, response.Res{
			Message: "success",
			Data:    data,
		})
	}
}

// Save receives a product and saves it. It returns the ID of the product saved.
func (wd *WarehouseDefault) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - read the body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request")
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
		warehouseRequest := WarehouseRequestJSON{}
		err = validate.CheckFieldExistance(warehouseRequest, bodyMap)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - unmarshal the body to a warehouseRequest struct for further processing
		err = json.Unmarshal(body, &warehouseRequest)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// process
		// - create a warehouseJSON
		warehouseJSON := WarehouseJSON{
			WarehouseCode:      warehouseRequest.WarehouseCode,
			Address:            warehouseRequest.Address,
			Telephone:          warehouseRequest.Telephone,
			MinimumCapacity:    warehouseRequest.MinimumCapacity,
			MinimumTemperature: warehouseRequest.MinimumTemperature,
			LocalityId:         warehouseRequest.LocalityId,
		}

		// - deserialize the warehouseSON
		wh := serializeWarehouse(warehouseJSON)
		err = validateWarehouseFields(&wh)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// - save the warehouse
		wh, err = wd.sv.Save(&wh)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrWarehouseServiceDuplicated):
				response.Error(w, http.StatusConflict, "warehouse already exists")
			case errors.Is(err, internal.ErrWarehouseServiceForeignKey):
				response.Error(w, http.StatusConflict, "foreign key error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// response
		// - deserialize the warehouse to JSON
		data := deserializeWarehouse(wh)
		// - return the warehouse as JSON
		response.JSON(w, http.StatusCreated, response.Res{
			Message: "success",
			Data:    data,
		})
	}
}

// Update receives a product and updates it. Returns an error if the product is not found.
func (wd *WarehouseDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// - get the warehouse with the id
		wh, err := wd.sv.Get(id)
		if err != nil {
			switch err {
			case internal.ErrWarehouseServiceNotFound:
				http.Error(w, "Warehouse not found", http.StatusNotFound)
			default:
				http.Error(w, "Unknown error", http.StatusInternalServerError)
			}
			return
		}

		// - deserialize the warehouse
		warehouseJSON := deserializeWarehouse(wh)

		// - read the body
		if err := request.JSON(r, &warehouseJSON); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request")
			return
		}

		// process
		// - set id (for cases where the id is different from the one in the URL)
		warehouseJSON.ID = id
		// - serialize the warehouse
		wh = serializeWarehouse(warehouseJSON)
		// - validate the warehouse
		err = validateWarehouseFields(&wh)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// - update the warehouse
		err = wd.sv.Update(&wh)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrWarehouseServiceNotFound):
				response.Error(w, http.StatusNotFound, "warehouse not found")
			case errors.Is(err, internal.ErrWarehouseServiceNothingToUpdate):
				response.Error(w, http.StatusConflict, "nothing to update")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// response
		// - deserialize the warehouse to JSON
		data := deserializeWarehouse(wh)
		// - return the warehouse as JSON
		response.JSON(w, http.StatusOK, response.Res{
			Message: "success",
			Data:    data,
		})

	}
}

// Delete deletes a product by ID. Returns an error if the product is not found.
func (wd *WarehouseDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - delete the warehouse
		err = wd.sv.Delete(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrWarehouseServiceNotFound):
				response.Error(w, http.StatusNotFound, "warehouse not found")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown error")
			}
			return
		}

		// response
		// - return the warehouse as JSON
		response.JSON(w, http.StatusNoContent, response.Res{
			Message: "success",
			Data:    nil,
		})
	}
}

// deserialize warehouse
func deserializeWarehouse(w internal.Warehouse) WarehouseJSON {
	return WarehouseJSON{
		ID:                 w.ID,
		WarehouseCode:      w.WarehouseCode,
		Address:            w.Address,
		Telephone:          w.Telephone,
		MinimumCapacity:    w.MinimumCapacity,
		MinimumTemperature: w.MinimumTemperature,
		LocalityId:         w.LocalityId,
	}
}

// serialize warehouse
func serializeWarehouse(w WarehouseJSON) internal.Warehouse {
	return internal.Warehouse{
		ID:                 w.ID,
		WarehouseCode:      w.WarehouseCode,
		Address:            w.Address,
		Telephone:          w.Telephone,
		MinimumCapacity:    w.MinimumCapacity,
		MinimumTemperature: w.MinimumTemperature,
		LocalityId:         w.LocalityId,
	}
}

// Validate zero values fields
func validateWarehouseFields(wh *internal.Warehouse) error {
	if wh.ID != 0 {
		return validate.ErrHandlerIdInRequest
	}
	if wh.Address == "" {
		return fmt.Errorf("%w: address", validate.ErrHandlerMissingField)
	}
	if wh.Telephone == "" {
		return fmt.Errorf("%w: telephone", validate.ErrHandlerMissingField)
	}
	if wh.WarehouseCode == "" {
		return fmt.Errorf("%w: warehousecode", validate.ErrHandlerMissingField)
	}
	if wh.MinimumCapacity == 0 {
		return fmt.Errorf("%w: minimum capacity", validate.ErrHandlerMissingField)
	}
	if wh.LocalityId == "" {
		return fmt.Errorf("%w: localityid", validate.ErrHandlerMissingField)
	}

	return nil
}
