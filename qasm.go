package quantut

import (
	"fmt"
	"os"
	"strings"
)

// === Create QASM file ===
// Create the content of QASM file

// Language found on https://www.quantum-inspire.com/kbase/cqasm/
// WARNING : CSWAP is not supported by QASMv1
func (c *QuantumCircuit) generateCQASM() string {
	var qasmContent string = "version 1.0\n\n"

	//Number of qubits
	qasmContent += fmt.Sprintf("qubits %d\n\n", c.numQubits)

	//Operations
	for _, op := range c.operations {
		if op.Gate().Id() == "INIT" || op.Gate().Id() == "CSWAP" {
			panic(fmt.Sprintf("Gate %s is not supported by QASM", op.Gate().Id()))
		}
		if len(op.Qubits()) == 1 {

			qasmContent += fmt.Sprintf("%s q[%d]\n", op.Gate().Id(), op.Qubits()[0])

		} else {
			if op.Gate().Id() == "MEASURE" {
				qasmContent += fmt.Sprintf("\nmeasure q[%d]\n", op.Qubits()[0])
			} else {
				if op.Gate().Id() == "CCNOT" {
					qasmContent += "Toffoli "
				} else {
					qasmContent += fmt.Sprintf("%s ", op.Gate().Id())
				}
				for i := 0; i < len(op.Qubits())-1; i++ {
					qasmContent += fmt.Sprintf("q[%d], ", op.Qubits()[i])
				}
				qasmContent += fmt.Sprintf("q[%d]\n", op.Qubits()[len(op.Qubits())-1])
			}
		}
	}

	return qasmContent

}

// Create a file compilable by qiskit
func (c *QuantumCircuit) generateOPENQASM() string {
	var qasmContent string = "OPENQASM 2.0;\ninclude \"qelib1.inc\";\n\n"

	//Number of qubits
	qasmContent += fmt.Sprintf("qreg q[%d];\n\n", c.numQubits)

	//Classical register
	qasmContent += fmt.Sprintf("creg c[%d];\n\n", c.NbClassicalBits())

	//Operations
	for _, op := range c.operations {
		if op.Gate().Id() == "INIT" {
			//TODO : vérifier
			panic(fmt.Sprintf("Gate %s is not supported by QASM", op.Gate().Id()))
		}
		if len(op.Qubits()) == 1 {
			qasmContent += fmt.Sprintf("%s q[%d];\n", strings.ToLower(op.Gate().Id()), op.Qubits()[0])

		} else {
			if op.Gate().Id() == "MEASURE" {
				qasmContent += fmt.Sprintf("\nmeasure q[%d] -> c[%d];\n", op.Qubits()[0], op.Qubits()[1])
			} else {
				switch op.Gate().Id() {
				case "CCNOT":
					qasmContent += "ccx "
				case "CNOT":
					qasmContent += "cx "
				default:
					qasmContent += fmt.Sprintf("%s ", strings.ToLower(op.Gate().Id()))
				}
				for i := 0; i < len(op.Qubits())-1; i++ {
					qasmContent += fmt.Sprintf("q[%d], ", op.Qubits()[i])
				}
				qasmContent += fmt.Sprintf("q[%d];\n", op.Qubits()[len(op.Qubits())-1])
			}
		}
	}

	return qasmContent

}

// Use previous function to write the ontent in a file
func (c *QuantumCircuit) ToQASM(filename string, version string) {
	var content string
	switch version {
	case "cQASM":
		content = c.generateCQASM()
	case "OPENQASM 2.0":
		content = c.generateOPENQASM()
	default:
		panic(fmt.Sprintf("Version %s is not supported for generating qasm file", version))
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}
