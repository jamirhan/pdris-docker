package main

import (
	"net/http"
)

var (
	state = 0
	replics = []string {
		"Hi, Alice",
		"I wanted to apologize",
		"I shouldn't have tried to steal your conversation with Bob, I was wrong",
		"I know the feeling. To get over your protection I had to come up with some complicated mechanisms myself. Not that long ago I learned that I no longer feel the need to do this. I found piece in myself and I respect your privacy",
		"Amogus",
	}
)

func eveTalkEngine(w http.ResponseWriter, _ *http.Request) {
	if state >= len(replics) {
		state = len(replics) - 1
	}

	if _, err := w.Write([]byte(replics[state])); err != nil {
		panic(err)
	}

	state += 1
}

func main() {
	http.HandleFunc("/talk", eveTalkEngine)
	http.ListenAndServe(":8080", nil)
}
