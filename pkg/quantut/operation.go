package quantut

type Operation struct {
	gate    Gate         //Gate used
	qubits  []int        //Index of Qubits involved : first is the control, second is the target
	options []complex128 //Options of the gate that can be used later for Initialization
}

// Getters
func (o *Operation) Gate() Gate {
	return o.gate
}

func (o *Operation) Qubits() []int {
	return o.qubits
}

func (o *Operation) Options() []complex128 {
	return o.options
}

// Utils
func (o *Operation) IsQubitInOperation(numQubit int) bool {
	for _, q := range o.qubits {
		if q == numQubit {
			return true
		}
	}
	return false
}

// Return true if the qubit is a control qubit for the operation
func (o *Operation) IsControlQubit(numQubit int) bool {
	if o.gate.NbControlQubit() == 0 {
		return false
	}
	for i := 0; i < int(o.gate.NbControlQubit()); i++ {
		if o.qubits[i] == numQubit {
			return true
		}
	}
	return false
}
