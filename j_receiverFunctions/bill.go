package main

import (
	"fmt"
	"os"
)

//It is better to pass pointers because everytime you go will create a new copy of the object to use
//inside the method
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b *bill) format() string {
	fmt.Println("Receive function")
	fs := "Bill breakdown \n"
	var total float64 = 0

	//list itemms
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) //-25 bcoz we want the spacing on the right, for left you need 25
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)
	total += b.tip
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}

// Receiver Functions with pointers
func (b *bill) updateTip(tip float64) {
	b.tip = tip //Automatic de-referencing for Structs
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save(){
	data:=[]byte(b.format()) //we have to write bytes to the file
	err:=os.WriteFile("bills/"+b.name+".txt",data,0644) //0644 is the permission
	if err!=nil {
		panic(err) //This is a method to stop the flow of the program and print out an error
	}	
	fmt.Println("Bill was saved to file")
}