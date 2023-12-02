package main

import (
	"fmt"
	"quantut"
)

func main() {

	circuit := quantut.NewQuantumCircuit(4) //création circuit à 2 qubits
	circuit.SetQubit(1, 0, 1)               //on initialise le qubit 1 à 1 immédiatement, avant la simulation
	circuit.InitClassicalRegister(2)        //on crée un registre classique de 2 bits
	//Remarque : InitClassicalRegister et SetQubit sont appliqués directement, pas au lancement de la simulation

	circuit.H(0)          //ajout porte H sur qubit 0
	circuit.H(0)          //retour à l'état initial
	circuit.H(0)          // obtent |+>
	circuit.Measure(0, 0) //on mesure le qubit 0 et on stocke le résultat dans le registre classique 0

	simulator := quantut.NewSimulator(circuit, 10)
	res := simulator.Run()
	fmt.Println(res)

}
