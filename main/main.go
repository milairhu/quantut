package main

import (
	"fmt"
	"quantut"
)

func main() {

	circuit := quantut.NewQuantumCircuit(4) //création circuit à 2 qubits
	circuit.InitClassicalRegister(3)        //on crée un registre classique de 2 bits
	//Remarque : InitClassicalRegister et SetQubit sont appliqués directement, pas au lancement de la simulation
	circuit.H(0) //on applique une porte Hadamard sur le qubit 0
	circuit.H(1) //on applique une porte Hadamard sur le qubit 1
	circuit.H(2)
	circuit.H(3)

	circuit.Measure(0, 0) //on mesure le qubit 0 et on stocke le résultat dans le registre classique 0
	circuit.Measure(1, 1) //on mesure le qubit 1 et on stocke le résultat dans le registre classique 1
	circuit.Measure(2, 2)

	simulator := quantut.NewSimulator(circuit, 1000)
	res := simulator.Run()
	fmt.Println(res)

}
