package device

import (
	"github.com/aguirre-matteo/mtp-tui/errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getJmtpfsOutput() ([]string, error) {
	cmd := exec.Command("jmtpfs", "-l")
	rawOutput, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}

	output := string(rawOutput)
	lines := strings.Split(output, "\n")

	start := -1
	var result []string
	for i, line := range lines {
		if strings.Contains(line, "Available devices (busLocation, devNum, productId, vendorId, product, vendor):") {
			start = i + 1
			break
		}
	}

	if start == -1 {
		return []string{}, nil
	}

	for _, line := range lines[start:] {
		if line != "" {
			result = append(result, line)
		}
	}
	return result, nil
}

func stringToDevice(str, mountpoint string) (Device, error) {
	fields := strings.Split(str, ", ")
	if len(fields) != 6 {
		return Device{}, errors.WrongDeviceStringFormatError(str)
	}

	dirname := fields[0] + "_" + fields[1]
	finalMnt := filepath.Join(mountpoint, dirname)

	device := Device{
		Bus:        fields[0],
		Id:         fields[1],
		Name:       fields[4],
		Mountpoint: finalMnt,
	}
	return device, nil
}

func getMountedDevices(mountpoint string) ([]Device, error) {
	files, err := os.ReadDir(mountpoint)
	if err != nil {
		return []Device{}, err
	}

	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	var devices []Device
	for _, dirname := range dirs {
		parts := strings.Split(dirname, "_")
		if len(parts) != 2 {
			return []Device{}, errors.WrongDirnameFormatError(dirname)
		}

		dirname := parts[0] + "_" + parts[1]
		finalMnt := filepath.Join(mountpoint, dirname)

		device := Device{
			Bus:        parts[0],
			Id:         parts[1],
			Name:       "Unknown",
			Mounted:    true,
			Mountpoint: finalMnt,
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func getJmtpfsDevices(mountpoint string) ([]Device, error) {
	output, err := getJmtpfsOutput()
	if err != nil {
		return []Device{}, err
	}

	var devices []Device
	for _, line := range output {
		device, err := stringToDevice(line, mountpoint)
		if err != nil {
			return []Device{}, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}
