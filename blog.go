package main

import (
	"flag"
	"fmt"
	"go/build"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/tools/blog"
	"golang.org/x/tools/playground/socket"
)

const (
	packagePath = "github.com/pietv/pietv.pub"
)

var (
	httpFlag   = flag.String("http", "localhost:8080", "HTTP Listen Address")
	originFlag = flag.String("origin", "", "web socket origin URL for playground (e.g. localhost)")
	baseFlag   = flag.String("base", "", "base path for articles and resources")
)

func main() {
	flag.Parse()

	if *baseFlag == "" {
		p, err := build.Default.Import(packagePath, "", build.FindOnly)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't find blog package: %v\n", err)
			os.Exit(1)
		}
		*baseFlag = p.Dir
	}

	ln, err := net.Listen("tcp", *httpFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	host, port, err := net.SplitHostPort(*httpFlag)
	if err != nil {
		log.Fatal(err)
	}

	srv, err := blog.NewServer(blog.Config{
		ContentPath:  filepath.Join(*baseFlag, "articles"),
		TemplatePath: filepath.Join(*baseFlag, "templates"),
		Hostname:     host,
		HomeArticles: 4,
		FeedArticles: 4,
		FeedTitle:    "FeedAtom",
		PlayEnabled:  true,
	})
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	origin := &url.URL{Scheme: "http"}
	if *originFlag != "" {
		origin.Host = net.JoinHostPort(*originFlag, port)
	} else if ln.Addr().(*net.TCPAddr).IP.IsUnspecified() {
		name, _ := os.Hostname()
		origin.Host = net.JoinHostPort(name, port)
	} else {
		origin.Host = *httpFlag
	}

	http.Handle("/resources/", http.FileServer(http.Dir(*baseFlag)))
	http.Handle("/socket", socket.NewHandler(origin))
	http.Handle("/", srv)
	log.Fatal(http.Serve(ln, nil))
}
