package main

import (
	"fmt"
	"sync"
)

func main() {
	workerNum := 10
	slice := make([]int, workerNum)
	wg := &sync.WaitGroup{}

	for i := 0; i < workerNum; i++ {
		j := i
		wg.Add(1)
		go func(){
			defer wg.Done()
			slice[j] = j
		}()
	}

	fmt.Println(slice)
}

