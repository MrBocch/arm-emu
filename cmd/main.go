package main

import (
	"os"
	"fmt"
	"github.com/MrBocch/arm-emu/cmd/assembler"
	// "github.com/MrBocch/arm-emu/cmd/vm"
)

func main(){
	if len(os.Args) == 1 {
		fmt.Println("How to use me")
		fmt.Println("arm-emu [filepath]")
		os.Exit(1)
	}

	path := os.Args[1]
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Print("Could not find: ")
		fmt.Println("[" + path + "]")
		os.Exit(1)
	}
	tokens := assembler.Lex(string(file))
	// assembler.PrintTokens(tokens)

	bin, err := assembler.Analyze(tokens)
	if err != nil {
		fmt.Println("Fix error: ", err)
		os.Exit(1)
	}

	fmt.Println("encoded: ")
	fmt.Println(bin)
	// vm.Run(bin)

}
