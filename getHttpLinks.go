package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"golang.org/x/net/html"
)

func getHref(t html.Token) (ok bool, href string) {
	//Goes over the attributes of the tokens until it finds a "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	return
}

//Gets all the http links from a webpage
func crawl(url string, ch chan string, chFinished chan bool) {
	resp, err = http.Get(url)
	defer func(){
		chFinished <- true //Notifies done running
	}
	if err != nil{ //Checks for non null error code
		fmt.Println("Error: Program failed to crawl "+url)
		return
	}

	b = resp.Body
	defer b.Close() //Close when the body is returned

	z := html.NewTokenizer(b)

	for{
		tt := z.Next()
		switch{
			case tt == html.ErrorToken: //End of the document
				return
			case tt == html.StartTagToken:
				t := z.Token()
				isAnchor := t.Data == "a"
				if !isAnchor{ //Checks to see if an anchor tag has been found
					continue
				}
				ok, url := getHref(t)
				if !ok{ //Checks to see if the
					continue
				}
				hasProto := strings.Index(url, "http") == 0 //Makes sure the url starts with http
				if hasProto{
					ch <- url
				}
		}
	}
}

func main(){
	founfUrls := make(map[string]bool)
	seedUrls := os.Args[1:]

	//Different chanels
	chUrl := make(chan string)
	chFinished := make(chan bool)

	//Starts concurrent crawl process
	for _, url := range seedUrls {
		go crawl(url, chUrls, chFinished)
	}

	//Sub to chanels
	for c := 0, c < len(seedUrls); {
		select {
		case url := <-chUrls:
			foundUrls[url] = true
		case <-chFinished:
			c++
		}
	}

	fmt.Println("\nFound", len(foundUrls), "unique urls:\n")

	//Prints the urls that were found
	for url, _ := range foundUrls {
		fmt.Println(" - " + url)
	}

	close(chUrls)  //Closes the chanel
}