package utils

import (
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
)

// CheckFieldExistance is a function that checks if the required fields exist in the fields map
func ValidateFieldExistence(data map[string]any, fields ...string) (err error) {
	isValid := true
	errStr := ""
	for _, field := range fields {
		if _, ok := data[field]; !ok {
			isValid = false
			errStr += "The field " + field + " is required\n"
		}
	}
	if !isValid {
		err = errors.New(errStr)
	}
	return
}

// Deprecated after v1.1.0
// ValidateAnyFieldExistence is a function that checks if any of the required fields exist in the fields map
func ValidateAnyFieldExistence(data map[string]any, fields ...string) (err error) {
	for _, field := range fields {
		if _, ok := data[field]; ok {
			return nil
		}
	}
	return errors.New("no valid fields found")
}

// ValidateBodyRequestWithStruct is a function that checks if the required fields exist in the fields map
func ValidateBodyRequestWithStruct(bytes []byte, jsonStruct any) (err error) {
	// Verify if the JSON fields match to map string/any
	jsonMap := make(map[string]any)
	err = json.Unmarshal(bytes, &jsonMap)
	if err != nil {
		err = errors.New("invalid datatype on JSON body request")

		return
	}

	// Verify if the JSON body is empty
	if len(jsonMap) == 0 {
		err = errors.New("empty JSON body request")

		return
	}

	// Deserialize JSON into an empty struct to get the keys and validate the fields types
	// from the bytes request and if match with the struct
	objType := reflect.TypeOf(jsonStruct)
	obj := reflect.New(objType).Interface()
	err = json.Unmarshal(bytes, &obj)
	if err != nil {
		err = errors.New("invalid datatype on JSON body request")

		return
	}

	// Check if the fields are present in the struct
	for key, value := range jsonMap {
		found := false
		for i := 0; i < objType.NumField(); i++ {
			fieldName := objType.Field(i).Tag.Get("json")
			fieldType := objType.Field(i).Type
			if key == fieldName {
				found = true
				// Perform explicit type conversion if necessary
				fieldValue := reflect.ValueOf(value)
				if fieldValue.Type() != fieldType {
					// If the types do not match, try to convert the value to the expected type
					convertedValue := reflect.New(fieldType).Elem()
					convertedValue.Set(fieldValue.Convert(fieldType))
				}
				break
			}
		}
		if !found {
			err = errors.New("Field " + key + " not found in the struct")
			return
		}
	}

	return
}

// Errors validation implementation
var (
	ErrInvalidFormatPhoneNumber = errors.New("invalid phone number format, must be 123-456-7890 or 1234567890 or 12345678")
)

// Validate phone number format
// format expected is 123-456-7890 or 1234567890 or 12345678
func ValidatePhoneNumber(phone string) (err error) {
	// 123-456-7890 or 1234567890 or 12345678
	phoneRegex := regexp.MustCompile(`^(\d{3}-\d{3}-\d{4}|\d{10}|\d{8}|\d{7})$`)
	if !phoneRegex.MatchString(phone) {
		err = ErrInvalidFormatPhoneNumber
	}
	return
}
