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
	currLexeme := ""

    for i := 0; i < len(code); i++ {
    	if false && old_line != line { fmt.Printf("\n"); fmt.Printf("[%03d]: ", line); old_line = line }

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
      	// handle the case you run into a "a string : that might contain \n"
      	// cant imagine it would be too hard, just like a comment just that you add to thing
       	case ';':
      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
      		currLexeme = ""
       	    for j := i+1; code[j] != '\n' && j < len(code); j++ { i = j }
        case '/':
        	if i + 1 >= len(code) { panic("error '/''") } // should have actual system for reporting errors 
          	if code[i+1] == '/' {
	      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
	      		currLexeme = ""
          		for j := i+1; code[j] != '\n' && j < len(code); j++ { i = j }
          		continue 
          	}
          	if code[i+1] == '*' {
	      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
	      		currLexeme = ""
          		state = "comment";
          		continue
          	}
          	panic(fmt.Sprintf("[%d]: unexpected token '/'", line))

        case '\n':
       		if len(tokens) == 0 {
       			line += 1
       			continue
       		}
       		if tokens[len(tokens)-1].Kind == NewLine {
       			line += 1
       			continue 
       		}

      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
      		currLexeme = ""
      		tokens = append(tokens, Token{Kind: NewLine, Line: line})
      		line += 1

      	case '.':
      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
      		currLexeme = ""
      		tokens = append(tokens, Token{Kind: Dot, Line: line})
      	case ',':
      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
      		currLexeme = ""
      		tokens = append(tokens, Token{Kind: Comma, Line: line})
      	case ':':
      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
      		currLexeme = ""
      		tokens = append(tokens, Token{Kind: Colon, Line: line})
      	case '#':
      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
      		currLexeme = ""
      		tokens = append(tokens, Token{Kind: Hash, Line: line})
      	case '-':
      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
      		currLexeme = ""
      		tokens = append(tokens, Token{Kind: Minus, Line: line})
      	case '+':
      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
      		currLexeme = ""
      		tokens = append(tokens, Token{Kind: Plus, Line: line})
      	case '[':
      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
      		currLexeme = ""
      		tokens = append(tokens, Token{Kind: RightBracket, Line: line})
      	case ']':
      		if currLexeme != "" { tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})}
      		currLexeme = ""
      		tokens = append(tokens, Token{Kind: LeftBracket, Line: line})
      	case ' ':
      		if currLexeme == "" { continue }
      		tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})
      		currLexeme = ""
      	case '\t':
      		if currLexeme == "" { continue }
      		tokens = append(tokens, Token{Kind: Identifier, Line: line, Lexeme: currLexeme})
      		currLexeme = ""

		default:
			currLexeme += string(byte) 
       		if false { fmt.Printf("%c", byte) }
       	}
       	
    }

	return tokens
}

func PrintTokens(tokens []Token) {
	typetoString := func (token Token) string {
		switch token.Kind {
		case Identifier: 	
			return "IDENTIFIER"
		case Number:
			return "NUMBER"
		case Plus:
			return "PLUS"
		case Minus:
			return "MINUS"
		case LeftBracket:
			return "LEFTBRACKET"
		case RightBracket:
			return "RIGHTBRACKET"
		case Dot:
			return "DOT"
		case Comma:
			return "COMMA"
		case Hash:
			return "HASH"
		case ZeroB:
			return "0Bit"
		case NewLine:
			return "NEWLINE"
		case Colon:  // because switch not exhaustive, ran into issue bc i forgor
			return "COLON"
		
		}
		return "err"
	}
	for _, t := range tokens {
		k := typetoString(t)
		l := t.Line
		lex := t.Lexeme 
		if lex != "" { fmt.Printf("(%s %s %d)\n", k, lex, l) }
		if lex == "" { fmt.Printf("(%s %d)\n", k, l) }
	}
}
