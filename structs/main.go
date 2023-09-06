package main

import "fmt"

type contactInfo struct {
    email   string
    zipCode int
}

type person struct {
    firstName string
    lastName  string
    //contact   contactInfo
    contactInfo
}

func main() {
    //alex := person{"Alex", "Anderson"}
    //alex := person{firstName: "Alex", lastName: "Anderson"}

    //var alex person
    //
    //alex.firstName = "Alex"
    //alex.lastName = "Anderson"
    //
    //fmt.Println(alex)
    //fmt.Printf("%+v", alex)

    jim := person{
        firstName: "Jim",
        lastName:  "Party",
        contactInfo: contactInfo{
            email:   "jim@gmail.com",
            zipCode: 94000,
        },
    }

    //jimPointer := &jim
    //jimPointer.updateName("Jimmy")
    //jim.print()

    jim.updateName("Jimmy")
    jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
    fmt.Printf("%+v", p)
}
