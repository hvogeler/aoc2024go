package main

import (
	"fmt"
	"math"
	"gonum.org/v1/gonum/optimize"
)

// Data
var indices = []float64{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
}

var values = []float64{
	2, 3, 3, 5, 8, 12, 22, 36, 43, 64, 108,
	173, 268, 365, 563, 877, 1327, 2087, 3030, 4566,
	7092, 10752, 16591, 24617, 37168, 57612, 86607,
	132428, 200383, 302093, 464293, 700640, 1067522,
	1622761, 2446758, 3746925, 5674664, 8611394,
	13117983, 19831861, 30246720, 45899885, 69592802,
	106032328, 160551878, 244255259, 371250148,
	562795064, 856649243, 1299173551,
}

// Exponential function
func exponential(params []float64, x float64) float64 {
	a, b, c := params[0], params[1], params[2]
	return a*math.Exp(b*x) + c
}

// Residual function for optimization
func residuals(params []float64) float64 {
	var sum float64
	for i := range indices {
		predicted := exponential(params, indices[i])
		error := predicted - values[i]
		sum += error * error
	}
	return sum
}

// Main
func main() {
	// Initial guesses for a, b, c
	initialParams := []float64{1.0, 0.1, 1.0}

	// Optimization using Gonum's minimize package
	result, err := optimize.Minimize(optimize.Problem{
		Func: residuals,
	}, initialParams, nil, nil)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Optimized parameters
	params := result.X
	fmt.Printf("Optimized Parameters: a=%.4f, b=%.4f, c=%.4f\n", params[0], params[1], params[2])

	// Print approximations
	fmt.Println("\nApproximations:")
	for i := range indices {
		approx := exponential(params, indices[i])
		fmt.Printf("Index %d: Approximation=%.2f, Actual=%.2f\n", i, approx, values[i])
	}

	fmt.Printf("75th Generation will have around %.0f stones", exponential(params, 75))
}