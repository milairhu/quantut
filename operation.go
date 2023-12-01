package quantut

type Operation struct {
	gate   Gate    //Gate used
	qubits []Qubit //Qubits involved : first is the control, second is the target
}

// Getters
func (o Operation) Gate() Gate {
	return o.gate
}

func (o Operation) Qubits() []Qubit {
	return o.qubits
}
