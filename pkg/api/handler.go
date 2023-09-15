package api

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.

/*
	Need this function to take in a map, and based on path in
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
*/

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		new_url, exists := pathsToUrls[r.URL.Path]
		if exists {
			http.Redirect(w, r, new_url, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	})
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

/*
	yaml := `
	- path: /urlshort
	 url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
	 url: https://github.com/gophercises/urlshort/tree/solution
	`
*/

type Config struct {
	Path string
	Url  string
}

func BuildMap(configs []Config) map[string]string {
	urlMap := map[string]string{}

	for _, config := range configs {
		urlMap[config.Path] = config.Url
	}

	return urlMap
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	var configs []Config

	err := yaml.Unmarshal(yml, &configs)
	if err != nil {
		return nil, err
	}

	pathMap := BuildMap(configs)
	return MapHandler(pathMap, fallback), nil
}

// Write SQLHandler function
func SQLHandler(strMap map[string]string, fallback http.Handler) (http.HandlerFunc, error) {

	// var configs []Config

	return nil, nil

}
