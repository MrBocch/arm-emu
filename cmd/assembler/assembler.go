package assembler

import (
	"fmt"
)

// only support ascii
func Lex(code string) []string {
	// too bad golang has no sumtypes
	// (normal | comment)
	state := "normal"
	old_line := 0
	line := 1
	var lexemes []string

    for i := 0; i < len(code); i++ {
    	if old_line != line {
    		fmt.Printf("\n")
    		fmt.Printf("[%03d]: ", line)
    		old_line = line 
    	}
		if state == "comment" {
			start := line
			for j := i +1; j < len(code)-1; j++ {
				if code[j] == '\n' { line += 1 }
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
      	case '\n':
      		line += 1
       	case ';':
       	    for j := i+1; code[j] != '\n' && j < len(code); j++ {
       	    	i = j
       	    }
        case '/':
        	if i + 1 >= len(code) { panic("error '/''") } // should have actual system for reporting errors 
          	if code[i+1] == '/' {
          		for j := i+1; code[j] != '\n' && j < len(code); j++ { i = j }
          		continue 
          	}
          	if code[i+1] == '*' { state = "comment"; continue }
          	panic(fmt.Sprintf("[%d]: unexpected token '/'", line))

		default:
       		fmt.Printf("%c", byte)
       	}
       	
    }

	return lexemes
}
