package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomAgent return one of theses agents randomly and to by pass waf
func RandomAgent() string {
	rand.Seed(time.Now().Unix())
	var ua = []string{
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 1.1.4322; 2345Explorer 5.0.0.14136)",
		"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.108 Safari/537.36 2345Explorer/7.1.0.12633",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3) Gecko/20090305 Firefox/3.1b3 GTB5",
		"Mozilla/5.0 (X11; Ubuntu; Linux armv7l; rv:17.0) Gecko/20100101 Firefox/17.0",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0",
	}

	return ua[rand.Intn(len(ua))]
}

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
