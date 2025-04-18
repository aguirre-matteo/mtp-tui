package device

import (
  "fmt"
  "os"
  "os/exec"
)



type Device struct {
  Bus        string
  Id         string
  Name       string
  Mounted    bool
  Mountpoint string
}



func (dev Device) Title() string {
  return dev.Name
}



func (dev Device) Description() string {
  s := fmt.Sprintf("Bus ID: %v", dev.Bus)
  s += fmt.Sprintf(" | Device ID: %v", dev.Id)
  if dev.Mounted {
    s += " | " + "Mounted at " + dev.Mountpoint
  }
  return s
}



func (dev Device) FilterValue() string {
  return dev.Name
}



func (dev *Device) Mount() error {
  if dev.Mounted {
    err := fmt.Errorf("Device %v (%v/%v) is already mounted", dev.Name, dev.Bus, dev.Id)
    return err
  }

  if _, err := os.Stat(dev.Mountpoint); !os.IsNotExist(err) {
    err = fmt.Errorf("Fatal: wanted to mount device %v (%v/%v), but mountpoint %v already exists", dev.Name, dev.Bus, dev.Id, dev.Mountpoint)
    return err
  }

  err := os.Mkdir(dev.Mountpoint, 0755)
  if err != nil {
    return err
  }
  
  devflag := "-device=" + dev.Bus + "," + dev.Id
  options := "default_permissions,allow_other"
  
  cmd := exec.Command("jmtpfs", devflag, "-o", options, dev.Mountpoint)
  err = cmd.Run()
  if err != nil {
    return err
  }

  dev.Mounted = true
  return nil
}



func (dev *Device) Umount() error {
  if !dev.Mounted {
    err := fmt.Errorf("Device %v (%v/%v) is not mounted", dev.Name, dev.Bus, dev.Id)
    return err
  }

  if _, err := os.Stat(dev.Mountpoint); os.IsNotExist(err) {
    err = fmt.Errorf("Fatal: wanted to umount device %v (%v/%v) but mountpoint %v does not exists", dev.Name, dev.Bus, dev.Id, dev.Mountpoint)
    return err
  }

  cmd := exec.Command("fusermount", "-u", dev.Mountpoint)
  err := cmd.Run()
  if err != nil {
    return err
  }

  err = os.Remove(dev.Mountpoint)
  if err != nil {
    return err
  }

  dev.Mounted = false
  return nil
}



func (dev *Device) ToggleMount() error {
  if dev.Mounted {
    err := dev.Umount()
    return err
  } else {
    err := dev.Mount()
    return err
  }
  return nil
}
