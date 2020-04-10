package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		for i, v := range pathsToUrls {
			if r.URL.String() == i {
				http.Redirect(w, r, v, http.StatusMovedPermanently)
			}
		}
		fallback.ServeHTTP(w, r)
	})
}

type PathtoURL struct {
	Path string
	Url  string
}

var pathToURL = map[string]string{}

func YAMLHandler(pathsToURLs map[string]string, fallback http.Handler) (http.HandlerFunc, error) {
	return MapHandler(pathsToURLs, fallback), nil
}

func YAMLParser(yml []byte) (map[string]string, error) {
	var p PathtoURL
	err := yaml.Unmarshal(yml, &p)
	if err != nil {
		return nil, err
	}
	pathsToURLs := buildMap(p)
	return pathsToURLs, nil
}

func buildMap(p PathtoURL) map[string]string {
	pathToURL[p.Path] = p.Url
	return pathToURL
}
