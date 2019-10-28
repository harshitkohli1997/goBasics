package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//hk := person{firstName: "HK", lastName: "Kohli"}
	//fmt.Println(hk)

	var hk person
	fmt.Printf("%+v", hk)
}
