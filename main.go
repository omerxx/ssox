package main

import (
	"fmt"
	"omerxx/ssox/components/tabs"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Tabs tabs.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.Tabs.Update(msg)
}

func (m model) View() string {
	return m.Tabs.View()
}

func main() {
	tabsModel := tabs.New()
	tabsModel.Tabs = []string{"Local Profiles", "Available Accounts", "Something Else"}
	tabsModel.TabContent = []string{"Lab", "Blush Tab", "Eye Shadow Tab", "Mascara Tab", "Foundation Tab"}
	m := model{Tabs: tabsModel}
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
