package main

import (
	"github.com/plouc/go-gitlab-client/hookci"
	"log"
	"net/http"
)

func main() {
	s, h := hookci.New("hookBuilder")

	http.Handle("/build/", s)
	go func() {
		for {
			hook, ok := <-h
			if ok {
				log.Println("Commit from: ", hook.UserName)
			} else {
				return
			}
		}
	}()

	log.Println("http://:9090/build/hookBuilder")
	log.Panic(http.ListenAndServe(":9090", nil))
}
