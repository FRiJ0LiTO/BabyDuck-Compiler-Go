package stack

import "fmt"

type Node struct {
	Value any
	Next  *Node
}

type Stack struct {
	Top    *Node
	Height int
}

// Push adds a new element to the top of the stack.
// It creates a new Node with the given value and sets it as the new top of the stack.
// The height of the stack is incremented by one.
//
// Parameters:
//
//	value (any): The value to be added to the stack.
func (stack *Stack) Push(value any) {
	newNode := &Node{Value: value, Next: stack.Top}
	stack.Top = newNode
	stack.Height++
}

// Pop removes and returns the top element of the stack.
// If the stack is empty, it returns nil.
// The height of the stack is decremented by one.
//
// Returns:
//
//	any: The value of the top element of the stack, or nil if the stack is empty.
func (stack *Stack) Pop() any {
	if stack.Height == 0 {
		return nil
	}
	value := stack.Top.Value
	stack.Top = stack.Top.Next
	stack.Height--
	return value
}

// Peek returns the value of the top element of the stack without removing it.
// If the stack is empty, it returns nil.
//
// Returns:
//   - any: The value of the top element of the stack, or nil if the stack is empty.
func (stack *Stack) Peek() any {
	if stack.Height == 0 {
		return nil
	}
	return stack.Top.Value
}

// ToStringSlice converts all stack values to a string slice.
// The values are returned in order from top to bottom.
// Note: This function assumes all values in the stack are strings.
// It will panic if any value cannot be type-asserted to string.
//
// Returns:
//   - []string: A slice containing all stack values as strings.
func (stack *Stack) ToStringSlice() []string {
	var result []string
	current := stack.Top
	for current != nil {
		result = append(result, current.Value.(string))
		current = current.Next
	}
	return result
}

// New creates and returns a new instance of the Stack struct.
// This function initializes an empty stack with no elements.
//
// Returns:
//   - *Stack: A pointer to the newly created Stack instance.
func New() *Stack {
	return &Stack{}
}

// PrintStack prints all elements in the stack from top to bottom.
// It iterates through the nodes starting from the top and prints each node's value.
func (stack *Stack) PrintStack() {
	temp := stack.Top
	for temp != nil {
		fmt.Println("Valor", temp.Value)
		temp = temp.Next
	}
}
