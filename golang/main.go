package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		time.Sleep(time.Millisecond * 500)
		close(ch)
	}()

	for {
		select {
		case d := <-ch:
			fmt.Printf("%v\n", d)
		}
	}
}
