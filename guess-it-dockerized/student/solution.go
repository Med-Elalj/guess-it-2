package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	numbers := []float64{}
	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			break
		}
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		numbers = append(numbers, num)
		min, max := predictNextRange(numbers)
		fmt.Printf("%.0f %.0f\n", min, max)
	}
}

func predictNextRange(numbers []float64) (float64, float64) {
	var X []float64
	var min, max float64
	if len(numbers) < 2 {
		return 0, 0
	}
	for i := 0; i < len(numbers); i++ {
		X = append(X, float64(i))
	}
	m, b, _ := calculateLRLandPCC(numbers)
	min = m*X[len(X)-1] + b - 20
	max = m*X[len(X)-1] + b + 20
	return min, max
}

func calculateLRLandPCC(numbers []float64) (float64, float64, float64) {
	n := float64(len(numbers))
	sumX, sumY, sumXY, sumX2, sumY2 := 0.0, 0.0, 0.0, 0.0, 0.0
	for i, v := range numbers {
		x := float64(i)
		y := v
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y
	}
	// Calculate slope (m) and intercept (b)
	m := ((n * sumXY) - sumX*sumY) / ((n * sumX2) - (sumX * sumX))
	b := (sumY - (m * sumX)) / n
	// Calculate Pearson Correlation Coefficient (r)
	r := (n*sumXY - sumX*sumY) / math.Sqrt((n*sumX2-sumX*sumX)*(n*sumY2-sumY*sumY))
	return m, b, r
}
