package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/SKB231/gopherex/ex4/linkParser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("2 args needed")
		os.Exit(1)
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println("Error oopening file")
	}
	result := linkParser.Parse(resp.Body)
	fmt.Println(result)
}
