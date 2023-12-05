package main

import (
	"fmt"

	"github.com/milairhu/quantut"
)

func main() {

	circuit := quantut.NewQuantumCircuit(3) //création circuit à 2 qubits
	circuit.InitClassicalRegister(3)        //on crée un registre classique de 2 bits
	//Remarque : InitClassicalRegister et SetQubit sont appliqués directement, pas au lancement de la simulation
	circuit.H(0) //on applique une porte Hadamard sur le qubit 0
	circuit.CNOT(1, 2)
	circuit.CCNOT(0, 1, 2)
	circuit.CSWAP(2, 0, 1)
	circuit.Measure(2, 2)

	circToCombine := quantut.NewQuantumCircuit(2) //on crée un circuit à 1 qubit
	circToCombine.InitClassicalRegister(2)        //on crée un registre classique de 1 bit
	circToCombine.CNOT(0, 1)                      //on applique une porte CNOT sur le qubit 0
	circToCombine.H(1)
	circToCombine.X(1)
	circToCombine.Y(1)
	circToCombine.Z(1)
	circToCombine.Measure(1, 1) //on mesure le qubit 0 et on stocke le résultat dans le registre classique 0

	circuit = circuit.Compose(circToCombine)    //on combine les deux circuits
	circuit.Measure(0, 0)                       //on mesure le qubit 0 et on stocke le résultat dans le registre classique 0
	circuit.Measure(1, 1)                       //on mesure le qubit 1 et on stocke le résultat dans le registre classique 1
	circuit.Measure(2, 2)                       //on mesure le qubit 2 et on stocke le résultat dans le registre classique 2
	circuit.ToQASM("test.qasm", "OPENQASM 2.0") //on écrit le circuit en QASM dans un fichier

	simulator := quantut.NewSimulator(circuit, 1000)
	res := simulator.Run()
	fmt.Println(res)

	//Test CNOT et CSWAP
	/*
		circuit2 := quantut.NewQuantumCircuit(3)
		circuit2.InitClassicalRegister(1)
		circuit2.H(0)
		circuit2.H(0)
		circuit2.CNOT(1, 2)
		circuit2.H(0)
		circuit2.H(0)
		circuit2.CSWAP(0, 1, 2)
		simulator2 := quantut.NewSimulator(circuit2, 1)
		res2 := simulator2.Run()
		fmt.Println(res2)
	*/

	/*
		// Test de la fonction tensorialProduct
		circ2 := quantut.NewQuantumCircuit(3)
		circ2.SetQubit(0, complex(1/math.Sqrt(2), 0), complex(1/math.Sqrt(2), 0))
		circ2.SetQubit(2, 0, 1)
		test := utils.TensorialProduct(circ2.Qubits())
		fmt.Println(test)
	*/

}
