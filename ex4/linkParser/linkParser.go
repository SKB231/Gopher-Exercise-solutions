package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func returnNodeType(node *html.Node) string {
	nodeTypeName := map[html.NodeType]string{0: "ErrorNode", 1: "TextNode", 2: "DocNode", 3: "ElementNode", 4: "CommentNode", 5: "DocTypeNode"}
	return nodeTypeName[node.Type]
}

// dfs method takes a single node, and a boolean and goes over every available node in the dfs Tree.
// Only set the childAnchor to true if the next node is the child of an anchor tag.
func dfs(node *html.Node, childOfAnchor bool, finalString []byte) []byte {
	if node == nil {
		// Should find a better way to do this for sure
		// Marks end of DFS.
		return finalString
	}
	if childOfAnchor && node.Type == 1 {
		finalText := strings.TrimLeftFunc(node.Data, func(c rune) bool { return c == ' ' || c == '\n' })
		// finalText = strings.TrimLeft(finalText, "\n")
		if len(finalText) > 0 {
			// fmt.Println("Node Type: ", returnNodeType(node), " and data: ", finalText, " of length ", len(finalText))
			finalTextAsBytes := []byte(finalText)
			// fmt.Println("Before ", string(finalString))
			finalString = append(finalString, finalTextAsBytes...)
			// fmt.Println("After ", string(finalString))

		}
	} else {
		// fmt.Println("Node Data but not child of href: ", node.Data, " of type ", returnNodeType(node))
	}
	var resultString []byte
	if !childOfAnchor {
		resultString = make([]byte, 1)
	} else {
		resultString = finalString
	}
	nextChildIsAnchor := false
	if node.Type == 3 && node.Data == "a" {
		// We have reached an anchor tag. Extract the available url and text element.
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				fmt.Println("Link is ", attr.Val)
			}
		}
		nextChildIsAnchor = true
	}

	// Iterate over all children
	if nextChildIsAnchor {
		resultString = dfs(node.FirstChild, childOfAnchor || nextChildIsAnchor, resultString)
	} else if childOfAnchor {
		// fmt.Println("Before sending to child ")
		// fmt.Println(finalString)
		finalString = dfs(node.FirstChild, childOfAnchor || nextChildIsAnchor, finalString)
	} else {
		dfs(node.FirstChild, childOfAnchor || nextChildIsAnchor, make([]byte, 0))
	}

	if nextChildIsAnchor {
		fmt.Println(resultString)
		fmt.Println(string(resultString))
	}

	// Iterate over all siblings. If nil, then will stop at the beginning of the function.
	// fmt.Println("Before sending to nextSibling ", string(finalString))
	finalString = dfs(node.NextSibling, childOfAnchor, finalString)
	// fmt.Println(string(finalString) + " XXX")
	return finalString
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(" Usage: \n ", os.Args[0], " ex<number>.html")
		return
	}
	exNum := os.Args[1]
	absPath, _ := filepath.Abs("./testHTML/ex" + exNum + ".html")
	fmt.Println(absPath)
	file, err := os.Open(absPath)
	if err != nil {
		fmt.Println("Error opening file. ")
		fmt.Println(err)
	}
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println("Error parsing document")
		fmt.Println(err)
	}
	fmt.Println("Root doc", returnNodeType(doc))
	dfs(doc, false, make([]byte, 0))
}
