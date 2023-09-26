package main

import "fmt"


type Person struct {
	Name string
	Age int
	Address string
	Height float32
}

type Node struct {
	Value string
	Next *Node
}

type Tree struct {
	Next *Tree
}

type Salary struct {
	Basic, HR, TA float64
}

type Employee struct {
	firstName, email string
	monthlySalary    []Salary
}


func main() {

	alice := Person{
		Name: "Alice",
		Age: 10,
		Address: "ABC",
		Height: 4.3,
	}
	fmt.Println(alice)
	fmt.Println(alice.Name)

	var johnPtr *Person = new(Person)
	fmt.Println(*johnPtr)
	fmt.Println(johnPtr.Age)

	describe(alice)

	mane := newPerson("mane", 10, "ABC", 5.2)
	mane.describe()
	mane.incrementAge()
	mane.describe()

	var employee = Employee{
		firstName:     "james",
		email:         "james@gmail.com",
		monthlySalary: []Salary{{Basic: 10, HR: 5, TA: 3}, {Basic: 12, HR: 6, TA: 2}},
	}

	employee.EmployeeInfo()
}

func describe(person Person) {
	fmt.Println(person)
}

func newPerson(name string, age int, address string, height float32) *Person {
	return &Person{
		Name: name,
		Age: age,
		Address: address,
		Height: height,
	}
}

func (p *Person) describe() {
	fmt.Println(*p)
}

func (p *Person) incrementAge() {
	p.Age++
}

func (e *Employee) EmployeeInfo() {
	fmt.Println(e.firstName)
	fmt.Println(e.email)

	for _, info := range e.monthlySalary {
		fmt.Println(info.Basic)
		fmt.Println(info.HR)
		fmt.Println(info.TA)
	}

}