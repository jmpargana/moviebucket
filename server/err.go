package main

type Err struct {
	trace      string
	status     int
	code       string `json:"code"`
	message    string `json:"message"`
	statusCode int    `json:"status"`
}

var errors = map[string]*Err{
	"movie_decode": &Err{
		status:     400,
		code:       "movie_decode",
		message:    "Could not decode movie payload",
		statusCode: 10001,
	},
}
