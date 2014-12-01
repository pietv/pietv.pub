package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"net/url"

	"golang.org/x/tools/blog"
	"golang.org/x/tools/playground/socket"
)

const hostname = "pietv.pub"

var (
	httpFlag = flag.String("http", "localhost:8080", "HTTP Listen Address")
)

func main() {
	flag.Parse()

	srv, err := blog.NewServer(blog.Config{
		ContentPath:  "articles",
		TemplatePath: "templates",
		Hostname:     hostname,
		HomeArticles: 4,
		FeedArticles: 4,
		FeedTitle:    "FeedAtom",
		PlayEnabled:  true,
	})
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	ln, err := net.Listen("tcp", *httpFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	url, _ := url.Parse("http://" + *httpFlag)
	http.Handle("/resources/", http.FileServer(http.Dir(".")))
	http.Handle("/socket", socket.NewHandler(url))
	http.Handle("/", srv)
	log.Fatal(http.Serve(ln, nil))
}
