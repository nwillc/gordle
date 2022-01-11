package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLetter(t *testing.T) {
	l := NewLetter('c')
	assert.Equal(t, l.letter, 'c')
	assert.Equal(t, l.score, NONE)
}

func Test_score(t *testing.T) {
	type args struct {
		word   string
		target string
	}
	tests := []struct {
		name string
		args args
		want []Letter
	}{
		{
			name: "Exact Match",
			args: args{
				word:   "amber",
				target: "amber",
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
			name: "No Matches",
			args: args{
				word:   "amber",
				target: "quick",
			},
			want: []Letter{
				{letter: 'a', score: RED},
				{letter: 'm', score: RED},
				{letter: 'b', score: RED},
				{letter: 'e', score: RED},
				{letter: 'r', score: RED},
			},
		},
		{
			name: "Amber Match",
			args: args{
				word:   "amber",
				target: "flask",
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
			name: "Mix Green Amber Red",
			args: args{
				word:   "amber",
				target: "alert",
			},
			want: []Letter{
				{letter: 'a', score: GREEN},
				{letter: 'm', score: RED},
				{letter: 'b', score: RED},
				{letter: 'e', score: AMBER},
				{letter: 'r', score: AMBER},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := score(tt.args.word, tt.args.target)
			assert.NoError(t, err)
			assert.Equal(t, len(tt.want), len(s))
			for i, c := range s {
				assert.Equal(t, tt.want[i], *c, "pos %d", i)
			}
		})
	}
}
