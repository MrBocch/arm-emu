package assembler

import (
	"fmt"
)

// only support ascii
func Lex(code string) []string {
	// too bad golang has no sumtypes
	// (normal | comment)
	state := "normal"
	// nested := 0
	var lexemes []string
    for i := 0; i < len(code); i++ {
		if state == "comment" {
			for j := i +1; j < len(code)-1; j++ {
				if code[j] == '*' && code[j+1] == '/' {
					i = j + 2 
					state = "normal"
					continue 
				}
			}
		} 
		
      	byte := code[i]
      	switch byte {
       	case ';':
       	    for j := i+1; code[j] != '\n' && j < len(code); j++ { i = j } ; continue 
        case '/':
        	if i + 1 >= len(code) { panic("unexpected token '/'") } // should have actual system for reporting errors 
          	if code[i+1] == '/' { for j := i+1; code[j] != '\n' && j < len(code); j++ { i = j } ; continue }
          	if code[i+1] == '*' { state = "comment"; continue }
          	panic("unexpected token '/'")
       	}

       	if state != "" {
		    fmt.Printf("Index %d: byte=%d char=%c\n", i, byte, byte)
       	}
    }

	return lexemes
}
