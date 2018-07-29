package utils

import (
	"fmt"
	"time"
)

const name = "Aiac0sz"
const version = "1.0"

// Red returns red color
func Red(str string) string {
	red := "\033[0;31m"
	reset := "\033[0m"
	return red + str + reset
}

// Green return green color
func Green(str string) string {
	green := "\033[0;32m"
	reset := "\033[0m"
	return green + str + reset
}

// Blue return blue color
func Blue(str string) string {
	blue := "\033[0;34m"
	reset := "\033[0m"
	return blue + str + reset
}

// Yellow return yellow color
func Yellow(str string) string {
	yellow := "\033[0;33m"
	reset := "\033[0m"
	return yellow + str + reset
}

//Banner prints out the banner of application
func Banner() {

	time.Sleep(1 * time.Second)
	fmt.Println("  ▄████  ▒█████    █████▒▄▄▄        ██████ ▄▄▄█████▓ ██▀███  ▓█████  ▄████▄   ▒█████   ███▄    █ ")
	fmt.Println("██▒ ▀█▒▒██▒  ██▒▓██   ▒▒████▄    ▒██    ▒ ▓█  ██▒ ▓▒▓██ ▒ ██▒▓█   ▀ ▒██▀ ▀█  ▒██▒  ██▒ ██ ▀█   ██")
	fmt.Println("▒██░▄▄▄░▒██░  ██▒▒████ ░▒██  ▀█▄  ░ ▓██▄   ▒ ▓██░ ▒░▓██ ░▄█ ▒▒███   ▒▓█    ▄ ▒██░  ██▒▓██  ▀█ ██▒")
	fmt.Println("░▓█  ██▓▒██   ██░░▓█▒  ░░██▄▄▄▄██   ▒   ██▒░ ▓██▓ ░ ▒██▀▀█▄  ▒▓█  ▄ ▒▓▓▄ ▄██▒▒██   ██░▓██▒  ▐▌██▒")
	fmt.Println("░▒▓███▀▒░ ████▓▒░░▒█░    ▓█   ▓██▒▒██████▒▒  ▒██▒ ░ ░██▓ ▒██▒░▒████▒▒ ▓███▀ ░░ ████▓▒░▒██░   ▓██░")
	time.Sleep(1 * time.Second)
	fmt.Println(" ░▒   ▒ ░ ▒░▒░▒░  ▒█░    ▒▒   ▓▒█░▒ ▒▓▒ ▒ ░  ▒ ░░   ░ ▒▓ ░▒▓░░░ ▒░ ░░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░   ▒ ▒ ")
	fmt.Println("  ░   ░   ░ ▒ ▒░  ░       ▒   ▒▒ ░░ ░▒  ░ ░    ░      ░▒ ░ ▒░ ░ ░  ░  ░  ▒     ░ ▒ ▒░ ░ ░░   ░ ▒░")
	fmt.Println("░ ░   ░ ░ ░ ░ ▒   ░ ░     ░   ▒   ░  ░  ░    ░        ░░   ░    ░   ░        ░ ░ ░ ▒     ░   ░ ░ ")
	fmt.Println("      ░     ░ ░               ░  ░      ░              ░        ░  ░░ ░          ░ ░           ░ ")
	fmt.Println(Yellow(name))
	fmt.Println(Red(version))
}
