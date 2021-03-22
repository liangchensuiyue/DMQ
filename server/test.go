package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 10
	ch <- 20
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(len(ch), cap(ch))

}
