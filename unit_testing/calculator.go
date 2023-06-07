package unit_testing

type Calculator struct {
	a int
	b int
}

type CalculatorInterface interface {
	Add() int
	Substract() int
	Multiply() int
	Divine() int
}

func (c *Calculator) Multiply() int {
	return c.a * c.b
}
func (c *Calculator) Add() int {
	return c.a + c.b
}
func (c *Calculator) Substract() int {
	return c.a - c.b
}
func (c *Calculator) Divine() int {
	return c.a / c.b
}
