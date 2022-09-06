package c

import (
	"fmt"
	"wire_best_practice/a"
	"wire_best_practice/b"
)

type (
	C struct {
		b  b.BInterface
		a  a.AInterface
		a1 int
		z  string
	}

	CInterface interface {
		View()
		A() a.AInterface
		B() b.BInterface
	}
)

func (c *C) View() {
	fmt.Println("view C")
}

func (c *C) A() a.AInterface {
	return c.a
}

func (c *C) B() b.BInterface {
	return c.b
}

func NewC(B b.BInterface, A a.AInterface) *C {
	var _ CInterface = (*C)(nil)
	return &C{
		b: B,
		a: A,
	}
}
