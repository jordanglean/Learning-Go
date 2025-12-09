// Package fileops contains utility functions to access float values in file
package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(balance float64, fileName string) {
	balanceText := fmt.Sprint(balance)
	_ = os.WriteFile(fileName, []byte(balanceText), 0o644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	balanceText := string(data)

	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("failed to convert file value to float")
	}

	return balance, nil
}
