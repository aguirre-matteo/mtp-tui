package root

import (
  "github.com/charmbracelet/bubbles/list"
  "github.com/charmbracelet/lipgloss"
  "github.com/aguirre-matteo/mtp-tui/device"
)



type Model struct {
  Devices []device.Device
  List    list.Model
  Keys    listKeyMap
  Style   lipgloss.Style
}



func NewModel(mountpoint string) (Model, error) {
  devices, err := device.GetDevices(mountpoint)
  if err != nil {
    return Model{}, err
  }

  keys := newListKeyMap()
  l := newList(devices, keys)
  style := newStyle()

  return Model{
    Devices: devices,
    List: l,
    Keys: keys,
    Style: style,
  }, nil
}



