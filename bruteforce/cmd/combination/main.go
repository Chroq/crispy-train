package main

import "bruteforce/lib"

func main() {
	password := "@kAl1"
	foundPassword := lib.FindPasswordByCombination(password)
	if foundPassword != "" {
		println("Password found:", foundPassword)
	} else {
		println("Password not found")
	}
}
