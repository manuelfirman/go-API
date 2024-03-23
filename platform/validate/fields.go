package validate

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

var (
	ErrFieldNotExists = fmt.Errorf("field not exists in data and is required")
)

type FieldError struct {
	Field string
	Msg   error
}

func (f *FieldError) Error() string {
	return fmt.Sprintf("field %s: %s", f.Field, f.Msg)
}

// CheckCompleteFields is a function that checks if all fields are complete
func CheckFieldExistance(s interface{}, data map[string]any) error {
	//get type of s
	t := reflect.TypeOf(s)
	//get fields of s
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Tag.Get("json")
		//fmt.Println(strings.ToLower(field))
		//check if field exists in data
		if _, ok := data[field]; !ok {
			if _, ok := data[strings.ToLower(field)]; !ok {
				return &FieldError{Field: field, Msg: ErrFieldNotExists}
			}
		}
	}
	return nil
}

// CheckCorrectField is a function that checks if a field is correct and exists in interface provided
func CheckCorrectsFields(s interface{}, data map[string]any) error {
	//instance slice of s fields
	fields := []string{}

	//get type of s
	t := reflect.TypeOf(s)
	//get fields of s
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Tag.Get("json"))
	}

	//check data fields and validate if exists in fields slice
	for key := range data {
		if !slices.Contains(fields, key) {
			return &FieldError{Field: key, Msg: fmt.Errorf("field not exists in struct")}
		}
	}

	return nil
}

// GetFields is a function that returns the field name of a struct
func GetFields(s interface{}, key string) (field string, err error) {
	//instance slice of s fields
	fields := []string{}

	//get type of s
	t := reflect.TypeOf(s)
	//get fields of s
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
	}

	for _, value := range fields {
		if strings.ToLower(value) == strings.ReplaceAll(key, "_", "") {
			return value, nil
		}
	}

	return "", fmt.Errorf("field not exists in struct")
}

// Decode is a function that decodes a map into a struct and returns the struct and an error
func Decode(m map[string]interface{}, s any) error {
	//convert to json
	jsonData, err := json.Marshal(m)
	if err != nil {
		return err
	}
	//convert to struct
	if err := json.Unmarshal(jsonData, &s); err != nil {
		return err
	}
	return nil
}
