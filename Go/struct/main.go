package main

import "fmt"

type person struct{
	firstName string
	lastName string
}

func main(){
	// structに型を入れている。
	alex := person{firstName:"Alex", lastName: "Anderson"}
	fmt.Println(alex)

}