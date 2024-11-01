package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/manuelfirman/go-API/internal"
	handler "github.com/manuelfirman/go-API/internal/handler/chi"
	"github.com/manuelfirman/go-API/internal/service"
	"github.com/manuelfirman/go-API/platform/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestEmployeeDefault_GetAll(t *testing.T) {
	t.Run("case 01 _ success[200]: get all employees", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the GetAll method
		expectedEmployees := []internal.Employee{
			{ID: 1, CardNumberID: 1234, FirstName: "John", LastName: "Doe", WarehouseID: 1},
			{ID: 2, CardNumberID: 5678, FirstName: "Jane", LastName: "Doe", WarehouseID: 2},
		}
		sv.On("GetAll").Return(expectedEmployees, nil)

		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		// - create a new http request with httptest
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)

		// - response
		// - create a new http response recorder
		res := httptest.NewRecorder()

		// - expecteds
		expectedBody := `{
			"message": "success",
			"data": [
				{
					"id": 1,
					"card_number_id": 1234,
					"first_name": "John",
					"last_name": "Doe",
					"warehouse_id": 1
				},
				{
					"id": 2,
					"card_number_id": 5678,
					"first_name": "Jane",
					"last_name": "Doe",
					"warehouse_id": 2
				}
			]
		}`

		// act
		hd.GetAll()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusOK, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - function called
		sv.AssertCalled(t, "GetAll")
		// - number of calls
		sv.AssertNumberOfCalls(t, "GetAll", 1)
		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 02 _ error[500]: internal server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the GetAll method
		sv.On("GetAll").Return([]internal.Employee{}, internal.ErrEmployeeServiceInternalError)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"status": "internal server error",
			"message": "internal server error"
		}`

		// act
		hd.GetAll()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())

		// - function called
		sv.AssertCalled(t, "GetAll")
		// - number of calls
		sv.AssertNumberOfCalls(t, "GetAll", 1)
		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 03 _ error[500]: unkown service error (from repository)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the GetAll function
		sv.On("GetAll").Return([]internal.Employee{}, internal.ErrEmployeeServiceUnknown)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		// - responde
		res := httptest.NewRecorder()

		// - expected body
		expectedBody := `{
			"status": "internal server error",
			"message": "unknown service error"
		}`

		// act
		hd.GetAll()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())

		// - function called
		sv.AssertCalled(t, "GetAll")
		// - number of calls
		sv.AssertNumberOfCalls(t, "GetAll", 1)
		// - expectations
		sv.AssertExpectations(t)

	})
	t.Run("case 04 _ error[500]: unknown server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the GetAll function
		sv.On("GetAll").Return([]internal.Employee{}, errors.New("unknown error"))
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		// - responde
		res := httptest.NewRecorder()

		// - expected body
		expectedBody := `{
			"status": "internal server error",
			"message": "unknown server error"
		}`

		// act
		hd.GetAll()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())

		// - function called
		sv.AssertCalled(t, "GetAll")
		// - number of calls
		sv.AssertNumberOfCalls(t, "GetAll", 1)
		// - expectations
		sv.AssertExpectations(t)
	})
}

func TestEmployeeDefault_Get(t *testing.T) {
	t.Run("case 01 _ success[200]: get an employee by id", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get function
		employee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}
		sv.On("Get", 1).Return(employee, nil)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
		// -- chi context for route
		chiCtx := chi.NewRouteContext()
		// -- add params to chiCtx context
		chiCtx.URLParams.Add("id", "1")
		// -- create a new url context
		newContext := context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx)
		// -- create a new request with the new context
		req = req.WithContext(newContext)

		// - response
		res := httptest.NewRecorder()

		// - expected body
		expectedBody := `{
			"message": "success",
			"data" : {
				"id": 1,
				"card_number_id": 1,
				"first_name": "John",
				"last_name": "Doe",
				"warehouse_id": 1
			}
		}`

		// act
		hd.Get()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusOK, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - function called
		sv.AssertCalled(t, "Get", 1)
		// - number of calls
		sv.AssertNumberOfCalls(t, "Get", 1)
		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 02 _ error[400]: invalid id", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/invalid", nil)
		// -- chi context
		chiCtx := chi.NewRouteContext()
		// -- add params to chi context
		chiCtx.URLParams.Add("id", "invalid")
		// -- new request context with the params
		newContext := context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx)
		// -- create a new request with the new context
		req = req.WithContext(newContext)
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "invalid id",
			"status": "bad request"
		}`

		// act
		hd.Get()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 03 _ error[404]: employee not found", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{}, internal.ErrEmployeeServiceNotFound)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "employee not found",
			"status": "not found"
		}`

		// act
		hd.Get()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusNotFound, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 04 _ error[500]: internal server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{}, internal.ErrEmployeeServiceInternalError)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "internal server error",
			"status": "internal server error"
		}`

		// act
		hd.Get()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 05 _ error[500]: unknown service error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{}, internal.ErrEmployeeServiceUnknown)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "unknown service error",
			"status": "internal server error"
		}`

		// act
		hd.Get()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 06 _ error[500]: unknown server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{}, errors.New("unknown error"))
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "unknown server error",
			"status": "internal server error"
		}`

		// act
		hd.Get()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
}

