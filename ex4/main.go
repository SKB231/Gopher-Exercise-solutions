package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/SKB231/gopherex/ex4/linkParser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("1 args needed")
		fmt.Printf("Usage: \n %v <url to parse> \n", os.Args[0])
		os.Exit(1)
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println("Error oopening file")
	}
	result, err := linkParser.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
