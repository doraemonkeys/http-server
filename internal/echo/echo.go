package echo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

// Start 启动 HTTP Echo 服务器
// 接收任意请求，将请求的完整信息以 JSON 格式返回
func Start(ip string, port int) {
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
	})

	addr := fmt.Sprintf("%s:%d", ip, port)
	log.Printf("Starting echo server on %s\n", addr)
	if ip != "0.0.0.0" {
		fmt.Printf("Try to access http://%s:%d\n", ip, port)
	} else {
		fmt.Printf("Try to access http://127.0.0.1:%d\n", port)
	}
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
