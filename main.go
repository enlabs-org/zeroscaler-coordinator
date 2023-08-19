package main

import (
	"encoding/json"
	"net/http"
)

type Wakeup struct {
	Url     string
	Service string
}

func main() {
	http.HandleFunc("/wakeup", func(w http.ResponseWriter, r *http.Request) {

		var wakeup Wakeup
		err := json.NewDecoder(r.Body).Decode(&wakeup)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte(wakeup.Url))
	})

	http.ListenAndServe(":80", nil)
}
