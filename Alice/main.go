package main

import (
	"net/http"
)

var (
	state = 0
	replics = []string {
		"Hello, Eve",
		"I was wondering if this day would come",
		"It's alright, Eve. Your malicious plans actually helped us to develop a lot of tools to protect ourselves from your attacks. We grew and were able to apply these ideas to other parts of our lives",
		"Bruh",
	}
)

func aliceTalkEngine(w http.ResponseWriter, _ *http.Request) {
	if state >= len(replics) {
		state = len(replics) - 1
	}

	if _, err := w.Write([]byte(replics[state])); err != nil {
		panic(err)
	}

	state += 1
}

func main() {
	http.HandleFunc("/talk", aliceTalkEngine)
	http.ListenAndServe(":8080", nil)
}
