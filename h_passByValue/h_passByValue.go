package main

import "fmt"

func main() {
	//Pass by value: string, int, float, boolean, array, structs 
	x:= "Old Value"

	updateX(x)
	fmt.Println(x)


	//Pass by reference: maps slices functions 
	menu := map[string]float64{
		"soup":  4.99,
		"salad": 7.99,
		"pie":   3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])
	updateMenu(menu);
	fmt.Println(menu) //coffee is addded in

}

func updateMenu(newMenu map[string]float64) {
	newMenu["coffee"] = 9.88
}

func updateX(y string){
	y="New Value"
}
