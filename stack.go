package main

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

func (stack *Stack[T]) Push(element T) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack[T]) Peek() T {
	var element T
	if !stack.IsEmpty() {
		element := stack.elements[len(stack.elements)-1]
		return element
	}
	return element
}

func (stack *Stack[T]) Pop() T {
	element := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return element
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.elements) == 0
}

func (stack *Stack[T]) Size() int {
	return len(stack.elements)
}
