package errors

import (
	"fmt"
)

var (
	JmtpfsNotFound = fmt.Errorf("Command 'jmtpfs' not found. Make sure it is installed on your system.")
	FuseNotFound   = fmt.Errorf("Command 'fusermount' not found. Make sure FUSE is installed on your system.")
)

func DeviceAlreadyMounted(name, bus, id string) error {
	msg := fmt.Sprintf("Device %v (%v/%v) is already mounted.\n", name, bus, id)
	return fmt.Errorf(msg)
}

func DeviceMountpointAlreadyExists(name, bus, id, mountpoint string) error {
	msg := fmt.Sprintf("Fatal: wanted to mount device %v (%v/%v), but its mountpoint already exists.\n", name, bus, id, mountpoint)
	msg += fmt.Sprintf("Check %v. If it's empty, remove it. Otherwise, try running 'fusermount -u %v'.", mountpoint, mountpoint)
	return fmt.Errorf(msg)
}

func DeviceNotMounted(name, bus, id string) error {
	msg := fmt.Sprintf("Device %v (%v/%v) is not mounted", name, bus, id)
	return fmt.Errorf(msg)
}

func DeviceMountpointDoesntExists(name, bus, id, mountpoint string) error {
	msg := fmt.Sprintf("Fatal: wanted to umount device %v (%v/%v), but mountpoint %v doesn't exists.", name, bus, id, mountpoint)
	return fmt.Errorf(msg)
}

func WrongDeviceStringFormat(str string) error {
	msg := fmt.Sprintf("Wrong device string format: %v\n", str)
	msg += "Expected format: *, *, *, *, *, *"
	return fmt.Errorf(msg)
}

func WrongDirnameFormat(dirname string) error {
	msg := fmt.Sprintf("Wrong directory name fomat: %v\n", dirname)
	msg += "Expected format: *_*"
	return fmt.Errorf(msg)
}

func ConfigPathRelative(path string) error {
	msg := "Expected an absolute path to config folder.\n"
	msg += fmt.Sprintf("Got %v\n", path)
	return fmt.Errorf(msg)
}

func ConfigFileNotFound(path string) error {
	msg := fmt.Sprintf("The config folder %v doens't exists.\n", path)
	return fmt.Errorf(msg)
}