func TestEmployeeDefault_Save(t *testing.T) {
	t.Run("case 01 _ success[201]: save an employee", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Save method
		sv.On("Save", mock.AnythingOfType("*internal.Employee")).Run(
			// - simulate the Save method behavior (update the employee ID)
			func(args mock.Arguments) {
				employee := args.Get(0).(*internal.Employee)
				employee.ID = 1
			}).Return(nil)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "success",
			"data": {
				"id": 1,
				"card_number_id": 1234,
				"first_name": "John",
				"last_name": "Doe",
				"warehouse_id": 1
			}
		}`

		// act
		hd.Save()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusCreated, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 02 _ error[400]: invalid body (ReadAll error)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()

		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - stub (ReadAll error)
		stub := utils.NewStubReadCloserWithErr(bodyRequest, errors.New("ReadAll error"))
		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", stub)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "invalid body: cannot read",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)

	})

	t.Run("case 03 _ error[400]: invalid body (unmarshal to map error)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid json)
		bodyRequest := strings.NewReader(`
			card_number_id: 1234,
			first_name: "John",
			last_name: "Doe",
			warehouse_id: 1
		`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "invalid body: cannot unmarshal to map",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 04 _ error[400]: invalid body (unmarshal to struct error)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": "invalid",
			"first_name": 1,
			"last_name": 1,
			"warehouse_id": "invalid"
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "invalid body: cannot unmarshal to struct",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 05 _ error[400]: invalid json keys (card_number_id)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing key: card_number_id not found",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 06 _ error[400]: invalid json keys (first_name)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing key: first_name not found",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 07 _ error[400]: invalid json keys (last_name)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing key: last_name not found",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 08 _ error[400]: invalid json keys (warehouse id)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe"
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing key: warehouse_id not found",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 09 _ error[400]: invalid json keys (id in request)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"id": 1,
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "id in request",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 10 _ error[400]: invalid employee values (card_number_id)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": -1,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing field: card_number_id",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 11 _ error[400]: invalid employee values (first_name)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing field: first_name",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 12 _ error[400]: invalid employee values (last_name)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing field: last_name",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 13 _ error[400]: invalid employee values (warehouse id)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": -1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing field: warehouse_id",
			"status": "bad request"
		}`

		// act
		hd.Save()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 14 _ error[409]: employee already exists", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Save method
		sv.On("Save", mock.AnythingOfType("*internal.Employee")).Return(internal.ErrEmployeeServiceDuplicated)

		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "employee already exists",
			"status": "conflict"
		}`

		// act
		hd.Save()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusConflict, res.Code)
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())

		// - expectations
		sv.AssertExpectations(t)

	})

	t.Run("case 15 _ error[500]: internal server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Save method
		sv.On("Save", mock.AnythingOfType("*internal.Employee")).Return(internal.ErrEmployeeServiceInternalError)

		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "internal server error",
			"status": "internal server error"
		}`

		// act
		hd.Save()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 16 _ error[500]: unknown service error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Save method
		sv.On("Save", mock.AnythingOfType("*internal.Employee")).Return(internal.ErrEmployeeServiceUnknown)

		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "unknown service error",
			"status": "internal server error"
		}`

		// act
		hd.Save()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 17 _ error[500]: unknown server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Save method
		sv.On("Save", mock.AnythingOfType("*internal.Employee")).Return(errors.New("unknown error"))

		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "unknown server error",
			"status": "internal server error"
		}`

		// act
		hd.Save()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())

		// - expectations
		sv.AssertExpectations(t)
	})
}

