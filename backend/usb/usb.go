package usb

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func DetectUSB(goos string) {
	fmt.Println("Detecting USB devices...")

	switch goos {
	case "linux":
		detectLinuxUSB()
	case "windows":
		detectWindowsUSB()
	default:
		fmt.Println("Operating system not currently supported:", goos)
	}
}

func detectLinuxUSB() {
	hasUSB, mounts := checkUSB()
	if hasUSB {
		fmt.Println("USB device(s) detected:")
		for _, mount := range mounts {
			authKeyPath := filepath.Join(mount, "auth_key")
			info, err := os.Stat(authKeyPath)
			if err != nil || !info.IsDir() {
				fmt.Println("  'auth_key' folder does not exist on", mount)
				continue
			}

			fmt.Println("Content of", authKeyPath, ":")
			files, err := listFiles(authKeyPath)
			if err != nil {
				fmt.Println("  Error reading:", err)
			} else if len(files) == 0 {
				fmt.Println("  Folder is empty")
			} else {
				for _, f := range files {
					fmt.Println("  -", f)
				}
			}
		}
	} else {
		fmt.Println("No USB device detected.")
	}
}

func detectWindowsUSB() {
	letters := []string{"E:", "F:", "G:", "H:", "I:", "J:"}
	found := false
	for _, letter := range letters {
		volume := letter + "\\"
		authKeyPath := filepath.Join(volume, "auth_key")
		info, err := os.Stat(authKeyPath)
		if err == nil && info.IsDir() {
			found = true
			fmt.Println("USB device detected on", volume)
			fmt.Println("Content of", authKeyPath, ":")
			files, err := listFiles(authKeyPath)
			if err != nil {
				fmt.Println("  Error reading:", err)
			} else if len(files) == 0 {
				fmt.Println("  Folder is empty")
			} else {
				for _, f := range files {
					fmt.Println("  -", f)
				}
			}
		}
	}
	if !found {
		fmt.Println("No USB device detected or no 'auth_key' folder present.")
	}
}

func checkUSB() (bool, []string) {
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

func listFiles(dirPath string) ([]string, error) {
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
