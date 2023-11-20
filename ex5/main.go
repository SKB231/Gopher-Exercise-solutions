package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/SKB231/gopherex/ex4/linkParser"
)

var visited map[string]bool = make(map[string]bool)

type Url struct {
	Loc string `xml:"loc"`
}
type urlset struct {
	Xmlns string `xml:"xmlns,attr"`
	Url   []Url  `xml:"url"`
}

// DFS function takes a string as the root node.
// It assumes that the URL already contains the correct origin. It performs a GET request on the link, extracts all the urls in the page, and runs the DFS on those methods as well.
// Using a map, it ensures that it doesen't visit the same link twice.
func dfs(root string, origin string, maxDepth int) {
	if root == "" || maxDepth < 0 {
		return
	}

	visited[root] = true

	resp, err := http.Get(root)
	if err != nil {
		fmt.Println("Error when reading page ", root)
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := linkParser.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing file")
		fmt.Println(err)
		os.Exit(1)
	}
	for _, linkStruct := range res {
		// Ignore all empty links
		if linkStruct.Link == "" {
			// fmt.Println("Empty Link" + linkStruct.Text)
			continue
		}

		finalLink := ""
		if linkStruct.Link[0] == byte('/') {
			// This is a relative path
			urlRes, _ := url.JoinPath(origin, linkStruct.Link)
			finalLink = urlRes
		} else {
			// This is an absolute path
			link := linkStruct.Link
			if len(link) < len(origin) || bytes.Compare([]byte(link[:len(origin)]), []byte(origin)) != 0 {
				// Ignore paths not from origin
				continue
			} else {
				finalLink = link
			}
		}

		if visited[finalLink] {
			// visited don't add to avoid cycle
			continue
		}

		dfs(finalLink, origin, maxDepth-1)
	}
}

func main() {
	startURL := "https://www.calhoun.io/"
	maxDepth := 3
	flag.StringVar(&startURL, "startURL", startURL, "The origin URL to return the sitemap for.")
	flag.IntVar(&maxDepth, "maxDepth", maxDepth, "The maximum depth to run the sitempap for.")
	flag.Parse()
	dfs(startURL, startURL, maxDepth)
	sitemap := urlset{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9", Url: make([]Url, 0)}
	for visitedURL := range visited {
		sitemap.Url = append(sitemap.Url, Url{Loc: visitedURL})
	}
	data, err := xml.MarshalIndent(sitemap, "", " ")
	if err != nil {
		fmt.Println("Error when parsing marshalling.")
		fmt.Println(err)
	}
	fmt.Println(xml.Header + string(data))
}
