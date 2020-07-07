package crawler

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func PrintLinks(links map[string]void) {
	fmt.Println("LINKS TO CRAWL ⛓")
	for link := range links {
		fmt.Printf("\t➡️  %v\n", link)
	}
	fmt.Println()
}

func PrintStaticAssets(staticAssets map[string]*[]string) {
	fmt.Println("ASSETS 🖼  👾 📹")
	for staticAsset, list := range staticAssets {
		fmt.Printf("\t%v\n", strings.ToUpper(staticAsset))
		for _, asset := range *list {
			fmt.Printf("\t\t➡️  %v\n", asset)
		}
	}
	fmt.Println("\n------------------------------------------------")
}

func OutputSitemap(assets []*html.Node) {
	formatted := FormatAssets(assets)
	PrintStaticAssets(formatted)
}
