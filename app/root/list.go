package root

import (
	"github.com/aguirre-matteo/mtp-tui/device"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

func newList(devices []device.Device, keys listKeyMap) list.Model {
	middlelist := make([]list.Item, len(devices))
	for i, dev := range devices {
		middlelist[i] = dev
	}

	l := list.New(middlelist, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Available MTP Devices"
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			keys.ToggleMount,
			keys.Refresh,
		}
	}

	return l
}
