package assembler

type TType int

const (
    Identifier TType = iota
    Number

    Plus
    Minus
    Colon    
    LeftBracket
    RightBracket
    LeftCurly
    RightCurly
    Dot
    Comma
    Hash
    ZeroB
    NewLine
    StringLiteral
)

type Token struct {
    Kind   TType
    Lexeme string
    Line   int 
}



