package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestGetHello(t *testing.T) {
	tt := []string{"john", "malcolm", "christina", "laksdj laksjd laksdfh laksdfh"}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("test get hello: %s", tc), func(t *testing.T) {
			is := is.New(t)
			server := newServer()
			// db, cleanup := connectTestDatabase()
			// defer cleanup()

			srv := httptest.NewServer(server.router)
			defer srv.Close()

			payload := []byte(fmt.Sprintf(`{"name": "%s"}`, tc))
			resp, err := http.Post(srv.URL+"/", "application/json", bytes.NewBuffer(payload))
			is.NoErr(err)
			is.Equal(resp.StatusCode, 200)
			data, err := ioutil.ReadAll(resp.Body)
			is.NoErr(err)
			is.Equal(string(data), fmt.Sprintf("\"hello %s\"\n", tc))
		})
	}

}
