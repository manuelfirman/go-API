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
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestEmployeeDefault_GetAll(t *testing.T) {
	t.Run("case 01 _ succes: get all employees", func(t *testing.T) {
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

	t.Run("case 02 _ error: internal server error", func(t *testing.T) {
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

	t.Run("case 03 _ error: unkown service error (from repository)", func(t *testing.T) {
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

	t.Run("case 04 _ error: unknown server error", func(t *testing.T) {
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
	t.Run("case 01 _ success: get an employee by id", func(t *testing.T) {
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

		// - function called
		sv.AssertCalled(t, "Get", 1)
		// - number of calls
		sv.AssertNumberOfCalls(t, "Get", 1)
		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 02 _ error: invalid id", func(t *testing.T) {
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

		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 03 _ error: employee not found", func(t *testing.T) {
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
		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 04 _ error: internal server error", func(t *testing.T) {
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
		// - expectations
		sv.AssertExpectations(t)
	})
	t.Run("case 05 _ error: unknown service error", func(t *testing.T) {
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
	t.Run("case 06 _ error: unknown server error", func(t *testing.T) {
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
		// - expectations
		sv.AssertExpectations(t)
	})
}

func TestEmployeeDefault_Save(t *testing.T) {
	t.Run("case 01 _ success: save an employee", func(t *testing.T) {
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

		// - expectations
		sv.AssertExpectations(t)
	})

	t.Run("case 02 _ error: invalid body (ReadAll error)", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("case 03 _ error: invalid body (unmarshal to map error)", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("case 04 _ error: invalid body (unmarshal to struct error)", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("case 05 _ error: invalid json keys", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("case 06 _ error: invalid employee values", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("case 07 _ error: employee already exists", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("case 08 _ error: internal server error", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("case 09 _ error: unknown service error", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("case 10 _ error: unknown server error", func(t *testing.T) {
		// arrange

		// act

		// assert
	})
}

func TestEmployeeDefault_Update(t *testing.T) {

}

func TestEmployeeDefault_Delete(t *testing.T) {

}
