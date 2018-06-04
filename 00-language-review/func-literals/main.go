package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println(p.firstName, p.lastName, "says, 'Good morning, James.'")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.firstName, sa.lastName, "says, 'Shaken not stirred.'")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	xi := []int{2, 3, 4, 5, 42}
	fmt.Println(xi)

	m := map[string]int{
		"Todd":    45,
		"Matthew": 31,
	}
	fmt.Println(m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	fmt.Println(p1)

	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()

	sa1.person.speak()

	// Polymorphism
	saySomething(p1)
	saySomething(sa1)
}
