version 1.0

qubits 3

H q[0]
CNOT q[1], q[2]
Toffoli q[0], q[1], q[2]
CNOT q[0], q[1]

measure q[1]
SWAP q[0], q[1]

measure q[0]

measure q[1]

measure q[2]
