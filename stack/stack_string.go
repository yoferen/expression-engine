package stack

import "errors"

type StackString interface {
	Stack
	Push(value string)
	Pop() (string, error)
	Peek() (string, error)
}

type stackString struct {
	stack []string
	index int
}

func NewStackString() StackString {
	startIndex := -1
	return &stackString{
		stack: []string{},
		index: startIndex,
	}
}

func (s *stackString) IsEmpty() bool {
	if s.index < 0 {
		return true
	}
	return false
}

func (s *stackString) GetLength() int {
	return s.index + 1
}

func (s *stackString) Push(value string) {
	s.index++
	if s.index == len(s.stack) {
		s.stack = append(s.stack, value)
	} else {
		s.stack[s.index] = value
	}
}

func (s *stackString) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("empty stack")
	}
	val := s.stack[s.index]
	s.index--
	return val, nil
}

func (s *stackString) Peek() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("empty stack")
	}
	return s.stack[s.index], nil
}
