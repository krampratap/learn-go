package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

// return two variables from function
func calc(num1, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func calc2(num1, num2 int) (out1, out2 int) {
	out1 = num1 + num2
	out2 = num1 - num2
	return
}

func main() {
	num1 := 5
	num2 := 4

	sum := add(num1, num2)
	sum, sub := calc(num1, num2)
	fmt.Println(sum)
	fmt.Println(sub)
}
