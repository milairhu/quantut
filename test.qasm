OPENQASM 2.0;
include "qelib1.inc";

qreg q[3];

creg c[3];

h q[0];
cx q[0], q[1];
ccx q[0], q[2], q[1];
swap q[0], q[1];
cswap q[0], q[1], q[2];
x q[0];
h q[1];
cx q[1], q[0];
y q[1];
z q[2];

measure q[0] -> c[0];

measure q[1] -> c[1];

measure q[2] -> c[2];
