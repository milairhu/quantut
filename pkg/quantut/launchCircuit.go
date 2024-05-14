package quantut

// Simulate a circuit
func (c *QuantumCircuit) LaunchCircuit() {
	//Apply operations
	for _, op := range c.operations {
		if len(c.operations) == 0 {
			panic("Operation is applied on no qubit")
		}

		// Particular cases : measure
		switch op.Gate().Id() {
		case "MEASURE":
			if len(op.Qubits()) != 2 {
				panic("Measure operation must have 1 qubit and 1 classical register")
			}
			//Measure the qubit in the classical register
			measureEffect(c, op.Qubits()[0], op.Qubits()[1])

		default:
			//For any other operation than measure
			switch op.Gate().Id() {
			case "X":
				c.globalState = xeffect(op.Qubits()[0], c.globalState, c.NumQubits())
			case "Y":
				c.globalState = yeffect(op.Qubits()[0], c.globalState, c.NumQubits())
			case "Z":
				c.globalState = zeffect(op.Qubits()[0], c.globalState, c.NumQubits())
			case "H":
				c.globalState = heffect(op.Qubits()[0], c.globalState, c.NumQubits())
			case "I":
				c.globalState = ieffect(op.Qubits()[0], c.globalState, c.NumQubits())
			case "S":
				c.globalState = seffect(op.Qubits()[0], c.globalState, c.NumQubits())
			case "T":
				c.globalState = teffect(op.Qubits()[0], c.globalState, c.NumQubits())
			case "CNOT":
				c.globalState = cnoteffect(op.Qubits()[0], op.Qubits()[1], c.globalState, c.NumQubits())
			case "SWAP":
				c.globalState = swapeffect(op.Qubits()[0], op.Qubits()[1], c.globalState, c.NumQubits())
			case "CS":
				c.globalState = cseffect(op.Qubits()[0], op.Qubits()[1], c.globalState, c.NumQubits())
			case "CCNOT":
				c.globalState = ccnoteffect(op.Qubits()[0], op.Qubits()[1], op.Qubits()[2], c.globalState, c.NumQubits())
			case "CSWAP":
				c.globalState = cswapeffect(op.Qubits()[0], op.Qubits()[1], op.Qubits()[2], c.globalState, c.NumQubits())

			default:
				panic("Unknown gate" + op.Gate().Id())
			}
		}
	}
}
