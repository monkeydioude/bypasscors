package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/monkeydioude/moon"
)

func main() {
	m := moon.Moon()
	m.AddHeader("Access-Control-Allow-Origin", "*")
	m.MakeRouter(
		moon.Get("/{url}", func(r *moon.Request) ([]byte, int, error) {
			resp, err := http.Get(r.Matches["url"])
			if err != nil {
				return []byte("ko"), 400, nil
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
				return []byte("ko"), 400, nil
			}
			return body, 200, nil
		}),
	)
	moon.ServerRun(":8080", m)
}
