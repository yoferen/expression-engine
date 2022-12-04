package main

import (
	expression "expression-engine/expression"
	"fmt"
)

func main() {
	infixExp := []string{"A", "&&", "(", "B", "||", "C", ")", "&&", "D"}
	operandValues := map[string]bool{
		"A": true,
		"B": false,
		"C": true,
		"D": true,
	}
	fmt.Println("Input infix expression: ", infixExp)
	fmt.Println("Input operand values: ", operandValues)
	postfixExp, _ := expression.ConvertInfixToPostfixLogicalExpression(infixExp)
	fmt.Println("Postfix expression: ", postfixExp)
	postfixEval, _ := expression.EvaluatePostfixLogicalExpression(postfixExp, operandValues)
	fmt.Println("Postfix evaluation: ", postfixEval)
}
