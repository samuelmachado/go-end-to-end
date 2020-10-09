package main

import "fmt"

type testiface interface {
	SayHello()
	Say(s string)
	Increment()
	GetInternalValue() int
}
type testConcreteImpl struct {
	i int
}

func NewTestConcretImpl() testiface {
	return new(testConcreteImpl)
}
func (tst *testConcreteImpl) SayHello() {
	fmt.Println("Hello")
}
func (tst *testConcreteImpl) Say(s string) {
	fmt.Println(s)
}
func (tst *testConcreteImpl) Increment() {
	tst.i++
}
func (tst *testConcreteImpl) GetInternalValue() int {
	return tst.i
}

type testEmbedding struct { // we walt this struct to have all features of *testConcretImpl
	*testConcreteImpl
}

func main() {
	var tiface testiface
	tiface = NewTestConcretImpl()
	var tiface2 testiface = NewTestConcretImpl()

	tiface.SayHello()
	tiface.Say("!!!!!")

	fmt.Println("hello!")
	tiface.Increment()
	fmt.Println(tiface.GetInternalValue())
	fmt.Println(tiface2.GetInternalValue())

}
