package main

import "fmt"

//Struct types using a method receiver. A method EmployeeInfo is added to the Employee struct.
type Salary struct {
	Basic, HR, TA float64
}

type Employee struct {
	firstName, email string
	monthlySalary    []Salary
}

func (e Employee) EmployeeInfo() string {
	fmt.Println(e.firstName)
	fmt.Println(e.email)

	for _, info := range e.monthlySalary {
		fmt.Println(info.Basic)
		fmt.Println(info.HR)
		fmt.Println(info.TA)
	}
	return ""
}

func main() {
	var employee = Employee{
		firstName:     "james",
		email:         "james@gmail.com",
		monthlySalary: []Salary{{Basic: 10, HR: 5, TA: 3}, {Basic: 12, HR: 6, TA: 2}},
	}

	fmt.Println(employee.EmployeeInfo())
}
