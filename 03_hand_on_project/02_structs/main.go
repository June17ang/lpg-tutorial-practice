package main

import "fmt"

type contactInfo struct {
	mobileNumber string
	email        string
}

type person struct {
	fname string
	lname string
	contactInfo
}

func main() {
	p1 := person{
		fname: "foo",
		lname: "bar",
		contactInfo: contactInfo{
			mobileNumber: "0123123123",
			email:        "test@test.com",
		},
	}

	fmt.Println(p1)
	p1Pointer := &p1
	p1Pointer.updateName("foob")
	p1.print()

	p1Pointer.updateName("test")
	p1.print()
}

//Avoiding pass by value
//Using pass by reference

func (pt *person) updateName(newFirstName string) {
	(*pt).fname = newFirstName
}

func (pt person) print() {
	fmt.Printf("%+v\n", pt)
}
