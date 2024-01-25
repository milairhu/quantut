# QuantUT : A Go package for Quantum Computing

Go package for Quantum Computing, implemented as a complement of a project led in the context of the course **IQ01** at the University of Technology of Compiègne (UTC), supervised by [Mr. Walter Schön](https://www.hds.utc.fr/~wschon/dokuwiki/fr/biographie) and [Mr. Ahmed Lounis](https://www.hds.utc.fr/~lounisah/dokuwiki/).

The package is based on the Qiskit and Circ libraries. It can be used to create, display and simulate quantum circuits.

## Functionalities

- Create, display and simulate quantum circuits. The package includes the composition of quantum gates, and the measurement of qubits.
- Export the circuit to a file in the QASM format. 2 formats are available.

## Getting Started

1. Install the package with the following command:

```bash
go get github.com/milairhu/quantut
```

2. Import the package in your code:

```go
import "github.com/milairhu/quantut/pkg/quantut"
```

## A simple program

```go
// Create a quantum circuit with 2 qubits and 2 classical bits
qc := quantut.NewQuantumCircuit(2)
qc.InitClassicalRegister(2)
qc.H(0)
qc.CNOT(0, 1)
qc.Measure(0, 0)
qc.Measure(1, 1)

// Display the circuit
qc.Draw()

// Simulate the circuit
sim := quantut.NewSimulator(qc, 1000)
res := sim.Run()
fmt.Println(res)

// Export the circuit to a file
qc.ToQASM("circuit.qasm", "OPENQASM 2.0")
```

## Areas of improvement

The package is functionnal. However, it can be improved in the following ways:

- Add **more quantum gates**. So far, only the most common gates are implemented. S and T gates are missing for example.
- Add **pre-built quantum circuits**. For example, the Grover algorithm, the Shor algorithm, Deutsch-Jozsa algorithm, etc.
- Improve the display of the circuit. The current display resembles the one of Cirq, but it can be improved. Especially, **some gates could be displayed on a single vertical line** to gain space.
- Improve the performance of the simulator. The simulator is currently quite slow. It could be improved by **using matrix product** instead of the current "hand-made" method.
- Read QASM file for circuits loading
  
