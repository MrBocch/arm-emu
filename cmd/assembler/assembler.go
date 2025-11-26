package assembler

import (
	"fmt"
)

func Lex(code string) []rune {
	var lexemes []rune
    for i := 0; i < len(code); i++ {
        byte := code[i]
        fmt.Printf("Index %d: byte=%d char=%c\n", i, byte, byte)
    }
	return lexemes
}
