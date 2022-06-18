package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// just writing contactInfo enought for contactInfo contactInfo
}

func main() {
	// yp := person{"yigit", "polat"}

	// defines zero value "" / 0 / false
	// var yp person
	// yp.firstName = "y"
	// yp.lastName = "p"

	yp := person{firstName: "y",
		lastName: "p",
		contact: contactInfo{
			email:   "yigit",
			zipCode: 123456,
		},
	}

	fmt.Println(yp)
	// ypPointer := &yp
	// ypPointer.updateName("Yigit")
	// shortcut just give the type of person
	yp.updateName("Yigit")
	yp.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
