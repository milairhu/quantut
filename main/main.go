package main

import (
	"fmt"
	"quantut"
)

func main() {

	circuit := quantut.NewQuantumCircuit(2) //création circuit à 2 qubits
	circuit.SetClassicalRegister(2)         //on crée un registre classique de 1 bit
	circuit.H(0)                            //ajout porte H sur qubit 0
	circuit.CNOT(0, 1)                      //ajout porte CNOT sur qubit 0 et 1
	circuit.Measure(0, 0)                   //mesure du qubit 0 dans le registre 0
	circuit.Measure(1, 1)                   //mesure du qubit 1 dans le registre 1
	fmt.Println(circuit.ClassicalRegister())

}
