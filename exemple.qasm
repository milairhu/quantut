OPENQASM 2.0;
include "qelib1.inc";

qreg q[4];

creg c[2];

h q[0];
cx q[0], q[2];
x q[2];
ccx q[0], q[3], q[2];
swap q[0], q[1];

measure q[0] -> c[0];

measure q[1] -> c[1];
