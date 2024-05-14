package quantut

import "math"

/**
* We differenciate Gate and Operations the same way than Cirq
**/
type Gate struct {
	id             string         //Type of the gate
	nbQubits       uint16         //Number of qubits involved
	nbControlQubit uint16         //Number of control qubits
	effect         [][]complex128 //Effect of the gate on the qubits
}

// Definitio of the gates (we can't use 'const' on a composed structure)
var (
	// 1 qubit gates
	h = Gate{id: "H", nbQubits: 1,
		effect: [][]complex128{{complex(1/math.Sqrt(2), 0), complex(1/math.Sqrt(2), 0)}, {complex(1/math.Sqrt(2), 0), complex(-1/math.Sqrt(2), 0)}}}
	x = Gate{id: "X", nbQubits: 1, effect: [][]complex128{{complex(0, 0), complex(1, 0)}, {complex(1, 0), complex(0, 0)}}}
	y = Gate{id: "Y", nbQubits: 1, effect: [][]complex128{{complex(0, 0), complex(0, -1)}, {complex(0, 1), complex(0, 0)}}}
	z = Gate{id: "Z", nbQubits: 1, effect: [][]complex128{{complex(1, 0), complex(0, 0)}, {complex(0, 0), complex(-1, 0)}}}
	i = Gate{id: "I", nbQubits: 1, effect: [][]complex128{{complex(1, 0), complex(0, 0)}, {complex(0, 0), complex(1, 0)}}}
	s = Gate{id: "S", nbQubits: 1, effect: [][]complex128{{complex(1, 0), complex(0, 0)}, {complex(0, 0), complex(0, 1)}}}
	t = Gate{id: "T", nbQubits: 1, effect: [][]complex128{{complex(1, 0), complex(0, 0)}, {complex(0, 0), complex(1/math.Sqrt(2), 1/math.Sqrt(2))}}} //TODO à tester

	// 2 qubits gates
	cnot = Gate{id: "CNOT",
		nbQubits:       2,
		nbControlQubit: 1,
		effect: [][]complex128{
			{complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0)},
			{complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0)},
		},
	}
	swap = Gate{id: "SWAP",
		nbQubits:       2,
		nbControlQubit: 0,
		effect: [][]complex128{
			{complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0)},
			{complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0)},
		},
	}

	cs = Gate{id: "CS",
		nbQubits:       2,
		nbControlQubit: 1,
		effect: [][]complex128{
			{complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 1)},
		},
	}

	// 3 qubits gates
	ccnot = Gate{id: "CCNOT",
		nbQubits:       3,
		nbControlQubit: 2,
		effect: [][]complex128{
			{complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0)},
		},
	}
	cswap = Gate{id: "CSWAP", nbQubits: 3,
		nbControlQubit: 1,
		effect: [][]complex128{
			{complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0), complex(0, 0), complex(0, 0)},
			{complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(0, 0), complex(1, 0)},
		},
	}

	// Edge case of MEASURE
	measure = Gate{id: "MEASURE", nbQubits: 2} //consume 1 qubit and 1 bit from the classical register

)

//Methods

//Getters
func (g Gate) Id() string {
	return g.id
}

func (g Gate) NbQubits() uint16 {
	return g.nbQubits
}

func (g Gate) Effect() [][]complex128 {
	return g.effect
}

func (g Gate) NbControlQubit() uint16 {
	return g.nbControlQubit
}
