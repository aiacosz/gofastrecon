package resolvers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"../utils"
)

type dns struct {
	Name  string `json: "name"`
	Ttl   string `json: "ttl"`
	Type  string `json: "type"`
	Value string `json: "value"`
}

func getDNS(body []byte) *dns {
	var d = new(dns)
	err := json.Unmarshal(body, &d)
	if err != nil {
		fmt.Println(utils.Red("[getDNS]"))
	}
	return d
}

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

func DNSApi(url string) {
	hostname := stripURL(url)
	req, err := http.Get("https://dns-api.org/NS/" + hostname)
	if err != nil {
		fmt.Println(utils.Red("[DNSApi] err"))
		log.Fatal(err)
	}
	defer req.Body.Close()

	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(utils.Red("[DNSApi] err"))
		log.Fatal(err)
	}
	teste := getDNS(bodyBytes)
	bodyString := string(bodyBytes)
	fmt.Println(teste.Name)
	fmt.Println(bodyString)
}
