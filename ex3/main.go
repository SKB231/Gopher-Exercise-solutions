package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type jsonText interface{}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []option `json:"options"`
}
type adventureBook map[string]arc

func ParseStory() adventureBook {
	storyJsonFile, err := os.ReadFile("gopher.json")
	if err != nil {
		fmt.Println("Error reading story file")
		fmt.Println(err)
		os.Exit(1)
	}
	var book adventureBook = make(adventureBook)

	json.Unmarshal(storyJsonFile, &book)
	return book
}

var book adventureBook

func handleBaseRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	t, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Println("Error when returning file => ", err)
	}
	f, err := os.ReadFile("style.css")
	if err != nil {
		fmt.Println("Error parsing CSS file")
		fmt.Println(err)
	}
	data := map[string]interface{}{
		"arcTitle": book["intro"].Title,
		"story":    book["intro"].Story,
		"options":  book["intro"].Options,
		"Style":    template.CSS(f),
	}
	fmt.Println(data)
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Error when returning file ", err)
	}
}

func main() {
	book = ParseStory()
	mux := http.NewServeMux() // Create a new Request handlers. Serve mux implements http.Handler since it has serve HTTP method. But also has the added benifit of matching the request URL path with the correct handler function.
	mux.Handle("/", http.HandlerFunc(handleBaseRequest))
	fmt.Println("Listening to port 8080")
	http.ListenAndServe(":8080", mux) // Listen and Serve Port 8080 using the serve mux we created above
}
