package main

import (
	"fmt"
	"github.com/flyinprogrammer/pkl-go-bananas/gen/earth"
	"reflect"
)

func NewPerson() earth.Person {
	return &earth.PersonImpl{
		FirstName: "Alan",
		LastName:  "Scherger",
	}
}
func main() {
	alanPerson := NewPerson()

	atlasMoth := earth.Bug{
		Name:  "Addison",
		Owner: &alanPerson,
	}

	fmt.Println("Type of alanPerson: ", reflect.TypeOf(alanPerson))
	fmt.Println("Type of atlasMoth: ", reflect.TypeOf(atlasMoth))
	me := *atlasMoth.Owner
	fmt.Println("Type of me: ", reflect.TypeOf(me))
	fmt.Println("atlasMoth Owner GetFirstName: ", (*atlasMoth.Owner).GetFirstName())
	fmt.Println("me PersonImpl FirstName: ", me.(*earth.PersonImpl).FirstName)

}
