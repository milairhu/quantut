package quantut

import "fmt"

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
		resChan := make(chan string)
		storeRes := make([]string, s.circuit.numQubits)
		for i := 0; i < s.circuit.numQubits; i++ {
			go s.Circuit().LaunchQubit(i, channelsMap[i], resChan)
		}
		//On récupère le résultat de chaque qubit
		for i := 0; i < s.circuit.numQubits; i++ {
			storeRes[i] = <-resChan //retourne un string "numQubit,valeur", ce qui nous permet de pouvoir ranger ces résultats par ordre alphabétique
		}

		//TODO : dois-je enregistrer les états finaux des qubits ou ceux des registres classiques ?

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

	}

	return res
}
