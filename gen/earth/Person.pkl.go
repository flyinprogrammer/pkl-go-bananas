// Code generated from Pkl module `flyinprogrammer.pkl_go_bananas.Earth`. DO NOT EDIT.
package earth

type Person interface {
	Being

	GetFirstName() string

	GetLastName() string
}

var _ Person = (*PersonImpl)(nil)

// A Person!
type PersonImpl struct {
	IsAlive bool `pkl:"isAlive"`

	// The person's first name
	FirstName string `pkl:"firstName"`

	// The person's last name
	LastName string `pkl:"lastName"`
}

func (rcv *PersonImpl) GetIsAlive() bool {
	return rcv.IsAlive
}

// The person's first name
func (rcv *PersonImpl) GetFirstName() string {
	return rcv.FirstName
}

// The person's last name
func (rcv *PersonImpl) GetLastName() string {
	return rcv.LastName
}
