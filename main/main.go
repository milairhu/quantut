package main

import (
	"fmt"

	"github.com/milairhu/quantut"
)

func main() {

	circuit := quantut.NewQuantumCircuit(4) //création circuit à 2 qubits
	circuit.InitClassicalRegister(2)        //création registre classique à 1 bit
	circuit.H(0)                            //porte de Hadamard sur le qubit 0
	circuit.CNOT(0, 2)                      //porte CNOT sur les qubits 0 et 1
	circuit.X(2)                            //porte X sur le qubit 2
	circuit.CCNOT(1, 2, 3)                  //porte CCNOT sur les qubits 0, 1 et 2

	circuit2 := quantut.NewQuantumCircuit(2) //création circuit à 2 qubits
	circuit2.InitClassicalRegister(2)        //création registre classique à 1 bit
	circuit2.SWAP(0, 1)
	circuit2.Measure(0, 0) //mesure du qubit 0 dans le registre classique 0
	circuit2.Measure(1, 1)

	circuit = circuit.Compose(circuit2) // composition des deux circuits

	simulator := quantut.NewSimulator(circuit, 1000) //création du simulateur à 1000 itérations
	res := simulator.Run()
	fmt.Println(res) //affichage des résultats

	circuit.ToQASM("exemple.qasm", "OPENQASM 2.0") //exportation du circuit en QASM

	//Test de l'affichage
	circuit.Display()

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
