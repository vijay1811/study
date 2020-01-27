package stack

import "errors"

// Stack ...
type Stack struct {
	arr []int
}

// ErrStackEmpty if given in push and pop method when stack is empty
var ErrStackEmpty = errors.New("stack is empty")

// NewStack gives a stack
func NewStack(size int) *Stack {
	return &Stack{
		arr: make([]int, 0, size),
	}
}

// Push pushes an element to stack
func (s *Stack) Push(e int) {
	s.arr = append(s.arr, e)
}

// Pop pops an element from stack
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrStackEmpty
	}
	e := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return e, nil
}

// Top returns top element from the stack
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, ErrStackEmpty
	}
	return s.arr[len(s.arr)-1], nil
}

// IsEmpty will return true if there is not element present is stack
func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}
