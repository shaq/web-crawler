package main

import (
	"crypto/tls"
	"fmt"
	crawler "github.com/shaq/web-crawler/crawler"
	"log"
	"net/http"
	"os"
)

var (
	config = &tls.Config{
		InsecureSkipVerify: true,
	}
	transport = &http.Transport{
		TLSClientConfig: config,
	}
	client = &http.Client{
		Transport: transport,
	}
)

func main() {
	fmt.Print("Welcome to Shaq's Web Crawler! 🥳\n\n\n")
	hostName, baseURL, err := crawler.CheckBaseURL(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

	assets:= crawler.Crawl(baseURL, client)
	links := crawler.ExtractLinks(assets, baseURL, hostName)
	crawler.PrintLinks(links)
	crawler.OutputSitemap(assets)

	for link := range links {
		a := crawler.Crawl(link, client)
		if a != nil {
			crawler.OutputSitemap(a)
		}
	}
}
