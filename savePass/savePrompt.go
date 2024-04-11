package savePass

import (
	"fmt"
	"unicode"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	inputs         []textinput.Model
	errVals        []string // Used for printing error or tip messages for each input field
	focused        int
	err            error
	tags           []string
	alternateNames []string
	passStrength   string
}

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("154"))
	dimStyle     = lipgloss.NewStyle().Foreground(darkGray).Align(lipgloss.Right)
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("178")).Align(lipgloss.Right)
	headingStyle = lipgloss.NewStyle().Foreground(hotPink)
	tagStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Background(lipgloss.Color("8")).PaddingLeft(1).PaddingRight(1)
	tipStyle     = lipgloss.NewStyle().Width(80).Height(4).Padding(1).Border(lipgloss.RoundedBorder(), true).Foreground(lipgloss.Color("205")).MarginTop(1)
)

type errMsg error

const (
	AliasName = iota
	Password
	AlternateNames
	Tags
	Desc
	Username
	Email
	URL
)

// initial
func initialSaveModel() model {
	var inputs []textinput.Model = make([]textinput.Model, 9)
	var errVals []string = make([]string, 9)

	inputs[AliasName] = textinput.New()
	inputs[AliasName].Placeholder = "Enter the name"
	inputs[AliasName].Focus()
	inputs[AliasName].CharLimit = 30
	inputs[AliasName].Width = 30
	inputs[AliasName].Prompt = ""
	inputs[AliasName].Validate = func(value string) error {
		err := isAlphanumeric(value)
		if err != nil {
			errVals[AliasName] = err.Error()
		} else {
			errVals[AliasName] = ""
		}
		return err
	}

	inputs[AlternateNames] = textinput.New()
	inputs[AlternateNames].Placeholder = "Enter alternate names (comma separated)"
	inputs[AlternateNames].CharLimit = 100
	inputs[AlternateNames].Width = 20
	inputs[AlternateNames].Prompt = ""
	// inputs[AlternateNames].Validate = func(value string) error {
	// 	if len(value) == 0 {
	// 		errVals[AlternateNames] = "alternate names cannot be empty"
	// 		return fmt.Errorf("alternate names cannot be empty")
	// 	} else {
	// 		errVals[AlternateNames] = ""
	// 	}
	// 	return nil
	// }

	inputs[Password] = textinput.New()
	inputs[Password].Placeholder = "Enter the password"
	inputs[Password].EchoMode = textinput.EchoPassword
	inputs[Password].CharLimit = 50
	inputs[Password].Width = 30
	inputs[Password].Prompt = ""
	inputs[Password].Validate = nil

	inputs[Tags] = textinput.New()
	inputs[Tags].Placeholder = "Use Commas to separate"
	inputs[Tags].CharLimit = 100
	inputs[Tags].Width = 15
	inputs[Tags].Prompt = ""
	inputs[Tags].Validate = nil
	inputs[Tags].SetSuggestions([]string{"tag1", "tag2", "tag3"})

	inputs[Desc] = textinput.New()
	inputs[Desc].Placeholder = "Enter description"
	inputs[Desc].CharLimit = 100
	inputs[Desc].Width = 40
	inputs[Desc].Prompt = ""
	inputs[Desc].Validate = nil

	// inputs[Strength] = textinput.New()
	// inputs[Strength].Placeholder = "Enter password strength"
	// inputs[Strength].CharLimit = 100
	// inputs[Strength].Width = 20
	// inputs[Strength].Prompt = ""
	// inputs[Strength].Validate = nil

	inputs[Username] = textinput.New()
	inputs[Username].Placeholder = "Username associated with password"
	inputs[Username].CharLimit = 100
	inputs[Username].Width = 32
	inputs[Username].Prompt = ""
	inputs[Username].Validate = nil

	inputs[Email] = textinput.New()
	inputs[Email].Placeholder = "Email associated with password"
	inputs[Email].CharLimit = 100
	inputs[Email].Width = 30
	inputs[Email].Prompt = ""
	inputs[Email].Validate = nil

	inputs[URL] = textinput.New()
	inputs[URL].Placeholder = "Enter the URL"
	inputs[URL].CharLimit = 100
	inputs[URL].Width = 20
	inputs[URL].Prompt = ""
	inputs[URL].Validate = nil

	return model{
		inputs:  inputs,
		focused: 0,
		errVals: errVals,
		err:     nil,
	}
}

