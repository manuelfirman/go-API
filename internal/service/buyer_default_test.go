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

// TestBuyerDefault_GetAll tests the GetAll method of the BuyerDefault service
func TestBuyerDefault_GetAll(t *testing.T) {
	t.Run("case 01 _ success: should return all buyers", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedBuyers := []internal.Buyer{
			{ID: 1, CardNumberID: 1, FirstName: "John", LastName: "Doe"},
			{ID: 2, CardNumberID: 2, FirstName: "Jane", LastName: "Doe"},
		}
		rp.On("GetAll").Return(expectedBuyers, nil)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		buyers, err := sv.GetAll()

		// assert
		// - error
		require.NoError(t, err)
		// - result
		require.Equal(t, expectedBuyers, buyers)

		// - function called
		rp.AssertCalled(t, "GetAll")
		// - number of calls
		rp.AssertNumberOfCalls(t, "GetAll", 1)
		// - expectation
		rp.AssertExpectations(t)

	})

	t.Run("case 02 _ error: should return an error (receive unknown error)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: unknown error", internal.ErrBuyerServiceUnknown)
		rp.On("GetAll").Return([]internal.Buyer{}, errors.New("unknown error"))
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		buyers, err := sv.GetAll()

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())
		// - result
		require.Empty(t, buyers)

		// - function called
		rp.AssertCalled(t, "GetAll")
		// - number of calls
		rp.AssertNumberOfCalls(t, "GetAll", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 03 _ error: should return an error (internal repository error)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: %v", internal.ErrBuyerService, internal.ErrBuyerRepository)
		rp.On("GetAll").Return([]internal.Buyer{}, internal.ErrBuyerRepository)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		buyers, err := sv.GetAll()

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())
		// - result
		require.Empty(t, buyers)

		// - function called
		rp.AssertCalled(t, "GetAll")
		// - number of calls
		rp.AssertNumberOfCalls(t, "GetAll", 1)
		// - expectation
		rp.AssertExpectations(t)
	})
}

// TestBuyerDefault_Get tests the Get method of the BuyerDefault service
func TestBuyerDefault_Get(t *testing.T) {
	t.Run("case 01 _ success: should return a buyer by ID", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedBuyer := internal.Buyer{ID: 1, CardNumberID: 1, FirstName: "John", LastName: "Doe"}
		rp.On("Get", 1).Return(expectedBuyer, nil)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		buyer, err := sv.Get(1)

		// assert
		// - error
		require.NoError(t, err)
		// - result
		require.Equal(t, expectedBuyer, buyer)

		// - function called
		rp.AssertCalled(t, "Get", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Get", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 02 _ error: buyer not found", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: %v", internal.ErrBuyerServiceNotFound, internal.ErrBuyerRepositoryNotFound)
		rp.On("Get", 1).Return(internal.Buyer{}, internal.ErrBuyerRepositoryNotFound)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		buyer, err := sv.Get(1)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())
		// - result
		require.Empty(t, buyer)

		// - function called
		rp.AssertCalled(t, "Get", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Get", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 03 _ error: internal repository error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: %v", internal.ErrBuyerService, internal.ErrBuyerRepository)
		rp.On("Get", 1).Return(internal.Buyer{}, internal.ErrBuyerRepository)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		buyer, err := sv.Get(1)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())
		// - result
		require.Empty(t, buyer)

		// - function called
		rp.AssertCalled(t, "Get", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Get", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 04 _ error: should return an error (receive unknown error)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: unknown error", internal.ErrBuyerServiceUnknown)
		rp.On("Get", 1).Return(internal.Buyer{}, errors.New("unknown error"))
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		buyer, err := sv.Get(1)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())
		// - result
		require.Empty(t, buyer)

		// - function called
		rp.AssertCalled(t, "Get", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Get", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

}

