package api

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

func newTS(fileName string) *httptest.Server {
	return httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.OpenFile(fileName, os.O_RDONLY, 0)
		if err != nil {
			panic(err)
		}
		reader := bufio.NewReader(f)
		resp, err := http.ReadResponse(reader, r)
		if err != nil {
			panic(err)
		}
		for k, v := range resp.Header {
			w.Header().Add(k, strings.Join(v, " "))
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(resp.StatusCode)

		w.Write(body)
	}))
}
