package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/SKB231/gopherex/ex4/linkParser"
)

var visited map[string]bool = make(map[string]bool)

// DFS function takes a string as the root node.
// It assumes that the URL already contains the correct origin. It performs a GET request on the link, extracts all the urls in the page, and runs the DFS on those methods as well.
// Using a map, it ensures that it doesen't visit the same link twice.
func dfs(root string, origin string) {
	if root == "" {
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

		dfs(finalLink, origin)
	}
}

func main() {
	startURL := "https://www.calhoun.io/"
	dfs(startURL, startURL)
	for visitedURL := range visited {
		fmt.Println(visitedURL)
	}
}
