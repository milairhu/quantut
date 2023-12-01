package quantut

type QuantumCircuit struct {
	// The number of qubits in the circuit
	numQubits  int
	operations []Operation
}

func NewQuantumCircuit(numQubits int) *QuantumCircuit {
	o := make([]Operation, 0)
	return &QuantumCircuit{numQubits: numQubits, operations: o}
}

// Add Operations to the circuit
func (c *QuantumCircuit) H(qubit Qubit) {
	c.operations = append(c.operations, Operation{gate: h, qubits: []int{qubit}})
}
