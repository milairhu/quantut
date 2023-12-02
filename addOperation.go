package quantut

import (
	"fmt"
	"quantut/utils"
)

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

// TOFFOLI
func (c *QuantumCircuit) TOFFOLI(control1 int, control2 int, target int) {
	if control1 >= c.numQubits || control1 < 0 || control2 >= c.numQubits || control2 < 0 || target >= c.numQubits || target < 0 {
		panic("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: toffoli, qubits: []int{control1, control2, target}})
}

// =============== Initialize Qubits ===============
func (c *QuantumCircuit) InitializeQubit(numQubit int, comp1 complex128, comp2 complex128) {
	if numQubit >= c.numQubits || numQubit < 0 {
		panic("Qubit number out of range")
	}
	if !utils.IsNormalized(comp1, comp2) {
		panic(fmt.Sprintf("Qubit value must be normalized : %f^2+%f^2 = %f", comp1, comp2, comp1*comp1+comp2*comp2))
	}

	c.operations = append(c.operations, Operation{gate: initialization, qubits: []int{numQubit}, options: []complex128{comp1, comp2}})

}
