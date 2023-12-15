package utils

import "fmt"

func ComplexMatrixProduct(a, b [][]complex128) [][]complex128 {
	if len(a[0]) != len(b) {
		panic(fmt.Sprintf("Matrix multiplication impossible : %d != %d", len(a[0]), len(b)))
	}

	res := make([][]complex128, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = make([]complex128, len(b[0]))
		for j := 0; j < len(b[0]); j++ {
			res[i][j] = 0
			for k := 0; k < len(a[0]); k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return res

}
