package main

import "fmt"

type Employee struct {
	id     int
	name   string
	age    int
	sex    bool
	class  string
	height int
	weight int
}

func newEmployee(builder *EmployeeBuilder) Employee {
	employee := Employee{
		id:     builder.id,
		name:   builder.name,
		age:    builder.age,
		sex:    builder.sex,
		class:  builder.class,
		height: builder.height,
		weight: builder.weight,
	}
	return employee
}

func (e Employee) toString() string {
	return fmt.Sprintf("{\"id\":%d, \"name\":\"%s\", \"age\":%d, \"sex\":%v,\"class\":\"%s\",\"height\":%d,\"weight\":%d}",
		e.id, e.name, e.age, e.sex, e.class, e.height, e.weight)
}

type EmployeeBuilder struct {
	id     int
	name   string
	age    int
	sex    bool
	class  string
	height int
	weight int
}

func NewEmplyeeBuilder() *EmployeeBuilder {
	return &EmployeeBuilder{}
}

func (builder *EmployeeBuilder) withId(id int) *EmployeeBuilder {
	builder.id = id
	return builder
}

func (builder *EmployeeBuilder) withName(name string) *EmployeeBuilder {
	builder.name = name
	return builder
}

func (builder *EmployeeBuilder) withAge(age int) *EmployeeBuilder {
	builder.age = age
	return builder
}

func (builder *EmployeeBuilder) withSex(sex bool) *EmployeeBuilder {
	builder.sex = sex
	return builder
}

func (builder *EmployeeBuilder) withClass(class string) *EmployeeBuilder {
	builder.class = class
	return builder
}

func (builder *EmployeeBuilder) withHeight(height int) *EmployeeBuilder {
	builder.height = height
	return builder
}

func (builder *EmployeeBuilder) withWeight(weight int) *EmployeeBuilder {
	builder.weight = weight
	return builder
}

func (builder *EmployeeBuilder) build() Employee {
	return newEmployee(builder)
}
