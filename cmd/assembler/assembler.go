package assembler

import (
	"fmt"
)

// only support ascii
func Lex(code string) []string {
	// too bad golang has no sumtypes
	// normal
	// comment
	state := "normal"
	// nested := 0
	var lexemes []string
    for i := 0; i < len(code); i++ {
      	byte := code[i]
      	switch byte {
       	case ';': state = "comment"; continue
        case '/':
        	if i + 1 >= len(code) {
         		panic("unended comment") // should have actual system for reporting errors
         	}
          	if code[i+1] == '/' { state = "comment"; continue }
       	}

        if state == "comment" && byte == '\n' { state = "normal"; continue }
        if state != "comment" {
	        fmt.Printf("Index %d: byte=%d char=%c\n", i, byte, byte)
        }
    }

	return lexemes
}
