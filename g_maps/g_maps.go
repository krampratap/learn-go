package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":  4.99,
		"salad": 7.99,
		"pie":   3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping through maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phonebook := map[int]string{
		122234555: "Tony",
		222222222: "Stark",
		999999999: "Captain",
	}
	fmt.Println(phonebook)

	fmt.Println(phonebook[222222222])
}