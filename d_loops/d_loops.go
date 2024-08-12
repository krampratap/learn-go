package main

import (
	"fmt"
	"math"
)

func main() {
	// for{
	// 	fmt.Println("Simple for")
	// }

	i := 1
	for i <= 5 {
		fmt.Println("printing ", i)
		i++
	}

	for j := 1; j <= 5; j++ {
		fmt.Println("printing j", j)
	}
	var num float64 = 9.5
	var result = math.Sqrt(num)
	fmt.Println(result)
	fmt.Printf("The result is %.2g Thank you", result)
}
