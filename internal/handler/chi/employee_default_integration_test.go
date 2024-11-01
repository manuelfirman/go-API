package handler_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	handler "github.com/manuelfirman/go-API/internal/handler/chi"
	"github.com/manuelfirman/go-API/internal/repository"
	"github.com/manuelfirman/go-API/internal/service"
	"github.com/stretchr/testify/require"
)

func init() {
	cfg := mysql.Config{
		User:      "root",
		Passwd:    "root",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "go_api_db_test",
		ParseTime: true,
	}
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}

func TestEmployeeDefault_Integration_GetAll(t *testing.T) {
	t.Run("case 01 _ success: employees found", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_GetAll__success_01")
		require.NoError(t, err)
		defer db.Close()
		// -- rollback database
		defer func(db *sql.DB) {
			_, err := db.Exec("DELETE FROM warehouses")
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM employees")
			require.NoError(t, err)
			_, err = db.Exec("ALTER TABLE employees AUTO_INCREMENT = 1")
			require.NoError(t, err)
			_, err = db.Exec("ALTER TABLE warehouses AUTO_INCREMENT = 1")
			require.NoError(t, err)
		}(db)

		// -- set up database
		err = func(db *sql.DB) error {
			_, err := db.Exec("INSERT INTO warehouses (warehouse_code, address, telephone, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)", "warehouse 1", "address 1", "2995654", 100, 10)
			if err != nil {
				return err
			}
			_, err = db.Exec("INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)", 1234, "John", "Doe", 1)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Get("/api/v1/employees", hd.GetAll())

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "success",
			"data": [
				{
					"id": 1,
					"card_number_id": 1234,
					"first_name": "John",
					"last_name": "Doe",
					"warehouse_id": 1
				}
			]
		}`

		// act
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusOK, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})

	t.Run("case 02 _ success: employees not found (empty table)", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_GetAll__success_02")
		require.NoError(t, err)
		defer db.Close()

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Get("/api/v1/employees", hd.GetAll())

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "success",
			"data": []
		}`

		// act
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusOK, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})

	t.Run("case 03 _ error: internal server error", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_GetAll__failure_03")
		require.NoError(t, err)
		// - close database to force an error (without defer)
		db.Close()

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Get("/api/v1/employees", hd.GetAll())

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "unknown service error",
			"status": "internal server error"
		}`

		// act
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})
}

func TestEmployeeDefault_Integration_Get(t *testing.T) {
	t.Run("case 01 _ success: employee found", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_Get__success_01")
		require.NoError(t, err)
		defer db.Close()
		// -- rollback database
		defer func(db *sql.DB) {
			_, err := db.Exec("DELETE FROM warehouses")
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM employees")
			require.NoError(t, err)
			_, err = db.Exec("ALTER TABLE employees AUTO_INCREMENT = 1")
			require.NoError(t, err)
			_, err = db.Exec("ALTER TABLE warehouses AUTO_INCREMENT = 1")
			require.NoError(t, err)
		}(db)

		// -- set up database
		err = func(db *sql.DB) error {
			_, err := db.Exec("INSERT INTO warehouses (warehouse_code, address, telephone, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)", "warehouse 1", "address 1", "2995654", 100, 10)
			if err != nil {
				return err
			}
			_, err = db.Exec("INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)", 1234, "John", "Doe", 1)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Get("/api/v1/employees/{id}", hd.Get())

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
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
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusOK, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})
	t.Run("case 02 _ error: bad request [invalid id]", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_Get__error_02")
		require.NoError(t, err)
		defer db.Close()

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Get("/api/v1/employees/{id}", hd.Get())

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/invalidID", nil)
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
				"message": "invalid id",
				"status": "bad request"
			}`

		// act
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})
	t.Run("case 03 _ error: employee not found", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_Get__error_03")
		require.NoError(t, err)
		defer db.Close()
		// -- rollback database
		defer func(db *sql.DB) {
			_, err := db.Exec("DELETE FROM warehouses")
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM employees")
			require.NoError(t, err)
			_, err = db.Exec("ALTER TABLE employees AUTO_INCREMENT = 1")
			require.NoError(t, err)
			_, err = db.Exec("ALTER TABLE warehouses AUTO_INCREMENT = 1")
			require.NoError(t, err)
		}(db)

		// -- set up database
		err = func(db *sql.DB) error {
			_, err := db.Exec("INSERT INTO warehouses (warehouse_code, address, telephone, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)", "warehouse 1", "address 1", "2995654", 100, 10)
			if err != nil {
				return err
			}
			_, err = db.Exec("INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)", 1234, "John", "Doe", 1)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Get("/api/v1/employees/{id}", hd.Get())

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/2", nil)
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
				"message": "employee not found",
				"status": "not found"
		}`

		// act
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusNotFound, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})
	t.Run("case 04 _ error: internal server error", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_GetAll__failure_03")
		require.NoError(t, err)
		// - close database to force an error (without defer)
		db.Close()

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Get("/api/v1/employees/{id}", hd.Get())

		// - request
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
		// - response
		res := httptest.NewRecorder()

		// - expected
		expectedBody := `{
			"message": "internal server error",
			"status": "internal server error"
		}`

		// act
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})

}

func TestEmployeeDefault_Integration_Save(t *testing.T) {
	t.Run("case 01 _ success: employee created", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_Save__success_01")
		require.NoError(t, err)
		defer db.Close()
		// -- rollback database
		defer func(db *sql.DB) {
			_, err := db.Exec("DELETE FROM warehouses")
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM employees")
			require.NoError(t, err)
			_, err = db.Exec("ALTER TABLE employees AUTO_INCREMENT = 1")
			require.NoError(t, err)
			_, err = db.Exec("ALTER TABLE warehouses AUTO_INCREMENT = 1")
			require.NoError(t, err)
		}(db)

		// - set up database
		err = func(db *sql.DB) error {
			_, err := db.Exec("INSERT INTO warehouses (warehouse_code, address, telephone, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)", "warehouse 1", "address 1", "2995654", 100, 10)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Post("/api/v1/employees", hd.Save())

		// - request
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)
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
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusCreated, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})
	t.Run("case 02 _ error: employee already exist", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_Save__error_02")
		require.NoError(t, err)
		defer db.Close()
		// -- rollback database
		defer func(db *sql.DB) {
			_, err := db.Exec("DELETE FROM warehouses")
			require.NoError(t, err)
			_, err = db.Exec("DELETE FROM employees")
			require.NoError(t, err)
			_, err = db.Exec("ALTER TABLE employees AUTO_INCREMENT = 1")
			require.NoError(t, err)
			_, err = db.Exec("ALTER TABLE warehouses AUTO_INCREMENT = 1")
			require.NoError(t, err)
		}(db)

		// -- set up database
		err = func(db *sql.DB) error {
			_, err := db.Exec("INSERT INTO warehouses (warehouse_code, address, telephone, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)", "warehouse 1", "address 1", "2995654", 100, 10)
			if err != nil {
				return err
			}
			_, err = db.Exec("INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)", 1234, "John", "Doe", 1)
			if err != nil {
				return err
			}
			return nil
		}(db)
		require.NoError(t, err)

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Post("/api/v1/employees", hd.Save())

		// - request
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)
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
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusConflict, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})
	t.Run("case 03 _ error: bad request invalid payload [unmarshal to struct]", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_Save__error_03")
		require.NoError(t, err)
		defer db.Close()

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Post("/api/v1/employees", hd.Save())

		// - request
		bodyRequest := strings.NewReader(`{
			"card_number_id": "invalid",
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)
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
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})

	t.Run("case 04 _ error: bad request invalid payload [unmarshal to map]", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_Save__error_04")
		require.NoError(t, err)
		defer db.Close()

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Post("/api/v1/employees", hd.Save())

		// - request
		// - malformed json
		bodyRequest := strings.NewReader(`
			"card_number_id": "invalid",
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)
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
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusBadRequest, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})

	t.Run("case 05 _ error: internal server error", func(t *testing.T) {
		// arrange
		// - database
		db, err := sql.Open("txdb", "TestEmployeeDefault_Integration_Save__error_04")
		require.NoError(t, err)
		// - close database to force an error (without defer)
		db.Close()

		// - repository
		rp := repository.NewEmployeeMySQL(db)
		// - service
		sv := service.NewEmployeeDefault(rp)
		// - handler
		hd := handler.NewEmployeeDefault(sv)
		// - router
		router := chi.NewRouter()
		router.Post("/api/v1/employees", hd.Save())

		// - request
		// - malformed json
		bodyRequest := strings.NewReader(`{
			"card_number_id": 1234,
			"first_name": "John",
			"last_name": "Doe",
			"warehouse_id": 1
		}`)
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
		router.ServeHTTP(res, req)

		// assert
		// - status code
		require.Equal(t, http.StatusInternalServerError, res.Code)
		// - body
		require.JSONEq(t, expectedBody, res.Body.String())
		// - content type
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})

}

func TestEmployeeDefault_Integration_Update(t *testing.T) {
	t.Run("case 01 _ success: employee updated", func(t *testing.T) {})

	t.Run("case 02 _ error: employee not found", func(t *testing.T) {})

	t.Run("case 03 _ error: bad request invalid payload", func(t *testing.T) {})

	t.Run("case 04 _ error: internal server error", func(t *testing.T) {

	})
}

func TestEmployeeDefault_Integration_Delete(t *testing.T) {

}
