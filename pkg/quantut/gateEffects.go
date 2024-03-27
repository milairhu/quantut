package quantut

import (
	"fmt"
	"math"
)

func convertIndToBinary(ind int, nbBits int) string {
	var res string = ""
	for i := 0; i < nbBits; i++ {
		res = fmt.Sprintf("%d", ind%2) + res
		ind = ind / 2
	}
	return res
}

// =========== 1 QUBIT GATES ===========
// All 4 look good
// Gate X
func xeffect(targetQubit int, generalState []complex128, nbQubit int) []complex128 {
	if targetQubit >= nbQubit {
		panic("Target qubit is not in the circuit")
	}
	stateMap := make(map[string]complex128)
	res := make([]complex128, len(generalState))
	for i := 0; i < len(generalState); i++ {
		stateMap[convertIndToBinary(i, nbQubit)] = 0
	}
	for i := 0; i < len(generalState); i++ {
		binary := convertIndToBinary(i, nbQubit)
		//Gate effect
		if binary[targetQubit] == '0' {
			binary = binary[:targetQubit] + "1" + binary[targetQubit+1:]
		} else {
			binary = binary[:targetQubit] + "0" + binary[targetQubit+1:]
		}
		stateMap[binary] += generalState[i]
	}
	//the map is set, we convert it to an array
	for i := 0; i < len(generalState); i++ {
		res[i] = stateMap[convertIndToBinary(i, nbQubit)]
	}

	return res
}

// Gate Y
func yeffect(targetQubit int, generalState []complex128, nbQubit int) []complex128 {
	if targetQubit >= nbQubit {
		panic("Target qubit is not in the circuit")
	}
	stateMap := make(map[string]complex128)
	res := make([]complex128, len(generalState))
	for i := 0; i < len(generalState); i++ {
		stateMap[convertIndToBinary(i, nbQubit)] = 0
	}
	for i := 0; i < len(generalState); i++ {
		binary := convertIndToBinary(i, nbQubit)
		//Gate effect
		if binary[targetQubit] == '0' {
			binary = binary[:targetQubit] + "1" + binary[targetQubit+1:]
			stateMap[binary] += generalState[i] * complex(0, 1)
		} else {
			binary = binary[:targetQubit] + "0" + binary[targetQubit+1:]
			stateMap[binary] += generalState[i] * complex(0, -1)
		}

	}
	//the map is set, we convert it to an array
	for i := 0; i < len(generalState); i++ {
		res[i] = stateMap[convertIndToBinary(i, nbQubit)]
	}

	return res
}

// Gate Z
func zeffect(targetQubit int, generalState []complex128, nbQubit int) []complex128 {
	if targetQubit >= nbQubit {
		panic("Target qubit is not in the circuit")
	}
	stateMap := make(map[string]complex128)
	res := make([]complex128, len(generalState))
	for i := 0; i < len(generalState); i++ {
		stateMap[convertIndToBinary(i, nbQubit)] = 0
	}
	for i := 0; i < len(generalState); i++ {
		binary := convertIndToBinary(i, nbQubit)
		//Gate effect
		if binary[targetQubit] == '0' {
			stateMap[binary] += generalState[i]
		} else {
			stateMap[binary] += generalState[i] * -1
		}

	}
	//the map is set, we convert it to an array
	for i := 0; i < len(generalState); i++ {
		res[i] = stateMap[convertIndToBinary(i, nbQubit)]
	}

	return res
}

