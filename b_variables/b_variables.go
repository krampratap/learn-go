package main

import "fmt"

func main() {
	var num int
	var num1, num2 int
	num1 = 5
	num2 = 6
	num1, num2 = 4, 11
	result := 9 // var result = 9

	const constant = 10
	fmt.Println(result)
	fmt.Println(num1 + num2)
	fmt.Println(num)
}

