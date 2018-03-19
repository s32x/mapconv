package mapconv

import (
	"fmt"
	"reflect"
	"strings"
)

// MtosJSON is a JSON specific helper method that will populate a
// structs data using struct field Names as well as JSON tags
func MtosJSON(m map[string]interface{}, out interface{}) error {
	return Mtos(m, out, "json")
}

// Mtos populates the passed out struct with all other structs
// and values corresponding in the passed map. Tags can also be passed
// to specify supported struct tag map keys
func Mtos(m map[string]interface{}, out interface{}, tags ...string) error {
	t := reflect.TypeOf(out).Elem()    // The type of the struct
	v := reflect.ValueOf(out).Elem()   // The value of the struct
	return setStruct(t, v, m, tags...) // Attempt to populate the struct data
}

// setStruct iterates over every field on the struct and attempts to
// find and set a corresponding map value using struct field names or
// any tags passed
func setStruct(t reflect.Type, v reflect.Value, m map[string]interface{}, tags ...string) error {
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i) // The structs reflect.StructField
		fieldValue := v.Field(i)  // The fields reflect.Value

		// Verify the field can be set
		if !fieldValue.CanSet() {
			fmt.Printf("Cannot set %s field value", fieldValue.String())
			continue
		}

		// Attempt to find the value in the map that corresponds to a
		// any of the passed struct tags values
		if len(tags) > 0 {
			for _, tag := range tags {
				if val, ok := m[structTagVal(structField.Tag, tag)]; ok {
					if err := setField(fieldValue.Type(), fieldValue, val, tags...); err != nil {
						return err
					}
					continue
				}
			}
		}

		// Find the value in the map that corresponds to the struct field name
		if val, ok := m[structField.Name]; ok {
			if err := setField(fieldValue.Type(), fieldValue, val); err != nil {
				return err
			}
			continue
		}
	}
	return nil
}

// setField assigns a statically typed map value to the reflected struct field Value
func setField(rt reflect.Type, rv reflect.Value, val interface{}, tags ...string) error {
	// If there's another tree of nested data recursively set
	// child struct data
	if nestedMap, ok := val.(map[string]interface{}); ok {
		return setStruct(rt, rv, nestedMap, tags...)
	}

	// Get the value of the field in the map
	v := reflect.ValueOf(val)

	// Verify the map and struct fields match in type before setting
	if valType := v.Type(); valType != rt {
		return fmt.Errorf("Provided map value type doesn't match struct field type : %s != %s", valType, rt)
	}

	// Set the value of the field on the struct and return
	rv.Set(v)
	return nil
}

// structTagVal recieves and parses a collection of struct tags and
// returns the cooresponding tags value
func structTagVal(tags reflect.StructTag, tag string) string {
	tagVal := tags.Get(tag)
	tagVal = strings.Split(tagVal, ",")[0] // Strip off optional params
	return tagVal
}
