package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Address
}

type Address struct {
	City    string
	State   string
	Pincode string
}

func main_string() {

	person := Person{
		Name: "Alic",
		Age:  23,
		Address: Address{
			City: "Hyderabad",
		},
	}

	var prinable Printable = person

	fmt.Println(prinable)

	person.UpdateName("Bob")

	fmt.Println(person)
	fmt.Println(person.Age)
	fmt.Println(person.City)

	var err error = person

	fmt.Println(err)

	// emtpy interface

	var someVal interface{}

	// someVal = 32
	someVal = "Veltris"
	someVal = false

	fmt.Println(someVal)

	switch v := someVal.(type) {

	case int:
		fmt.Println("v", v)
	case string:
	}

}

func (p Person) UpdateName(uname string) {
	p.Name = uname
}

// set of methods
// type implement it satishfy the intfer

type Printable interface {
	PrintStruct()
}

func (p Person) PrintStruct() {
	fmt.Println(p.Name, p.Age)
}

func (p Person) Error() string {
	return fmt.Sprintf("explicit error")
}
