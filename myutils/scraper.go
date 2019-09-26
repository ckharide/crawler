package myutils

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
)

func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}

func ScrapeURL(url string, urlchannel chan string, chFinished chan bool) {
	fmt.Println("Scraping URL ", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to crawl the url ", url)
		return
	}

	defer func() {
		// Notify that we're done after this function
		chFinished <- true
	}()

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		//bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//bodyString := string(bodyBytes)
		//fmt.Println(bodyString)
	}

	fmt.Println("Scrape Started ")
	z := html.NewTokenizer(resp.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, url := getHref(t)
			//fmt.Println("Scrape URL Got ", url)
			if !ok {
				continue
			}

			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				urlchannel <- url
			}

		}
	}
}
