package assembler

import (
    "fmt"
    "strings"
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
      		if currLexeme != "" {
      			ttype := identifyLex(currLexeme)
      		    addToken(&tokens, ttype, currLexeme, line)
      		}
      		currLexeme = ""
       	    for j := i+1; code[j] != '\n' && j < len(code); j++ { i = j }
        case '/':
        	if i + 1 >= len(code) { panic("error '/''") } // should have actual system for reporting errors 
          	if code[i+1] == '/' {
          	    if currLexeme != "" {
          	    	ttype := identifyLex(currLexeme)
          	    	addToken(&tokens, ttype, currLexeme, line)
          	    }
          	    currLexeme = ""
          		for j := i+1; code[j] != '\n' && j < len(code); j++ { i = j }
          		continue 
          	}
          	if code[i+1] == '*' {
          	    if currLexeme != "" {
          	    	ttype := identifyLex(currLexeme)
          	    	addToken(&tokens, ttype, currLexeme, line)
          	    }
          	    currLexeme = ""
          	    state = "comment"
          	    continue
          	 }
          	panic(fmt.Sprintf("[%d]: unexpected token '/'", line))

        case '\n':
      		if currLexeme != "" {
      			ttype := identifyLex(currLexeme)
      			addToken(&tokens, ttype, currLexeme, line)
      		}
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

      	case '+', '-', ':', '[', ']', '{', '}','.', ',', '#':
			if currLexeme != "" {
				ttype := identifyLex(currLexeme)
				addToken(&tokens, ttype, currLexeme, line)
			}
			t := tokenTypeFromByte(byte)
			currLexeme = ""
			addToken(&tokens, t, "", line)

      	case ' ', '\t':
      		if currLexeme == "" { continue }
      		ttype := identifyLex(currLexeme)
      		addToken(&tokens, ttype, currLexeme, line)
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
		case LeftCurly:
			return "LEFTCURLY"
		case RightCurly:
			return "RIGHTCURLY"
		case Dot:
			return "DOT"
		case Comma:
			return "COMMA"
		case Hash:
			return "HASH"
		case NewLine:
			return "NEWLINE"
		case Colon:
	        return "COLON"
	    case StringLiteral:
	    	return "STRINGLITERAL"
	    case Instruction:
	    	return "OP"
	    case Register:
	    	return "REG"
	    case BitNumber:
	    	return "BITNUMBER"
	    case HexNumber:
	    	return "HexNumber"
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
    // would make more sense if it were named
    // opening and closing braces
    case '[':
        return LeftBracket
    case ']': 
        return RightBracket
    case '{':
    	return LeftCurly
    case '}':
    	return RightCurly
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

func identifyLex(lex string) TType {
	if len(lex) > 2 && lex[0:2] == "0b" { return BitNumber }
	if len(lex) > 2 && lex[0:2] == "0x" { return HexNumber }
	if isOp(lex) { return Instruction }
	if isDigit(lex[0]) { return Number } // if first char is a number, assume its a number
	if isRegister(lex) { return Register }

	return Identifier
}

func isDigit(s byte) bool {
	switch s {
	case '0','1','2', '3','4','5','6','7','8','9': { return true }
	}
	return false 
}

func isOp(op string) bool {
	switch strings.ToLower(op) {
	// its bad to combinar code + data, get this from somewhere else, dont want to
	// track several lists of ops 
	case "mov", "add", "sub","str","ldr","cmp", "beq", "b", "bgt", "push", "pop","bl", "adds","halt": { return true }
	}
	return false 
}

func isRegister(reg string) bool {
	switch strings.ToLower(reg) {
	case "r0","r1","r2","r3","r4", "r5", "r6", "r7","r8", "r9", "r10", "r11", "r12",
		 "r13","r14","r15","lr","sp","pc": { return true }
	}
	return false 
}
