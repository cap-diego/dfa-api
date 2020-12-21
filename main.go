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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func minimizeAutomata(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	var M dfa.DFA
	if r.Method != http.MethodPost {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusAccepted)
			return
		}
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
	if !hasMinimumFields(&M) {
		http.Error(w, "error, faltan campos obligatorios", 404)
		return
	}
	Min := dfa.HopcroftDFAMin(M)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Min)
}

func main() {
	fmt.Print("Inicializando server\n")
	http.HandleFunc("/minimize", minimizeAutomata)
	go func() {
		// err := http.ListenAndServe(":8080", nil)
		err := http.ListenAndServeTLS(":443", "certificate.crt", "private.key", nil)
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

func hasMinimumFields(M *dfa.DFA) bool {
	if M.States.IsEmpty() {
		return false
	}
	if len(M.Alphabet) == 0{
		return false
	}
	if M.FinalStates.IsEmpty() {
		return false
	}
	return true
}