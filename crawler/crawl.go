package crawler

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
)

func CheckBaseURL(args []string) (string, string, error) {
	if len(args) > 1 {
		webURL := args[1]
		hostName, err := validateURL(webURL)
		if err != nil {
			return "", "", err
		}
		return hostName, webURL, err
	}
	return "", "", fmt.Errorf("no URL passed to crawl")
}

func validateURL(rawURL string) (string, error) {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return "", fmt.Errorf("%v is an invalid URL", rawURL)
	}
	return u.Hostname(), nil
}

func getDOM(body io.Reader) (*html.Node, error) {
	parentNode, err := html.Parse(body)
	if err != nil {
		log.Fatalln(err)
	}
	return parentNode, nil
}

func Crawl(url string, client *http.Client) ([]*html.Node, error) {
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Unable to get URL %v; error status code: %d\n%s\n", url, resp.StatusCode, resp.Status)
		return nil, nil
	}

	fmt.Printf("\nCrawling %v ... ⤵️\n", url)

	dom, err := getDOM(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	return ExtractAssets(dom), nil
}