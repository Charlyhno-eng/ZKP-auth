package usb

import (
	"bufio"
	"os"
	"strings"
)

func CheckUSB() (bool, []string) {
	var usbMounts []string
	file, err := os.Open("/proc/mounts")
	if err != nil {
		return false, usbMounts
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			mountPoint := fields[1]
			if strings.HasPrefix(mountPoint, "/media/") || strings.HasPrefix(mountPoint, "/run/media/") {
				usbMounts = append(usbMounts, mountPoint)
			}
		}
	}
	return len(usbMounts) > 0, usbMounts
}

func ListFiles(dirPath string) ([]string, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, entry := range entries {
		files = append(files, entry.Name())
	}
	return files, nil
}
