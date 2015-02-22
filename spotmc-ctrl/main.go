package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/goura/spotmc-ctrl/handlers"
)

func checkEnv() (err error) {
	ac := handlers.NewAppConfig()
	if ac.AutoScalingGroupName == "" || ac.AWSRegion == "" || ac.BasicUsername == "" || ac.BasicPassword == "" {
		return fmt.Errorf("Set the neccessary env vars")
	}
	return
}

func logAndAuth(h http.HandlerFunc) (wrapped http.HandlerFunc) {
	wrapped = func(w http.ResponseWriter, r *http.Request) {
		handlers.AuthWrapper(h)(w, r)
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
	}
	return
}

func main() {
	err := checkEnv()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	http.HandleFunc("/", logAndAuth(handlers.TopPageHandler))
	http.HandleFunc("/start", logAndAuth(handlers.StartPageHandler))
	http.HandleFunc("/stop", logAndAuth(handlers.StopPageHandler))
	http.ListenAndServe("0.0.0.0:5000", nil)
}
