package go_bert_tokenizer

import (
	"strings"
)

const DefaultMaxWordChars = 200
const DefaultUnknownToken = "[UNK]"

type WordpieceTokenizer struct {
	vocab        *Vocab
	maxWordChars int
	unknownToken string
}

func NewWordpieceTokenizer(voc *Vocab) *WordpieceTokenizer {
	return &WordpieceTokenizer{
		vocab:        voc,
		maxWordChars: DefaultMaxWordChars,
		unknownToken: DefaultUnknownToken,
	}
}

func (wp *WordpieceTokenizer) Tokenize(text string) []string {

	var toks []string

	for _, tok := range tokenizeWhitespace(text) {
		char := strings.Split(tok, "")
		if len(char) > wp.maxWordChars {
			toks = append(toks, wp.unknownToken)
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

				if _, ok := wp.vocab.GetToken()[substr]; ok {
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
			toks = append(toks, wp.unknownToken)
		} else {
			toks = append(toks, sub_tokens...)
		}
	}

	return toks
}

func (wp *WordpieceTokenizer) SetMaxWordChars(c int) {
	wp.maxWordChars = c
}

func (wp *WordpieceTokenizer) SetUnknownToken(tok string) {
	wp.unknownToken = tok
}
