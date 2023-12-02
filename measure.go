package quantut

import (
	"fmt"
	"math/cmplx"
	"math/rand"
	"time"
)

// =============== Measure ===============

// Measure() add the operation for measuring
func (c *QuantumCircuit) Measure(qubit int, register int) {
	if qubit >= c.numQubits || qubit < 0 {
		panic("Qubit number out of range")
	}
	if register >= len(c.classicalRegister) || register < 0 {
		panic("Register number out of range")
	}
	c.operations = append(c.operations, Operation{gate: measure, qubits: []int{qubit, register}})
}

// measureEffect() is the effect of the measure operation : store the result in the classical register
func measureEffect(c *QuantumCircuit, qubit int, register int) {
	if qubit >= c.numQubits || qubit < 0 {
		panic(fmt.Sprintf("Qubit number out of range : %d", qubit))
	}
	if register >= len(c.classicalRegister) || register < 0 {

		panic(fmt.Sprintf("Register number out of range : %d", register))
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Float64() // generate a random number between 0 and 1

	var resMeasure int
	if random <= cmplx.Abs(c.qubitsValues[qubit][0]) { // Compare to the probability of the qubit to be 0
		c.classicalRegister[register] = 0
		resMeasure = 0
	} else {
		c.classicalRegister[register] = 1
		resMeasure = 1
	}

	//TODO vérifier que la projection se passe comme ça
	c.qubitsValues[qubit].Init(complex(1-float64(resMeasure), 0), complex(0, float64(resMeasure)))
}
