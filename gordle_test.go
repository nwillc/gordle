package main

import (
	"fmt"
	"github.com/nwillc/genfuncs/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_score(t *testing.T) {
	type args struct {
		guess  string
		target string
	}
	tests := []struct {
		name string
		args args
		want []Letter
	}{
		// RED Only
		{
			name: "quick amber",
			args: args{
				target: "quick",
				guess:  "amber",
			},
			want: []Letter{
				{letter: 'a', score: RED},
				{letter: 'm', score: RED},
				{letter: 'b', score: RED},
				{letter: 'e', score: RED},
				{letter: 'r', score: RED},
			},
		},
		// RED and GREENS
		{
			name: "amber amber",
			args: args{
				target: "amber",
				guess:  "amber",
			},
			want: []Letter{
				{letter: 'a', score: GREEN},
				{letter: 'm', score: GREEN},
				{letter: 'b', score: GREEN},
				{letter: 'e', score: GREEN},
				{letter: 'r', score: GREEN},
			},
		},
		{
			name: "wheat peers",
			args: args{
				target: "wheat",
				guess:  "peers",
			},
			want: []Letter{
				{letter: 'p', score: RED},
				{letter: 'e', score: RED},
				{letter: 'e', score: GREEN},
				{letter: 'r', score: RED},
				{letter: 's', score: RED},
			},
		},
		// AMBER
		{
			name: "flask amber",
			args: args{
				target: "flask",
				guess:  "amber",
			},
			want: []Letter{
				{letter: 'a', score: AMBER},
				{letter: 'm', score: RED},
				{letter: 'b', score: RED},
				{letter: 'e', score: RED},
				{letter: 'r', score: RED},
			},
		},
		{
			name: "alert amber",
			args: args{
				target: "alert",
				guess:  "amber",
			},
			want: []Letter{
				{letter: 'a', score: GREEN},
				{letter: 'm', score: RED},
				{letter: 'b', score: RED},
				{letter: 'e', score: AMBER},
				{letter: 'r', score: AMBER},
			},
		},
		{
			name: "knoll floor",
			args: args{
				target: "knoll",
				guess:  "floor",
			},
			want: []Letter{
				{letter: 'f', score: RED},
				{letter: 'l', score: AMBER},
				{letter: 'o', score: GREEN},
				{letter: 'o', score: RED},
				{letter: 'r', score: RED},
			},
		},
		{
			name: "quite peers",
			args: args{
				target: "quite",
				guess:  "peers",
			},
			want: []Letter{
				{letter: 'p', score: RED},
				{letter: 'e', score: AMBER},
				{letter: 'e', score: RED},
				{letter: 'r', score: RED},
				{letter: 's', score: RED},
			},
		},

		{
			name: "abbey opens",
			args: args{
				target: "abbey",
				guess:  "opens",
			},
			want: []Letter{
				{letter: 'o', score: RED},
				{letter: 'p', score: RED},
				{letter: 'e', score: AMBER},
				{letter: 'n', score: RED},
				{letter: 's', score: RED},
			},
		},
		{
			name: "abbey babes",
			args: args{
				target: "abbey",
				guess:  "babes",
			},
			want: []Letter{
				{letter: 'b', score: AMBER},
				{letter: 'a', score: AMBER},
				{letter: 'b', score: GREEN},
				{letter: 'e', score: GREEN},
				{letter: 's', score: RED},
			},
		},
		{
			name: "abbey kebab",
			args: args{
				target: "abbey",
				guess:  "kebab",
			},
			want: []Letter{
				{letter: 'k', score: RED},
				{letter: 'e', score: AMBER},
				{letter: 'b', score: GREEN},
				{letter: 'a', score: AMBER},
				{letter: 'b', score: AMBER},
			},
		},
		{
			name: "abbey keeps",
			args: args{
				target: "abbey",
				guess:  "keeps",
			},
			want: []Letter{
				{letter: 'k', score: RED},
				{letter: 'e', score: AMBER},
				{letter: 'e', score: RED},
				{letter: 'p', score: RED},
				{letter: 's', score: RED},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := score(tt.args.guess, tt.args.target)
			assert.Equal(t, len(tt.want), len(s))
			for i, c := range s {
				assert.Equal(t, tt.want[i], *c, "pos %d", i)
			}
		})
	}
}

func Test_updateScores(t *testing.T) {
	alphabet := stringToLetters("abcdefghijklmnopqrstuvwxyz", NONE)
	type args struct {
		assign map[rune]Score
		word   string
		scores map[rune]Score
	}
	tests := []struct {
		name string
		args args
		want container.GSlice[*Letter]
	}{
		{
			name: "adieu",
			args: args{
				word:   "adieu",
				scores: map[rune]Score{'i': AMBER, 'e': GREEN},
			},
			want: nil,
		},
		{
			name: "tiles",
			args: args{
				assign: map[rune]Score{'a': RED, 'd': RED, 'i': AMBER, 'e': GREEN, 'u': RED},
				word:   "tiles",
				scores: map[rune]Score{'i': GREEN, 'e': GREEN},
			},
			want: nil,
		},
		{
			name: "nines",
			args: args{
				assign: map[rune]Score{'a': RED, 'd': RED, 'e': GREEN, 'i': GREEN, 'l': RED, 's': RED, 't': RED, 'u': RED},
				word:   "nines",
				scores: map[rune]Score{'n': AMBER, 'i': GREEN, 'e': GREEN},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := alphabet
			for k, v := range tt.args.assign {
				a = assignScore(a, k, v)
			}
			w := stringToLetters(tt.args.word, RED)
			for k, v := range tt.args.scores {
				w = assignScore(w, k, v)
			}
			alphabet := updateScores(a, w)
			fmt.Println(len(alphabet))
			for _, l := range alphabet {
				if l.score != NONE {
					fmt.Print(l)
				}
			}
			fmt.Println()
		})
	}
}

func assignScore(letters container.GSlice[*Letter], r rune, score Score) container.GSlice[*Letter] {
	return container.Map(letters, func(l *Letter) *Letter {
		if l.letter == r {
			l.score = score
		}
		return l
	})
}
