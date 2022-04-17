package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var UnixTime struct {
	HumanReadable string `json:"ansic"`
	Unixmilli     int64  `json:"unixmilli"`
}
var CmdArguments = os.Args

func uTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	utime := UnixTime

	utime.HumanReadable = time.Now().Format(time.ANSIC)
	utime.Unixmilli = time.Now().UnixMilli()

	json.NewEncoder(w).Encode(utime)
}

func handleRequests() {
	http.HandleFunc("/test", uTime)

	if len(os.Args) <= 1 {
		fmt.Printf("\nRunning http server on localhost:8000\n\nIf you need to run on a different ip address or port press crtl+c \nrun time-api <ip address>:<port> \n ")
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	} else {
		fmt.Println("Running http server on  ", CmdArguments[1])
		log.Fatal(http.ListenAndServe(CmdArguments[1], nil))
	}
}

func main() {
	handleRequests()
}
