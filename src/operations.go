package main

import "fmt"

// Operation implements a common function for a polymorphic structure of calculations
type Operation interface {
	Calculate() int
}

// Sum type implements a sums simulation
type Sum struct {
	operand1, operand2 int
}

// Subtraction type implements a subtractions simulation
type Subtraction struct {
	operand1, operand2 int
}

func main() {
	operations := make([]Operation, 4)
	operations[0] = Sum{10, 20}
	operations[1] = Subtraction{30, 15}
	operations[2] = Subtraction{10, 50}
	operations[3] = Sum{5, 2}

	accumulator := 0
	for _, op := range operations {
		value := op.Calculate()
		fmt.Printf("%v = %d\n", op, value)
		accumulator += value
	}

	fmt.Println("Accumulated value: = ", accumulator)
}

// Calculate - implements the sum calculation
func (s Sum) Calculate() int {
	return s.operand1 + s.operand2
}

func (s Sum) String() string {
	return fmt.Sprintf("%d + %d", s.operand1, s.operand2)
}

// Calculate - implements the subtraction calculation
func (s Subtraction) Calculate() int {
	return s.operand1 - s.operand2
}

func (s Subtraction) String() string {
	return fmt.Sprintf("%d - %d", s.operand1, s.operand2)
}
