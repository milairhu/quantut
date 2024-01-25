package main

import (
	"fmt"
	"time"

	quantut "github.com/milairhu/quantut/pkg/quantut"
)

func executeCirc1() {

	circuit := quantut.NewQuantumCircuit(3)
	circuit.InitClassicalRegister(3)
	circuit.H(0)
	circuit.CNOT(0, 1)
	circuit.CCNOT(0, 2, 1)
	circuit.CSWAP(0, 1, 2)

	circuit2 := quantut.NewQuantumCircuit(2)
	circuit2.X(0)
	circuit2.H(1)
	circuit2.CNOT(1, 0)

	circuit = circuit.Compose(circuit2)

	circuit.Y(1)
	circuit.Z(2)

	circuit.Measure(0, 0)
	circuit.Measure(1, 1)
	circuit.Measure(2, 2)

	circuit.Draw()

	sim := quantut.NewSimulator(circuit, 1000)
	res := sim.Run()
	fmt.Println(res)

}

func main() {
	const NB_EXECUTION = 1
	startTime := time.Now()

	// Appeler la fonction à mesurer
	for i := 0; i < NB_EXECUTION; i++ {
		executeCirc1()
	}

	// Enregistrez le temps de fin
	endTime := time.Now()

	// Calculez la durée totale
	duration := endTime.Sub(startTime)

	// Affichez le temps d'exécution
	fmt.Printf("La fonction a mis %s pour s'exécuter %d fois.\n", duration, NB_EXECUTION)

}
