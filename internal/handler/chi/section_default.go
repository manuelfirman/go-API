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

// SectionJSON is the JSON representation of a section
type SectionJSON struct {
	// ID is the unique identifier of the section
	ID int `json:"id"`
	// SectionNumber is the number of the section
	SectionNumber int `json:"section_number"`
	// CurrentTemperature is the current temperature of the section
	CurrentTemperature float64 `json:"current_temperature"`
	// MinimumTemperature is the minimum temperature that can be maintained in the section
	MinimumTemperature float64 `json:"minimum_temperature"`
	// CurrentCapacity is the current capacity of the section
	CurrentCapacity int `json:"current_capacity"`
	// MinimumCapacity is the minimum capacity of the section
	MinimumCapacity int `json:"minimum_capacity"`
	// MaximumCapacity is the maximum capacity of the section
	MaximumCapacity int `json:"maximum_capacity"`
	// WarehouseID is the unique identifier of the warehouse to which the section belongs
	WarehouseID int `json:"warehouse_id"`
	// ProductTypeID is the unique identifier of the type of product stored in the section
	ProductTypeID int `json:"product_type_id"`
}

// NewSectionDefault creates a new instance of the section handler
func NewSectionDefault(sv internal.SectionService) *SectionDefault {
	return &SectionDefault{
		sv: sv,
	}
}

// SectionDefault is the default implementation of the section handler
type SectionDefault struct {
	sv internal.SectionService
}

// GetAll returns all sections. Returns an error if the operation fails.
func (h *SectionDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sections, err := h.sv.GetAll()
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrSectionService):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrSectionServiceUnkown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}
		}
		// serialize the sections
		data := make([]SectionJSON, len(sections))
		for i, section := range sections {
			data[i] = serializeSection(section)
		}
		// return the response
		response.JSON(w, http.StatusOK, Response{
			Message: "success",
			Data:    data,
		})
	}
}

func (h *SectionDefault) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from url
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		section, err := h.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrSectionServiceNotFound):
				response.Error(w, http.StatusNotFound, "section not found")
			case errors.Is(err, internal.ErrSectionService):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrSectionServiceUnkown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// - serialize data to response
		data := serializeSection(section)

		// response
		response.JSON(w, http.StatusOK, Response{
			Message: "succes",
			Data:    data,
		})
	}
}

func (h *SectionDefault) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - read the body in []byte
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body: cannot read")
			return
		}
		// - unmarshal body to map for validations
		var bodyMap map[string]any
		if err = json.Unmarshal(body, &bodyMap); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body: cannot unmarshal to map")
			return
		}
		// - validate
		if err = validateKeyExistance(bodyMap, "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// - unmarshal to struct
		var sectionJSON SectionJSON
		if err = json.Unmarshal(body, &sectionJSON); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body: cannot unmarshal to struct")
			return
		}
		// - validate zero values
		if err = validateSectionZeroValues(sectionJSON); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// - deserialize
		section := deserializeSection(sectionJSON)

		// process
		err = h.sv.Save(&section)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrSectionServiceDuplicated):
				response.Error(w, http.StatusBadRequest, "section already exists")
			case errors.Is(err, internal.ErrSectionService):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrSectionServiceUnkown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// - serialize data to response
		data := serializeSection(section)

		// response
		response.JSON(w, http.StatusCreated, Response{
			Message: "success",
			Data:    data,
		})
	}
}

func (h *SectionDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from url
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// - get section by id
		section, err := h.sv.Get(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrSectionServiceNotFound):
				response.Error(w, http.StatusNotFound, "section not found")
			case errors.Is(err, internal.ErrSectionService):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrSectionServiceUnkown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// - deserialize section to JSON
		sectionJSON := serializeSection(section)

		// - map JSON to the section
		if err := request.JSON(r, &sectionJSON); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - validate zero values
		if err := validateSectionZeroValues(sectionJSON); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// - deserialize sectionJSON to internal section
		section = deserializeSection(sectionJSON)

		// process
		err = h.sv.Update(&section)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrSectionServiceDuplicated):
				response.Error(w, http.StatusBadRequest, "section already exists")
			case errors.Is(err, internal.ErrSectionService):
				response.Error(w, http.StatusInternalServerError, "internal server error")
			case errors.Is(err, internal.ErrSectionServiceUnkown):
				response.Error(w, http.StatusInternalServerError, "unknown service error")
			default:
				response.Error(w, http.StatusInternalServerError, "unknown server error")
			}

			return
		}

		// - serialize data to response
		data := serializeSection(section)

		// response
		response.JSON(w, http.StatusOK, Response{
			Message: "success",
			Data:    data,
		})
	}
}

// serializeSection serializes a section into a SectionJSON
func serializeSection(section internal.Section) SectionJSON {
	return SectionJSON{
		ID:                 section.ID,
		SectionNumber:      section.SectionNumber,
		CurrentTemperature: section.CurrentTemperature,
		MinimumTemperature: section.MinimumTemperature,
		CurrentCapacity:    section.CurrentCapacity,
		MinimumCapacity:    section.MinimumCapacity,
		MaximumCapacity:    section.MaximumCapacity,
		WarehouseID:        section.WarehouseID,
		ProductTypeID:      section.ProductTypeID,
	}
}

// deserializeSection deserializes a SectionJSON into a section
func deserializeSection(section SectionJSON) internal.Section {
	return internal.Section{
		ID:                 section.ID,
		SectionNumber:      section.SectionNumber,
		CurrentTemperature: section.CurrentTemperature,
		MinimumTemperature: section.MinimumTemperature,
		CurrentCapacity:    section.CurrentCapacity,
		MinimumCapacity:    section.MinimumCapacity,
		MaximumCapacity:    section.MaximumCapacity,
		WarehouseID:        section.WarehouseID,
		ProductTypeID:      section.ProductTypeID,
	}
}

// validateSectionZeroValues
func validateSectionZeroValues(section SectionJSON) error {
	// validate that id does has send in the request
	if section.ID != 0 {
		return fmt.Errorf("%w: %v", ErrHandlerMissingField, "id in request")
	}
	if section.SectionNumber == 0 {
		return fmt.Errorf("%w: %v", ErrHandlerMissingField, "section_number")
	}
	if section.MinimumCapacity < 0 {
		return fmt.Errorf("%w: %v", ErrHandlerMissingField, "minimum_capacity")
	}
	if section.MaximumCapacity < 0 {
		return fmt.Errorf("%w: %v", ErrHandlerMissingField, "maximum_capacity")
	}
	if section.WarehouseID == 0 {
		return fmt.Errorf("%w: %v", ErrHandlerMissingField, "warehouse_id")
	}
	if section.ProductTypeID == 0 {
		return fmt.Errorf("%w: %v", ErrHandlerMissingField, "product_type_id")
	}

	return nil
}
