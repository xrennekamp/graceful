package graceful_test

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/xrennekamp/graceful"
)

func Example() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		fmt.Fprintln(w, "world")
	})

	srv := &http.Server{Addr: "localhost:8080"}

	fin := finish.New()
	fin.Add(srv)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	fin.Wait()
}
