package main

import (
	"net/http"
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"time"
)

var startedAt = time.Now()

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	w.Write([]byte("Hello " + name + " you are " + age + " years old"))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	w.Write([]byte("hello " + username + " your password is " + password))
}

func Family(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/myfamily/members.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Fprintf(w, string(data))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)

	if duration.Seconds() < 10 || duration.Seconds() > 30 {
		w.WriteHeader(500)
		w.Write([]byte("Not healthy"))
	}else {
		w.WriteHeader(200)
		w.Write([]byte("Healthy"))
	}
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/family", Family)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":80", nil)
}