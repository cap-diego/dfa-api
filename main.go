package main


import (
	"time"
	"os"
	"fmt"
	"context"
	"os/signal"
	"net/http"
	"encoding/json"
	"github.com/cap-diego/dfa-minimization-algorithm"
)


func minimizeAutomata(w http.ResponseWriter, r *http.Request) {
	var M dfa.DFA
	if r.Method != http.MethodPost {
		http.Error(w, "error, method post expected", 400)
		return
	}
	if r.Body == nil {
		http.Error(w, "error, body expected", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&M)
	if err != nil {
		http.Error(w, fmt.Sprintf("error while decoding dfa, %s\n", err.Error()), 404)
		return
	}
	Min := dfa.HopcroftDFAMin(M)
	json.NewEncoder(w).Encode(Min)
}

func main() {
	fmt.Print("Inicializando server\n")
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
    context.WithTimeout(context.Background(), 10 * time.Second)
}