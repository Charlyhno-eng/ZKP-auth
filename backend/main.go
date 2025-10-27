package main

import (
	"go-zkp/usb"
	"runtime"
)

func main() {
	usb.DetectUSB(runtime.GOOS)
}
