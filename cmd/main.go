package main

import (
	"os"
	"fmt"
	"github.com/MrBocch/arm-emu/cmd/assembler"
	"github.com/MrBocch/arm-emu/cmd/vm"
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

	//vm.RunTui(bin)
	vm.RunGui(bin)

}

func printBinary(bin []uint32) {
	for _, instr := range bin {
	    b := fmt.Sprintf("%032b", instr)
	    fmt.Printf("%s %s %s %s %s\n", b[0:8], b[8:12], b[12:16], b[16:20], b[20:32])
	}
}

func printBinOp(bin []uint32) {
	for _, istr := range(bin){
		pack, err := (assembler.Decode(istr))
		if err != nil {
			panic("error on decoding")
		}
		fmt.Println(pack)
	}
}
