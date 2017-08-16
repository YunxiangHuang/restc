package main

import (
	"log"
	"net/http"

	"github.com/thbourlove/restc"
)

func main() {
	client := restc.NewClient()
	usernames := []string{}
	req, err := http.NewRequest("GET", "https://api.github.com/users", nil)
	if err != nil {
		log.Fatalf("new request %v", err)
	}
	client.FetchJsonDataWithPath(req, &usernames, "$[*].login")
	log.Println(usernames)
}