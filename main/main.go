package main

import (
	"fmt"

	"github.com/milairhu/quantut"
)

func main() {

	circuit := quantut.NewQuantumCircuit(4) //création circuit à 2 qubits
	circuit := quantut.NewQuantumCircuit(4) //création circuit à 2 qubits
	circuit.InitClassicalRegister(2)        //création registre classique à 1 bit
	circuit.H(0)                            //porte de Hadamard sur le qubit 0
	circuit.CNOT(0, 1)                      //porte CNOT sur les qubits 0 et 1
	circuit.Measure(0, 0)                   //mesure du qubit 0 dans le registre classique 0

	circuit.Measure(1, 1) //mesure du qubit 1 dans le registre classique 1

	sim := quantut.NewSimulator(circuit, 1000) //création du simulateur
	res := sim.Run()                           //lancement de la simulation
	fmt.Println(res)                           //affichage des résultats

}
