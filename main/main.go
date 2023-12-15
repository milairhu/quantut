package main

import (
	"fmt"

	"github.com/milairhu/quantut"
)

func main() {

	circuit := quantut.NewQuantumCircuit(3) //création du circuit
	circuit.InitClassicalRegister(3)
	circuit.H(0) //ajout d'une porte de Hadamard sur le premier qubit
	circuit.CNOT(0, 1)
	circuit.CCNOT(0, 2, 1)
	circuit.SWAP(0, 1)
	circuit.CSWAP(0, 1, 2)
	circuit.Y(1)
	circuit.Z(2)
	circuit.X(0)
	circuit.H(2)
	circuit.Measure(0, 0)
	circuit.Measure(1, 1)
	circuit.Measure(2, 2)

	circuit.Draw()

	sim := quantut.NewSimulator(circuit, 1000) //création du simulateur
	res := sim.Run()                           //lancement de la simulation
	fmt.Println(res)                           //affichage des résultats

	circuit.ToQASM("test.qasm", "OPENQASM 2.0")
}