func TestEmployeeDefault_Update(t *testing.T) {
	t.Run("case 01 _ success[200]: update an employee by id", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{
			ID:           1,
			CardNumberID: 1234,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}, nil)
		// - mock the Update method
		sv.On("Update", mock.AnythingOfType("*internal.Employee")).Return(nil)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 4321,
			"first_name": "JohnJohn",
			"last_name": "DoeDoe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "success",
			"data": {
				"id": 1,
				"card_number_id": 4321,
				"first_name": "JohnJohn",
				"last_name": "DoeDoe",
				"warehouse_id": 1
			}
		}`

		// act
		hd.Update()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusOK, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 02 _ error[400]: invalid id", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
				"card_number_id": 4321,
				"first_name": "JohnJohn",
				"last_name": "DoeDoe",
				"warehouse_id": 1
			}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/invalid", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "invalid")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
				"message": "invalid id",
				"status": "bad request"
			}`

		// act
		hd.Update()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 03 _ error[404]: employee not found", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{}, internal.ErrEmployeeServiceNotFound)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
				"card_number_id": 4321,
				"first_name": "JohnJohn",
				"last_name": "DoeDoe",
				"warehouse_id": 1
			}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
				"message": "employee not found",
				"status": "not found"
			}`

		// act
		hd.Update()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusNotFound, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 04 _ error[400]: invalid body", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{
			ID:           1,
			CardNumberID: 1234,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}, nil)

		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`
			card_number_id: 4321,
			first_name: "JohnJohn",
			last_name: "DoeDoe",
			warehouse_id: 1
		`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "invalid body",
			"status": "bad request"
		}`

		// act
		hd.Update()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 05 _ error[400]: invalid employee values (card_number_id)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{
			ID:           1,
			CardNumberID: 1234,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}, nil)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": -1,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing field: card_number_id",
			"status": "bad request"
		}`

		// act
		hd.Update()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 06 _ error[400]: invalid employee values (first_name)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{
			ID:           1,
			CardNumberID: 1234,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}, nil)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing field: first_name",
			"status": "bad request"
		}`

		// act
		hd.Update()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 07 _ error[400]: invalid employee values (last_name)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{
			ID:           1,
			CardNumberID: 1234,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}, nil)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing field: last_name",
			"status": "bad request"
		}`

		// act
		hd.Update()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 08 _ error[400]: invalid employee values (warehouse id)", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{
			ID:           1,
			CardNumberID: 1234,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}, nil)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": -1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "missing field: warehouse_id",
			"status": "bad request"
		}`

		// act
		hd.Update()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 09 _ error[500]: internal server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{
			ID:           1,
			CardNumberID: 1234,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}, nil)
		// - mock the Update method
		sv.On("Update", mock.AnythingOfType("*internal.Employee")).Return(internal.ErrEmployeeServiceInternalError)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "internal server error",
			"status": "internal server error"
		}`

		// act
		hd.Update()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 10 _ error[500]: unknown service error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{
			ID:           1,
			CardNumberID: 1234,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}, nil)
		// - mock the Update method
		sv.On("Update", mock.AnythingOfType("*internal.Employee")).Return(internal.ErrEmployeeServiceUnknown)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "unknown service error",
			"status": "internal server error"
		}`

		// act
		hd.Update()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 11 _ error[500]: unknown server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Get method
		sv.On("Get", 1).Return(internal.Employee{
			ID:           1,
			CardNumberID: 1234,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}, nil)
		// - mock the Update method
		sv.On("Update", mock.AnythingOfType("*internal.Employee")).Return(errors.New("unknown error"))
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - body request (invalid values for the employee)
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", bodyRequest)
		req.Header.Set("Content-Type", "application/json")
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "unknown server error",
			"status": "internal server error"
		}`

		// act
		hd.Update()(res, req)
		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 12 _ error[500]: get internal server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		sv.On("Get", 1).Return(internal.Employee{}, internal.ErrEmployeeServiceInternalError)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "get internal server error",
			"status": "internal server error"
		}`

		// act
		hd.Update()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)

	})

	t.Run("case 13 _ error[500]: get unknown service error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		sv.On("Get", 1).Return(internal.Employee{}, internal.ErrEmployeeServiceUnknown)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "get unknown service error",
			"status": "internal server error"
		}`

		// act
		hd.Update()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 14 _ error[500]: get unknown server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		sv.On("Get", 1).Return(internal.Employee{}, errors.New("unknown error"))
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "get unknown server error",
			"status": "internal server error"
		}`

		// act
		hd.Update()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
}

func TestEmployeeDefault_Delete(t *testing.T) {
	t.Run("case 01 _ success[200]: delete an employee by id", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Delete method
		sv.On("Delete", 1).Return(nil)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "success",
			"data": null
		}`

		// act
		hd.Delete()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusNoContent, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 02 _ error[400]: invalid id", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/invalid", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "invalid")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "invalid id",
			"status": "bad request"
		}`

		// act
		hd.Delete()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 03 _ error[404]: employee not found", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Delete method
		sv.On("Delete", 1).Return(internal.ErrEmployeeServiceNotFound)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "employee not found",
			"status": "not found"
		}`

		// act
		hd.Delete()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusNotFound, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 04 _ error[500]: internal server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Delete method
		sv.On("Delete", 1).Return(internal.ErrEmployeeServiceInternalError)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "internal server error",
			"status": "internal server error"
		}`

		// act
		hd.Delete()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 05 _ error[500]: unknown service error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Delete method
		sv.On("Delete", 1).Return(internal.ErrEmployeeServiceUnknown)
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "unknown service error",
			"status": "internal server error"
		}`

		// act
		hd.Delete()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 06 _ error[500]: unknown server error", func(t *testing.T) {
		// arrange
		// - service mock
		sv := service.NewEmployeeMock()
		// - mock the Delete method
		sv.On("Delete", 1).Return(errors.New("unknown error"))
		// - handler
		hd := handler.NewEmployeeDefault(sv)

		// - request
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "unknown server error",
			"status": "internal server error"
		}`

		// act
		hd.Delete()(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body response
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))

		// - expectations
		sv.AssertExpectations(t)
	})
}
