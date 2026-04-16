package requests

type Method string

var Methods = methods{
	OPTIONS: "OPTIONS",
	GET:     "GET",
	HEAD:    "HEAD",
	POST:    "POST",
	PUT:     "PUT",
	DELETE:  "DELETE",
	TRACE:   "TRACE",
	CONNECT: "CONNECT",
}

type methods struct {
	OPTIONS Method
	GET     Method
	HEAD    Method
	POST    Method
	PUT     Method
	DELETE  Method
	TRACE   Method
	CONNECT Method
}
