package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (h *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Goodbye")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "IT'S BROKEN!!!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Goodbye %s", d)
}
