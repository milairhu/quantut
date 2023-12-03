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

//Définition des portes (on ne peut pas utiliser const sur une structure composée)
var (
	// Portes à 1 qubit
	h = Gate{id: "H", nbQubits: 1, effect: [][]complex128{{complex(1/math.Sqrt(2), 0), complex(1/math.Sqrt(2), 0)}, {complex(1/math.Sqrt(2), 0), complex(-1/math.Sqrt(2), 0)}}}
	x = Gate{id: "X", nbQubits: 1, effect: [][]complex128{{complex(0, 0), complex(1, 0)}, {complex(1, 0), complex(0, 0)}}}
	y = Gate{id: "Y", nbQubits: 1, effect: [][]complex128{{complex(0, 0), complex(0, -1)}, {complex(0, 1), complex(0, 0)}}}
	z = Gate{id: "Z", nbQubits: 1, effect: [][]complex128{{complex(1, 0), complex(0, 0)}, {complex(0, 0), complex(-1, 0)}}}

	// Portes à 2 qubits
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

	// Portes à 3 qubits
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

	//Measure (cas particulier)
	measure = Gate{id: "MEASURE", nbQubits: 2} //Cas particulier : prend 1 qubit et 1 bit de registre classique

	//Initialisation (cas particulier)
	initialization = Gate{id: "INIT", nbQubits: 1}
)

//Méthods

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
