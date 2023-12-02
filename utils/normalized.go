package utils

import (
	"math"
)

/*
	func Round(num float64, accuracy uint) int {
		factor := 1
		for i := uint(0); i < accuracy; i++ {
			factor *= 10
		}
		return int(num * float64(accuracy))
	}
*/
func IsNormalized(comp1, comp2 complex128) bool {
	module := comp1*comp1 + comp2*comp2
	re := math.Round(real(module))
	im := math.Round(imag(module))
	return (re == 1 || re == -1) && im == 0
}
