package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {

	port := flag.Int("port", 8080, "port to serve on")
	ip := flag.String("ip", "0.0.0.0", "ip address to serve on")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		requestInfo := map[string]interface{}{
			"Method":           r.Method,
			"URL":              r.URL.String(),
			"Proto":            r.Proto,
			"Header":           r.Header,
			"Body":             string(requestDump),
			"ContentLength":    r.ContentLength,
			"Host":             r.Host,
			"RemoteAddr":       r.RemoteAddr,
			"RequestURI":       r.RequestURI,
			"TLS":              r.TLS,
			"TransferEncoding": r.TransferEncoding,
			"Form":             r.Form,
			"PostForm":         r.PostForm,
			"MultipartForm":    r.MultipartForm,
			"Trailer":          r.Trailer,
			"RequestTime":      time.Now().Format(time.RFC3339),
		}

		prettyJSON, err := json.MarshalIndent(requestInfo, "", "  ")
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		fmt.Fprintf(w, "%s", prettyJSON)
		// log.Printf("Received request:\n%s", prettyJSON)
	})

	addr := fmt.Sprintf("%s:%d", *ip, *port)
	log.Printf("Starting server on %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
