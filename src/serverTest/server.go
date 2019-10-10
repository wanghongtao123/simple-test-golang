package servertest

import (
	"net/http"
	"fmt"
)

// 实现方法和interface的分离
type PlayerServer struct {
    store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    player := r.URL.Path[len("/players/"):]
    score := p.store.GetPlayerScore(player)
    if score == 0 {
        w.WriteHeader(http.StatusNotFound)
    }
    fmt.Fprint(w, score)
}

type PlayerStore interface {
    GetPlayerScore(name string) int
}
