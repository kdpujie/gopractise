package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// Generate a 6×6 matrix of random values.
	data := make([]float64, 36)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	a := mat.NewDense(6, 6, data)
	a.Apply()
	fmt.Print(a)
}
