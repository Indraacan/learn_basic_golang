package example

//Exampel Data for new Keyword
type Exampel struct {
	Name string
	Age  int
}

//NewExample function
func NewExample(name string, age int) *Exampel {

	e := new(Exampel)
	e.Name = name
	e.Age = age

	return e
}
