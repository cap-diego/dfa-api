package main


import (
	"time"
	"context"
	"os/signal"
	"os"
	"fmt"
	"net/http"
	dfa "github.com/cap-diego/dfa-minimization-algorithm"
)


func minimizeAutomata(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	fmt.Print("Inicializando server\n")
	fmt.Printf("%#v\n", dfa)
	http.HandleFunc("/minimize", minimizeAutomata)
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan // Block 
	fmt.Print("Terminate: ", sig)
    context.WithTimeout(context.Background(), 30 * time.Second)
}