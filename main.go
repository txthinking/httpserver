package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	s := ":8080"
	if len(os.Args) == 2 {
		if strings.Contains(os.Args[1], "h") {
			fmt.Println("$ httpserver [:8080]")
			return
		}
		s = os.Args[1]
	}
	fmt.Println("listen http on", s)
	log.Println(http.ListenAndServe(s, http.FileServer(http.Dir("."))))
}
