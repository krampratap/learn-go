package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	bufio - for input read
	os - is the source from where the bufio should read the input from
	we get value, error from bufio ReadString().. here we have to pass until which character we have to read the values
*/

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	/*reader := bufio.NewReader(os.Stdin)
	fmt.Println("Create a new bill name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("The name entered is: ", name)*/

	reader := bufio.NewReader(os.Stdin)

	input, _ := getInput("Get Bill name:", reader)
	b := newBill(input)
	fmt.Println("Created bill with name:", input)
	return b
}

func promptOptions(b *bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Select option (a - add item, s - save bill, t - add tip):", reader)
	switch opt {
	case "a":
		name, _ := getInput("Get the item name", reader)
		price, _ := getInput("Get the Price", reader) //Automatically the inputs are strings, so covert to floats

		p, err := strconv.ParseFloat(price, 64) //64 is float 64
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println(name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Get Tip", reader)
		t, err := strconv.ParseFloat(tip, 64) //64 is float 64
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("The bill is saved: ", b.name)
	default:
		fmt.Println("That was not a valid option")
		promptOptions(b)
	}
}
