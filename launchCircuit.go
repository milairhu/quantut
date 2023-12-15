package quantut

// Simulate a circuit
func (c *QuantumCircuit) LaunchCircuit() {

	/*var test []complex128 = []complex128{complex(1, 0), complex(0, 1), complex(0, 0), complex(0, 0)}
	fmt.Println(test)
	fmt.Println("X : ", xeffect(1, test, 2))
	fmt.Println("Y : ", yeffect(1, test, 2))
	fmt.Println("Z : ", zeffect(1, test, 2))
	fmt.Println("H : ", heffect(1, test, 2))

	fmt.Println("CNOT : ", cnoteffect(0, 1, test, 2))
	fmt.Println("CNOT : ", cnoteffect(1, 0, test, 2))

	fmt.Println("SWAP : ", swapeffect(0, 1, test, 2))
	fmt.Println("SWAP : ", swapeffect(1, 0, test, 2))

	test2 := []complex128{complex(1, 0), complex(0, 0), complex(0, 0), complex(1, 0),
		complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0)}

	fmt.Println("CCNOT : ", ccnoteffect(0, 1, 2, test2, 3))
	fmt.Println("CCNOT : ", ccnoteffect(1, 0, 2, test2, 3))
	fmt.Println("CCNOT : ", ccnoteffect(0, 2, 1, test2, 3))

	fmt.Println("CSWAP : ", cswapeffect(0, 1, 2, test2, 3))
	fmt.Println("CSWAP : ", cswapeffect(1, 0, 2, test2, 3))
	fmt.Println("CSWAP : ", cswapeffect(0, 2, 1, test2, 3))
	*/
	//Apply operations
	for _, op := range c.operations {
		if len(c.operations) == 0 {
			panic("Operation is applied on no qubit")
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
			switch op.Gate().Id() {
			case "X":
				c.globalState = xeffect(op.Qubits()[0], c.globalState, c.NumQubits())
			case "Y":
				c.globalState = yeffect(op.Qubits()[0], c.globalState, c.NumQubits())
			case "Z":
				c.globalState = zeffect(op.Qubits()[0], c.globalState, c.NumQubits())
			case "H":
				c.globalState = heffect(op.Qubits()[0], c.globalState, c.NumQubits())
			case "CNOT":
				c.globalState = cnoteffect(op.Qubits()[0], op.Qubits()[1], c.globalState, c.NumQubits())
			case "SWAP":
				c.globalState = swapeffect(op.Qubits()[0], op.Qubits()[1], c.globalState, c.NumQubits())
			case "CCNOT":
				c.globalState = ccnoteffect(op.Qubits()[0], op.Qubits()[1], op.Qubits()[2], c.globalState, c.NumQubits())
			case "CSWAP":
				c.globalState = cswapeffect(op.Qubits()[0], op.Qubits()[1], op.Qubits()[2], c.globalState, c.NumQubits())

			default:
				panic("Unknown gate" + op.Gate().Id())
			}
		}
		//fmt.Print("Global state after ", op.Gate().Id(), " : ")
		//c.DisplayGlobalState()
	}
}
