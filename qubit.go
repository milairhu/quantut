package quantut

import (
	"fmt"
	"math"
	"quantut/utils"
)

type Qubit [2]complex128 //Qubit is a slice of  complexes

func (q *Qubit) Init(comp1 complex128, comp2 complex128) {
	if !utils.IsNormalized(comp1, comp2) {
		panic(fmt.Sprintf("Qubit value must be normalized : %f^2+%f^2 = %f", comp1, comp2, comp1*comp1+comp2*comp2))
	}
	q[0] = comp1
	q[1] = comp2
}

func (q *Qubit) Vector() [][]complex128 {
	return [][]complex128{{q[0]}, {q[1]}}
}

func (q *Qubit) IsNormalized() bool {
	module := q[0]*q[0] + q[1]*q[1]
	re := math.Round(real(module))
	im := math.Round(imag(module))
	return (re == 1 || re == -1) && im == 0
}
