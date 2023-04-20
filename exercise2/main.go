package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/jezpoz/gophercises/exercise2/handler"
)

func main() {
	filePtr := flag.String("file", "redirects.yaml", "-f [relative path to .yaml file]")
	flag.Parse()

	file, err := os.Open(*filePtr)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	bytes := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bytes)
	if err != nil {
		panic(err)
	}

	mux := defaultMux()

	// Fallback redirects
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yamlHandler, err := handler.YAMLHandler(bytes, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}
