package main

type Greeter interface {
	sayHello()
}

type Employee struct {
	name string
	age  int
}

func (e Employee) sayHello() {
	println("Hello")
}
