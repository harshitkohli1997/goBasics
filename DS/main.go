package main

import (
	"fmt"
)

type contactInfo struct {
	ph  int
	add string
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	hk := person{firstName: "HK"}
	fmt.Println(hk)
	hkptr := &hk
	hk.contact.add = "dilshad garden"

	hk.print()
	hkptr.update("DUDE")
	hk.print()

	hk.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) update(newFirstName string) {
	(*p).firstName = newFirstName
}
