package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/nwillc/genfuncs/gentype"
)

type (
	Score  string
	Letter struct {
		letter rune
		score  Score
	}
)

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

func main() {
	var (
		alphabet = stringToLetters("abcdefghijklmnopqrstuvwxyz")
		words    = strings.Split(dict, "\n")
		wordMap  = gentype.NewMapSet[string]()
		rnd      = rand.New(rand.NewSource(time.Now().Unix()))
		target   = ""
		isGreen  = func(l *Letter) bool { return l.score == GREEN }
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
		scores := score(word, target)
		if scores.All(isGreen) {
			fmt.Println("Got it in", attempt)
			return
		}
		display(scores)
		alphabet = update(alphabet, scores)
		display(alphabet)
	}

	fmt.Println("The word was:", target)
}

func score(word, target string) gentype.Slice[*Letter] {
	result := stringToLetters(word)
	var t gentype.Slice[rune] = []rune(target)
	for i, l := range result {
		l.score = RED
		if l.letter == t[i] {
			l.score = GREEN
		} else if t.Any(func(r rune) bool { return r == l.letter }) {
			l.score = AMBER
		}
	}
	return result
}

func display(letters gentype.Slice[*Letter]) {
	for _, l := range letters {
		fmt.Print(colorMap[l.score], string(l.letter))
	}
	fmt.Println(color.Reset)
}

func update(alphabet, scores gentype.Slice[*Letter]) gentype.Slice[*Letter] {
	return gentype.Map(alphabet, func(l *Letter) *Letter {
		return gentype.Fold(scores, l, func(l *Letter, sl *Letter) *Letter {
			if l.score == GREEN || sl.letter != l.letter {
				return l
			}
			return sl
		})
	})
}

func stringToLetters(s string) gentype.Slice[*Letter] {
	return gentype.Map([]rune(s), func(r rune) *Letter { return &Letter{letter: r, score: NONE} })
}
