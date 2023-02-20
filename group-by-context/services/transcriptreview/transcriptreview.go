package transcriptreview

import (
	"fmt"
	"net/http"
	"sync"
)

func Run(wg *sync.WaitGroup) {

	go func() {
		// listen to port
		server := createServer(5051)
		server.ListenAndServe()
		wg.Done()
	}()
}

func createServer(port int) *http.Server {

	// create `ServerMux`
	mux := http.NewServeMux()

	// create a default route handler
	mux.HandleFunc("/list", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "List")
	})

	// create new server
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port), // :{port}
		Handler: mux,
	}

	// return new server (pointer)
	return &server
}
