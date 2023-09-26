package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.State)
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Person2 struct {
	Name string
	Age  int
	Address // Embedded Address
}

func main() {
	person := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street: "123 Main St",
			City:   "New York",
			State:  "NY",
		},
	}
	fmt.Println("Street:", person.Address.Street)

	person2 := Person2{
		Name: "Jane Doe",
		Age:  28,
		Address: Address{
			Street: "456 Elm St",
			City:   "Los Angeles",
			State:  "CA",
		},
	}
	fmt.Println("City:", person2.City) // Direct access to the embedded field
	fmt.Println("Full Address:", person2.FullAddress()) // Using the method from the embedded struct
}
