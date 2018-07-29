package crawler

import (
	"fmt"
	"net/http"

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
