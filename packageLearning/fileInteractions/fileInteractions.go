package fileInteractions

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filePath string) (float64, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return 1000, errors.New("Failed to read file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(filePath string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filePath, []byte(valueText), 0644)
}
