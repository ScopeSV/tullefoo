package foomodel

type FooModel struct{}

func (fm *FooModel) GetFoo() string {
	return "foo"
}

func NewFooModel() *FooModel {
	return &FooModel{}
}
