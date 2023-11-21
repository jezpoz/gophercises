package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	cyoa "github.com/jezpoz/gophercises/exercise3"
)

func main() {
	port := flag.Int("port", 3000, "the port the app is going to use")
	file := flag.String("file", "gopher.json", "the JSON with the CYOA story")
	flag.Parse()
	fmt.Printf("Using story %s\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	tpl := template.Must(template.New("").Parse(storyTemplate))
	h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl), cyoa.WithPathFn(pathFun))
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Start the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFun(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var storyTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
    <title>Choose your own adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
      {{range .Options}}
      <li>
        <a href="/story/{{.Chapter}}">{{.Text}}</a>
      </li>
      {{end}}
    </ul>
  </body>
</html>`
