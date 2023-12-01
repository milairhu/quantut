package quantut

import (
	"errors"
	"fmt"
	"math/cmplx"
	"math/rand"
	"time"
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
func (c *QuantumCircuit) SetQubit(numQubit int, comp1 complex128, comp2 complex128) error {
	if numQubit >= c.numQubits || numQubit < 0 {
		return errors.New("Qubit number out of range")
	}
	if comp1*comp1+comp2*comp2 != 1 {
		return errors.New("Qubit value must be normalized")
	}

	c.qubitsValues[numQubit].Init(comp1, comp2)
	return nil
}

// Définie le nombre de registres classics
func (c *QuantumCircuit) SetClassicalRegister(numRegister int) {
	c.classicalRegister = make([]int, numRegister)
}

// =============== Add gates to the circuit ===============
// Hadmard
func (c *QuantumCircuit) H(numQubit int) error {
	if numQubit >= c.numQubits || numQubit < 0 {
		return errors.New("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: h, qubits: []int{numQubit}})
	return nil
}

// Pauli X (NOT)
func (c *QuantumCircuit) X(numQubit int) error {
	if numQubit >= c.numQubits || numQubit < 0 {
		return errors.New("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: x, qubits: []int{numQubit}})
	return nil
}

// Pauli Y
func (c *QuantumCircuit) Y(numQubit int) error {
	if numQubit >= c.numQubits || numQubit < 0 {
		return errors.New("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: y, qubits: []int{numQubit}})
	return nil
}

// Pauli Z
func (c *QuantumCircuit) Z(numQubit int) error {
	if numQubit >= c.numQubits || numQubit < 0 {
		return errors.New("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: z, qubits: []int{numQubit}})
	return nil
}

// CNOT
func (c *QuantumCircuit) CNOT(control int, target int) error {
	if control >= c.numQubits || control < 0 || target >= c.numQubits || target < 0 {
		return errors.New("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: cnot, qubits: []int{control, target}})
	return nil
}

// SWAP
func (c *QuantumCircuit) SWAP(qubit1 int, qubit2 int) error {
	if qubit1 >= c.numQubits || qubit1 < 0 || qubit2 >= c.numQubits || qubit2 < 0 {
		return errors.New("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: swap, qubits: []int{qubit1, qubit2}})
	return nil
}

// CCNOT
func (c *QuantumCircuit) CCNOT(control1 int, control2 int, target int) error {
	if control1 >= c.numQubits || control1 < 0 || control2 >= c.numQubits || control2 < 0 || target >= c.numQubits || target < 0 {
		return errors.New("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: ccnot, qubits: []int{control1, control2, target}})
	return nil
}

// CSWAP
func (c *QuantumCircuit) CSWAP(control int, target1 int, target2 int) error {
	if control >= c.numQubits || control < 0 || target1 >= c.numQubits || target1 < 0 || target2 >= c.numQubits || target2 < 0 {
		return errors.New("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: cswap, qubits: []int{control, target1, target2}})
	return nil
}

// TOFFOLI
func (c *QuantumCircuit) TOFFOLI(control1 int, control2 int, target int) error {
	if control1 >= c.numQubits || control1 < 0 || control2 >= c.numQubits || control2 < 0 || target >= c.numQubits || target < 0 {
		return errors.New("Qubit number out of range")
	}
	c.operations = append(c.operations, Operation{gate: toffoli, qubits: []int{control1, control2, target}})
	return nil
}

// =============== Measure ===============
func (c *QuantumCircuit) Measure(qubit int, register int) error {
	if qubit >= c.numQubits || qubit < 0 {
		return errors.New("qubit number out of range")
	}
	if register >= len(c.classicalRegister) || register < 0 {
		return errors.New("register number out of range")
	}
	// The result of he measure is 0 or 1 depending on the value of the qubit.
	// The result is stored in the classical register

	rand.Seed(time.Now().UnixNano())
	// We generate a random number between 0 and 1
	random := rand.Float64()
	// We compare it to the probability of the qubit to be 0
	fmt.Println("random : ", random)
	fmt.Println("qubit : ", c.qubitsValues[qubit])
	fmt.Println("module : ", cmplx.Abs(c.qubitsValues[qubit][0]))
	if random <= cmplx.Abs(c.qubitsValues[qubit][0]) {
		c.classicalRegister[register] = 0
	} else {
		c.classicalRegister[register] = 1
	}

	//TODO : modifier le qubit mesuré pour qu'il soit projeté sur l'état mesuré

	return nil
}
