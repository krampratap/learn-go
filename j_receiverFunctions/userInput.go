package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println(opt)
}
