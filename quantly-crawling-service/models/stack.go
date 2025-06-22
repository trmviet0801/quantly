package models

type Stack[T any] struct {
	data []T
}

func (stack *Stack[T]) Push(value T) {
	stack.data = append(stack.data, value)
}

func (stack *Stack[T]) Pop() (T, bool) {
	var result T
	if len(stack.data) == 0 {
		return result, false
	}
	last := len(stack.data) - 1
	val := stack.data[last]
	stack.data = stack.data[:last]
	return val, true
}

func (stack *Stack[T]) Peek() (T, bool) {
	var result T
	if len(stack.data) == 0 {
		return result, false
	}
	return stack.data[len(stack.data)-1], true
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.data) == 0
}
