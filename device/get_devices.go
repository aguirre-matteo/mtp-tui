package device

import (
	//"fmt"
	"slices"
)

func GetDevices(mountpoint string) ([]Device, error) {
	jdevs, err := getJmtpfsDevices(mountpoint)
	if err != nil {
		return []Device{}, err
	}

	mdevs, err := getMountedDevices(mountpoint)
	if err != nil {
		return []Device{}, err
	}

	var devices []Device
	if len(jdevs) == 0 {
		devices = mdevs
	} else {
		var appendedDevicesIndexes []int
		for _, jdev := range jdevs {
			for i, mdev := range mdevs {
				if jdev.Id == mdev.Id && jdev.Bus == mdev.Bus {
					jdev.Mounted = true
					appendedDevicesIndexes = append(appendedDevicesIndexes, i)
					continue
				}

				if !slices.Contains(appendedDevicesIndexes, i) {
					devices = append(devices, mdev)
					appendedDevicesIndexes = append(appendedDevicesIndexes, i)
				}
			}
			devices = append(devices, jdev)
		}
	}
	return devices, nil
}
