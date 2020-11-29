package Mapper

type TodosMapper struct {
	Title  string
	Status int
	base   //继承
}

type base struct {
	Id int
}
