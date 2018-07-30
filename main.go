package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"./cms"
	"./crawler"
	"./resolvers"
	"./utils"
)

func validadeInput(url string) {
	if url == "" {
		fmt.Println("Provide URL to start !")
		fmt.Println(utils.Red("go run main.gon -url domain.com"))
		os.Exit(0)
	}
}

func validadeURL(url string) {
	if !strings.Contains(url, "http://") && !strings.Contains(url, "https://") {
		fmt.Println(utils.Red("Provide a correct schema.. https://example.com or http://example.com"))
		os.Exit(0)
	}
}

func main() {

	url := flag.String("url", "", "url from target")
	flag.Parse()

	validadeInput(*url)
	validadeURL(*url)

	start := time.Now()
	fmt.Println(start.Format("3:04PM"))
	utils.Banner()
	resolvers.ResolveURL(*url)
	resolvers.LookupHost(*url)
	// resolvers.DNSNS(*url) buggie
	crawler.Resolve(*url)
	crawler.GetCommonHeaders(*url)
	resolvers.DNSApi(*url)
	cms.WPScan(*url)
}
