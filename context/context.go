package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
}

func NewHandler(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}
