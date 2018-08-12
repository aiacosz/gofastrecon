package crawler

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"../utils"
)

func sendRequests(url string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   3 * time.Second,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Print(utils.Red("Sucess: "))
		fmt.Println(utils.Yellow(url))
	}
}

//CrawlerCommonFiles make requests to endpoints
func CrawlerCommonFiles(url string) {

	var commonFiles = []string{
		"robots.txt",
		"sitemap.xml",
		"sitemap.xml.gz",
		"crossdomain.xml",
		"phpinfo.php",
		"test.php",
		"elmah.axd",
		"graphql",
		"server-status",
		"jmx-console",
		"admin-console",
		"web-console",
		"swagger.json",
		"api.json",
		"trace.axd",
		"admin",
		"app.js",
		"aws.js",
	}

	for _, endpoint := range commonFiles {
		schema := url + "/" + endpoint
		sendRequests(schema)
		continue

	}

}

// GetCommonHeaders get headers of response
func GetCommonHeaders(url string) {
	fmt.Println(utils.Yellow("[+] Try to get some headers"))
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   3 * time.Second,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	headers := resp.Header
	for k, v := range headers {

		if strings.Contains(k, "X-Powered-By") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println()
		}

		if strings.Contains(k, "Server") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println()
		}

		if strings.Contains(k, "Access-Control-Allow-Origin") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println()
		}

		if strings.Contains(k, "Access-Control-Allow-Credential") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println()
		}

		if strings.Contains(k, "Access-Control-Allow-Methods") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println()
		}

		if strings.Contains(k, "X-XSS-Protection") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println(utils.Blue("Maybe some XSS atacks not be effective.."))
			fmt.Println()
		}

		if strings.Contains(k, "strict-transport-security") {
			fmt.Printf("%s  :  %s", k, v)
			fmt.Println(utils.Blue("HSTS attacks.. not be effective.."))
			fmt.Println()
		}

	}

}
