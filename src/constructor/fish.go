package constructor

type Fish struct {
	Name  string
	Count int
}

// setter
func (f *Fish) SetName(name string) {
	f.Name = name
}

// setter
func (f *Fish) SetCount(count int) {
	f.Count = count
}

// getter
func (f *Fish) GetName() string {
	return f.Name
}

// getter
func (f *Fish) GetCount() int {
	return f.Count
}

// constructor
func NewFish(name string, count int) Fish {
	return Fish{
		Name:  name,
		Count: count,
	}
}
