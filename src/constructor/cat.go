package constructor

type Cat struct {
	Name        string
	Age         int
	HomeAddress string
	Sex         string
}

// Java constructor equals function with struct
func NewCat(name, address, sex string, age int) Cat {
	return Cat{
		Name:        name,
		HomeAddress: address,
		Sex:         sex,
		Age:         age,
	}
}
