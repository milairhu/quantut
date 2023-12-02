package quantut

import (
	"fmt"
	"quantut/utils"
)

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
			measureEffect(c, op.Qubits()[0], op.Qubits()[1])

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
				//fmt.Println("Value before ", op.Gate().Id(), " : ", c.qubitsValues[numQubit])
				c.SetQubit(numQubit, calc[0][0], calc[1][0])
				//fmt.Println("Value after ", op.gate.id, " : ", c.qubitsValues[numQubit])
			} else {

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
