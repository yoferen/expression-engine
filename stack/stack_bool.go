package stack

import "errors"

type StackBool interface {
	Stack
	Push(value bool)
	Pop() (bool, error)
	Peek() (bool, error)
}

type stackBool struct {
	stack []bool
	index int
}

func NewStackBool() StackBool {
	startIndex := -1
	return &stackBool{
		stack: []bool{},
		index: startIndex,
	}
}

func (s *stackBool) IsEmpty() bool {
	if s.index < 0 {
		return true
	}
	return false
}

func (s *stackBool) GetLength() int {
	return s.index + 1
}

func (s *stackBool) Push(value bool) {
	s.index++
	if s.index == len(s.stack) {
		s.stack = append(s.stack, value)
	} else {
		s.stack[s.index] = value
	}
}

func (s *stackBool) Pop() (bool, error) {
	if s.IsEmpty() {
		return false, errors.New("empty stack")
	}
	val := s.stack[s.index]
	s.index--
	return val, nil
}

func (s *stackBool) Peek() (bool, error) {
	if s.IsEmpty() {
		return false, errors.New("empty stack")
	}
	return s.stack[s.index], nil
}
