package main

import "bruteforce/lib"

func main() {
	password := "@k4l1_l0l"
	foundPassword := lib.FindPasswordByCombinationAsync(password)
	if foundPassword != "" {
		println("Password found:", foundPassword)
	} else {
		println("Password not found")
	}
}
