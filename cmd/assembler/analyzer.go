package assembler

import (
    "fmt"
)

func Analyze(t []Token) {
	var line [] Token
	for _, k := range t {
		if k.Kind != NewLine { line = append(line, k) }
		if k.Kind == NewLine {
			fmt.Printf("Analyze %v\n", line)
			line = line[:0]
		}
	}
}

