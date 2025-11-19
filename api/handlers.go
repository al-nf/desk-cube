package main

import(
	"net/http"
)

func HandleSlack(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}

func HandleDiscord(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}

func HandleGmail(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}
