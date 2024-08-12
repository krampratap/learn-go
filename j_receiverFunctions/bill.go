package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 12.99, "cake": 13.99},
		tip:   0,
	}
	return b
}

func (b bill) format() string {
	fmt.Println("Receive function")
	fs := "Bill breakdown \n"
	var total float64 = 0

	//list itemms
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) //-25 bcoz we want the spacing on the right, for left you need 25
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}
