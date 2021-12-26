package main

import "fmt"

func main() {
	c := make(chan int, 0)

		go func(){
			fmt.Println(1)
			 <-c
		}()
		go func(){
			fmt.Println(1)
			<-c
		}()
		go func(){
			fmt.Println(1)
			<-c
		}()
		go func(){
			fmt.Println(1)
			<-c
		}()
		go func(){
			fmt.Println(1)
			<-c
		}()
		go func(){
			fmt.Println(1)
			<-c
		}()
		go func(){
			fmt.Println(1)
			<-c
		}()
		go func(){
			fmt.Println(1)
			<-c
		}()
		go func(){
			fmt.Println(1)
			<-c
		}()
		go func(){
			fmt.Println(1)
			<-c
 		}()
	 for i:=0;i<10;i++{
		 c <- 0
	 }
}