package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	platform := GetPlatform()
	// Command to clear the screen depended on the platform
	// Determine the command to clear the screen based on the platform
	c := ""
	switch platform {
	case "windows":
		c = "cls"
	case "mac":
		c = "clear"
	case "linux":
		c = "clear"
	default:
		fmt.Println("Unknown platform, cannot clear screen.")
		return
	} // switch
	cmd := exec.Command(c)
	cmd.Stdout = os.Stdout
	cmd.Run()
} // clearScreen

func GetPlatform() string {
	os := runtime.GOOS
	var platform string
	switch os {
	case "windows":
		fmt.Println("Windows")
		platform = "windows"
	case "darwin":
		fmt.Println("MAC operating system")
		platform = "mac"
	case "linux":
		fmt.Println("Linux")
		platform = "linux"
	default:
		fmt.Printf("%s.\n", os)
		platform = "unknown"
	}
	return platform
} // getPlatform
