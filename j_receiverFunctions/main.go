package main

import "fmt"

// To run: go run main.go bill.go
func main() {
	mybill := newBill("Dominos bill")
	mybill.updateTip(10)
	mybill.addItem("onion soup", 10.50)
	fmt.Println(mybill.format())
}
