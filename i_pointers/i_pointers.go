package main

import "fmt"

func main() {
	m:="Tina"
	fmt.Println("Memory address of Tina:",&m)

	n:=&m
	fmt.Println("Memory address of Tina:",n)
	fmt.Println("Dereferencing the Pointer:", *n) //print is giving gap after "* Pointer:" wihtout specifying

	updateName(n)
	fmt.Println("Update name of Tina:", *n)
	
}

func updateName(x *string){//We are accepting in a pointer to whatever value s stored in that location
	*x="Rina"  //* means, get the value of the address at x and de reference it 
}