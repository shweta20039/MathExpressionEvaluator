package main

import "strings"

type Tokenizer struct {
	tokens []rune
	pos    int
}

func NewTokenizer(s string) *Tokenizer {
	return &Tokenizer{tokens: []rune(strings.ReplaceAll(s, " ", "")), pos: 0}
}

func (t *Tokenizer) hasNext() bool {
	return t.pos < len(t.tokens)
}

func (t *Tokenizer) current() rune {
	return t.tokens[t.pos]
}

func (t *Tokenizer) next() {
	if t.hasNext() {
		t.pos++
	}
}
