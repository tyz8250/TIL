package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct{
	firstName string
	lastName string
	contactInfo
}

func main(){
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.updateName("Tommy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFistName string){
	(*pointerToPerson).firstName = newFistName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}
