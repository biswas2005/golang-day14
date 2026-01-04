package interfaces

import "fmt"

type Calculation interface {
	Execute(a, b int) int
}

type Addtion struct{}

func (Addtion) Execute(a, b int) int {
	return a + b
}

type Multiply struct{}

func (Multiply) Execute(a, b int) int {
	return a * b
}

func Calculate(op Calculation, a, b int) {
	result := op.Execute(a, b)
	fmt.Println("Result:", result)
}

func Calculator2() {
	add := Addtion{}
	mul := Multiply{}

	Calculate(add, 2, 3)
	Calculate(mul, 3, 4)
}
