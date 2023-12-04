package quantut

import (
	"fmt"

	"github.com/milairhu/quantut/utils"
)

type QuantumCircuit struct {
	numQubits         int          //Number of qubits involved
	operations        []Operation  //Operations to apply on the circuit
	qubitsValues      []Qubit      //Values of the qubits
	classicalRegister []int        //Values of the classical register
	globalState       []complex128 //Global state of the circuit. Represente the tensorial product of all qubits
}

func NewQuantumCircuit(numQubits int) *QuantumCircuit {
	o := make([]Operation, 0)
	qv := make([]Qubit, numQubits)
	//Init qubits values to 0
	for i := 0; i < numQubits; i++ {
		qv[i].Init(1, 0)
	}
	r := make([]int, 0)
	var nbComposante = 2
	for i := 0; i < numQubits; i++ {
		nbComposante *= 2
	}
	gs := make([]complex128, nbComposante)
	//Global state is initialized to 0...0
	gs[0] = 1
	return &QuantumCircuit{numQubits: numQubits, operations: o, qubitsValues: qv, classicalRegister: r, globalState: gs}
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

func (c *QuantumCircuit) GlobalState() []complex128 {
	return c.globalState
}

func (c *QuantumCircuit) ClassicalRegisterValue() []int {
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

func (c *QuantumCircuit) SetGlobalState(newState []complex128) {
	if len(newState) != len(c.globalState) {
		panic("New state must have the same length as the global state")
	}
	c.globalState = newState
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

// Return an array of all qubit values
func (c *QuantumCircuit) Qubits() [][]complex128 {
	res := make([][]complex128, c.numQubits)
	for i := 0; i < c.numQubits; i++ {
		res[i] = c.qubitsValues[i].ToArrComplex128()
	}
	return res
}

// ===== Combine circuits =====
// Add the operations of the circuit given in paramter to the current circuit
// A tester
func (c *QuantumCircuit) Compose(circuit *QuantumCircuit) *QuantumCircuit {

	if c.numQubits < circuit.numQubits {
		panic("The circuit to add has more qubits than the current circuit")
	}
	if len(c.classicalRegister) < len(circuit.classicalRegister) {
		panic("The circuit to add has more bits in the classical register than the current circuit")
	}

	resCirc := NewQuantumCircuit(c.numQubits)                          //Create a new circuit with the same number of qubits as the current circuit
	resCirc.operations = append(c.operations, c.Operations()...)       //Add the operations of the current circuit to the new circuit
	resCirc.operations = append(c.operations, circuit.Operations()...) //Add the operations of the circuit to add to the new circuit

	resCirc.InitClassicalRegister(uint8(len(c.classicalRegister))) //Init the classical register of the new circuit
	return resCirc
}
