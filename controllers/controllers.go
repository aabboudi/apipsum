package controllers

import (
	"math/rand"
	"strconv"
	"time"
)

func randomDate(startYear, endYear int) time.Time {
	rng := rand.New(rand.NewSource(int64(rand.Intn(100))))

	diff := time.Date(endYear, time.December, 31, 23, 59, 59, 0, time.UTC).Sub(
		time.Date(startYear, time.January, 1, 0, 0, 0, 0, time.UTC)).Seconds()
	randomSeconds := rng.Int63n(int64(diff) + 1)

	randomDuration := time.Duration(randomSeconds) * time.Second
	return time.Date(startYear, time.January, 1, 0, 0, 0, 0, time.UTC).Add(randomDuration)
}

// Must start with capital letter to be used outside this package
func GenerateData(schema map[string]string) map[string]interface{} {
	data := make(map[string]interface{})

	for key, valueType := range schema {
		switch valueType {
		case "string":
			data[key] = "name_" + strconv.Itoa(rand.Intn(1000))
		case "email":
			data[key] = "email_" + strconv.Itoa(rand.Intn(10000)) + "@domain.com"
		case "int":
			data[key] = rand.Intn(100)
		case "float":
			data[key] = rand.Float64() * 100
		case "bool":
			data[key] = rand.Intn(2) == 1
		case "datetime":
			data[key] = randomDate(1900, 2024)
		default:
			data[key] = nil
		}
	}

	return data
}
