package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/dimfeld/httptreemux"
)

func main() {
	router := httptreemux.New()
	router.GET("/view/*file", staticHandler)

	fmt.Println("server up ...")
	go http.ListenAndServe(":8080", router)

	waitSignal()
}

func waitSignal() {
	c := make(chan int)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		for range signalChan {
			c <- 0
		}
	}()

	<-c
	fmt.Println("... server down !")
}

func staticHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	prefix := "../../../../../timeline/"
	file := params["file"]
	fmt.Println("file requested : " + file)

	data, err := ioutil.ReadFile(prefix + file)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if strings.HasSuffix(file, ".css") {
		w.Header().Set("Content-Type", "text/css")
	}

	w.Write(data)
}
