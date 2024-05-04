package constructor

type Horse struct {
	Name string
	Age  int
}

func NewHorse(name string, age int) Horse {
	return Horse{Name: name, Age: age}
}
