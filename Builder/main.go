package main

import "fmt"

func main() {
	employeeBuilder := NewEmplyeeBuilder()
	employeeBuilder.withAge(18).
		withName("liming").
		withSex(false).
		withHeight(180).
		withWeight(110).
		withId(1).
		withClass("English")

	employee := employeeBuilder.build()
	fmt.Println(employee.toString())
}