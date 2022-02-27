package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time employee"
}

func (tEmployee TemporatyEmployee) getMessage() string {
	return "Temporary employee"
}

type TemporatyEmployee struct {
	Person
	Employee
	taxRate int
}

type Printinfo interface {
	getMessage() string
}

func getMessage(p Printinfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Emiliano"
	ftEmployee.age = 22
	ftEmployee.id = 5

	tEmployee := TemporatyEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
