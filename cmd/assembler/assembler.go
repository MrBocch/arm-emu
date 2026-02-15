package assembler

import (
	"fmt"
)

// only support ascii
func Lex(code string) []Token {
	// too bad golang has no sumtypes
	// (normal | comment)
	state := "normal"
	line := 1
	var tokens []Token 
	currLexeme := ""

    for i := 0; i < len(code); i++ {
		if state == "comment" {
			start := line
			for j := i + 1; j < len(code)-1; j++ {
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
      		if currLexeme != "" { addToken(&tokens, Identifier, currLexeme, line) }
      		currLexeme = ""
       	    for j := i+1; code[j] != '\n' && j < len(code); j++ { i = j }
        case '/':
        	if i + 1 >= len(code) { panic("error '/''") } // should have actual system for reporting errors 
          	if code[i+1] == '/' {
          	    if currLexeme != "" { addToken(&tokens, Identifier, currLexeme, line) }
          	    currLexeme = ""
          		for j := i+1; code[j] != '\n' && j < len(code); j++ { i = j }
          		continue 
          	}
          	if code[i+1] == '*' {
          	    if currLexeme != "" { addToken(&tokens, Identifier, currLexeme, line) }
          	    currLexeme = ""
          	    state = "comment"
          	    continue
          	 }
          	panic(fmt.Sprintf("[%d]: unexpected token '/'", line))

        case '\n':
      		if currLexeme != "" { addToken(&tokens, Identifier, currLexeme, line) }
      		currLexeme = ""
      		if len(tokens) == 0 {
       			line += 1
       			continue
       		}
       		if tokens[len(tokens)-1].Kind == NewLine {
       			line += 1
       			continue 
       		}
      		addToken(&tokens, NewLine, "", line)
      		line += 1

      	case '+', '-', ':', '[', ']', '.', ',', '#':
			if currLexeme != "" { addToken(&tokens, Identifier, currLexeme, line)}
			t := tokenTypeFromByte(byte)
			currLexeme = ""
			addToken(&tokens, t, "", line)

      	case ' ', '\t':
      		if currLexeme == "" { continue }
      		addToken(&tokens, Identifier, currLexeme, line)
      		currLexeme = ""

      	case '"':
      		// currLexeme should be empty? idk
        	stringLiteral := ""
			for j := i + 1; j < len(code)-1; j++ {
				// TODO add escaping strings support \"
				i = j
				if code[j] == '"' { break }
				stringLiteral += string(code[j])
			}
        	if i + 1 >= len(code) { panic("unclosed string") } // should have actual system for reporting errors 
			addToken(&tokens, StringLiteral, stringLiteral, line)

		default:
			currLexeme += string(byte) 
       		if false { fmt.Printf("%c", byte) }
       	}
       	
    }

	return tokens
}

func addToken(tokens *[]Token, kind TType, lexeme string, line int) {
    *tokens = append(*tokens, Token{Kind: kind, Line: line, Lexeme: lexeme})
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
		case Colon:
	        return "COLON"
	    case StringLiteral:
	    	return "STRINGLITERAL"
		
		}
		return "err"
	}
	for _, t := range tokens {
		k := typetoString(t)
		l := t.Line
		lex := t.Lexeme 
		if t.Kind == StringLiteral {
			fmt.Println("* STRING LITERAL")
			fmt.Printf("* [%s]", lex)
			fmt.Println("")
		} else {
			if lex != "" { fmt.Printf("(%s %s %d)\n", k, lex, l) }
			if lex == "" { fmt.Printf("(%s %d)\n", k, l) }
			if k == "NEWLINE" { fmt.Println() }		
		}

	}
}

func tokenTypeFromByte(b byte) TType {
    switch b {
    case '+':
        return Plus
    case '-':
        return Minus
    case ':':
        return Colon
    case '[':
        return LeftBracket
    case ']':
        return RightBracket
    case '.':
        return Dot
    case ',':
        return Comma
    case '#':
        return Hash
    }
    // should return err in this case
    return Identifier
}

