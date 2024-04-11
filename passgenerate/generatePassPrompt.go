package passgenerate

import (
	"fmt"
	"os"
	"pass/functions"
	"strconv"

	// "pass/algos"

	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices   []string
	cursor    int
	selected  map[int]struct{}
	page      int
	textinput textinput.Model
}

var (
	length      string
	pass        []string
	helperStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7d7d7d"))
)

var style = lipgloss.NewStyle().Foreground(lipgloss.Color("#7d7d7d"))

func initialPassModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter the length of password"
	ti.Focus()
	ti.Prompt = "Length: "
	ti.CharLimit = 2
	ti.Width = 20
	return model{
		choices:   []string{"alphabets", "numbers", "special characters", "use (Copy to clipboard, don't print it here)"},
		selected:  make(map[int]struct{}),
		textinput: ti,
		page:      0,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tickMsg:
		m.page++

	case tea.KeyMsg:
		if m.page == 0 {
			switch msg.Type {
			case tea.KeyCtrlC, tea.KeyEsc:
				fmt.Println("Exiting")
				return m, tea.Quit
			case tea.KeyEnter:
				m.page++
				length = m.textinput.Value()
			}
			m.textinput, cmd = m.textinput.Update(msg)
			return m, cmd
		} else if m.page == 1 {
			switch msg.String() {

			case "ctrl+c", "q", "esc":
				return m, tea.Quit

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case "t":
				return m, tick()

			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}

			// The "enter" key and the spacebar (a literal space) toggle
			// the selected state for the item that the cursor is pointing at.
			case " ":
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}

			case "enter":
				m.page++
				lengthInt, _ := strconv.Atoi(length)
				_, isNum := m.selected[1]
				_, isSpecial := m.selected[2]

				pass = append(pass, GenerateRandomString(lengthInt, isNum, isSpecial))
				pass = append(pass, GenerateRandomString(lengthInt, isNum, isSpecial))
				pass = append(pass, GenerateRandomString(lengthInt, isNum, isSpecial))
			}
		} else if m.page == 2 {
			switch msg.String() {
			case "ctrl+c", "q", "esc":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < 2 {
					m.cursor++
				}
			case "enter":
				functions.CopyToClipboard(pass[m.cursor])
				return m, tea.Quit
			}
		} else {
			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	return m, nil
}

// func laj() tea.Msg {
// 	return tea.KeyMsg{
// 		Type:    tea.KeyRunes,
// 		Runes:   []rune{'l', 'a', 'j'},
// 		Modifiers: tea.ModAlt,
// 	}
// }

type tickMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m model) View() string {
	if m.page == 0 {
		helper := "••••••••••••••••••••••••••••••••••••••••••••••\n↑/k up • ↓/j down • enter choose • / filter • q quit • ? more"
		return fmt.Sprintf("%s \n\n\n\n\n\n\n\n\n\n\n\n\n%s", m.textinput.View(), helperStyle.Render(helper))
	} else if m.page == 1 {
		// The header
		s := "Select the things in password of length " + length + "\n\n"

		// Iterate over our choices
		for i, choice := range m.choices {

			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}

			// Is this choice selected?
			checked := " " // not selected
			if _, ok := m.selected[i]; ok {
				checked = "\033[33m*\033[0m" // selected! (colored green)
			}
			// Render the row
			if _, ok := m.selected[i]; ok {
				s += fmt.Sprintf("\033[36m%s [%s\033[36m] %s\033[0m\n", cursor, checked, choice) // selected! (colored green)
			} else {
				s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
			}

			// Render the row
			// s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
		}

		s += fmt.Sprint(helperStyle.Render("\n\n\n\n\n\n\n••••••••••••••••••••••••••••••••••••••••••••••\n↑/k up • ↓/j down • enter choose • / filter • q quit • ? more"))
		// The footer
		// s += "\nPress Esc or Ctrl+C to quit.\n"

		// Send the UI for rendering
		return s
	} else {
		s := "Select(Up/Down) & Enter to copy to clipboard\n\n"
		// cursor := " "
		for i := 0; i < 3; i++ {
			if m.cursor == i {
				s += fmt.Sprintf(">  %s\n", pass[i])
			} else {
				s += fmt.Sprintf("   %s\n", style.Render(pass[i]))
			}

		}
		return s
	}
}

func GeneratePassPrompt() {
	p := tea.NewProgram(initialPassModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
