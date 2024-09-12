package patterns

type Human interface {
	Act()
}

type American struct {
	Human
}

func (American) Act() {
	println("I m normal citizen")
}

type DrugDealer struct {
	American
}

func (d DrugDealer) Act() {
	d.American.Act()
	println("I m selling drugs now!")
}
