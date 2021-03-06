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
	Init("127.0.0.1", "timeline", "events", "tags")
	StartServer()
	waitSignal()
}

func StartServer() {
	router := httptreemux.New()
	router.GET("/view/*file", staticHandler)
	router.GET("/ws/events/:start/:end/tags/*tags", GetEventsHandler)
	router.GET("/ws/events/:start/:end/tags", GetEventsHandler)
	router.GET("/ws/events/:start/:end", GetEventsHandler)
	router.POST("/ws/events", AddEventsHandler)
	router.GET("/ws/events/byname/:name", SearchEventsHandler)
	router.GET("/ws/tags", GetTagsHandler)
	router.GET("/ws/tags/:query", GetTagsHandler)
	router.POST("/ws/tags", AddTagHandler)

	fmt.Println("server up ...")
	go http.ListenAndServe(":8080", router)
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
	prefix := "./web/"
	file := params["file"]
	fmt.Println("File requested : ", file)

	data, err := ioutil.ReadFile(prefix + file)
	if err != nil {
		fmt.Println("not found ", file)
		http.NotFound(w, r)
		return
	}

	configureMimeType(w, file)
	w.Write(data)
}

func configureMimeType(w http.ResponseWriter, fileName string) {
	type t struct {
		extension string
		mimeType  string
	}

	l := []t{
		t{".css", "text/css"},
		t{".js", "text/javascript"},
		t{".html", "text/html"},
	}

	for _, t := range l {
		if strings.HasSuffix(fileName, t.extension) {
			w.Header().Set("Content-Type", t.mimeType)
			return
		}
	}
}
