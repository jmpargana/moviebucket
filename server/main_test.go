package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/matryer/is"
)

func beforeEach(t *testing.T) (*is.I, *httptest.Server, sqlmock.Sqlmock, func()) {
	is := is.New(t)
	db, mock, err := sqlmock.New()
	is.NoErr(err)
	server := newServer(db)
	srv := httptest.NewServer(server.router)
	cleanup := func() {
		db.Close()
		srv.Close()
	}
	return is, srv, mock, cleanup
}

func TestGetHello(t *testing.T) {
	tt := []string{"john", "malcolm", "christina", "laksdj laksjd laksdfh laksdfh"}

	is, srv, mock, cleanup := beforeEach(t)
	defer cleanup()

	for _, tc := range tt {
		t.Run(fmt.Sprintf("test get hello: %s", tc), func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec("UPDATE counts (.+) WHERE id = 1").WillReturnResult(sqlmock.NewResult(0, 1))

			payload := []byte(fmt.Sprintf(`{"name": "%s"}`, tc))
			resp, err := http.Post(srv.URL+"/", "application/json", bytes.NewBuffer(payload))
			is.NoErr(err)
			is.Equal(resp.StatusCode, 200)
			data, err := ioutil.ReadAll(resp.Body)
			is.NoErr(err)
			is.Equal(string(data), fmt.Sprintf("\"hello %s\"\n", tc))
			is.NoErr(mock.ExpectationsWereMet())
		})
	}
}

func TestAddMovie(t *testing.T) {
	is, srv, mock, cleanup := beforeEach(t)
	defer cleanup()

	mock.ExpectExec(fmt.Sprintf("INSERT INTO movies (.+) VALUES ('%s')", "Harry Potter"))

	payload := []byte(fmt.Sprintf(`{"name":"%s"}`, "Harry Potter"))
	resp, err := http.Post(srv.URL+"/movies", "application/json", bytes.NewBuffer(payload))

	is.NoErr(err)
	is.Equal(resp.StatusCode, 200)
	is.NoErr(mock.ExpectationsWereMet())
}
