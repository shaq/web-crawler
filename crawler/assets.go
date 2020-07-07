package crawler

import (
	"golang.org/x/net/html"
	"strings"
)

type void struct{}

var ValidAssetTags = map[string]void{
	"a":      {},
	"script": {},
	"img":    {},
	"image":  {},
}

func ExtractAssets(domNode *html.Node) []*html.Node {
	if _, isValidTag := ValidAssetTags[domNode.Data]; domNode.Type == html.ElementNode && isValidTag {
		return []*html.Node{domNode}
	}

	var assets []*html.Node
	for childNode := domNode.FirstChild; childNode != nil; childNode = childNode.NextSibling {
		assets = append(assets, ExtractAssets(childNode)...)
	}

	return assets
}

// For each asset type, store a (pointer to) slice of strings containing each asset.
func FormatAssets(assets []*html.Node) map[string]*[]string {
	formatted := make(map[string]*[]string)
	for _, asset := range assets {
		for _, attr := range asset.Attr {
			key := attr.Key
			tag := asset.Data
			val := attr.Val
			if tag != "a" && key == "src" && len(val) > 0 {
				if _, exists := formatted[tag]; exists {
					*formatted[tag] = append(*formatted[tag], attr.Val)
				} else {
					formatted[tag] = &[]string{attr.Val}
				}
			}
		}
	}
	return formatted
}

func ExtractLinks(assets []*html.Node, baseURL, hostName string) map[string]void {
	links := make(map[string]void)
	for _, asset := range assets {
		for _, attr := range asset.Attr {
			if attr.Key == "href" {
				link, sameDomain := handleLinks(attr.Val, baseURL, hostName)
				if _, exists := links[attr.Val]; !exists && sameDomain {
					links[link] = void{}
				}
			}
		}
	}
	return links
}

func handleLinks(link, baseURL, hostName string) (string, bool) {
	if strings.HasPrefix(link, "/") {
		return baseURL + link, true
	} else if !strings.HasPrefix(link, "mail") {
		return link, strings.Contains(link, hostName)
	} else if strings.HasPrefix(link, "data") || strings.HasPrefix(link, "#") {
		return "", false
	}
	return link, false
}
