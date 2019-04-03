package hard

import (
	"fmt"
	"math"
	"testing"
)

func TestMatrixMul(t *testing.T) {
	mul := matrixMul
	pow := matrixPow

	m := [4]int{1, 1, 3, 1}

	r := mul(m, pow(m, 2))
	printMatrix(r)

	fmt.Println("----")

	r = mul(pow(m, 2), m)
	printMatrix(r)
}

func TestFib(t *testing.T) {
	fmt.Printf("%v\n", Fib(0))
	fmt.Printf("%v\n", Fib(1))
	fmt.Printf("%v\n", Fib(2))
	fmt.Printf("%v\n", Fib(3))
	fmt.Printf("%v\n", Fib(4))
	fmt.Printf("%v\n", Fib(5))
	fmt.Printf("%v\n", Fib(6))
	fmt.Printf("%v\n", Fib(7))
}

func BenchmarkFibIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibIterate(1000)
	}
	fmt.Println(math.MaxInt64)
	fmt.Println(FibIterate(1000))
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(1000)
	}
}

func printMatrix(m [4]int) {
	fmt.Printf(`[%v	%v]
[%v	%v]
`, m[0], m[1], m[2], m[3])
}
