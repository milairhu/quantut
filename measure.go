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
	var proba0 float64
	for i := 0; i < len(c.globalState); i++ {
		binary := convertIndToBinary(i, c.numQubits)
		if binary[qubit] == '0' {
			proba0 += cmplx.Abs(c.globalState[i]) * cmplx.Abs(c.globalState[i])
		}
	}
	if random <= proba0 { // Compare to the probability of the qubit to be 0
		c.classicalRegister[register] = 0
		resMeasure = 0
	} else {
		c.classicalRegister[register] = 1
		resMeasure = 1
	}

	newStateMap := make(map[string]complex128)
	for i := 0; i < len(c.globalState); i++ {
		newStateMap[convertIndToBinary(i, c.numQubits)] = 0
	}
	var r byte
	if resMeasure == 0 {
		r = '0'
	} else {
		r = '1'
	}
	for i := 0; i < len(c.globalState); i++ {
		binary := convertIndToBinary(i, c.numQubits)

		// on remplit le map avec les valeurs des nouveaux états
		if binary[qubit] == r {
			//Si le qubit mesuré correspond, on change rien
			newStateMap[binary] += c.globalState[i]
		}
	}
	//On a projetté les états sur le résultat de la mesure, il faut normaliser
	//On calcule la norme
	var norm complex128
	for _, val := range newStateMap {
		norm += val * val
	}
	norm = cmplx.Sqrt(norm)

	//On normalise
	for key, val := range newStateMap {
		newStateMap[key] = val / norm
	}

	//On remplace l'ancien état par le nouveau
	for i := 0; i < len(c.globalState); i++ {
		c.globalState[i] = newStateMap[convertIndToBinary(i, c.numQubits)]
	}
}
