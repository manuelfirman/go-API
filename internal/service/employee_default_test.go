package service_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/manuelfirman/go-API/internal"
	"github.com/manuelfirman/go-API/internal/repository"
	"github.com/manuelfirman/go-API/internal/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TestEmployeeDefault_GetAll tests the GetAll method of the EmployeeDefault service
func TestEmployeeDefault_GetAll(t *testing.T) {
	t.Run("case 01 _ success: get all employees", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expected result
		expectedEmployees := []internal.Employee{
			{ID: 1, CardNumberID: 1, FirstName: "John", LastName: "Doe", WarehouseID: 1},
			{ID: 2, CardNumberID: 2, FirstName: "Jane", LastName: "Doe", WarehouseID: 1},
		}
		rp.On("GetAll").Return(expectedEmployees, nil)
		// - service
		s := service.NewEmployeeDefault(rp)

		// act
		employees, err := s.GetAll()

		// assert
		// - error
		require.NoError(t, err)
		// - result
		require.NotNil(t, employees)
		require.Equal(t, expectedEmployees, employees)

		// - function called
		rp.AssertCalled(t, "GetAll")
		// - number of calls
		rp.AssertNumberOfCalls(t, "GetAll", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 02 _ error: repository error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - set expectations
		expectedError := fmt.Errorf("%w: %v", internal.ErrEmployeeServiceInternalError, internal.ErrEmployeeRepository)
		rp.On("GetAll").Return([]internal.Employee{}, internal.ErrEmployeeRepository)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		employees, err := sv.GetAll()

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())
		// - result
		require.Empty(t, employees)

		// - function called
		rp.AssertCalled(t, "GetAll")
		// - number of calls
		rp.AssertNumberOfCalls(t, "GetAll", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 03 _ error: unknown error", func(t *testing.T) {
		// - arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: unknown error", internal.ErrEmployeeServiceUnknown)
		rp.On("GetAll").Return([]internal.Employee{}, errors.New("unknown error"))
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		employees, err := sv.GetAll()

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())
		// - result
		require.Empty(t, employees)

		// - function called
		rp.AssertCalled(t, "GetAll")
		// - number of calls
		rp.AssertNumberOfCalls(t, "GetAll", 1)
		// - expectations
		rp.AssertExpectations(t)
	})
}

// TestEmployeeDefault_Get tests the Get method of the EmployeeDefault service
func TestEmployeeDefault_Get(t *testing.T) {
	t.Run("case 01 _ success: get employee by ID", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}
		rp.On("Get", 1).Return(expectedEmployee, nil)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		employee, err := sv.Get(1)

		// assert
		// - error
		require.NoError(t, err)
		// - result
		require.Equal(t, expectedEmployee, employee)

		// - function called
		rp.AssertCalled(t, "Get", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Get", 1)
		// - expectations
		rp.AssertExpectations(t)

	})

	t.Run("case 02 _ error: employee not found", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: %v", internal.ErrEmployeeServiceNotFound, internal.ErrEmployeeRepositoryNotFound)
		rp.On("Get", 1).Return(internal.Employee{}, internal.ErrEmployeeRepositoryNotFound)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		employee, err := sv.Get(1)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())
		// - result
		require.Empty(t, employee)

		// - function called
		rp.AssertCalled(t, "Get", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Get", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 03 _ error: repository error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: %v", internal.ErrEmployeeServiceInternalError, internal.ErrEmployeeRepository)
		rp.On("Get", 1).Return(internal.Employee{}, internal.ErrEmployeeRepository)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		employee, err := sv.Get(1)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())
		// - result
		require.Empty(t, employee)

		// - function called
		rp.AssertCalled(t, "Get", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Get", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 04 _ error: unknown error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: unknown error", internal.ErrEmployeeServiceUnknown)
		rp.On("Get", 1).Return(internal.Employee{}, errors.New("unknown error"))
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		employee, err := sv.Get(1)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())
		// - result
		require.Empty(t, employee)

		// - function called
		rp.AssertCalled(t, "Get", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Get", 1)
		// - expectations
		rp.AssertExpectations(t)
	})
}

// TestEmployeeDefault_Save tests the Save method of the EmployeeDefault service
func TestEmployeeDefault_Save(t *testing.T) {
	t.Run("case 01 _ succes: save employee", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		newEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}
		rp.On("Save", mock.AnythingOfType("*internal.Employee")).Return(nil)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Save(&newEmployee)

		// assert
		// - error
		require.NoError(t, err)

		// - function called
		rp.AssertCalled(t, "Save", &newEmployee)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Save", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 02 _ error: employee already exists", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: %v", internal.ErrEmployeeServiceDuplicated, internal.ErrEmployeeRepositoryDuplicated)
		newEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}
		rp.On("Save", mock.AnythingOfType("*internal.Employee")).Return(internal.ErrEmployeeRepositoryDuplicated)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Save(&newEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Save", &newEmployee)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Save", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 03 _ error: repository error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: %v", internal.ErrEmployeeServiceInternalError, internal.ErrEmployeeRepository)
		newEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}
		rp.On("Save", mock.AnythingOfType("*internal.Employee")).Return(internal.ErrEmployeeRepository)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Save(&newEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Save", &newEmployee)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Save", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 04 _ error: unknown error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: unknown error", internal.ErrEmployeeServiceUnknown)
		newEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}
		rp.On("Save", mock.AnythingOfType("*internal.Employee")).Return(errors.New("unknown error"))
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Save(&newEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Save", &newEmployee)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Save", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 05 _ error: invalid employee (first name)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: first name", internal.ErrEmployeeServiceFieldRequired)
		newEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			// no first name
			LastName:    "Doe",
			WarehouseID: 1,
		}

		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Save(&newEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())
		require.Equal(t, 1, newEmployee.WarehouseID)

		// - function called
		rp.AssertNotCalled(t, "Save", &newEmployee)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 06 _ error: invalid employee (last name)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: last name", internal.ErrEmployeeServiceFieldRequired)
		newEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			// no last name
			WarehouseID: 0,
		}

		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Save(&newEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertNotCalled(t, "Save", &newEmployee)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 07 _ error: invalid employee (warehouse id)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: warehouse id", internal.ErrEmployeeServiceFieldRequired)
		newEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			// no warehouse id
		}

		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Save(&newEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertNotCalled(t, "Save", &newEmployee)
		// - expectations
		rp.AssertExpectations(t)
	})
}

