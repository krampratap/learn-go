package main

import "fmt"

// To run: go run main.go bill.go
func main() {
	mybill := newBill("Dominos bill")
	mybill.updateTip(10)
	mybill.addItem("onion soup", 10.50)
	mybill.addItem("pie", 12.99)
	mybill.addItem("cake", 13.99)
	fmt.Println(mybill.format())
	fmt.Println()
	newBill := createBill()
	promptOptions(&newBill)
	fmt.Println(newBill)
}
