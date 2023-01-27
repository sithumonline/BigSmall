package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
)

const keyServerAddr = "serverAddr"

func getSmall(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /small request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Small is beautiful!\n")
}

func getBig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// print 1 to 50000
	for i := 0; i < 50000; i++ {
		fmt.Printf("%d, ", i)
	}

	fmt.Printf("%s: got /big request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Big is beautiful!\n")
}

func main() {
	// print number of cpu cores
	fmt.Printf("Number of CPU cores: %d\n", runtime.NumCPU())
	 fmt.Printf("GOMAXPROCS is set to: %d\n", runtime.GOMAXPROCS(0))

	http.HandleFunc("/small", getSmall)
	http.HandleFunc("/big", getBig)

	http.ListenAndServe(":3000", nil)
}

// GOMAXPROCS=1 go run main.go 
// wrk --rate=20 --duration=5s "http://localhost:8090/big" & wrk --rate=20 --duration=5s "http://localhost:8090/small"
