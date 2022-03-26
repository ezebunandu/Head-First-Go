package main

import (
	"fmt"
	"head_first_go/magazine"
)

func main() {
	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	employee.Street = "456 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
	fmt.Println(employee.Address)
}
