package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/nwillc/genfuncs/gentype"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Score string

const (
	RED   Score = "R"
	AMBER Score = "A"
	GREEN Score = "G"
	NONE  Score = " "
)

var (
	//go:embed data/dict.txt
	dict     string
	colorMap = map[Score]string{
		RED:   color.Red,
		AMBER: color.Yellow,
		GREEN: color.Green,
		NONE:  color.Reset,
	}
)

type Letter struct {
	letter rune
	score  Score
}

func main() {
	var (
		words    = strings.Split(dict, "\n")
		wordMap  = gentype.NewMapSet[string]()
		alphabet = gentype.Map(
			[]rune("abcdefghijklmnopqrstuvwxyz"),
			func(r rune) *Letter { return NewLetter(r) })
		rnd     = rand.New(rand.NewSource(time.Now().Unix()))
		match   = true
		target  = ""
		isGreen = func(l *Letter) bool { return l.score == GREEN }
	)

	for _, word := range words {
		wordMap.Add(word)
	}

	target = words[rnd.Intn(len(words))]

	display(alphabet)
	var reader = bufio.NewReader(os.Stdin)

	for attempt := 1; attempt < 7; attempt++ {
		fmt.Printf("Guess %d: ", attempt)
		word, _ := reader.ReadString('\n')
		word = word[0 : len(word)-1]
		if !wordMap.Contains(word) {
			fmt.Printf("%s not in word list.\n", word)
			attempt--
			continue
		}
		scores, err := score(word, target)
		if err != nil {
			fmt.Println(err)
			attempt--
			continue
		}
		match = scores.All(isGreen)
		display(scores)
		if match {
			fmt.Println("Got it in", attempt)
			break
		}
		update(alphabet, scores)
		display(alphabet)
	}

	if !match {
		fmt.Println("The word was", target)
	}
}

func NewLetter(r rune) *Letter {
	return &Letter{letter: r, score: NONE}
}

func score(word, target string) (gentype.Slice[*Letter], error) {
	if len(word) != len(target) {
		return nil, fmt.Errorf("length mismatch")
	}
	result := make([]*Letter, len(word))
	w := []rune(word)
	t := []rune(target)
	for i, r := range w {
		result[i] = NewLetter(r)
		result[i].score = RED
		if w[i] == t[i] {
			result[i].score = GREEN
		} else {
			for _, tc := range t {
				if tc == w[i] {
					result[i].score = AMBER
					break
				}
			}
		}
	}
	return result, nil
}

func display(letters gentype.Slice[*Letter]) {
	for _, l := range letters {
		fmt.Print(colorMap[l.score], string(l.letter))
	}
	fmt.Println(color.Reset)
}

func update(alphabet, scores []*Letter) {
	for _, l := range scores {
		offset := l.letter - rune('a')
		if alphabet[offset].score != GREEN {
			alphabet[offset].score = l.score
		}
	}
}