// TestEmployeeDefault_Save_Invalid tests the Save method of the EmployeeDefault service with an invalid employee
func TestEmployeeDefault_Update(t *testing.T) {

	t.Run("case 01 _ succes: update employee by ID", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}
		rp.On("Update", mock.AnythingOfType("*internal.Employee")).Return(nil)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Update(&expectedEmployee)

		// assert
		// - error
		require.NoError(t, err)

		// - function called
		rp.AssertCalled(t, "Update", &expectedEmployee)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Update", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 02 _ error: employee not found", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		updatedEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}
		expectedError := fmt.Errorf("%w: %v", internal.ErrEmployeeServiceNotFound, internal.ErrEmployeeRepositoryNotFound)
		rp.On("Update", mock.AnythingOfType("*internal.Employee")).Return(internal.ErrEmployeeRepositoryNotFound)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Update(&updatedEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Update", &updatedEmployee)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Update", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 03 _ error: repository error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		updatedEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}
		expectedError := fmt.Errorf("%w: %v", internal.ErrEmployeeServiceInternalError, internal.ErrEmployeeRepository)
		rp.On("Update", mock.AnythingOfType("*internal.Employee")).Return(internal.ErrEmployeeRepository)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Update(&updatedEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Update", &updatedEmployee)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Update", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 04 _ error: unknown error", func(t *testing.T) {
		// assert
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		updatedEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  1,
		}
		expectedError := fmt.Errorf("%w: unknown error", internal.ErrEmployeeServiceUnknown)
		rp.On("Update", mock.AnythingOfType("*internal.Employee")).Return(errors.New("unknown error"))
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Update(&updatedEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Update", &updatedEmployee)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Update", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 05 _ error: invalid employee (first name)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: first name", internal.ErrEmployeeServiceFieldRequired)
		newEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			// no first name
			LastName:    "Doe",
			WarehouseID: 1,
		}

		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Update(&newEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertNotCalled(t, "Update", &newEmployee)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 06 _ error: invalid employee (last name)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: last name", internal.ErrEmployeeServiceFieldRequired)
		newEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			// no last name
			WarehouseID: 1,
		}

		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Update(&newEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertNotCalled(t, "Update", &newEmployee)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 07 _ error: invalid employee (warehouse id)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: warehouse id", internal.ErrEmployeeServiceFieldRequired)
		newEmployee := internal.Employee{
			ID:           1,
			CardNumberID: 1,
			FirstName:    "John",
			LastName:     "Doe",
			// no warehouse id
		}

		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Update(&newEmployee)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertNotCalled(t, "Update", &newEmployee)
		// - expectations
		rp.AssertExpectations(t)
	})

}

// TestEmployeeDefault_Delete tests the Delete method of the EmployeeDefault service
func TestEmployeeDefault_Delete(t *testing.T) {
	t.Run("case 01 _ succes: delete employee by ID", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		rp.On("Delete", 1).Return(nil)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Delete(1)

		// assert
		// - error
		require.NoError(t, err)

		// - function called
		rp.AssertCalled(t, "Delete", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Delete", 1)
		// = expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 02 _ error: employee not found", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expectations
		expectedError := fmt.Errorf("%w: %v", internal.ErrEmployeeServiceNotFound, internal.ErrEmployeeRepositoryNotFound)
		rp.On("Delete", 1).Return(internal.ErrEmployeeRepositoryNotFound)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Delete(1)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Delete", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Delete", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 03 _ error: repository error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: %v", internal.ErrEmployeeServiceInternalError, internal.ErrEmployeeRepository)
		rp.On("Delete", 1).Return(internal.ErrEmployeeRepository)
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Delete(1)

		// assert
		// - errpr
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function callled
		rp.AssertCalled(t, "Delete", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Delete", 1)
		// - expectations
		rp.AssertExpectations(t)
	})

	t.Run("case 04 _ error: unknown error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewEmployeeMock()
		// - expecteds
		expectedError := fmt.Errorf("%w: unknown error", internal.ErrEmployeeServiceUnknown)
		rp.On("Delete", 1).Return(errors.New("unknown error"))
		// - service
		sv := service.NewEmployeeDefault(rp)

		// act
		err := sv.Delete(1)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Delete", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Delete", 1)
		// - expectations
		rp.AssertExpectations(t)
	})
}