// TestBuyerDefault_Save tests the Save method of the BuyerDefault service
func TestBuyerDefault_Save(t *testing.T) {
	t.Run("case 01 _ success: save a buyer", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		newBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			FirstName:    "John",
			LastName:     "Doe",
		}
		rp.On("Save", mock.AnythingOfType("*internal.Buyer")).Return(nil)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Save(newBuyer)

		// assert
		// - error
		require.NoError(t, err)

		// - function called
		rp.AssertCalled(t, "Save", newBuyer)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Save", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 02 _ error: buyer already exists", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: %v", internal.ErrBuyerServiceDuplicated, internal.ErrBuyerRepositoryDuplicated)
		newBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			FirstName:    "John",
			LastName:     "Doe",
		}
		rp.On("Save", mock.AnythingOfType("*internal.Buyer")).Return(internal.ErrBuyerRepositoryDuplicated)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Save(newBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Save", newBuyer)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Save", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 03 _ error: internal repository error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: %v", internal.ErrBuyerService, internal.ErrBuyerRepository)
		newBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			FirstName:    "John",
			LastName:     "Doe",
		}
		rp.On("Save", mock.AnythingOfType("*internal.Buyer")).Return(internal.ErrBuyerRepository)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Save(newBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Save", newBuyer)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Save", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 04 _ error: should return an error (receive unknown error)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: unknown error", internal.ErrBuyerServiceUnknown)
		newBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			FirstName:    "John",
			LastName:     "Doe",
		}
		rp.On("Save", mock.AnythingOfType("*internal.Buyer")).Return(errors.New("unknown error"))
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Save(newBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Save", newBuyer)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Save", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 05 _ error: invalid buyer (card number id)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set expectation
		expectedError := fmt.Errorf("%w: card_number_id", internal.ErrBuyerServiceFieldRequired)
		newBuyer := &internal.Buyer{
			ID: 1,
			// no card number id
			FirstName: "John",
			LastName:  "Doe",
		}
		// ? dont need to set the expectation because the function will not be called
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Save(newBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function not called (validation error)
		rp.AssertNotCalled(t, "Save")
		// - number of calls
		rp.AssertNumberOfCalls(t, "Save", 0)
	})

	t.Run("case 06 _ error: invalid buyer (first name)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set expectation
		expectedError := fmt.Errorf("%w: first_name", internal.ErrBuyerServiceFieldRequired)
		newBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			// no first name
			LastName: "Doe",
		}
		// ? dont need to set the expectation because the function will not be called
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Save(newBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function not called (validation error)
		rp.AssertNotCalled(t, "Save")
		// - number of calls
		rp.AssertNumberOfCalls(t, "Save", 0)
	})

	t.Run("case 07 _ error: invalid buyer (last name)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set expectation
		expectedError := fmt.Errorf("%w: last_name", internal.ErrBuyerServiceFieldRequired)
		newBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			FirstName:    "John",
			// no last name
		}
		// ? dont need to set the expectation because the function will not be called
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Save(newBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function not called (validation error)
		rp.AssertNotCalled(t, "Save")
		// - number of calls
		rp.AssertNumberOfCalls(t, "Save", 0)
	})
}

