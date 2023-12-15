package quantut

import (
	"fmt"
)

/*
* The simulator is the object that will run the circuit
 */
type Simulator struct {
	circuit *QuantumCircuit //Circuit to run
	shots   uint            //Number of shots to simulate
}

// Constructor
func NewSimulator(circuit *QuantumCircuit, shots uint) *Simulator {
	return &Simulator{circuit: circuit, shots: shots}
}

// Getters
func (s *Simulator) Circuit() *QuantumCircuit {
	return s.circuit
}

func (s *Simulator) Shots() uint {
	return s.shots
}

// Setters
func (s *Simulator) SetCircuit(circuit *QuantumCircuit) {
	s.circuit = circuit
}

func (s *Simulator) SetShots(shots uint) {
	s.shots = shots
}

// Run the circuit and return the result
func (s *Simulator) Run() (res map[string]uint) {

	//The simulator apply the operations of the circuit on the qubits
	//result map
	res = make(map[string]uint)

	initClassicalRegister := make([]int, len(s.circuit.classicalRegister))
	copy(initClassicalRegister, s.circuit.classicalRegister)
	initGlobalState := make([]complex128, len(s.circuit.globalState))
	copy(initGlobalState, s.circuit.globalState)

	for numShot := 0; numShot < int(s.shots); numShot++ {
		//On simule un shot

		s.Circuit().LaunchCircuit()

		// on enregistre les états finaux des registres classiques :
		var resShot string
		for i := len(s.circuit.classicalRegister) - 1; i >= 0; i-- {
			resShot += fmt.Sprintf("%d", s.circuit.classicalRegister[i])
		}
		//on stocke le résultat dans le map
		_, exists := res[resShot]
		if !exists {
			res[resShot] = 1
		} else {
			res[resShot]++
		}
		//On remet le registre classique à son état initial
		for i := 0; i < len(s.circuit.classicalRegister); i++ {
			s.circuit.classicalRegister[i] = initClassicalRegister[i]
		}
		//On remet l'état général à son état initial
		for i := 0; i < len(s.circuit.globalState); i++ {
			s.circuit.globalState[i] = initGlobalState[i]
		}
	}
	return res
}
