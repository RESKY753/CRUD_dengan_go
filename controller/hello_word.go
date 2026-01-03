package controller

import "net/http"

func NewHelloWordController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello word"))
	}
}