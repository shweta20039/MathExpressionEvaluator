package main

func NewSolution() *Solution {
	return &Solution{}
}

func main() {
	expression := "3+(2*2)-4/2"
	sol := NewSolution()
	result := sol.Calculate(expression)
	println("Result:", result)
}