// Gate H
func heffect(targetQubit int, generalState []complex128, nbQubit int) []complex128 {
	if targetQubit >= nbQubit {
		panic("Target qubit is not in the circuit")
	}
	stateMap := make(map[string]complex128)
	res := make([]complex128, len(generalState))
	for i := 0; i < len(generalState); i++ {
		stateMap[convertIndToBinary(i, nbQubit)] = 0
	}
	for i := 0; i < len(generalState); i++ {
		binary := convertIndToBinary(i, nbQubit)
		//Gate effect
		if binary[targetQubit] == '0' {
			binary1 := binary[:targetQubit] + "1" + binary[targetQubit+1:]
			stateMap[binary] += generalState[i] * complex(1/math.Sqrt(2), 0)
			stateMap[binary1] += generalState[i] * complex(1/math.Sqrt(2), 0)

		} else {
			binary0 := binary[:targetQubit] + "0" + binary[targetQubit+1:]
			stateMap[binary] -= generalState[i] * complex(1/math.Sqrt(2), 0)
			stateMap[binary0] += generalState[i] * complex(1/math.Sqrt(2), 0)
		}
	}
	//the map is set, we convert it to an array
	for i := 0; i < len(generalState); i++ {
		res[i] = stateMap[convertIndToBinary(i, nbQubit)]
	}

	return res
}

// =========== 2 QUBIT GATES ===========

// Gate CNOT
func cnoteffect(controlQubit int, targetQubit int, generalState []complex128, nbQubit int) []complex128 {
	if targetQubit >= nbQubit {
		panic("Target qubit is not in the circuit for cnot")
	}
	if controlQubit >= nbQubit {
		panic("Control qubit is not in the circuit for cnot")
	}
	if controlQubit == targetQubit {
		panic("Control qubit and target qubit must be different for CNOT")
	}

	stateMap := make(map[string]complex128)
	res := make([]complex128, len(generalState))
	for i := 0; i < len(generalState); i++ {
		stateMap[convertIndToBinary(i, nbQubit)] = 0
	}
	for i := 0; i < len(generalState); i++ {
		binary := convertIndToBinary(i, nbQubit)
		//Gate effect
		if binary[controlQubit] == '0' {
			// don't' modify the state
			stateMap[binary] += generalState[i]
		} else {
			// if control qubit is 1, apply X gate on target qubit
			if binary[targetQubit] == '0' {
				binary = binary[:targetQubit] + "1" + binary[targetQubit+1:]
			} else {
				binary = binary[:targetQubit] + "0" + binary[targetQubit+1:]
			}
			stateMap[binary] += generalState[i]
		}

	}
	//the map is set, we convert it to an array
	for i := 0; i < len(generalState); i++ {
		res[i] = stateMap[convertIndToBinary(i, nbQubit)]
	}

	return res
}

// Gate SWAP
func swapeffect(targetQubit1 int, targetQubit2 int, generalState []complex128, nbQubit int) []complex128 {
	if targetQubit1 >= nbQubit {
		panic("Target1 qubit is not in the circuit for swap")
	}
	if targetQubit2 >= nbQubit {
		panic("Target2 qubit is not in the circuit for swap")
	}
	if targetQubit1 == targetQubit2 {
		panic("Target1 qubit and target2 qubit must be different for SWAP")
	}
	stateMap := make(map[string]complex128)
	res := make([]complex128, len(generalState))
	for i := 0; i < len(generalState); i++ {
		stateMap[convertIndToBinary(i, nbQubit)] = 0
	}
	for i := 0; i < len(generalState); i++ {
		binary := convertIndToBinary(i, nbQubit)
		//Gate effect
		var newValTarget1, newValTarget2 string
		if binary[targetQubit1] == '0' {
			newValTarget2 = "0"
		} else {
			newValTarget2 = "1"
		}
		if binary[targetQubit2] == '0' {
			newValTarget1 = "0"
		} else {
			newValTarget1 = "1"
		}
		if targetQubit1 < targetQubit2 {
			binary = binary[:targetQubit1] + newValTarget1 + binary[targetQubit1+1:targetQubit2] + newValTarget2 + binary[targetQubit2+1:]
		} else {
			binary = binary[:targetQubit2] + newValTarget2 + binary[targetQubit2+1:targetQubit1] + newValTarget1 + binary[targetQubit1+1:]
		}
		stateMap[binary] += generalState[i]
	}
	//the map is set, we convert it to an array
	for i := 0; i < len(generalState); i++ {
		res[i] = stateMap[convertIndToBinary(i, nbQubit)]
	}

	return res
}

