package quantut

import "errors"

type Qubit [2]complex128 //Qubit is a slice of  complexes

func (q *Qubit) Init(comp1 complex128, comp2 complex128) error {
	if comp1*comp1+comp2*comp2 != 1 {
		return errors.New("values not normalized for initializing the qubit")
	}
	q[0] = comp1
	q[1] = comp2
	return nil
}
