package main

import (
	"os"
	"fmt"
	"github.com/MrBocch/arm-emu/cmd/assembler"
	//tea "github.com/charmbracelet/bubbletea"
	//"github.com/charmbracelet/lipgloss"
)

//type model struct {}
//func initialModel() model {return model{}}
// func (m model) Init() tea.Cmd {return nil}
// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {return m, nil}
//func (m model) View() string {return style.Render(s)}

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
	assembler.Lex(string(file))

}
