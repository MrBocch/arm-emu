package vm

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
)

type Computer struct {
	registers []int
	mem       []int 
}

type model struct {
	counter int 
}

func initialModel() model {
	return model{ counter: 0, }
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    // Is it a key press?
    case tea.KeyPressMsg:
        // Cool, what was the actual key pressed?
        switch msg.String() {
        // These keys should exit the program.
        case "q":
            return m, tea.Quit
        case "up", "k":
			m.counter++
        case "down", "j":
			m.counter--
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() tea.View {
    s := fmt.Sprintf("\n\ncount: #%d", m.counter)
    s += fmt.Sprintf("\n(press q to quit)")
	
    // Send the UI for rendering
    return tea.NewView(s)
}

func Run() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
