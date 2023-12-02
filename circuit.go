package quantut

import (
	"fmt"
	"quantut/utils"
)

type QuantumCircuit struct {
	numQubits         int         //Number of qubits involved
	operations        []Operation //Operations to apply on the circuit
	qubitsValues      []Qubit     //Values of the qubits
	classicalRegister []int       //Values of the classical register
}

func NewQuantumCircuit(numQubits int) *QuantumCircuit {
	o := make([]Operation, 0)
	qv := make([]Qubit, numQubits)
	//Init qubits values to 0
	for i := 0; i < numQubits; i++ {
		qv[i].Init(1, 0)
	}
	r := make([]int, 0)
	return &QuantumCircuit{numQubits: numQubits, operations: o, qubitsValues: qv, classicalRegister: r}
}

// ===== Getters =====
func (c *QuantumCircuit) NumQubits() int {
	return c.numQubits
}

func (c *QuantumCircuit) Operations() []Operation {
	return c.operations
}

func (c *QuantumCircuit) QubitsValues() []Qubit {
	return c.qubitsValues
}

func (c *QuantumCircuit) ClassicalRegister() []int {
	return c.classicalRegister
}

// === Setters ===
func (c *QuantumCircuit) SetQubit(numQubit int, comp1 complex128, comp2 complex128) {
	if numQubit >= c.numQubits || numQubit < 0 {
		panic(fmt.Sprintf("Qubit number out of range : %d", numQubit))
	}
	if !utils.IsNormalized(comp1, comp2) {
		panic(fmt.Sprintf("Qubit value must be normalized : %f^2+%f^2 = %f", comp1, comp2, comp1*comp1+comp2*comp2))
	}

	c.qubitsValues[numQubit].Init(comp1, comp2)
}

// Define number of bits in the classical register
func (c *QuantumCircuit) InitClassicalRegister(numRegister uint8) {

	c.classicalRegister = make([]int, numRegister)
}

// Set value of a bit in the classical register
func (c *QuantumCircuit) SetClassicalRegister(numRegister int, value int) {
	if numRegister >= len(c.classicalRegister) || numRegister < 0 {
		panic("Register number out of range")
	}
	if value != 0 && value != 1 {
		panic("Register value must be 0 or 1")
	}
	c.classicalRegister[numRegister] = value
}