// init
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyDown:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			} else if m.focused == 3 {
				if m.inputs[m.focused].Value() != "" {
					m.tags = append(m.tags, m.inputs[m.focused].Value())
					m.inputs[m.focused].Reset()
				} else {
					m.nextInput()
				}
			} else if m.focused == 2 {
				if m.inputs[m.focused].Value() != "" {
					m.alternateNames = append(m.alternateNames, m.inputs[m.focused].Value())
					m.inputs[m.focused].Reset()
				} else {
					m.nextInput()
				}
			} else {
				m.nextInput()
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP, tea.KeyUp:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		// case tea.KeySpace:
		// 	if m.focused == 4 {
		// 		m.tags = append(m.tags, m.inputs[m.focused].Value())
		// 		m.inputs[m.focused].Reset()
		// 	}
		case tea.KeyCtrlQ:
			if m.focused == 1 {
				if m.inputs[m.focused].EchoMode == textinput.EchoPassword {
					m.inputs[m.focused].EchoMode = textinput.EchoNormal
				} else {
					m.inputs[m.focused].EchoMode = textinput.EchoPassword
				}
			}
		}
		switch msg.String() {
		case ",":
			if m.focused == 3 {
				m.tags = append(m.tags, m.inputs[m.focused].Value())
				m.inputs[m.focused].Reset()
				return m, tea.Batch(cmds...)
			} else if m.focused == 2 {
				m.alternateNames = append(m.alternateNames, m.inputs[m.focused].Value())
				m.inputs[m.focused].Reset()
				return m, tea.Batch(cmds...)
			}
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	// Input Validation Conditions
	if m.inputs[AliasName].Value() == "" {
		m.errVals[AliasName] = "!Alias name cannot be empty"
	} else {
		m.errVals[AliasName] = ""
	}

	if m.inputs[Password].Value() != "" {
		m.errVals[Password] = "Use Ctrl+Q to view/hide password"
	} else {
		m.errVals[Password] = errorStyle.Render("!Password cannot be empty")
	}

	// Update focused input

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)

	// Handle focused input
}

// view

func (m model) View() string {
	// var b strings.Builder

	b := fmt.Sprintf(
		"%s\n  %s %s %s\n  %s: %s %s\n\n%s\n  %s: %s %s",
		headingStyle.Render("Required Values:"),
		inputStyle.Render("Alias Name:"),
		m.inputs[AliasName].View(),
		errorStyle.Render(m.errVals[AliasName]),
		inputStyle.Render("Password"),
		m.inputs[Password].View(),
		m.errVals[Password],
		headingStyle.Render("Optional Values: (Ctrl+Enter to skip)"),
		inputStyle.Render("Alternate Names"),
		m.inputs[AlternateNames].View(),
		dimStyle.Render(m.errVals[AlternateNames]),
		// inputStyle.Render("Tags"),
		// m.inputs[Tags].View(),
	)

	for _, err := range m.alternateNames {
		b += fmt.Sprintf(" %s", tagStyle.Render(err))
	}

	b += fmt.Sprintf("\n  %s: %s", inputStyle.Render("Tags"), m.inputs[Tags].View())

	for _, err := range m.tags {
		b += fmt.Sprintf(" %s", tagStyle.Render(err))
	}
	b += fmt.Sprintf(
		"\n  %s: %s\n  %s: %s\n  %s: %s\n  %s: %s",
		inputStyle.Render("Description"),
		m.inputs[Desc].View(),
		// inputStyle.Render("Strength"),
		// m.inputs[Strength].View(),
		inputStyle.Render("Username"),
		m.inputs[Username].View(),
		inputStyle.Render("Email"),
		m.inputs[Email].View(),
		inputStyle.Render("URL"),
		m.inputs[URL].View(),
	)

	// tags := strings.Join(m.tags, ", ")
	// tagOutput := tagStyle.Render(tags)
	// b += fmt.Sprintf("\n%s: %s", inputStyle.Render("Tags"), tagOutput)
	// for i, input := range m.inputs {
	// 	if i == m.focused {
	// 		b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(input.View()))
	// 	} else {
	// 		b.WriteString(input.View())
	// 	}
	// 	b.WriteString("\n\n")
	// }
	b += tipStyle.Render("Press Ctrl+Enter to skip optional values")
	return b
	// return b.String()
}

// savePrompt
func SavePrompt() {
	p := tea.NewProgram(initialSaveModel())
	if err := p.Start(); err != nil {
		panic(err)
	}
}

func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}

func isAlphanumeric(input string) error {
	for _, char := range input {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return fmt.Errorf("input must be alphanumeric")
		}
	}
	return nil
}