// TestBuyerDefault_Update tests the Update method of the BuyerDefault service
func TestBuyerDefault_Update(t *testing.T) {
	t.Run("case 01 _ success: update a buyer", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		updatedBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			FirstName:    "John",
			LastName:     "Doe",
		}
		rp.On("Update", mock.AnythingOfType("*internal.Buyer")).Return(nil)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Update(updatedBuyer)

		// assert
		// - error
		require.NoError(t, err)

		// - function called
		rp.AssertCalled(t, "Update", updatedBuyer)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Update", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 02 _ error: buyer already exists", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: %v", internal.ErrBuyerServiceDuplicated, internal.ErrBuyerRepositoryDuplicated)
		updatedBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			FirstName:    "John",
			LastName:     "Doe",
		}
		rp.On("Update", mock.AnythingOfType("*internal.Buyer")).Return(internal.ErrBuyerRepositoryDuplicated)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Update(updatedBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Update", updatedBuyer)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Update", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 03 _ error: internal repository error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: %v", internal.ErrBuyerService, internal.ErrBuyerRepository)
		updatedBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			FirstName:    "John",
			LastName:     "Doe",
		}
		rp.On("Update", mock.AnythingOfType("*internal.Buyer")).Return(internal.ErrBuyerRepository)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Update(updatedBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Update", updatedBuyer)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Update", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 04 _ error: should return an error (receive unknown error)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: unknown error", internal.ErrBuyerServiceUnknown)
		updatedBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			FirstName:    "John",
			LastName:     "Doe",
		}
		rp.On("Update", mock.AnythingOfType("*internal.Buyer")).Return(errors.New("unknown error"))
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Update(updatedBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function called
		rp.AssertCalled(t, "Update", updatedBuyer)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Update", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 05 _ error: invalid buyer (card number id)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set expectation
		expectedError := fmt.Errorf("%w: card_number_id", internal.ErrBuyerServiceFieldRequired)
		newBuyer := &internal.Buyer{
			ID: 1,
			// no card number id
			FirstName: "John",
			LastName:  "Doe",
		}

		// ? dont need to set the expectation because the function will not be called

		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Update(newBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function not called (validation error)
		rp.AssertNotCalled(t, "Update")
		// - number of calls
		rp.AssertNumberOfCalls(t, "Update", 0)
	})

	t.Run("case 06 _ error: invalid buyer (first name)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set expectation
		expectedError := fmt.Errorf("%w: first_name", internal.ErrBuyerServiceFieldRequired)
		newBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			// no first name
			LastName: "Doe",
		}

		// ? dont need to set the expectation because the function will not be called

		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Update(newBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function not called (validation error)
		rp.AssertNotCalled(t, "Update")
		// - number of calls
		rp.AssertNumberOfCalls(t, "Update", 0)
	})

	t.Run("case 07 _ error: invalid buyer (last name)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set expectation
		expectedError := fmt.Errorf("%w: last_name", internal.ErrBuyerServiceFieldRequired)
		newBuyer := &internal.Buyer{
			ID:           1,
			CardNumberID: 23,
			FirstName:    "John",
			// no last name
		}

		// ? dont need to set the expectation because the function will not be called

		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Update(newBuyer)

		// assert
		// - error
		require.Error(t, err)
		require.EqualError(t, err, expectedError.Error())

		// - function not called (validation error)
		rp.AssertNotCalled(t, "Update")
		// - number of calls
		rp.AssertNumberOfCalls(t, "Update", 0)
	})
}

// TestBuyerDefault_Delete tests the Delete method of the BuyerDefault service
func TestBuyerDefault_Delete(t *testing.T) {
	t.Run("case 01 _ success: delete a buyer", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		rp.On("Delete", 1).Return(nil)
		// - service
		sv := service.NewBuyerDefault(rp)

		// act
		err := sv.Delete(1)

		// assert
		// - error
		require.NoError(t, err)

		// - function called
		rp.AssertCalled(t, "Delete", 1)
		// - number of calls
		rp.AssertNumberOfCalls(t, "Delete", 1)
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 02 _ error: buyer not found", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: %v", internal.ErrBuyerServiceNotFound, internal.ErrBuyerRepositoryNotFound)
		rp.On("Delete", 1).Return(internal.ErrBuyerRepositoryNotFound)
		// - service
		sv := service.NewBuyerDefault(rp)

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
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 03 _ error: internal repository error", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: %v", internal.ErrBuyerService, internal.ErrBuyerRepository)
		rp.On("Delete", 1).Return(internal.ErrBuyerRepository)
		// - service
		sv := service.NewBuyerDefault(rp)

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
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 04 _ error: foreign key constraint", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: %v", internal.ErrBuyerServiceFK, internal.ErrBuyerRepositoryFK)
		rp.On("Delete", 1).Return(internal.ErrBuyerRepositoryFK)
		// - service
		sv := service.NewBuyerDefault(rp)

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
		// - expectation
		rp.AssertExpectations(t)
	})

	t.Run("case 05 _ error: should return an error (receive unknown error)", func(t *testing.T) {
		// arrange
		// - repository mock
		rp := repository.NewBuyerMock()
		// - set the expectation
		expectedError := fmt.Errorf("%w: unknown error", internal.ErrBuyerServiceUnknown)
		rp.On("Delete", 1).Return(errors.New("unknown error"))
		// - service
		sv := service.NewBuyerDefault(rp)

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
		// - expectation
		rp.AssertExpectations(t)
	})
}
