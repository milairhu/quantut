package quantut

import (
	"fmt"
	"sync"
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
	/*Idée : on lance 1 go routine par qubit.
	On met un channel entre chque qubit pour synchroniser les opérations quand il y a des portes à plusieurs qubits
	Un channel ppermet d'envoyer sa valeur actuelle à d'autres qubits
	*/

	//result map
	res = make(map[string]uint)

	//Store initial values of the qubits and classical registers
	initQubitsValues := make([]Qubit, s.circuit.numQubits)
	for i := 0; i < s.circuit.numQubits; i++ {
		initQubitsValues[i] = s.circuit.qubitsValues[i]
	}
	initClassicalRegister := make([]int, len(s.circuit.classicalRegister))
	copy(initClassicalRegister, s.circuit.classicalRegister)

	//Création des channels : un par paire de qubits
	var nbChannels int = (s.circuit.numQubits * (s.circuit.numQubits - 1)) / 2

	channels := make([]chan Qubit, nbChannels)
	for i := 0; i < nbChannels; i++ {
		channels[i] = make(chan Qubit)
	}
	//Dans un map, on stocke les channels auxquels chaque qubit a accès
	channelsMap := make(map[int][]chan Qubit, s.circuit.numQubits)
	for i := 0; i < s.circuit.numQubits; i++ {
		channelsMap[i] = make([]chan Qubit, 0)
	}
	//On remplit le map avec les channels
	var nextChannel int = 0
	for i := 0; i < s.circuit.numQubits-1; i++ {
		//Pour chaque qubit sauf le dernier
		for j := i + 1; j < s.circuit.numQubits; j++ {
			//pour chaque qubit suivant
			channelsMap[i] = append(channelsMap[i], channels[nextChannel])
			channelsMap[j] = append(channelsMap[j], channels[nextChannel])
			nextChannel++
		}
	}

	for numShot := 0; numShot < int(s.shots); numShot++ {
		//On simule un shot
		//resChan := make(chan string)
		wg := sync.WaitGroup{}

		//storeRes := make([]string, s.circuit.numQubits)
		for i := 0; i < s.circuit.numQubits; i++ {
			wg.Add(1)

			go func(i int) {
				//Problème : toutes les go routines ne sont pas envoyées...
				defer wg.Done()
				s.Circuit().LaunchQubit(i, channelsMap[i])
			}(i)
		}
		wg.Wait()
		/*
			//On récupère le résultat de chaque qubit
			for i := 0; i < s.circuit.numQubits; i++ {
				storeRes[i] = <-resChan //retourne un string "numQubit,valeur", ce qui nous permet de pouvoir ranger ces résultats par ordre alphabétique
				<-resChan
			}*/

		//Si on enregistre les états finaux des qubits :
		/*
			//Sort storeRes
			sort.Strings(storeRes)

			var resShot string
			for i := 0; i < s.circuit.numQubits; i++ {
				storeRes[i] = strings.Split(storeRes[i], ",")[1]
				resShot += storeRes[i]
			}
		*/

		//Si on enregistre les états finaux des registres classiques :
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

		//On remet les qubits à leur état initial
		for i := 0; i < s.circuit.numQubits; i++ {
			s.circuit.qubitsValues[i] = initQubitsValues[i]
		}
		//On remet le registre classique à son état initial
		for i := 0; i < len(s.circuit.classicalRegister); i++ {
			s.circuit.classicalRegister[i] = initClassicalRegister[i]
		}

	}

	return res
}
