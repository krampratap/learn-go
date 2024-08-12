package main

import "fmt"

func main(){
	a()
}

func a(){
	fmt.Println("In a first line")
	defer b() 
	fmt.Println("In a second line")
}

func b(){
	fmt.Println("In b")
}
