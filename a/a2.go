package a

import "fmt"

type (
	A2 struct {
		b1 int
		c1 C
	}
)

func (a *A2) Show() {
	fmt.Println("Show A2")
}

func NewA2() *A2 {
	var _ AInterface = (*A2)(nil)
	return &A2{}
}
