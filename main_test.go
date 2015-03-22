package main
import (
    "testing"
    "net/http"
)

func TestCanAccessStaticFile(t *testing.T) {
    StartServer()

    resp, err := http.Get("http://localhost:8080/view/app/index.html")
    if err != nil {
        t.Error("une erreur est intervenue", err)
    }

    if resp.StatusCode != 200 {
        t.Error("Le statut devrait être 200 mais : ", resp.StatusCode)
    }
}


func TestCanCatchNotFound(t *testing.T) {
    StartServer()

    resp, err := http.Get("http://localhost:8080/view/toto")
    if err != nil {
        t.Error("une erreur est intervenue", err)
    }

    if resp.StatusCode != 404 {
        t.Error("Le statut devrait être 404 mais : ", resp.StatusCode)
    }
}