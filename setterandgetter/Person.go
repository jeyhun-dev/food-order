package setterandgetter

type Person struct {
	name    string
	age     int
	address string
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func (p *Person) SetAddress(address string) {
	p.address = address
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) GetAddress() string {
	return p.address
}
