package agendor

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func StructToQueryString(data interface{}) (string, error) {
	queryParams := url.Values{}

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("input is not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		tag := fieldType.Tag.Get("query")
		if tag == "" || tag == "-" {
			tag = fieldType.Name
		}

		switch field.Kind() {
		case reflect.String:
			if field.String() != "" {
				queryParams.Add(tag, field.String())
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.Int() != 0 {
				queryParams.Add(tag, strconv.FormatInt(field.Int(), 10))
			}
		case reflect.Bool:
			if field.Bool() {
				queryParams.Add(tag, strconv.FormatBool(field.Bool()))
			}
		case reflect.Slice:
			for j := 0; j < field.Len(); j++ {
				queryParams.Add(tag, fmt.Sprintf("%v", field.Index(j)))
			}
		}
	}

	return queryParams.Encode(), nil
}

// UnmarshalJSON converte discount de string ou float64 para Float64OrString
func (f *Discount) UnmarshalJSON(data []byte) error {
	// Tenta primeiro converter o dado para float64
	var floatVal float64
	if err := json.Unmarshal(data, &floatVal); err == nil {
		*f = Discount(floatVal)
		return nil
	}

	// Se falhar, tenta converter o dado como string e depois para float64
	var strVal string
	if err := json.Unmarshal(data, &strVal); err == nil {
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			return err
		}
		*f = Discount(floatVal)
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into Discount", string(data))
}
