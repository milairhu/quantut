package quantut

type Qubit [2]complex128 //Qubit is a slice of  complexes

func (q *Qubit) Init(comp1 complex128, comp2 complex128) {
	if comp1*comp1+comp2*comp2 != 1 {
		panic("Qubit value must be normalized")
	}
	q[0] = comp1
	q[1] = comp2
}
