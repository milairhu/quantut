package quantut

type QuantumCircuit struct {
	// The number of qubits in the circuit
	numQubits  int
	operations []Gate
}

func NewQuantumCircuit(numQubits int) *QuantumCircuit {
	return &QuantumCircuit{numQubits: numQubits}
}
