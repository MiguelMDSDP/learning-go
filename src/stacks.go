package main

import (
	"errors"
	"fmt"
)

func main() {
	stack := Stack{}
	fmt.Println("Created stack size: ", stack.Size())
	fmt.Println("Empty? ", stack.Empty())

	stack.Push("Go")
	stack.Push(2009)
	stack.Push(3.14)
	stack.Push("End")
	fmt.Println("Size after pushing 4 elements: ", stack.Size())
	fmt.Println("Empty? ", stack.Empty())

	for !stack.Empty() {
		v, _ := stack.Pop()
		fmt.Println("Poping: ", v)
		fmt.Println("Size: ", stack.Size())
		fmt.Println("Empty? ", stack.Empty())
	}

	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

// Stack - new data type to simulate stacks
type Stack struct {
	values []interface{}
}

// Size returns the size of the stack
func (stack Stack) Size() int {
	return len(stack.values)
}

// Empty tells us if the stack is empty or not
func (stack Stack) Empty() bool {
	return len(stack.values) == 0
}

// Push pushes a value into the stack
func (stack *Stack) Push(value interface{}) {
	stack.values = append(stack.values, value)
}

// Pop verifies if the stack is empty, if it is, returns an error, if not, pops an element
func (stack *Stack) Pop() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("empty stack")
	}
	value := stack.values[stack.Size()-1]
	stack.values = stack.values[:stack.Size()-1]
	return value, nil
}
