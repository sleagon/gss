package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	dir  = "."
	port = int64(6789)
)

func parseArgs() {
	if len(os.Args) < 2 {
		return
	}
	i := 1
	for i < len(os.Args)-1 {
		if os.Args[i] == "-p" {
			port, _ = strconv.ParseInt(os.Args[i+1], 10, 64)
			i += 2
			continue
		}
		break
	}
	// unknown arguments
	if i < len(os.Args)-1 {
		panic("unknown arguments")
	}
	// dir argument
	if i == len(os.Args)-1 && os.Args[len(os.Args)-1] != "" {
		dir = os.Args[len(os.Args)-1]
	}
}

func main() {
	if err := recover(); err != nil {
		log.Println("Usage\r\n [-p port] [dir]")
		log.Fatal(err)
	}
	// parse arguments
	parseArgs()
	fmt.Printf("start to serve %s over port %d\n", dir, port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir))))
}
