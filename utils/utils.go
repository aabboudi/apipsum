package utils

import (
	"errors"
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
