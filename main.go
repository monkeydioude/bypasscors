package main

import (
	"io"
	"log"
	"net/http"

	"github.com/monkeydioude/moon"
)

func main() {
	m := moon.Moon()
	m.AddHeader("Access-Control-Allow-Origin", "*")
	m.MakeRouter(
		moon.Get("/bypasscors/{url}", func(r *moon.Request) ([]byte, int, error) {
			url := r.Matches["url"]
			if r.HTTPRequest.URL.RawQuery != "" {
				url += "?" + r.HTTPRequest.URL.RawQuery
			}
			resp, err := http.Get(url)

			if err != nil {
				log.Fatalln(err)
				return []byte("ko"), 400, nil
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
				return []byte("ko"), 400, nil
			}
			log.Printf("OK call to %s\n", url)
			return body, 200, nil
		}),
	)
	moon.ServerRun(":8080", m)
}
