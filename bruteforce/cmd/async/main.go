package main

import "bruteforce/lib"

func main() {
	password := "@k4l114"
	foundPassword := lib.FindPasswordAsync(password)
	if foundPassword != "" {
		println("Password found:", foundPassword)
	} else {
		println("Password not found")
	}
}
