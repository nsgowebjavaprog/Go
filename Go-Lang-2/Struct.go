package main

import "fmt"

var pf = fmt.Println

/*
type customer struct {
	name    string
	address string
	bal     float64
}

func get_custo_info(c customer) {
	fmt.Printf("%s owns us %f.2\n", c.name, c.bal)
}

func new_custo_add(c *customer, address string) {
	c.address = address
}

func main() {
	var ts customer
	ts.name = "Name_of_custo"
	ts.address = "Bijapur"
	ts.bal = 2323456.34

	get_custo_info(ts) //Name_of_custo owns us 2323456.340000.2

	//new_custo_add(&ts)

}
*/

type rectangle struct {
	len, hei float64
}

//The code defines a rectangle struct with two fields:
//len and hei, both of type float64.
// These fields represent the length and height of the rectangle.

func (r rectangle) Area() float64 {
	return r.len * r.hei
}

func main() {
	reac := rectangle{10, 20}
	pf("RectAngle is: ", reac.Area()) // RectAngle is: 200
}
