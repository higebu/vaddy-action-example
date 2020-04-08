package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	verificationCode := os.Getenv("VADDY_VERIFICATION_CODE")
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8888"
	}

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}
	h3 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, verificationCode)
	}
	xss := func(w http.ResponseWriter, req *http.Request) {
		query := req.URL.Query()
		if q, ok := query["q"]; ok {
			io.WriteString(w, fmt.Sprintf("<html><div>%s</div></html>", q))
		} else {
			io.WriteString(w, "<html>no params</html>")
		}
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)
	http.HandleFunc(fmt.Sprintf("/vaddy-%s.html", verificationCode), h3)
	http.HandleFunc("/xss.html", xss)

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
