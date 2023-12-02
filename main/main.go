package main

import (
	"fmt"
	"quantut"
)

func main() {

	circuit := quantut.NewQuantumCircuit(4) //création circuit à 2 qubits
	circuit.SetQubit(1, 0, 1)               //on initialise le qubit 1 à 1
	circuit.InitClassicalRegister(2)        //on crée un registre classique de 2 bits
	circuit.H(0)                            //ajout porte H sur qubit 0
	circuit.CNOT(0, 1)                      //ajout porte CNOT sur qubit 0 et 1
	circuit.Measure(0, 0)                   //mesure du qubit 0 dans le registre 0
	circuit.Measure(1, 1)                   //mesure du qubit 1 dans le registre 1
	fmt.Println(circuit.ClassicalRegister())

	simulator := quantut.NewSimulator(circuit, 10)
	res := simulator.Run()
	fmt.Println(res)

}
