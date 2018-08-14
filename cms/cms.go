package cms

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
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
		Timeout:   5 * time.Second,
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", utils.RandomAgent())

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("[+] %s -- %d \n", url, resp.StatusCode)

	if resp.Header.Get("Location") != "" {
		fmt.Printf("[+] Redirect found: %s", url)
	}

}

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
		"/wp-config.php",
		"/wp-login",
	}

	for _, v := range defaultPaths {
		uri := url + v
		sendRequests(uri)

	}
}

// JoomScan searching for file like joomla
func JoomScan(url string) {
	var defaultPaths = [...]string{
		"/joomla/administrator",
		"/joomla/admin",
		"/configuration.php",
		"/administrator/index.php",
		"/joomla.xml",
		"/htaccess.txt",
		"/administrator/templates/hathor/LICENSE.txt",
		"/web.config.txt",
		"/robots.txt.dist",
		"/LICENSE.txt",
		"/media/jui/fonts/icomoon-license.txt",
		"/media/editors/tinymce/jscripts/tiny_mce/license.txt",
		"/media/editors/tinymce/jscripts/tiny_mce/plugins/style/readme.txt",
		"/libraries/idna_convert/ReadMe.txt",
		"/libraries/simplepie/README.txt",
		"/libraries/simplepie/LICENSE.txt",
		"/libraries/simplepie/idn/ReadMe.txt",
	}

	for _, v := range defaultPaths {
		uri := url + v
		sendRequests(uri)

	}
}

// DrupScan searching for common files drupal
func DrupScan(url string) {
	var defaultPaths = [...]string{
		"/modules/simpletest/files/php-1.txt",
		"/modules/simpletest/files/sql-1.txt",
		"/modules/simpletest/files/html-1.txt",
		"/modules/simpletest/tests/common_test_info.txt",
		"/modules/filter/tests/filter.url-output.txt",
		"/modules/filter/tests/filter.url-input.txt",
		"/modules/search/tests/UnicodeTest.txt",
		"/themes/README.txt",
		"/themes/stark/README.txt",
		"/sites/README.txt",
		"/sites/all/modules/README.txt",
		"/sites/all/themes/README.txt",
		"/modules/simpletest/files/html-2.html",
		"/modules/color/preview.html",
		"/themes/bartik/color/preview.html",
		"/UPGRADE.txt",
		"/CHANGELOG.txt",
		"/INSTALL.sqlite.txt",
		"/LICENSE.txt",
		"/INSTALL.txt",
		"/COPYRIGHT.txt",
		"/web.config",
		"/modules/README.txt",
		"/modules/simpletest/files/README.txt",
		"/modules/simpletest/files/javascript-1.txt",
	}

	for _, v := range defaultPaths {
		uri := url + v
		sendRequests(uri)

	}
}

// NodeCMSCommons common nodeJS files
func NodeCMSCommons(url string) {
	var defaultPaths = [...]string{
		"/api/controllers",
		"/node_modules",
		"/hooks",
		"/middleware",
		"/routes.json",
		"/api/controllers/index.js",
		"/api/controllers/index.js",
		"/packages.json",
		"/requests.json",
	}

	for _, v := range defaultPaths {
		uri := url + v
		sendRequests(uri)

	}

}
