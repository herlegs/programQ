package main

import "fmt"

func main() {
	//ip := make([]int, 2)
	//addr := add(ip)
	//fmt.Println(addr == &ip)

	x := test{}
	y := test{}
	fmt.Println(x == y)
}

type test struct {
	a int
}

func add(a []int) *[]int {
	return &a
}
