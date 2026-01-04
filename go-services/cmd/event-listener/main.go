package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", getHealth)
	mux.HandleFunc("POST /url-link-clicked", UrlLinkClickedEventHandler)

	// TODO get port from env
	err := http.ListenAndServe(":3081", mux)

	fmt.Println("Server started on port 3081")

	if err != nil {
		panic(err)
	}

}

type UrlLinkClickedEvent struct {
	UserId         string `json:"user_id"` /* uuid */
	Url            string `json:"url"`
	Timestamp      int64  `json:"timestamp"`
	EventEmittedAt string `json:"event_emitted_at"` /* ISO 8601 */
	Placement      string `json:"placement"`        /* header_link, footer_link, etc */
	BrowserName    string `json:"browser_name"`     /* eg Chrome */
	BrowserVersion string `json:"browser_version"`
	OsName         string `json:"os_name"`
	OsVersion      string `json:"os_version"`
	DeviceType     string `json:"device_type"`
	UserAgent      string `json:"user_agent"`
}

func UrlLinkClickedEventHandler(w http.ResponseWriter, r *http.Request) {
	var event UrlLinkClickedEvent

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	fmt.Println("event: ")
	fmt.Println(event)

	// TODO send event to kafka
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
