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

// Check if the input string is valid integer
// Underscore char allowed as thousand separator
// ASCII codes of numbers are 48...57
// ASCII code of underscore char is 95
// ASCII Table -- Printable Characters:
// https://condor.depaul.edu/sjost/it236/documents/ascii.htm
func IsUint(s string) bool {
	ClearScreen()
	is_valid := false
	nu := GetNumOfUnderscores(s)
	switch nu {

	case 0: // no underscores in input string
		for i, c := range s {
			i++
			if c >= '0' && c <= '9' {
				is_valid = true
			} else {
				is_valid = false
				break
			}
		} // for
		return is_valid

	default:
		// There are underscores in input string.
		// Remove leading and trailing underscores.
		fmt.Println(s)
	jmp1:
		if s[0] == '_' {
			s = s[1:]
		}
		if s[len(s)-1] == '_' {
			s = s[:len(s)-1]
		}
		if s[0] == '_' || s[len(s)-1] == '_' {
			goto jmp1
		}

		s = ReverseString(s)
		fmt.Println(s)
		ucount := 0 // num of underscores in string
		nu = GetNumOfUnderscores(s)
		strlen := len(s)  // string length with underscores
		ns := strlen - nu // string length without underscores
		ureq := ns / 4    // required quantity of underscores
		is_valid = false
		if ureq != nu {
			return false
		}

		for i, c := range s {
			i++
			if c == '_' {
				ucount++
				remainder := i % 4
				if remainder == 0 { // multiple of 4
					is_valid = true // valid thousand separator
				} else {
					is_valid = false // invalid thousand separator
					break
				} // else
			}
		} // for
	} // switch
	return is_valid
} // IsUint

// Get number of underscores in string
func GetNumOfUnderscores(s string) int {
	counter := 0
	for _, c := range s { // index, char
		if c == '_' {
			counter++
		}
	}
	return counter
} // GetNumOfUnderscores

// Check if string only contains digits
func HasDigitsOnly(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
} // HasDigitsOnly

// This function returns reverse input string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
} // ReverseString

// Foreground colors
// Foreground text colors
var (
	FgBlack   = "\x1b[30m"
	FgRed     = "\x1b[31m"
	FgGreen   = "\x1b[32m"
	FgYellow  = "\x1b[33m"
	FgBlue    = "\x1b[34m"
	FgMagenta = "\x1b[35m"
	FgCyan    = "\x1b[36m"
	FgWhite   = "\x1b[37m"
	Reset     = "\x1b[0m" // reset
)

// Bright Foreground colors
var (
	BrightFgBlack   = "\x1b[90m"
	BrightFgRed     = "\x1b[91m"
	BrightFgGreen   = "\x1b[92m"
	BrightFgYellow  = "\x1b[93m"
	BrightFgBlue    = "\x1b[94m"
	BrightFgMagenta = "\x1b[95m"
	BrightFgCyan    = "\x1b[96m"
	BrightFgWhite   = "\x1b[97m"
)

// Foreground Colors Combined with Background Colors
// After each color change needed reset
var (
	FgBlackBgGreen   = "\x1b[30m\x1b[42m" // fg: black, bg: green\x1b[0m"); // OK
	FgBlackBgYellow  = "\x1b[30m\x1b[43m" // fg: black, bg: yellow\x1b[0m"); // OK
	FgBlackBgCyan    = "\x1b[30m\x1b[46m" // fg: black, bg: cyan\x1b[0m"); // OK
	FgBlackBgWhite   = "\x1b[30m\x1b[47m" // fg: black, bg: white\x1b[0m"); // OK
	FgWhiteBgBlue    = "\x1b[37m\x1b[44m" // fg: white, bg: blue\x1b[0m"); // OK
	FgWhiteBgMagenta = "\x1b[37m\x1b[45m" // fg: white, bg: magenta\x1b[0m"); // OK
	FgRedBgWhite     = "\x1b[31m\x1b[47m" // fg: red, bg: white\x1b[0m"); // OK
)
