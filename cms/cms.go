package cms

import (
	"fmt"
	"log"
	"net/http"

	"../utils"
)

// WPScan crawlling common paths and dorks for wordpress
func WPScan(url string) {
	var defaultPaths = [...]string{
		"/readme.html",
		"/license.txt",
		"/xmlrpc.php",
		"/wp-config-sample.php",
		"/wp-includes/images/crystal/license.txt",
		"/wp-includes/images/crystal/license.txt",
		"/wp-includes/js/plupload/license.txt",
		"/wp-includes/js/plupload/changelog.txt",
		"/wp-includes/js/tinymce/license.txt",
		"/wp-includes/js/tinymce/plugins/spellchecker/changelog.txt",
		"/wp-includes/js/swfupload/license.txt",
		"/wp-includes/ID3/license.txt",
		"/wp-includes/ID3/readme.txt",
		"/wp-includes/ID3/license.commercial.txt",
		"/wp-content/themes/twentythirteen/fonts/COPYING.txt",
		"/wp-content/themes/twentythirteen/fonts/LICENSE.txt",
		"/wp-config",
		"/wp-login",
	}

	for _, v := range defaultPaths {
		uri := url + v
		req, err := http.NewRequest("GET", uri, nil)
		if err != nil {
			fmt.Println(utils.Red("[WPScan] err"))
			log.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(utils.Red("[WPScan] err"))
			log.Fatal(err)
		}

		if resp.StatusCode == 200 {
			fmt.Printf("[+] found: %s \n", uri)
		}

	}
}
