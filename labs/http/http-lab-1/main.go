package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]time.Time)
	tz := r.URL.Query().Get("tz")
	hasError := false
	if tz == "" {
		m["current_time"] = time.Now().UTC()
	} else if locations := strings.Split(tz, ","); len(locations) > 1 {
		for _, l := range locations {
			loc, err := time.LoadLocation(l)
			if err != nil {
				hasError = true
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
			} else {
				m[l] = time.Now().In(loc)
			}
		}
	} else {
		loc, _ := time.LoadLocation(tz)
		m["current_time"] = time.Now().In(loc)
	}

	if !hasError {
		w.Header().Add("Content-Type", "application/json")

		json.NewEncoder(w).Encode(m)
	}
}

func main() {
	http.HandleFunc("/api/time", handler)

	http.ListenAndServe(":8080", nil)
}
