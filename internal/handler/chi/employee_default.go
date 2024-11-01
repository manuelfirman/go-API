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

// EmployeeJSON is the json response of a employee
type EmployeeJSON struct {
	ID           int    `json:"id" example:"1"`
	CardNumberID int    `json:"card_number_id" example:"1234"`
	FirstName    string `json:"first_name" example:"John"`
	LastName     string `json:"last_name" example:"Doe"`
	WarehouseID  int    `json:"warehouse_id" example:"1"`
}

// NewEmployeeDefault creates a new instance of the employee handler
func NewEmployeeDefault(sv internal.EmployeeService) *EmployeeDefault {
	return &EmployeeDefault{
		sv: sv,
	}
}

// EmployeeDefault is the default implementation of the employee handler
type EmployeeDefault struct {
	sv internal.EmployeeService
}

func (h *EmployeeDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := h.sv.GetAll()
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrEmployeeServiceInternalError):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrEmployeeServiceUnknown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		data := make([]EmployeeJSON, len(employees))
		for i, v := range employees {
			data[i] = serializeEmployee(v)
		}

		response.JSON(w, http.StatusOK, response.Res{
			Message: "success",
			Data:    data,
		})
	}
}

// Get returns an employee by ID
func (h *EmployeeDefault) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - get employee by id
		employee, err := h.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrEmployeeServiceNotFound):
				response.Error(w, http.StatusNotFound, "employee not found")
			case errors.Is(err, internal.ErrEmployeeServiceInternalError):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrEmployeeServiceUnknown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// - serialize
		data := serializeEmployee(employee)

		// response
		response.JSON(w, http.StatusOK, response.Res{
			Message: "success",
			Data:    data,
		})
	}
}

// Save saves the given employee
func (h *EmployeeDefault) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - read body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body: cannot read")
			return
		}

		// - unmarshal to map for validation
		var bodyMap map[string]any
		if err = json.Unmarshal(body, &bodyMap); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body: cannot unmarshal to map")
			return
		}
		// - validate the body keys
		if err = validate.KeyExistance(bodyMap, "card_number_id", "first_name", "last_name", "warehouse_id"); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// - unmarshal the body to a EmployeeJSON
		var employeeJSON EmployeeJSON
		if err = json.Unmarshal(body, &employeeJSON); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body: cannot unmarshal to struct")
			return
		}
		// - validate the employee id
		if err = ValidateEmployeeID(employeeJSON); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		// - validate the body values
		if err = validateEmployeeZeroValues(employeeJSON); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// process
		// - deserialize EmployeeJSON to an internal employee
		employee := deserializeEmployee(employeeJSON)
		// - save employee
		err = h.sv.Save(&employee)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrEmployeeServiceDuplicated):
				response.Error(w, http.StatusConflict, "employee already exists")
			case errors.Is(err, internal.ErrEmployeeServiceInternalError):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrEmployeeServiceUnknown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// - serialize the employee to response
		data := serializeEmployee(employee)

		// response
		response.JSON(w, http.StatusCreated, response.Res{
			Message: "success",
			Data:    data,
		})
	}
}

// Update updates the given employee
func (h *EmployeeDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// - get employee by id
		employee, err := h.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrEmployeeServiceNotFound):
				response.Error(w, http.StatusNotFound, "employee not found")
			case errors.Is(err, internal.ErrEmployeeServiceInternalError):
				response.Error(w, http.StatusInternalServerError, "get internal server error")
			case errors.Is(err, internal.ErrEmployeeServiceUnknown):
				response.Error(w, http.StatusInternalServerError, "get unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "get unknown server error")
			}

			return
		}

		// - serialize the employee
		employeeJSON := serializeEmployee(employee)

		// - map the json to employee
		if err = request.JSON(r, &employeeJSON); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		employeeJSON.ID = id

		// process
		// - validate the employee
		if err = validateEmployeeZeroValues(employeeJSON); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		// - deserialize the employee
		employee = deserializeEmployee(employeeJSON)
		// - update the employee
		err = h.sv.Update(&employee)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrEmployeeServiceInternalError):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrEmployeeServiceUnknown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// - serialize the employee to response
		data := serializeEmployee(employee)

		// response
		response.JSON(w, http.StatusOK, response.Res{
			Message: "success",
			Data:    data,
		})
	}
}

// Delete deletes the employee with the given ID
func (h *EmployeeDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - delete employee by id
		err = h.sv.Delete(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrEmployeeServiceNotFound):
				response.Error(w, http.StatusNotFound, "employee not found")
			case errors.Is(err, internal.ErrEmployeeServiceInternalError):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrEmployeeServiceUnknown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// response
		response.JSON(w, http.StatusNoContent, response.Res{
			Message: "success",
			Data:    nil,
		})

	}
}

// serializeEmployee creates a new json from the given employee
func serializeEmployee(e internal.Employee) EmployeeJSON {
	return EmployeeJSON{
		ID:           e.ID,
		CardNumberID: e.CardNumberID,
		FirstName:    e.FirstName,
		LastName:     e.LastName,
		WarehouseID:  e.WarehouseID,
	}
}

// deserializeEmployee creates a new employee from the given json
func deserializeEmployee(e EmployeeJSON) internal.Employee {
	return internal.Employee{
		ID:           e.ID,
		CardNumberID: e.CardNumberID,
		FirstName:    e.FirstName,
		LastName:     e.LastName,
		WarehouseID:  e.WarehouseID,
	}
}

// validateEmployee validates the employee fields
func validateEmployeeZeroValues(e EmployeeJSON) error {
	if e.CardNumberID <= 0 {
		return fmt.Errorf("%w: card_number_id", validate.ErrHandlerMissingField)
	}

	if e.FirstName == "" {
		return fmt.Errorf("%w: first_name", validate.ErrHandlerMissingField)
	}

	if e.LastName == "" {
		return fmt.Errorf("%w: last_name", validate.ErrHandlerMissingField)
	}

	if e.WarehouseID <= 0 {
		return fmt.Errorf("%w: warehouse_id", validate.ErrHandlerMissingField)
	}

	return nil
}

func ValidateEmployeeID(e EmployeeJSON) error {
	if e.ID != 0 {
		return validate.ErrHandlerIdInRequest
	}

	return nil
}
