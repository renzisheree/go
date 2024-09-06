package main

import "fmt"
type contactInfo struct {
	email  string
	zipcode int
}

type person struct{
	firstName string
	lastName string
	 contactInfo
}
func main() {
	 Hoa := person{
		firstName: "Hoa",
		lastName: "Le",
		contactInfo: contactInfo{
			email: "Hoa@gmail.com",
			zipcode: 130703,
		},
	 }
	 Hoa.updateName("Vo iu")
	 Hoa.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	 	 fmt.Printf("%+v",p)

}