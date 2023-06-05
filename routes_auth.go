package main

import (
	"fmt"
	"net/http"
)

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to Goflix")
	}
}
