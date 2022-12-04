package expression

import (
	"errors"
	stackLib "expression-engine/stack"
)

var (
	isLogicalExpressionSymbol = map[SymbolExpression]bool{
		SymbolNot:          true,
		SymbolAnd:          true,
		SymbolOr:           true,
		SymbolOpenBracket:  true,
		SymbolCloseBracket: true,
	}
	isLogicalOperator = map[LogicalOperator]bool{
		LogicalNot: true,
		LogicalAnd: true,
		LogicalOr:  true,
	}
)

type LogicalOperator string

const (
	LogicalNot LogicalOperator = LogicalOperator(SymbolNot)
	LogicalAnd LogicalOperator = LogicalOperator(SymbolAnd)
	LogicalOr  LogicalOperator = LogicalOperator(SymbolOr)
)

func (lo LogicalOperator) GetPriority() int {
	switch lo {
	case LogicalNot:
		return 3
	case LogicalAnd:
		return 2
	case LogicalOr:
		return 1
	}
	return 0
}

func ConvertInfixToPostfixLogicalExpression(expression []string) ([]string, error) {
	res := []string{}
	stack := stackLib.NewStackString()
	for _, element := range expression {
		if !isLogicalExpressionSymbol[SymbolExpression(element)] {
			res = append(res, element)
		} else {
			symbol := SymbolExpression(element)
			if symbol == SymbolOpenBracket {
				stack.Push(element)
			} else if symbol == SymbolCloseBracket {
				for {
					val, err := stack.Peek()
					if err != nil {
						return nil, errors.New("invalid infix expression")
					}
					if val == string(SymbolOpenBracket) {
						_, err := stack.Pop()
						if err != nil {
							return nil, errors.New("invalid infix expression")
						}
						break
					}
					val, err = stack.Pop()
					if err != nil {
						return nil, errors.New("invalid infix expression")
					}
					res = append(res, val)
				}
			} else {
				if stack.IsEmpty() {
					stack.Push(element)
					continue
				}
				val, err := stack.Peek()
				if err != nil {
					return nil, errors.New("invalid infix expression")
				}
				if LogicalOperator(symbol).GetPriority() > LogicalOperator(val).GetPriority() {
					stack.Push(element)
				} else {
					for {
						val, _ := stack.Peek()
						if LogicalOperator(symbol).GetPriority() <= LogicalOperator(val).GetPriority() {
							val, err = stack.Pop()
							if err != nil {
								return nil, errors.New("invalid infix expression")
							}
							res = append(res, val)
						} else {
							stack.Push(element)
							break
						}
					}
				}
			}
		}
	}

	for {
		if stack.IsEmpty() {
			break
		}
		val, err := stack.Pop()
		if err != nil || val == string(SymbolOpenBracket) || val == string(SymbolCloseBracket) {
			return nil, errors.New("invalid infix expression")
		}
		res = append(res, val)
	}

	if len(res) == 0 {
		return nil, errors.New("empty infix expression")
	}

	if !isLogicalOperator[LogicalOperator(res[len(res)-1])] {
		return nil, errors.New("invalid infix expression")
	}

	return res, nil
}

func EvaluatePostfixLogicalExpression(expression []string, values map[string]bool) (bool, error) {
	var err error
	if len(expression) == 0 {
		return false, errors.New("empty postfix logical expression")
	}
	stack := stackLib.NewStackBool()
	for _, element := range expression {
		if !isLogicalExpressionSymbol[SymbolExpression(element)] {
			stack.Push(values[element])
		} else {
			operator := LogicalOperator(element)
			var first, second, res bool
			switch operator {
			case LogicalNot:
				first, err = stack.Pop()
				res = !first
			case LogicalAnd:
				first, err = stack.Pop()
				second, err = stack.Pop()
				res = first && second
			case LogicalOr:
				first, err = stack.Pop()
				second, err = stack.Pop()
				res = first || second
			}
			if err != nil {
				return false, errors.New("invalid postfix logical expression")
			}
			stack.Push(res)
		}
	}
	res, err := stack.Pop()
	if err != nil || !stack.IsEmpty() {
		return false, errors.New("invalid postfix logical expression")
	}
	return res, nil
}
