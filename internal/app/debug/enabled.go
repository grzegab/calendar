//go:build debug

package debug

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

const port = ":6060"

func Start() {
	go func() {
		fmt.Printf("starting pprof on %s\n", port)
		if err := http.ListenAndServe(port, nil); err != nil {
			fmt.Printf("pprof server error: %v\n", err)
		}
	}()
}

func StartMemProfile() {
	fmt.Println("starting memory profiling")
}

func StopMemProfile() {
	fmt.Println("stopping memory profiling")
}
