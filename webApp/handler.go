package webApp

import (
	"net/http"
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *Request)
}
