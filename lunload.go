package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func usage() {
	myName := os.Args[0]
	fmt.Println("Usage:\n", myName, "URL to file")
}

func main() {
	if len(os.Args) > 1 {
		var (
			fileName    string
			fullURLFile string
		)
		fullURLFile = os.Args[1]
		fileURL, err := url.Parse(fullURLFile)
		if err != nil {
			log.Fatal(err)
		}
		path := fileURL.Path
		segments := strings.Split(path, "/")
		fileName = segments[len(segments)-1]

		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		client := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}

		resp, err := client.Get(fullURLFile)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		size, err := io.Copy(file, resp.Body)
		defer file.Close()
		fmt.Printf("Downloaded a file %s with size %d", fileName, size)
	} else {
		usage()
	}
}
