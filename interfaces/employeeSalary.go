package interfaces

import "fmt"

type Calculator interface {
	Calculate()
}

type FullTimeEmployee struct {
	Name   string
	Salary int
}

func (s *FullTimeEmployee) Calculate() {
	s.Name = "Akash"
	s.Salary = 35000

}

type ContractEmployee struct {
	Name   string
	Salary int
}

func (c *ContractEmployee) Calculate() {
	c.Name = "Muj"
	c.Salary = 20000

}

func Company() {
	fte := &FullTimeEmployee{}
	ce := &ContractEmployee{}

	company := []Calculator{fte, ce}

	for _, v := range company {
		v.Calculate()
		fmt.Println(v)
	}
}
