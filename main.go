package main

import (
	"calculator/prices"
)

func main() {
	// prices {10, 20, 30}
	// taxRates {0, 0.07, 0.1, 0.15}
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
