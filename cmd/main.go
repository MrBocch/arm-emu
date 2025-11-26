package main

import (
    // "os"
	"fmt"
	"github.com/MrBocch/arm-emu/cmd/assembler"
	//"github.com/charmbracelet/lipgloss"
	//tea "github.com/charmbracelet/bubbletea"
)

//type model struct {}
//func initialModel() model {return model{}}
// func (m model) Init() tea.Cmd {return nil}
// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {return m, nil}
//func (m model) View() string {return style.Render(s)}

func main(){
	fmt.Println("hello")
	assembler.Test()
    // p := tea.NewProgram(initialModel())
    // if _, err := p.Run(); err != nil {
}
