/*
 * Copyright (c) 2022, nwillc@gmail.com
 *
 * Permission to use, copy, modify, and/or distribute this software for any purpose with or
 * without fee is hereby granted, provided that the above copyright notice and this permission
 * notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
 * REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
 * AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
 * INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
 * LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
 * OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
 * PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
)

type (
	Score  string
	Letter struct {
		letter rune
		score  Score
	}
	ColorFunc func(...interface{})
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
	colorMap = map[Score]ColorFunc{
		RED:   color.New(color.FgBlack, color.BgRed).PrintFunc(),
		AMBER: color.New(color.FgBlack, color.BgYellow).PrintFunc(),
		GREEN: color.New(color.FgBlack, color.BgGreen).PrintFunc(),
		NONE:  color.New(color.FgWhite, color.BgBlack).PrintFunc(),
	}
)

func main() {
	var (
		alphabet                         = stringToLetters("abcdefghijklmnopqrstuvwxyz")
		words    container.Slice[string] = strings.Split(dict, "\n")
		wordSet                          = container.ToSet(words)
		target                           = words.Random()
		isGreen                          = func(l *Letter) bool { return l.score == GREEN }
		input                            = bufio.NewReader(os.Stdin)
		masks                            = container.NewDeque[container.Slice[*Letter]]()
	)

	display(alphabet)

	for attempt := 1; attempt < 7; attempt++ {
		fmt.Printf("Guess %d: ", attempt)
		word, _ := input.ReadString('\n')
		word = word[0 : len(word)-1]
		if len(word) == 0 || !wordSet.Contains(word) {
			fmt.Printf("%s not in word list.\n", word)
			attempt--
			continue
		}
		scores := score(word, target)
		masks.Add(container.Map(scores, func(letter *Letter) *Letter { return &Letter{letter: '_', score: letter.score} }))
		if scores.All(isGreen) {
			fmt.Printf("Got it in %d/6.\n", attempt)
			for _, m := range masks.Values() {
				display(m)
			}
			return
		}
		display(scores)
		alphabet = updateScores(alphabet, scores)
		display(alphabet)
	}
	fmt.Println("The word was:", target)
}

func score(word, target string) container.Slice[*Letter] {
	result := stringToLetters(word)
	var t container.Slice[rune] = []rune(target)
	for i, l := range result {
		l.score = RED
		if l.letter == t[i] {
			l.score = GREEN
		} else if t.Any(genfuncs.IsEqualComparable(l.letter)) {
			l.score = AMBER
		}
	}
	return result
}

func display(letters container.Slice[*Letter]) {
	for _, l := range letters {
		colorMap[l.score](string(l.letter))
	}
	fmt.Println()
}

func updateScores(alphabet, scores container.Slice[*Letter]) container.Slice[*Letter] {
	return container.Map(alphabet, func(l *Letter) *Letter {
		return container.Fold(scores, l, func(l *Letter, sl *Letter) *Letter {
			if l.score == GREEN || sl.letter != l.letter {
				return l
			}
			return sl
		})
	})
}

func stringToLetters(s string) container.Slice[*Letter] {
	return container.Map([]rune(s), func(r rune) *Letter { return &Letter{letter: r, score: NONE} })
}
