package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	apiHost = "http://localhost:3001"
	apiPath = "/portal-api/oas-doc"
	apiURL  = "http://localhost:3001/portal-api/oas-doc"
	token   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQcm92aWRlcklEIjoiMCIsIlVzZXJJRCI6IiQyYSQxMCRodEtZYUphOHRxdVFMcWV2VWdLWThPaER5dUc2RndXeVZCV0k5d0ZrNFRVSHdyL3ltZy4xQyJ9.xl3idg-g7WF4cNNMWXKWwEsMRHQ_b-l29ool3Uq85zI"
)

func main() {
	mux := http.NewServeMux()

	// serve static files
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	// serve the oas document generated
	mux.HandleFunc("/docs/oas", func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}

		req, _ := http.NewRequest("GET", apiURL, nil)
		req.Header.Add("Authorization", token)

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			log.Fatal("ServeHTTP:", err)
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})

	if err := http.ListenAndServe(":3007", mux); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}
}
