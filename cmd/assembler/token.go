package assembler

type TType int

const (
    Identifier TType = iota
    Number

    Plus
    Minus
    Colon    
    LeftBracket
    RightBraket
    Dot
    Comma
    Hash
    ZeroB
    Dash
    NewLine
)

type Token struct {
    Kind   TType
    Lexeme string
    Line   int 
}



