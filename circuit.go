package quantut

import (
	"fmt"
	"math/cmplx"
	"math/rand"
	"quantut/utils"
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

// =============== Measure ===============
func (c *QuantumCircuit) Measure(qubit int, register int) {
	if qubit >= c.numQubits || qubit < 0 {
		panic("Qubit number out of range")
	}
	if register >= len(c.classicalRegister) || register < 0 {
		panic("Register number out of range")
	}
	c.operations = append(c.operations, Operation{gate: measure, qubits: []int{qubit, register}})
}

func (c *QuantumCircuit) MeasureEffect(qubit int, register int) {
	if qubit >= c.numQubits || qubit < 0 {
		panic("Qubit number out of range")
	}
	if register >= len(c.classicalRegister) || register < 0 {
		panic("Register number out of range")
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
	var resMeasure int
	if random <= cmplx.Abs(c.qubitsValues[qubit][0]) {
		c.classicalRegister[register] = 0
		resMeasure = 0
	} else {
		c.classicalRegister[register] = 1
		resMeasure = 1
	}

	//TODO vérifier que la projection se passe comme ça
	c.qubitsValues[qubit].Init(complex(1-float64(resMeasure), 0), complex(0, float64(resMeasure)))
}

// Launch operations on a qubit
func (c *QuantumCircuit) LaunchQubit(numQubit int, channels []chan Qubit, resChan chan string) {
	//Create map for easier use of channels
	qubitChanMap := make(map[int]chan Qubit, len(channels))
	for i := 0; i < len(channels); i++ {
		if i < numQubit {
			qubitChanMap[i] = channels[i]
		} else {
			qubitChanMap[i+1] = channels[i]
		}
	}

	//Apply operations
	for _, op := range c.operations {
		if len(c.operations) == 0 {
			panic("Operartion is applied on no qubit")
		}

		// Check if the qubit is involved in the operation
		var qubitInOperation bool = op.IsQubitInOperation(numQubit)
		if !qubitInOperation {
			//If not, next operation
			continue
		}

		// Particular cases : measure and initialization
		switch op.Gate().Id() {

		case "MEASURE":
			if len(op.Qubits()) != 2 {
				panic("Measure operation must have 1 qubit and 1 classical register")
			}
			//Measure the qubit in the classical register
			c.MeasureEffect(op.Qubits()[0], op.Qubits()[1])

		case "INIT":
			if len(op.Qubits()) != 1 {
				panic("Initialization operation must have 1 qubit")
			}
			if len(op.Options()) != 2 {
				panic("Initialization operation must have 2 values for the qubit")
			}
			//Initialize the qubit
			c.SetQubit(op.Qubits()[0], op.Options()[0], op.Options()[1])

		default:
			//For any other operation than measure and initialization
			var nbQubits int = len(op.Qubits())
			if nbQubits == 1 {
				//If only one qubit is involved, apply the gate
				qubitToVector := c.qubitsValues[numQubit].Vector()
				var calc = utils.ComplexMatrixMultiply(op.Gate().Effect(), qubitToVector)
				/*
					fmt.Println("Op : ", op.Gate().Effect())
					fmt.Println("Value init : ", c.qubitsValues[numQubit])
					fmt.Println("calc : ", calc)
				*/
				c.SetQubit(numQubit, calc[0][0], calc[1][0])

			}

		}
	}
	//Test synchronisation
	res := 0
	if numQubit%2 == 0 {
		res = 1
	}
	resChan <- fmt.Sprintf("%d,%d", numQubit, res)
}
