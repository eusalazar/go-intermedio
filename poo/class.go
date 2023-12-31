package main

import "fmt"

type Employee struct {
	id   int
	name string
}

// lo q indicamos que la struct Employee ahora va a poseer un metodo SetId
func (e *Employee) SetId(id int) { //recibe un parametro de tipo entero
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	//fmt.Printf("%v", e)
	e.id = 1
	e.name = "Name"
	//fmt.Printf("%v", e)
	e.SetId(5)
	e.SetName("Name 2")
	//fmt.Printf("%v", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
