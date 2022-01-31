package reflect

import (
	"fmt"
	"reflect"
)

func AssignValue(fldV interface{}, fldO reflect.Value) (err error) {
	var fldA interface{}
	var ok bool

	err = fmt.Errorf("Invalid specification of struct field")

	if reflect.TypeOf(fldV).Kind() == reflect.Struct {
		fldA = struct{}{}
		ok = true
	} else {
		switch fldV.(type) {
		case string:
			fldA, ok = fldV.(string)
		case int:
			fldA, ok = fldV.(int)
		case float64:
			fldA, ok = fldV.(float64)
		default:
			return err
		}
	}

	if !ok {
		return err
	}

	if fldO.Type().AssignableTo(reflect.TypeOf(fldA)) {
		fldO.Set(reflect.ValueOf(fldA))
		return nil
	}

	return err
}

func AssignToStruct(in interface{}, values map[string]interface{}) error {
	inValue := reflect.ValueOf(in)

	if inValue.Kind() != reflect.Ptr {
		return fmt.Errorf("Invalid specification of \"in\", should be pointer")
	}

	inValue = inValue.Elem()

	if inValue.Kind() != reflect.Struct {
		return fmt.Errorf("Invalid specification of \"&in\", should be struct")
	}

	for key, value := range values {
		for i := 0; i < inValue.Type().NumField(); i++ {
			fld := inValue.Type().Field(i)
			if key == fld.Name {
				fldO := inValue.FieldByName(fld.Name)

				if err := AssignValue(value, fldO); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
