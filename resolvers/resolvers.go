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
	TTL   string `json: "ttl"`
	Type  string `json: "type"`
	Value string `json: "value"`
}

func stripNullJSON(body string) []string {
	newArray := strings.Split(body, "[")
	//fmt.Println(newArray)
	return newArray
}

func getDNS(body []byte) {
	// var d = new(dns)
	var t []dns
	err := json.Unmarshal(body, &t)
	if err != nil {
		log.Println(err)
		fmt.Println(utils.Red("[getDNS]"))
	}
	fmt.Printf("[+] DNS INFO: %v\n", t)
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
	fmt.Println(hostname)
	IPAddr, err := net.ResolveIPAddr("ip", hostname)
	if err != nil {
		fmt.Println(utils.Red("Error in resolving IP !!"))
		log.Print(err)
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
		log.Print(err)
	}
	for i := 0; i < len(adrs); i++ {
		segments := strings.SplitAfter(adrs[i], " ")
		fmt.Print(utils.Yellow("IP:"))
		fmt.Printf("%s \n", segments)
	}
}

//DNSApi go to dns-api and retrieve someinfo about dns
func DNSApi(url string) {

	hostname := stripURL(url)
	req, err := http.Get("https://dns-api.org/NS/" + hostname)
	if err != nil {
		fmt.Println(utils.Red("[DNSApi] err"))
		log.Print(err)
	}
	defer req.Body.Close()

	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(utils.Red("[DNSApi] err"))
		log.Print(err)
	}
	stripNullJSON(string(bodyBytes))
	jsonData := []byte(string(bodyBytes))
	getDNS(jsonData)
}
