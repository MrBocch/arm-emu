package assembler

import (
	"fmt"
)

// only support ascii
func Lex(code string) []Token {
	// too bad golang has no sumtypes
	// (normal | comment)
	state := "normal"
	old_line := 0
	line := 1
	var tokens []Token 
	// currLexeme := ""

    for i := 0; i < len(code); i++ {
    	if false && old_line != line {
    		fmt.Printf("\n")
    		fmt.Printf("[%03d]: ", line)
    		old_line = line 
    	}
		if state == "comment" {
			start := line
			for j := i +1; j < len(code)-1; j++ {
				if code[j] == '\n' { line += 1 }
				if code[j] == '/' && code[j+1] == '*' {
					panic(fmt.Sprintf("[%d]: Can't nest multiline comments", line))
				}
				if code[j] == '*' && code[j+1] == '/' {
					i = j + 2 
					state = "normal"
					continue 
				}
			}
			if state != "normal" { panic(fmt.Sprintf("[%d]: unclosed comment", start)) }
		} 
		
      	byte := code[i]
      	switch byte {
       	case ';':
       	    for j := i+1; code[j] != '\n' && j < len(code); j++ {
       	    	i = j
       	    }

       	case '\n':
       		if len(tokens) == 0 {
       			line += 1
       			continue
       		}
       		if tokens[len(tokens)-1].Kind == NewLine {
       			line += 1
       			continue 
       		}
      		tokens = append(tokens, Token{Kind: NewLine, Line: line})
      		line += 1

      	case ',':
      		tokens = append(tokens, Token{Kind: Comma, Line: line})
      	case ':':
      		tokens = append(tokens, Token{Kind: Colon, Line: line})
      	case '#':
      		tokens = append(tokens, Token{Kind: Hash, Line: line})
      	case '-':
      		tokens = append(tokens, Token{Kind: Minus, Line: line})
      	case '+':
      		tokens = append(tokens, Token{Kind: Minus, Line: line})

        case '/':
        	if i + 1 >= len(code) { panic("error '/''") } // should have actual system for reporting errors 
          	if code[i+1] == '/' {
          		for j := i+1; code[j] != '\n' && j < len(code); j++ { i = j }
          		continue 
          	}
          	if code[i+1] == '*' { state = "comment"; continue }
          	panic(fmt.Sprintf("[%d]: unexpected token '/'", line))

		default:
       		if false { fmt.Printf("%c", byte) }
       	}
       	
    }

	return tokens
}
