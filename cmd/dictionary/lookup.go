package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var ApiUrlBase = "https://api.dictionaryapi.dev/api/v2/entries/en/"

type Meaning struct {
	PartOfSpeech string
}

type Entry struct {
	Word     string
	Meanings []Meaning
}

func main() {
	input := os.Args[1]
	f, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
			return
		}
		word := scanner.Text()
		_, err := lookup(word)
		if err != nil {
			fmt.Println("REJECTED:", word)
		} else {
			fmt.Println("ACCEPTED:", word)
		}
	}
}

func lookup(word string) ([]Entry, error) {
	url := ApiUrlBase + word
	var response *http.Response
	var err error
	for {
		response, err = http.Get(url)
		if err != nil {
			return nil, err
		}
		if response.StatusCode != 429 {
			break
		}
		log.Printf("429 %s", response.Status)
		if body, err := ioutil.ReadAll(response.Body); err == nil {
			log.Print(string(body))
		}

		time.Sleep(time.Minute)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf(response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var payload []Entry
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
