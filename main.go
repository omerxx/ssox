package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle      = lipgloss.NewStyle().Margin(1, 2)
	quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type mainModel struct {
	list    list.Model
	profile string
	region  string
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			log.Print("enter")
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.profile = i.Title()
				m.region = "eu-west-1"
				// var cmd tea.Cmd
				// return NewAccountModel(m.region), cmd
				command := exec.Command("aws", "sso", "login", "--profile", m.profile)
				command.Run()
			}
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m mainModel) View() string {
	if m.profile != "" {
		log.Print(m.profile)
		return ""
	}
	return "\n" + m.list.View()
}

func getProfiles() (items []list.Item) {
	home := os.Getenv("HOME")
	f := fmt.Sprintf("%s/.aws/config", home)
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile(`^\[.*\]$`)
	rregion, _ := regexp.Compile(`^region.*$`)
	var item item
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			if item.title != "" {
				items = append(items, item)
				item.desc = ""
			}

			result := strings.Trim(scanner.Text(), "[]")
			split := strings.Split(result, " ")
			item.title = split[1]
		} else if rregion.Match([]byte(scanner.Text())) {
			item.desc = item.desc + scanner.Text()
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	defer f.Close()

	items := getProfiles()
	myList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	myList.SetFilteringEnabled(true)
	myList.SetShowFilter(true)
	myList.SetShowStatusBar(false)
	myList.SetShowTitle(false)
	m := mainModel{list: myList}
	m.list.Title = "Configured Profiles"

	p := tea.NewProgram(m, tea.WithAltScreen())
	log.Println("starting...")

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
