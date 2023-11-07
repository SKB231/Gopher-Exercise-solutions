package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	returnFunction := func(w http.ResponseWriter, r *http.Request) {
		requestURL := r.URL

		if finalURL, ok := pathsToUrls[requestURL.String()]; ok {
			http.Redirect(w, r, finalURL, 302)
		} else {
			fallback.ServeHTTP(w, r) // The request URL doesn't exist in the pathsToURLs. Thus we will have to use the fallback URL to do the job for us.
		}
	}
	return returnFunction
}

type urlKeyValuePair struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
type yamlArr []urlKeyValuePair

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	var yamlInput yamlArr
	err := yaml.Unmarshal(yml, &yamlInput)
	if err != nil {
		fmt.Println("Error processing the yaml file")
		fmt.Println(err)
		return nil, err
	}
	parsedData := make(map[string]string)
	for _, pair := range yamlInput {
		parsedData[pair.Path] = pair.Url
	}
	finalHandler := MapHandler(parsedData, fallback)
	return finalHandler, err
}
