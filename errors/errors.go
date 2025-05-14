package errors

import (
	"fmt"
)

var (
	JmtpfsNotFoundError = fmt.Errorf("Command 'jmtpfs' not found. Make sure it is installed on your system.")
	FuseNotFoundError   = fmt.Errorf("Command 'fusermount' not found. Make sure FUSE is installed on your system.")
)

func DeviceAlreadyMountedError(name, bus, id string) error {
	msg := fmt.Sprintf("Device %v (%v/%v) is already mounted.\n", name, bus, id)
	return fmt.Errorf(msg)
}

func DeviceMountpointAlreadyExistsError(name, bus, id, mountpoint string) error {
	msg := fmt.Sprintf("Fatal: wanted to mount device %v (%v/%v), but its mountpoint already exists.\n", name, bus, id, mountpoint)
	msg += fmt.Sprintf("Check %v. If it's empty, remove it. Otherwise, try running 'fusermount -u %v'.", mountpoint, mountpoint)
	return fmt.Errorf(msg)
}

func DeviceNotMountedError(name, bus, id string) error {
	msg := fmt.Sprintf("Device %v (%v/%v) is not mounted", name, bus, id)
	return fmt.Errorf(msg)
}

func DeviceMountpointDoesntExistsError(name, bus, id, mountpoint string) error {
	msg := fmt.Sprintf("Fatal: wanted to umount device %v (%v/%v), but mountpoint %v doesn't exists.", name, bus, id, mountpoint)
	return fmt.Errorf(msg)
}

func WrongDeviceStringFormatError(str string) error {
	msg := fmt.Sprintf("Wrong device string format: %v\n", str)
	msg += "Expected format: *, *, *, *, *, *"
	return fmt.Errorf(msg)
}

func WrongDirnameFormatError(dirname string) error {
	msg := fmt.Sprintf("Wrong directory name fomat: %v\n", dirname)
	msg += "Expected format: *_*"
	return fmt.Errorf(msg)
}

func ConfigPathRelativeError(path string) error {
	msg := "Expected an absolute path to config folder.\n"
	msg += fmt.Sprintf("Got %v\n", path)
	return fmt.Errorf(msg)
}

func ConfigFileNotFoundError(path string) error {
	msg := fmt.Sprintf("The config folder %v doens't exists.\n", path)
	return fmt.Errorf(msg)
}
