package resolvers

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"../utils"
)

func stripURL(url string) string {
	if strings.Contains(url, "https://") {
		newURL := strings.Replace(url, "https://", "", -1)
		return newURL
	}
	if strings.Contains(url, "http://") {
		newURL := strings.Replace(url, "http://", "", -1)
		return newURL
	}
	//goHorse x)
	return url
}

// ResolveURL resolv if true goes.. on if not..
func ResolveURL(url string) {
	hostname := stripURL(url)
	IPAddr, err := net.ResolveIPAddr("ip", hostname)
	if err != nil {
		fmt.Println(utils.Red("Error in resolving IP !!"))
		log.Fatal(err)
		os.Exit(1)
	}

	network := IPAddr.Network()

	fmt.Printf("Address : %s \nNetwork name : %s \n", IPAddr.String(), network)

}

// LookupHost resolv a url
func LookupHost(url string) {
	hostname := stripURL(url)
	// lookupHost using local dns entries
	fmt.Printf(utils.Yellow("Resolving.. "))
	fmt.Println(hostname)

	adrs, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Println("[LookupHost] errr..")
		log.Fatal(err)
	}
	for i := 0; i < len(adrs); i++ {
		segments := strings.SplitAfter(adrs[i], " ")
		fmt.Print(utils.Yellow("IP:"))
		fmt.Printf("%s \n", segments)
	}
}
