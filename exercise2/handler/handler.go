package handler

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYaml)

	return MapHandler(pathMap, fallback), err
}

type Redirects struct {
	Path string
	Url  string
}

func parseYAML(data []byte) ([]Redirects, error) {
	parsed := []Redirects{}

	err := yaml.Unmarshal(data, &parsed)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func buildMap(parsedYaml []Redirects) map[string]string {
	mapped := make(map[string]string)
	for _, m := range parsedYaml {
		mapped[m.Path] = m.Url
	}
	return mapped
}
