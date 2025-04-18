package app

import (
  "github.com/aguirre-matteo/mtp-tui/app/root"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/bubbles/key"
  "github.com/spf13/viper"
)



type model struct {
  mountpoint string
  root       root.Model
  errors     []error
  width      int 
  height     int
}



func InitialModel(mountpoint string) (model, error) {
  var errors []error

  rt, err := root.NewModel(mountpoint)
  if err != nil {
    return model{}, nil
  }

  return model{
    mountpoint: mountpoint,
    root: rt,
    errors: errors,
  }, nil
  
}



func (m *model) Refresh() error {
  var err error
  m.root, err = root.NewModel(m.mountpoint)
  if err != nil {
    return err
  }

  m.root.List.SetSize(m.width, m.height)
  return nil
}



func (m model) Init() tea.Cmd {
  return nil
}



func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:

    switch {
    case key.Matches(msg, m.root.Keys.Refresh):
      err := m.Refresh()
      if err != nil {
        panic(err)
      }

    case key.Matches(msg, m.root.Keys.ToggleMount):
      if len(m.root.Devices) > 0 {
        index := m.root.List.Index()
        err := m.root.Devices[index].ToggleMount()
        if err != nil {
          panic(err)
        }

        err = m.Refresh()
        if err != nil {
          panic(err)
        }
      }
    }

  case tea.WindowSizeMsg:
	  h, v := m.root.Style.GetFrameSize()
    fwidth := msg.Width-h
    fheight := msg.Height-v

		m.root.List.SetSize(fwidth, fheight) 
	  m.width = fwidth
    m.height = fheight
  }

  var cmd tea.Cmd
	m.root.List, cmd = m.root.List.Update(msg)
	return m, cmd
}


func (m model) View() string {
  s := m.root.Style.Render(m.root.List.View())
  return s
}



func Run() error {
  mountpoint := viper.GetString("mount.point")
  initial, err := InitialModel(mountpoint)
  if err != nil {
    return err
  }

  p := tea.NewProgram(initial)
  if err != nil {
    return err
  }

  if _, err := p.Run(); err != nil {
    return err
  }
  return nil
}
