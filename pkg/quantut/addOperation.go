package quantut

// =============== Add gates to the circuit ===============

// Hadmard
func (c *QuantumCircuit) H(numQubit int) {
	if numQubit >= c.numQubits || numQubit < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: h, qubits: []int{numQubit}})
}

// Pauli X (NOT)
func (c *QuantumCircuit) X(numQubit int) {
	if numQubit >= c.numQubits || numQubit < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: x, qubits: []int{numQubit}})
}

// Pauli Y
func (c *QuantumCircuit) Y(numQubit int) {
	if numQubit >= c.numQubits || numQubit < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: y, qubits: []int{numQubit}})
}

// Pauli Z
func (c *QuantumCircuit) Z(numQubit int) {
	if numQubit >= c.numQubits || numQubit < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: z, qubits: []int{numQubit}})
}

// Identity
func (c *QuantumCircuit) I(numQubit int) {
	if numQubit >= c.numQubits || numQubit < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: i, qubits: []int{numQubit}})
}

// S gate
func (c *QuantumCircuit) S(numQubit int) {
	if numQubit >= c.numQubits || numQubit < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: s, qubits: []int{numQubit}})
}

// T gate
func (c *QuantumCircuit) T(numQubit int) {
	if numQubit >= c.numQubits || numQubit < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: t, qubits: []int{numQubit}})
}

//CONVENTION : control qubits are the first ones in the list of parameters

// CNOT
func (c *QuantumCircuit) CNOT(control int, target int) {
	if control >= c.numQubits || control < 0 || target >= c.numQubits || target < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: cnot, qubits: []int{control, target}})
}

// SWAP
func (c *QuantumCircuit) SWAP(qubit1 int, qubit2 int) {
	if qubit1 >= c.numQubits || qubit1 < 0 || qubit2 >= c.numQubits || qubit2 < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: swap, qubits: []int{qubit1, qubit2}})
}

// CS
func (c *QuantumCircuit) CS(control int, target int) {
	if control >= c.numQubits || control < 0 || target >= c.numQubits || target < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: cs, qubits: []int{control, target}})
}

// CCNOT
func (c *QuantumCircuit) CCNOT(control1 int, control2 int, target int) {
	if control1 >= c.numQubits || control1 < 0 || control2 >= c.numQubits || control2 < 0 || target >= c.numQubits || target < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: ccnot, qubits: []int{control1, control2, target}})
}

// CSWAP
func (c *QuantumCircuit) CSWAP(control int, target1 int, target2 int) {
	if control >= c.numQubits || control < 0 || target1 >= c.numQubits || target1 < 0 || target2 >= c.numQubits || target2 < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: cswap, qubits: []int{control, target1, target2}})
}
