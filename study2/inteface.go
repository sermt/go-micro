package main

import "fmt"

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

func genericFunction[T any](item T) T {
	switch v := any(item).(type) {
	case string:
		fmt.Println("It's a string:", v)
	case int:
		fmt.Println("It's an integer, doubled:", v*2)
	case float64:
		fmt.Println("It's a float64, doubled:", v*2)
	default:
		fmt.Println("Unhandled type")
	}
	return item
}
