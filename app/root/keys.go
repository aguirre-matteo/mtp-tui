package root

import (
  "github.com/charmbracelet/bubbles/key"
)



type listKeyMap struct {
  ToggleMount key.Binding
  Refresh     key.Binding
}



func newListKeyMap() listKeyMap {
  return listKeyMap{
    ToggleMount: key.NewBinding(
      key.WithKeys("enter"),
      key.WithHelp("󰌑", "mount/umount"),
    ),
    Refresh: key.NewBinding(
      key.WithKeys("r"),
      key.WithHelp("r", "refresh"),
    ),
  }
}
