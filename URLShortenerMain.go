package main

import (
	"fmt"
	"log"
	"net/http"

	"clayclaw.me/urlshortener/services"
)

const port = "8080"

func main() {

	fmt.Printf("GCI, made with love <3\n")
	services.RegisterURL("boss", "https://www.jboss.org/")
	services.RegisterURL("google", "https://www.google.com/")
	services.RegisterURL("duckduckgo", "https://duckduckgo.com/")
	services.RegisterURL("illegalsite", "https://thisisILLEGAL.youhaveneverseenbefore/")

	services.RegisterBanURL("illegalsite")
	// fmt.Printf("TEST1: %t", services.IsURLRegistered("illegalsite"))
	// fmt.Printf("TEST2: %t", services.IsBanned("illegalsite"))

	http.HandleFunc("/", redirect)

	fmt.Printf("Http server initialized")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func redirect(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("\nFOUND: " + r.URL.Path[1:])

	if services.IsURLRegistered(r.URL.Path[1:]) {

		if !services.IsBanned(r.URL.Path[1:]) {
			http.Redirect(w, r, services.GetURL(r.URL.Path[1:]), 301)
		} else {
			http.Redirect(w, r, "https://imgur.com/HNWaaAr", 301)
		}

	} else {
		http.Redirect(w, r, "https://imgur.com/qTpzOFU", 301)
	}
	

}
