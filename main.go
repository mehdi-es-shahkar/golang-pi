package main

import (
	"fmt"
	"math"
)

// Integral approximates ∫ₐᵇ f(x)dx using right Riemann sum
func Integral(f func(float64) float64, a, b float64, n int64) float64 {
	dx := (b - a) / float64(n)
	sum := 0.0

	for i := int64(1); i <= n; i++ {
		x := a + float64(i)*dx
		sum += f(x) * dx
	}
	return float64(4) * sum
}

func main() {
	// circule function for calculate integral --- > ∫₀¹sqrt(1 - x²) dx
	circule := func(x float64) float64 { return math.Sqrt(1 - x*x) }
	fmt.Println("pi number is:", Integral(circule, 0, 1, 1000000))

}
