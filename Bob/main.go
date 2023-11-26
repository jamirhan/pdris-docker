package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type speakerType string

const (
	alice speakerType = "Alice"
	eve speakerType = "Eve"
)

func (s *speakerType) swap() {
	if *s == alice {
		*s = eve
		return
	} 
	*s = alice
}

var (
	dialogue []string
	eveURL string
	aliceURL string
	currentSpeaker speakerType
)

func spy(w http.ResponseWriter, _ *http.Request) {
	client := http.Client{}

	fuckUp := func(err error) {
		w.Write([]byte("Sorry, we fucked up. Ask the admin to check the logs. 500 or whatever"))
		panic(err)
	}

	var urlBase string
	switch currentSpeaker {
	case alice:
		urlBase = aliceURL
	case eve:
		urlBase = eveURL
	}

	finalURL, err := url.JoinPath(urlBase, "talk")
	if err != nil {
		fuckUp(err)
	}

	resp, err := client.Get(finalURL)
	if err != nil {
		fuckUp(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fuckUp(err)
	}

	fmt.Println(body)

	if resp.StatusCode != http.StatusOK {
		fuckUp(errors.New(""))
	}

	dialogue = append(dialogue, fmt.Sprintf("%s: %s", currentSpeaker, (body)))
	currentSpeaker.swap()

	_, err = w.Write([]byte(strings.Join(dialogue, "\n")))
	if err != nil {
		fuckUp(err)
	}
}

func main() {
	var found bool
	eveURL, found = os.LookupEnv("EVE_URL")
	if !found {
		panic("EVE_URL env not found")
	}
	aliceURL, found = os.LookupEnv("ALICE_URL")
	if !found {
		panic("ALICE_URL env not found")
	}
	currentSpeaker = eve

	http.HandleFunc("/spy", spy)
	http.ListenAndServe(":8080", nil)
}
