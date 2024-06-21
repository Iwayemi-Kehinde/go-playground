package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("%d is %d years old\n", p.Name, p.Age)
}
func main() {
	p := Person{"Iwayemi Kehinde", 15}
	p.Greet()
	mybill := newBill("mario's bill")
	mybill.addItem("onion soup", 4.50)
	mybill.addItem("veg pie", 8.95)
	mybill.addItem("toffee pudding", 4.95)
	mybill.addItem("coffee", 3.25)
	mybill.updateTip(10)
	fmt.Println(mybill.format())
}
