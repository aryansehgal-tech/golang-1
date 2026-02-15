package main 

import (
	"fmt"     
	"log"      
	"log/slog" // Structured logging (key-value style logs, newer than log)
	"net/http" 
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHello)
	mux.HandleFunc("/goodbye", handleGoodbye)
	// nil means "use the default ServeMux" (the default router)
	// log.Fatal will log an error AND stop the program if server fails
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// w = ResponseWriter (used to send data back to client)
// r = Request (contains request info like headers, body, etc.)
// We ignore r here using _ because we don’t need it
func handleHello(w http.ResponseWriter, _ *http.Request) {

	// Write sends bytes as HTTP response to the client
	// []byte converts string to byte slice (Write needs bytes)
	// wc = number of bytes written
	// err = error if something went wrong
	wc, err := w.Write([]byte("Hello World\n"))

	// Always check errors when writing responses
	if err != nil {
		slog.Error("Error writing response", "err", err)
		return 
	}
	fmt.Printf("%d bytes written\n", wc)
}

func handleGoodbye(w http.ResponseWriter, _ *http.Request) {
	wc, err := w.Write([]byte("GoodBye"))

	if err != nil {
		slog.Error("Error writing response", "err", err)
		return
	}

	fmt.Printf("%d bytes written\n", wc)
}