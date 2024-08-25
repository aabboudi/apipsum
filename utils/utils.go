package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomInt(optional ...int) (int, error) {
	min, max := 0, 100

	switch len(optional) {
	case 0:
		break
	case 1:
		max = optional[0]
	case 2:
		min = optional[0]
		max = optional[1]
	default:
		return 0, errors.New("too many arguemnts")
	}

	if min > max {
		return 0, errors.New("min cannot be greater than max")
	}

	return rand.Intn(max-min+1) + min, nil
}

func RandomFloat(optional ...float64) (float64, error) {
	min, max := 0.0, 100.0

	switch len(optional) {
	case 0:
		break
	case 1:
		max = optional[0]
	case 2:
		min = optional[0]
		max = optional[1]
	default:
		return 0, errors.New("too many arguemnts")
	}

	if min > max {
		return 0, errors.New("min cannot be greater than max")
	}

	return min + rand.Float64()*(max-min), nil
}

func RandomDate(startYear, endYear int) (time.Time, error) {
	rng := rand.New(rand.NewSource(int64(rand.Intn(100))))

	diff := time.Date(endYear, time.December, 31, 23, 59, 59, 0, time.UTC).Sub(
		time.Date(startYear, time.January, 1, 0, 0, 0, 0, time.UTC)).Seconds()
	randomSeconds := rng.Int63n(int64(diff) + 1)

	randomDuration := time.Duration(randomSeconds) * time.Second
	return time.Date(startYear, time.January, 1, 0, 0, 0, 0, time.UTC).Add(randomDuration), nil
}

func RandomString(optional ...int) (string, error) {
	var length int

	switch len(optional) {
	case 0:
		length = rand.Intn(8) + 3
	case 1:
		length = optional[0]
		if length <= 0 {
			return "", errors.New("length must be greater than 0")
		}
	default:
		return "", errors.New("too many arguments")
	}

	vowels := "aeiou"
	consonants := "bcdfghjklmnpqrstvwxyz"

	var word strings.Builder
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			word.WriteByte(consonants[rand.Intn(len(consonants))])
		} else {
			word.WriteByte(vowels[rand.Intn(len(vowels))])
		}
	}

	return word.String(), nil
}

func RandomEmail() (string, error) {
	username, err := RandomString()
	if err != nil {
		return "", err
	}

	domain, err := RandomString()
	if err != nil {
		return "", err
	}

	tld, err := RandomString(3)
	if err != nil {
		return "", err
	}

	return username + "@" + domain + "." + tld, nil
}

func RandomUUID(optional ...int) (string, error) {
	version, variant := 4, 2

	if len(optional) > 0 {
		if optional[0] < 1 || optional[0] > 5 {
			return "", fmt.Errorf("invalid UUID version: %d\nmust be between 1 and 5", optional[0])
		} else {
			version = optional[0]
		}
	}

	if len(optional) == 2 {
		variant = optional[1]
	} else if len(optional) > 2 {
		return "", errors.New("too many arguments")
	}

	uuid := make([]byte, 16)
	for i := range uuid {
		uuid[i] = byte(rand.Intn(256))
	}

	uuid[6] = (uuid[6] & 0x0f) | byte(version<<4)

	switch variant {
	case 0: // NCS backward compatibility
		uuid[8] = (uuid[8] & 0x7f)
	case 2: // RFC 4122/DCE 1.1
		uuid[8] = (uuid[8] & 0x3f) | 0x80
	case 6: // Microsoft Corporation GUID
		uuid[8] = (uuid[8] & 0x1f) | 0xc0
	case 7: // Future use
		uuid[8] = (uuid[8] & 0x1f) | 0xe0
	default:
		return "", fmt.Errorf("invalid UUID variant: %d\nplease choose between 0, 2, 6, and 7", variant)
	}

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func RandomPhoneNumber(optional ...string) (string, error) {
	formats := map[string]string{
		"us1":         "(XXX) XXX-XXXX",
		"us2":         "XXX-XXX-XXXX",
		"intl":        "+XX XXXXXXXXXX",
		"intl_dashed": "+XX-XXXX-XXXXXX",
		"intl_area":   "+XX (XXX) XXX-XXXX",
		"simple":      "XXXXX-XXXXX",
		"dotted":      "XXX.XXX.XXXX",
		"plain":       "XXXXXXXXXX",
	}

	var format_key string
	if len(optional) == 0 {
		format_key = "us1"
	} else if len(optional) == 1 {
		format_key = optional[0]
	} else {
		return "", errors.New("too many arguments")
	}

	if format, exists := formats[format_key]; exists {
		var result strings.Builder
		for _, char := range format {
			if char == 'X' {
				result.WriteByte(byte(rand.Intn(10) + '0'))
			} else {
				result.WriteByte(byte(char))
			}
		}
		return result.String(), nil
	} else {
		return "", fmt.Errorf("invalid format code: %s\nplease select a valid format", format_key)
	}
}
