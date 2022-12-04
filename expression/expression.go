package expression

type SymbolExpression string

const (
	SymbolOpenBracket  SymbolExpression = "("
	SymbolCloseBracket SymbolExpression = ")"
	SymbolNot          SymbolExpression = "!"
	SymbolAnd          SymbolExpression = "&&"
	SymbolOr           SymbolExpression = "||"
)
