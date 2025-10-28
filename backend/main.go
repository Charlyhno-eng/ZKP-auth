package main

import (
	"go-zkp/internal/usb"
	"runtime"
)

func main() {
	usb.DetectUSB(runtime.GOOS)
}
