package controllers

import (
	"apipsum/utils"
	"errors"
	"strings"
	"time"
)

func GenerateData(schema map[string]interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	var err error

	for key, field := range schema {
		fieldMap := field.(map[string]interface{})
		fieldType := strings.ToLower(fieldMap["type"].(string))

		switch fieldType {
		case "bool":
			data[key], err = handleBool()

		case "int":
			data[key], err = handleInt(fieldMap)

		case "float":
			data[key], err = handleFloat(fieldMap)

		case "string":
			data[key], err = handleString(fieldMap)

		case "email":
			data[key], err = handleEmail()

		case "date":
			data[key], err = handleDate()

		case "uuid":
			data[key], err = handleUUID(fieldMap)

		case "phone_number":
			data[key], err = handlePhoneNumber(fieldMap)

		default:
			return nil, errors.New("invalid data type")
		}

		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func handleBool() (bool, error) {
	return utils.RandomBool()
}

func handleString(fieldMap map[string]interface{}) (string, error) {
	var maxLength int

	if val, ok := fieldMap["max_length"]; ok {
		switch v := val.(type) {
		case float64:
			maxLength = int(v)
		case int:
			maxLength = v
		default:
			return "", errors.New("max_length must be a valid numbern")
		}
	}

	if maxLength > 0 {
		return utils.RandomString(maxLength)
	} else {
		return utils.RandomString()
	}
}

func handlePhoneNumber(fieldMap map[string]interface{}) (string, error) {
	var format string
	if val, ok := fieldMap["format"]; ok {
		format = val.(string)
	}

	if format == "" {
		return utils.RandomPhoneNumber()
	} else {
		return utils.RandomPhoneNumber(format)
	}
}

func handleUUID(fieldMap map[string]interface{}) (string, error) {
	var version, variant int
	if val, ok := fieldMap["version"]; ok {
		version = int(val.(float64))
	}
	if val, ok := fieldMap["variant"]; ok {
		variant = int(val.(float64))
	} else {
		variant = 2
	}

	return utils.RandomUUID(version, variant)
}

func handleEmail() (string, error) {
	return utils.RandomEmail()
}

func handleInt(fieldMap map[string]interface{}) (int, error) {
	var min, max float64
	if val, ok := fieldMap["min"]; ok {
		min = val.(float64)
	}
	if val, ok := fieldMap["max"]; ok {
		max = val.(float64)
	}

	return utils.RandomInt(int(min), int(max))
}

func handleFloat(fieldMap map[string]interface{}) (float64, error) {
	var min, max float64
	if val, ok := fieldMap["min"]; ok {
		min = val.(float64)
	}
	if val, ok := fieldMap["max"]; ok {
		max = val.(float64)
	}

	return utils.RandomFloat(min, max)
}

func handleDate() (time.Time, error) {
	return utils.RandomDate(1900, 2024)
}
