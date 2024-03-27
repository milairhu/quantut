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

	// Call the function to be benchmarked
	for i := 0; i < NB_EXECUTION; i++ {
		executeCirc1()
	}

	// Record time spent
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	// Display recorded time
	fmt.Printf("The function took %s to execute %d times.\n", duration, NB_EXECUTION)

}
