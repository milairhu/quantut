package utils

import "fmt"

// TensorialProduct compute the tensorial product of a slice of qubits
func TensorialProduct(qubits [][]complex128) []complex128 {
	if len(qubits) == 0 {
		panic("Tensorial product impossible : no qubits")
	}
	var nbComposante = 2
	for i := 0; i < len(qubits)-1; i++ {
		nbComposante *= 2
	}
	res := make([]complex128, nbComposante)
	//res is filled to get tensorial product
	// use a technique based on bit representation of the index of the vector
	for curr := 0; curr < nbComposante; curr++ {
		//convert index of result into bit representation
		var bitRepresentation = fmt.Sprintf("%b", curr)
		//complete bit representation with 0
		for len(bitRepresentation) < len(qubits) {
			bitRepresentation = "0" + bitRepresentation
		}
		fmt.Println(bitRepresentation)

		//compute tensorial product for the index
		res[curr] = 1
		for i := 0; i < len(qubits); i++ {
			if bitRepresentation[i] == '0' {
				res[curr] *= qubits[i][0]
			} else {
				res[curr] *= qubits[i][1]
			}
		}
	}

	return res
}
