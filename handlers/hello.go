package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) SayHello(w http.ResponseWriter, r *http.Request) {
	h.l.Println("hola mundo genial")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "IT'S BROKEN!!!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", d)
}
