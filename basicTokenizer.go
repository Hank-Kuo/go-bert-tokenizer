package go_bert_tokenizer

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

type BasicTokenizer struct {
	Lower bool
}

func NewBasicTokenizer(lower bool) *BasicTokenizer {
	return &BasicTokenizer{Lower: lower}
}

func (bt *BasicTokenizer) Tokenize(text string) []string {
	text = clean(text)
	text = padChinese(text)

	var toks []string
	for _, tok := range tokenizeWhitespace(text) {
		if bt.Lower {
			tok = strings.ToLower(tok)
			tok = stripAccents(tok)
		}
		toks = append(toks, splitPunc(tok)...)
	}

	toks = tokenizeWhitespace(strings.Join(toks, " "))

	return toks
}

func clean(text string) string {
	var b strings.Builder
	for _, c := range text {

		if c == 0 || c == 0xfffd || isControl(c) {
			continue
		} else if isWhitespace(c) {
			b.WriteRune(' ')
		} else {
			b.WriteRune(c)
		}
	}

	return b.String()
}

func stripAccents(text string) string {
	var b strings.Builder
	for _, c := range norm.NFD.String(text) {
		if !unicode.Is(unicode.Mn, c) {
			b.WriteRune(c)
		}
	}
	return b.String()
}

func splitPunc(text string) []string {
	var toks []string
	var b strings.Builder
	for _, c := range text {
		if isPunctuation(c) {
			toks = append(toks, b.String())
			toks = append(toks, string(c))
			b.Reset()
		} else {
			b.WriteRune(c)
		}
	}
	if b.Len() > 0 {
		toks = append(toks, b.String())
	}
	return toks
}

func tokenizeWhitespace(text string) []string {
	split := strings.Split(text, " ")
	var toks []string
	for _, tok := range split {
		if tok != "" {
			toks = append(toks, tok)
		}
	}
	return toks
}

func padChinese(text string) string {
	var b strings.Builder
	for _, c := range text {
		if isChinese(c) {
			b.WriteRune(' ')
			b.WriteRune(c)
			b.WriteRune(' ')
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}
