package stack

type Stack struct {
	depth int
}

func New() *Stack {
	return &Stack{}
}
