package crawler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"../utils"
)

var commonFiles = []string{
	"robots.txt",
	"sitemap.xml",
	"sitemap.xml.gz",
	"crossdomain.xml",
	"phpinfo.php",
	"test.php",
	"elmah.axd",
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
	"../../etc/passwd",
}

//Resolve GOHORSE
func Resolve(url string) {
	sendRequests(commonFiles, url)
}

//SendRequests recive array of strings and make requests to endpoints
func sendRequests(endpoints []string, url string) {

	for _, endpoint := range endpoints {
		schema := url + "/" + endpoint

		req, err := http.Get(schema)
		if err != nil {
			fmt.Println(utils.Red("[sendRequests] err.."))
			panic(err)
		}
		defer req.Body.Close()

		//bodyBytes, err := ioutil.ReadAll(req.Body)
		//if err != nil {
		//	fmt.Println(utils.Red("[sendRequests] err.."))
		//	panic(err)
		//}

		if req.StatusCode == 200 {
			fmt.Print(utils.Yellow("[+] Found! \n"))
			fmt.Printf(utils.Green(schema + "\n"))
			//bodyString := string(bodyBytes)
			//fmt.Println(bodyString)
		}
	}

}

// GetCommonHeaders get headers of response
func GetCommonHeaders(url string) {
	fmt.Println(utils.Yellow("[+] Try to get some headers"))
	req, err := http.Get(url)
	if err != nil {
		fmt.Println(utils.Red("[GetCommonHeaders] err"))
		log.Fatal(err)
	}

	headers := req.Header
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
