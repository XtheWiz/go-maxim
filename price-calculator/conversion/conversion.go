package conversion

import (
	"errors"
	"strconv"
)

func StringToFloats(inputStrings []string) ([]float64, error) {
	var floats []float64
	for _, str := range inputStrings {
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}
