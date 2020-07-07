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

	assets, err := crawler.Crawl(baseURL, client)
	if err != nil {
		log.Println(err)
	}

	links := crawler.ExtractLinks(assets, baseURL, hostName)
	crawler.PrintLinks(links)
	crawler.OutputSitemap(assets)

	for link := range links {
		a, err := crawler.Crawl(link, client)
		if err != nil {
			log.Fatalln(err)
		}
		if a != nil {
			crawler.OutputSitemap(a)
		}
	}
}
