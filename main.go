package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	t_start := time.Now()
	d := integral(circle, 0.0, 26.0, 126000.0, 26.0)
	fmt.Println(d)
	//t := time.Since(t_start)
	t := time.Since(t_start)
	fmt.Printf("total of run time is %d\n", t)

}

func circle(x float64, r float64) float64 {
	return math.Sqrt(math.Pow(r, 2) - math.Pow(x, 2))
}

type operation func(float64, float64) float64

func integral(op operation, a float64, b float64, n float64, r float64) float64 {
	s := 0.0
	dx := (b - a) / n
	fmt.Println(dx)
	for i := 1.0; i <= n; i++ {
		x := a + (i * dx)
		h := op(x, r)
		s += dx * h
	}
	fmt.Printf("n is : %f\n", n)
	return (4 * s) / (math.Pow(r, 2))
	//return 4 * s

}
