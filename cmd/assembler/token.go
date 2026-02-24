package assembler

type TType int

const (
	Identifier TType = iota
	Instruction
	Register
	BitNumber
	HexNumber
	Number

	StringLiteral
	// one char tokens
	Plus
	Minus
	Colon
	LeftBracket
	RightBracket
	LeftCurly
	RightCurly

	Dot // maybe this should not be its own token

	Comma
	Hash
	NewLine
)

type Token struct {
	Kind   TType
	Lexeme string
	Line   int
}
