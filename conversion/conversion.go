package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloats(lines []string) ([]float64, error) {
	var floats []float64
	for _, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Converting price to float failed. ", err)
			return nil, err
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}
