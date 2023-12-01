package quantut

type Operation struct {
	gate   Gate  //Gate used
	qubits []int //Index of Qubits involved : first is the control, second is the target
}

// Getters
func (o Operation) Gate() Gate {
	return o.gate
}

func (o Operation) Qubits() []int {
	return o.qubits
}