// =========== 3 QUBIT GATES ===========

// Gate CCNOT
func ccnoteffect(controlQubit1 int, controlQubit2 int, targetQubit int, generalState []complex128, nbQubit int) []complex128 {
	if targetQubit >= nbQubit {
		panic("Target qubit is not in the circuit for ccnot")
	}
	if controlQubit1 >= nbQubit {
		panic("Control1 qubit is not in the circuit for ccnot")
	}
	if controlQubit2 >= nbQubit {
		panic("Control2 qubit is not in the circuit for ccnot")
	}
	if controlQubit1 == targetQubit || controlQubit2 == targetQubit || controlQubit1 == controlQubit2 {
		panic("Control qubits and target qubit must be different for CCNOT")
	}

	stateMap := make(map[string]complex128)
	res := make([]complex128, len(generalState))
	for i := 0; i < len(generalState); i++ {
		stateMap[convertIndToBinary(i, nbQubit)] = 0
	}
	for i := 0; i < len(generalState); i++ {
		binary := convertIndToBinary(i, nbQubit)
		//Gate effect
		if binary[controlQubit1] == '0' || binary[controlQubit2] == '0' {
			// dont modify the state
			stateMap[binary] += generalState[i]
		} else {
			// if control qubits are 1, apply X gate on target qubit
			if binary[targetQubit] == '0' {
				binary = binary[:targetQubit] + "1" + binary[targetQubit+1:]
			} else {
				binary = binary[:targetQubit] + "0" + binary[targetQubit+1:]
			}
			stateMap[binary] += generalState[i]
		}

	}
	//the map is set, we convert it to an array
	for i := 0; i < len(generalState); i++ {
		res[i] = stateMap[convertIndToBinary(i, nbQubit)]
	}

	return res
}

// Gate CSWAP
func cswapeffect(controlQubit int, targetQubit1 int, targetQubit2 int, generalState []complex128, nbQubit int) []complex128 {
	if targetQubit1 >= nbQubit {
		panic("Target1 qubit is not in the circuit for cswap")
	}
	if targetQubit2 >= nbQubit {
		panic("Target2 qubit is not in the circuit for cswap")
	}
	if controlQubit >= nbQubit {
		panic("Control qubit is not in the circuit for cswap")
	}
	if targetQubit1 == targetQubit2 || targetQubit1 == controlQubit || targetQubit2 == controlQubit {
		panic("Target qubits and control qubit must be different for CSWAP")
	}
	stateMap := make(map[string]complex128)
	res := make([]complex128, len(generalState))
	for i := 0; i < len(generalState); i++ {
		stateMap[convertIndToBinary(i, nbQubit)] = 0
	}
	for i := 0; i < len(generalState); i++ {
		binary := convertIndToBinary(i, nbQubit)
		//Gate effect
		if binary[controlQubit] == '0' {
			// don't modify the state
			stateMap[binary] += generalState[i]
		} else {
			var newValTarget1, newValTarget2 string
			if binary[targetQubit1] == '0' {
				newValTarget2 = "0"
			} else {
				newValTarget2 = "1"
			}
			if binary[targetQubit2] == '0' {
				newValTarget1 = "0"
			} else {
				newValTarget1 = "1"
			}
			if targetQubit1 < targetQubit2 {
				binary = binary[:targetQubit1] + newValTarget1 + binary[targetQubit1+1:targetQubit2] + newValTarget2 + binary[targetQubit2+1:]
			} else {
				binary = binary[:targetQubit2] + newValTarget2 + binary[targetQubit2+1:targetQubit1] + newValTarget1 + binary[targetQubit1+1:]
			}
			stateMap[binary] += generalState[i]
		}
	}
	//the map is set, we convert it to an array
	for i := 0; i < len(generalState); i++ {
		res[i] = stateMap[convertIndToBinary(i, nbQubit)]
	}

	return res
}
