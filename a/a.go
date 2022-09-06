package a

import "fmt"

type (
	A struct {
		b1 int
		c1 C
	}

	C struct {
		d string
	}

	AInterface interface {
		Show()
	}
)

func (a *A) Show() {
	fmt.Println("Show A")
}

func NewA() *A {
	var _ AInterface = (*A)(nil)
	return &A{}
}
