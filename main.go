package main

import (
	"fmt"
)




func main() {
	fmt.Println("Hello, playground")
	
	f := func() <-chan int {
		n := 10
		out := make(chan int)
		go func() {
			for i:=0; i < n;i++ {
				out <- i
			}
			close(out)
		}()

		return out
	}
	
	gen := f()
	fmt.Println(<-gen)
	fmt.Println(<-gen)
	for v := range gen {
		fmt.Println(v)
	}
	
}
