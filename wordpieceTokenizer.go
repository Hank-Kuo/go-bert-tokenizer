package go_bert_tokenizer

import (
	"strings"
)

const DefaultMaxWordChars = 200
const DefaultUnknownToken = "[UNK]"

type WordpieceTokenizer struct {
	Vocab        *Vocab
	MaxWordChars int
	UnknownToken string
}

func NewWordpieceTokenizer(voc *Vocab) *WordpieceTokenizer {
	return &WordpieceTokenizer{
		Vocab:        voc,
		MaxWordChars: DefaultMaxWordChars,
		UnknownToken: DefaultUnknownToken,
	}
}

func (tkz *WordpieceTokenizer) Tokenize(text string) []string {

	var toks []string

	for _, tok := range tokenizeWhitespace(text) {
		char := strings.Split(tok, "")
		if len(char) > tkz.MaxWordChars {
			toks = append(toks, tkz.UnknownToken)
			continue
		}

		isBad := false
		start := 0
		sub_tokens := []string{}
		for start < len(char) {
			end := len(char)
			curSubStr := ""

			for start < end {
				substr := strings.Join(char[start:end], "")
				if start > 0 {
					substr = "##" + substr
				}

				if _, ok := tkz.Vocab.GetToken()[substr]; ok {
					curSubStr = substr
					break
				}
				end -= 1
			}
			if curSubStr == "" {
				isBad = true
				break
			}
			sub_tokens = append(sub_tokens, curSubStr)
			start = end
		}

		if isBad {
			toks = append(toks, tkz.UnknownToken)
		} else {
			toks = append(toks, sub_tokens...)
		}
	}

	return toks
}

func (tkz *WordpieceTokenizer) SetMaxWordChars(c int) {
	tkz.MaxWordChars = c
}

func (tkz *WordpieceTokenizer) SetUnknownToken(tok string) {
	tkz.UnknownToken = tok
}
