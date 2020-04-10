package main

import (
	"fmt"
	"log"
	"net/http"
	"url-handler/urlshort"
)

var pathsToURLs = map[string]string{
	"/eg":         "https://gobyexample.com/",
	"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	"/tour":       "https://tour.golang.org/welcome/1",
}

var p1 = `
        path: /blog
        url: https://blog.golang.org/
    `
var p2 = `
        path: /http
        url: https://golang.org/pkg/net/http/
    `
var pathToURLsYAML = []string{p1, p2}

func main() {
	mux := defaultMux()
	mapHandler := urlshort.MapHandler(pathsToURLs, mux)
	yamlHandler, err := getYAMLHandler(mapHandler)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", printWelcome)
	return mux
}

func printWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Welcome to URL Shortner")
}

func getYAMLHandler(mapHandler http.HandlerFunc) (yamlHandler http.HandlerFunc, err error) {
	var URLs = map[string]string{}
	for _, pathToURL := range pathToURLsYAML {
		var err error
		URLs, err = urlshort.YAMLParser([]byte(pathToURL))
		if err != nil {
			log.Println(err)
		}
	}
	yamlHandler, err = urlshort.YAMLHandler(URLs, mapHandler)
	if err != nil {
		return nil, err
	}
	return yamlHandler, err
}
