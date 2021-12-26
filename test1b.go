package main

import "fmt"

func main() {
c := make(chan int,0)
	go func() {
		fmt.Println("下山的路又堵起了")
		<-c
	}()
	c <- 0
}