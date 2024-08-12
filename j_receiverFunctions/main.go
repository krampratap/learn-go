package main

import "fmt"

//To run: go run main.go bill.go
func main() {
	mybill := newBill("Dominos bill")
	fmt.Println(mybill.format())
}
