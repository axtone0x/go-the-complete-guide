package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Failed to find file.")
	}

	valueText := string(data)
	balance, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, errors.New("Failed to parse stored value.")
	}

	return balance, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value) // This gives us a string
	os.WriteFile(fileName, []byte(valueText), 0644)
}