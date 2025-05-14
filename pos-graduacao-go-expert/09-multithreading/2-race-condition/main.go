package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		// m.Unlock()
		atomic.AddUint64(&number, 1)
		time.Sleep(300 * time.Second)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}
