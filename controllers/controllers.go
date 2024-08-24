package controllers

import (
	"apipsum/utils"
)

func GenerateData(schema map[string]interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	var err error

	for key, field := range schema {
		fieldMap := field.(map[string]interface{})
		fieldType := fieldMap["type"].(string)

		var maxLength int
		if val, ok := fieldMap["max_length"]; ok {
			maxLength = int(val.(float64))
		}

		var min, max float64
		if val, ok := fieldMap["min"]; ok {
			min = val.(float64)
		}
		if val, ok := fieldMap["max"]; ok {
			max = val.(float64)
		}

		switch fieldType {
		case "bool":
			data[key] = utils.RandomBool()
		case "int":
			data[key], err = utils.RandomInt(int(min), int(max))
		case "float":
			data[key], err = utils.RandomFloat(min, max)
		case "string":
			if maxLength > 0 {
				data[key], err = utils.RandomString(maxLength)
			} else {
				data[key], err = utils.RandomString()
			}
		case "email":
			data[key], err = utils.RandomEmail()
		case "date":
			data[key], err = utils.RandomDate(1900, 2024)
		default:
			data[key] = nil
		}

		if err != nil {
			return nil, err
		}
	}

	return data, nil
}
