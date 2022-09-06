package b

import "wire_best_practice/a"

type (
	B struct {
		A  a.AInterface
		X1 int
	}

	BInterface interface {
		Show()
	}
)

func (b *B) Show() {
	b.A.Show()
}

func NewB(A a.AInterface) *B {
	var _ BInterface = (*B)(nil)
	return &B{
		A: A,
	}
}
