package quantut

import (
	"fmt"
	"quantut/utils"
)

// Launch operations on a qubit
func (c *QuantumCircuit) LaunchQubit(numQubit int, channels []chan Qubit) {
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
			panic("Operation is applied on no qubit")
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
				var calc = utils.ComplexMatrixProduct(op.Gate().Effect(), qubitToVector)
				c.SetQubit(numQubit, calc[0][0], calc[1][0])
			} else {
				// If more than one qubit is involved, qubit must be synchronized with the others
				// Get the other qubits involved in the operation
				othersQubits := make([]int, nbQubits-1)
				for i := 0; i < nbQubits; i++ {
					if op.Qubits()[i] != numQubit {
						othersQubits[i] = op.Qubits()[i]
					}
				}
				//Convention : control qubits are given first in the operation
				//Case 1 : the qubit is a control qubit
				if op.IsControlQubit(numQubit) {
					//Qubit sends its value to the target qubit(s) //TODO : checker CSWAP, 2 cibles
					for _, q := range othersQubits {
						qubitChanMap[q] <- c.qubitsValues[numQubit]
					}
				} else { //Case 2 : the qubit is a target qubit
					//Qubit waits for the control qubit(s) to send its value
					receivedValues := make([]Qubit, op.gate.nbControlQubit)
					for i, q := range othersQubits {
						receivedValues[i] = <-qubitChanMap[q]
					}

					//Apply the gate
					//TODO
					fmt.Println("Qubit", numQubit, "received values", receivedValues)
				}
			}

		}
	}
}
