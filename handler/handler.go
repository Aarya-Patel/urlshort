package handler

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
	handler := func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		if redirectUrl, ok := pathsToUrls[urlPath]; ok {
			fmt.Printf("Redirecting URL from %s -> %s\n", urlPath, redirectUrl)
			http.Redirect(w, r, redirectUrl, http.StatusFound)
			return
		}

		fmt.Printf("Default handler!\n")
		fallback.ServeHTTP(w, r)

	}

	return handler
}

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
	pathsToUrls, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	mapHandler := MapHandler(pathsToUrls, fallback)
	return mapHandler, nil
}

type pathToUrlMapping struct {
	Path string
	Url  string
}

func parseYAML(yml []byte) (map[string]string, error) {
	// Read the YAML and store it in the array
	var mappings []pathToUrlMapping
	err := yaml.Unmarshal(yml, &mappings)
	if err != nil {
		return nil, err
	}

	// Create a map with path -> url
	pathsToUrls := make(map[string]string)
	for _, mapping := range mappings {
		pathsToUrls[mapping.Path] = mapping.Url
	}

	return pathsToUrls, nil
}
