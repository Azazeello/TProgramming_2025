package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, x := 7.2, 1.3, 1.56
	y := math.Sqrt((a + b*x) / math.Pow(math.Log10(x), 3))
	fmt.Printf("y = %.4f\n", y)
}
