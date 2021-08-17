package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var stringReplacements = map[string]string{
	"abn":       "ABN Amro",
	"ing":       "ING Bank",
	"rabo":      "Rabobank",
	"triodos":   "Triodos Bank",
	"volksbank": "de Volksbank",
}

type Input struct {
	Text string `json:"text"`
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", inputParser)

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

// This functions receives an Input struct, from which it uses the Text field.
// The text is then split into a slice, word by word. Each word is then inspected.

func replaceText(text Input) string {
	wordsSlice := strings.Split(text.Text, " ")
	var storageSlice []string

	for _, word := range wordsSlice {
		// We only want words which contain only alphabetical characters.
		reg, err := regexp.Compile("[^a-zA-Z]+")
		if err != nil {
			log.Fatal(err)
		}
		// We do this in order to catch corner cases where we could have punctuation marks at the end of a word we'd want to pick up; e.g. "I worked at ING."
		processedString := strings.ToLower(reg.ReplaceAllString(word, ""))

		if val, ok := stringReplacements[processedString]; ok {
			storageSlice = append(storageSlice, val)
		} else {
			storageSlice = append(storageSlice, word)
		}
	}
	result := strings.Join(storageSlice, " ")
	return result
}

func inputParser(w http.ResponseWriter, r *http.Request) {
	// Declare a new Input struct.
	var input Input

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with the new text which has replaced the strings.
	w.Write([]byte(replaceText(input)))
}
