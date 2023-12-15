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
	if numQubits < 1 {
		panic("Number of qubits must be greater than 0")
	}
	o := make([]Operation, 0)
	qv := make([]Qubit, numQubits)
	//Init qubits values to 0
	for i := 0; i < numQubits; i++ {
		qv[i].Init(1, 0)
	}
	r := make([]int, 0)
	var nbComposante = 2
	for i := 1; i < numQubits; i++ {
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

func (c *QuantumCircuit) NbClassicalBits() int {
	return len(c.classicalRegister)
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
	//Update global state

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

// == Display circuit == //

func fillGapWithSpace(lengthToReach, initialLength int) string {
	var str string
	for i := initialLength; i < lengthToReach; i++ {
		str += " "
	}
	return str
}

func fillGapWithDash(lengthToReach, initialLength int) string {
	var str string
	for i := initialLength; i < lengthToReach; i++ {
		str += "-"
	}
	return str
}

func decideIfLinkNeeded(links []int, currQubit int) bool {
	//On retourne true si des | doivent descendre du qubit

	var existsLower bool
	var existsHigherOrEqual bool
	for _, qubit := range links {
		if qubit > currQubit {
			existsLower = true
		} else {
			existsHigherOrEqual = true
		}
		if existsLower && existsHigherOrEqual {
			return true
		}
	}
	return false
}

func (c *QuantumCircuit) Draw() {
	//We need to spot the largest gate name
	var maxLength int
	for _, op := range c.operations {
		if len(op.Gate().id) > maxLength {
			maxLength = len(op.Gate().id)
		}
	}

	//We save the indexes of the operations that necessitate links
	links := make(map[int][]int)
	for i, op := range c.operations {
		if op.Gate().Id() != "MEASURE" && len(op.Qubits()) > 1 {
			links[i] = op.Qubits()
		}
	}

	mat1 := make([][]string, len(c.operations))

	for i := 0; i < len(c.operations); i++ {
		op := c.operations[i]
		mat1[i] = make([]string, c.numQubits)
		for ind, qubit := range op.Qubits() {
			if ind < int(op.Gate().nbControlQubit) {
				mat1[i][qubit] = "@"
			} else {
				mat1[i][qubit] = op.Gate().id
			}
		}
	}

	//We display the transposed matrix
	matRes := make([][]string, len(mat1[0]))
	for i := 0; i < len(mat1[0]); i++ {
		matRes[i] = make([]string, len(mat1))
		for j := 0; j < len(mat1); j++ {
			matRes[i][j] = mat1[j][i]
		}
	}
	const nbSpaceBetweenLines = 3
	//We display the matrix
	var str string
	var lineLength int = 5 + maxLength*len(matRes[0])
	for indQubit := 0; indQubit < len(matRes); indQubit++ {

		//First, display the qubit number
		str += fmt.Sprintf("q%d | ", indQubit)

		//Display the operations applied on the qubit
		for indOp, op := range matRes[indQubit] {
			if op == "" {
				str += fillGapWithDash(maxLength, 0)
			} else {
				str += matRes[indQubit][indOp]
				str += fillGapWithDash(maxLength, len(matRes[indQubit][indOp]))
			}

		}
		if indQubit != len(matRes)-1 {
			for i := 0; i < nbSpaceBetweenLines; i++ {
				str += "\n   | "
				var numOp int
				for indCol := 5; indCol < lineLength; indCol += maxLength {
					tab, ok := links[numOp]
					if ok && decideIfLinkNeeded(tab, indQubit) {
						str += "|" + fillGapWithSpace(maxLength, 1)
					} else {
						str += fillGapWithSpace(maxLength, 0)
					}
					numOp++
				}

			}
		}
		str += "\n"

	}

	fmt.Println(str)

}

// ===== Display general state =====
func (c *QuantumCircuit) DisplayGlobalState() {
	var str string
	for i := 0; i < len(c.globalState); i++ {
		if c.globalState[i] != 0 {
			str += fmt.Sprintf("%f|%s> + ", c.globalState[i], convertIndToBinary(i, c.numQubits))
		}
	}
	fmt.Println(str[:len(str)-3])
}
