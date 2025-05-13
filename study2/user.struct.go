package main

import "fmt"

type User struct {
	Name    string
	Age     int
	City    string
	Country string
}

func (u User) PrintName() {
	fmt.Printf("Name: %s\n", u.Name)
}

func (u User) PrintAge() {
	fmt.Printf("Age: %d\n", u.Age)
}

func (u User) PrintAddress() {
	fmt.Printf("Country: %s City: %s\n", u.Country, u.City)
}
